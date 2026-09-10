use flate2::read::GzDecoder;
use onig::{Regex, SearchOptions};
use serde::{Deserialize, Serialize};
use sha2::{Digest, Sha256};
use std::{
    collections::BTreeMap,
    error::Error,
    fs,
    hint::black_box,
    io::{self, BufRead, Read, Write},
    path::{Path, PathBuf},
    sync::Arc,
    time::{Duration, Instant},
};
use tokenizers::{
    pre_tokenizers::split::{Split, SplitPattern},
    OffsetReferential, OffsetType, PreTokenizedString, PreTokenizer, SplitDelimiterBehavior,
};

type Result<T> = std::result::Result<T, Box<dyn Error + Send + Sync>>;
#[derive(Deserialize)]
struct Pattern {
    pattern: String,
    pattern_sha256: String,
    fixture_sha256: String,
    #[serde(default)]
    fixture_pattern: Option<String>,
    #[serde(default)]
    full_coverage: bool,
}
#[derive(Deserialize)]
struct Manifest {
    schema: u32,
    patterns: BTreeMap<String, Pattern>,
    corpora: Vec<Corpus>,
}
#[derive(Deserialize)]
struct Corpus {
    id: String,
    file: String,
    sha256: String,
    compressed_sha256: String,
    bytes: usize,
    batches: Vec<BatchSpec>,
}
#[derive(Deserialize)]
struct BatchSpec {
    size: usize,
    file: String,
    sha256: String,
    requests: usize,
}
#[derive(Deserialize)]
struct Request {
    start: usize,
    end: usize,
}
struct Batch {
    data: Arc<[u8]>,
    requests: Vec<(usize, usize)>,
    input_sha256: String,
    capacity: usize,
    matches: usize,
    verified: bool,
    bytes: usize,
    spans: usize,
}
#[derive(Deserialize)]
struct Command {
    op: String,
    #[serde(default)]
    corpus: String,
    #[serde(default)]
    size: usize,
    #[serde(default)]
    engine: String,
    #[serde(default = "default_ms")]
    min_ms: u64,
}
fn default_ms() -> u64 {
    500
}
#[derive(Serialize)]
struct Sample {
    iterations: u64,
    elapsed_ns: u128,
    checksum: u64,
    bytes: usize,
    requests: usize,
    spans: usize,
}
fn digest(data: &[u8]) -> String {
    format!("{:x}", Sha256::digest(data))
}
fn read_checked(path: &Path, expected: &str) -> Result<Vec<u8>> {
    let data = fs::read(path)?;
    if digest(&data) != expected {
        return Err(format!("hash mismatch: {}", path.display()).into());
    }
    Ok(data)
}
fn gunzip(path: &Path, expected: &str) -> Result<Vec<u8>> {
    let compressed = read_checked(path, expected)?;
    let mut data = Vec::new();
    GzDecoder::new(&compressed[..]).read_to_end(&mut data)?;
    Ok(data)
}
fn hf_spans(split: &Split, src: &str) -> Result<Vec<(usize, usize)>> {
    let mut pre = PreTokenizedString::from(src);
    split.pre_tokenize(&mut pre)?;
    Ok(pre
        .get_splits(OffsetReferential::Original, OffsetType::Byte)
        .iter()
        .map(|(_, span, _)| *span)
        .collect())
}
fn onig_spans(re: &Regex, src: &str) -> Vec<(usize, usize)> {
    let mut out = Vec::new();
    let mut previous = 0;
    for (a, b) in re.find_iter(src) {
        if a > previous {
            out.push((previous, a))
        }
        out.push((a, b));
        previous = b;
    }
    if previous < src.len() {
        out.push((previous, src.len()))
    }
    out
}
fn at_spans(re: &Regex, src: &str) -> Result<Vec<(usize, usize)>> {
    let mut out = Vec::new();
    let mut pos = 0;
    while pos < src.len() {
        let n = re
            .match_with_options(src, pos, SearchOptions::SEARCH_OPTION_NONE, None)
            .ok_or("match-at failed")?;
        if n == 0 {
            return Err("empty match".into());
        }
        out.push((pos, pos + n));
        pos += n;
    }
    Ok(out)
}
fn equal(
    engine: &str,
    request: &str,
    src: &str,
    got: &[(usize, usize)],
    want: &[(usize, usize)],
) -> Result<()> {
    if got != want {
        let at = (0..got.len().min(want.len()))
            .find(|&i| got[i] != want[i])
            .unwrap_or(got.len().min(want.len()));
        return Err(format!("{engine}: {request}: first differing span {at}, got {:?}, expected {:?}, input prefix {:?}",got.get(at),want.get(at),src.chars().take(120).collect::<String>()).into());
    }
    Ok(())
}
#[derive(Deserialize)]
struct Fixture {
    input: String,
    matches: Vec<FixtureSpan>,
    split: Vec<FixtureSpan>,
}
#[derive(Deserialize)]
struct FixtureSpan {
    #[serde(rename = "Start")]
    start: usize,
    #[serde(rename = "End")]
    end: usize,
}
fn main() -> Result<()> {
    let args: Vec<String> = std::env::args().collect();
    if args.len() != 4 {
        return Err("usage: pretok-reference CORPUS_DIR ORIGINAL_FIXTURES_GZ PATTERN_ID".into());
    }
    let root = PathBuf::from(&args[1]);
    let manifest: Manifest = serde_json::from_slice(&fs::read(root.join("manifest.json"))?)?;
    if manifest.schema != 3 {
        return Err("unsupported corpus schema".into());
    }
    let selected = manifest
        .patterns
        .get(&args[3])
        .ok_or("pattern missing from corpus manifest")?;
    if digest(selected.pattern.as_bytes()) != selected.pattern_sha256 {
        return Err("pattern hash mismatch".into());
    }
    let re = Regex::new(&selected.pattern)?;
    let split = Split::new(
        SplitPattern::Regex(selected.pattern.clone()),
        SplitDelimiterBehavior::Isolated,
        false,
    )?;
    let match_split = Split::new(
        SplitPattern::Regex(selected.pattern.clone()),
        SplitDelimiterBehavior::Removed,
        true,
    )?;
    // The wheel oracle and the newly resolved native build must agree before any timing.
    let guard_pattern = selected
        .fixture_pattern
        .as_deref()
        .unwrap_or(&selected.pattern);
    let guard_re = Regex::new(guard_pattern)?;
    let guard_split = Split::new(
        SplitPattern::Regex(guard_pattern.to_owned()),
        SplitDelimiterBehavior::Isolated,
        false,
    )?;
    let fixture_data = gunzip(Path::new(&args[2]), &selected.fixture_sha256)?;
    let mut fixture_count = 0;
    for line in fixture_data
        .split(|b| *b == b'\n')
        .filter(|s| !s.is_empty())
    {
        let f: Fixture = serde_json::from_slice(line)?;
        let want: Vec<_> = f.split.iter().map(|s| (s.start, s.end)).collect();
        equal(
            "hf-split",
            "original fixture",
            &f.input,
            &hf_spans(&guard_split, &f.input)?,
            &want,
        )?;
        let matches: Vec<_> = f.matches.iter().map(|s| (s.start, s.end)).collect();
        equal(
            "onig-iter",
            "original fixture",
            &f.input,
            &guard_re.find_iter(&f.input).collect::<Vec<_>>(),
            &matches,
        )?;
        equal(
            "onig-at",
            "original fixture",
            &f.input,
            &at_spans(&guard_re, &f.input)?,
            &matches,
        )?;
        fixture_count += 1;
    }
    let mut batches = BTreeMap::new();
    for corpus in manifest.corpora {
        let data = gunzip(&root.join(&corpus.file), &corpus.compressed_sha256)?;
        if data.len() != corpus.bytes || digest(&data) != corpus.sha256 {
            return Err("corpus mismatch".into());
        }
        let data: Arc<[u8]> = data.into();
        let text = std::str::from_utf8(&data)?;
        for spec in corpus.batches {
            let cases: Vec<Request> =
                serde_json::from_slice(&read_checked(&root.join(&spec.file), &spec.sha256)?)?;
            if cases.len() != spec.requests {
                return Err("request count mismatch".into());
            }
            let mut requests = Vec::new();

            let mut byte_count = 0;
            let mut previous = 0;
            for case in cases {
                let src = text
                    .get(case.start..case.end)
                    .ok_or("invalid request range")?;
                if case.start != previous {
                    return Err("non-contiguous request ranges".into());
                }
                previous = case.end;
                byte_count += src.len();
                requests.push((case.start, case.end));
            }
            if previous != text.len() {
                return Err("incomplete request coverage".into());
            }
            batches.insert(
                (corpus.id.clone(), spec.size),
                Batch {
                    data: Arc::clone(&data),
                    requests,
                    input_sha256: corpus.sha256.clone(),
                    capacity: 0,
                    matches: 0,
                    verified: false,
                    bytes: byte_count,
                    spans: 0,
                },
            );
        }
    }
    let stdout = io::stdout();
    let mut output = stdout.lock();
    writeln!(
        output,
        "{}",
        serde_json::json!({"ready":true,"fixtures":fixture_count,"fixture_pattern":guard_pattern,"full_coverage":selected.full_coverage,"batches":batches.len(),"tokenizers":"0.22.2"})
    )?;
    output.flush()?;
    for line in io::stdin().lock().lines() {
        let command: Command = serde_json::from_str(&line?)?;
        if command.op == "quit" {
            break;
        }
        if command.op == "reference" {
            let key = (command.corpus, command.size);
            let batch = batches.get_mut(&key).ok_or("unknown corpus/size")?;
            let result = reference(
                batch,
                &re,
                &split,
                &match_split,
                &key.0,
                key.1,
                &selected.pattern_sha256,
                selected.full_coverage,
            );
            match result {
                Ok(value) => writeln!(output, "{}", serde_json::to_string(&value)?)?,
                Err(err) => writeln!(output, "{}", serde_json::json!({"error":err.to_string()}))?,
            }
            output.flush()?;
            continue;
        }
        let result = if command.engine == "onig-at" && !selected.full_coverage {
            Err("match-at diagnostic requires a fully covering pattern".into())
        } else if command.op == "compile" {
            compile_sample(&selected.pattern, &command.engine, command.min_ms)
        } else {
            let batch = batches
                .get(&(command.corpus, command.size))
                .ok_or("unknown corpus/size")?;
            sample(batch, &re, &split, &command.engine, command.min_ms)
        };
        match result {
            Ok(value) => writeln!(output, "{}", serde_json::to_string(&value)?)?,
            Err(err) => writeln!(output, "{}", serde_json::json!({"error":err.to_string()}))?,
        }
        output.flush()?;
    }
    Ok(())
}
#[derive(Serialize)]
struct ReferenceRequest {
    start: usize,
    end: usize,
    matches: Vec<(usize, usize)>,
    split: Vec<(usize, usize)>,
}
#[derive(Serialize)]
struct Reference {
    corpus: String,
    size: usize,
    pattern_sha256: String,
    input_sha256: String,
    requests: Vec<ReferenceRequest>,
}
fn reference(
    batch: &mut Batch,
    re: &Regex,
    split: &Split,
    match_split: &Split,
    corpus: &str,
    size: usize,
    pattern_sha: &str,
    full_coverage: bool,
) -> Result<Reference> {
    batch.verified = false;
    let mut requests = Vec::new();
    let mut total = 0;
    let mut match_count = 0;
    let mut capacity = 0;
    for (i, &(start, end)) in batch.requests.iter().enumerate() {
        let src = std::str::from_utf8(&batch.data[start..end])?;
        let matches = hf_spans(match_split, src)?;
        let spans = hf_spans(split, src)?;
        let label = format!("{corpus}/{size}/{i}");
        equal("onig-split", &label, src, &onig_spans(re, src), &spans)?;
        equal(
            "onig-iter",
            &label,
            src,
            &re.find_iter(src).collect::<Vec<_>>(),
            &matches,
        )?;
        // Match-at is currently a diagnostic for the fully covering P01/P04/P05 plans.
        if full_coverage {
            equal("onig-at", &label, src, &at_spans(re, src)?, &matches)?;
        }
        total += spans.len();
        match_count += matches.len();
        capacity = capacity.max(spans.len());
        requests.push(ReferenceRequest {
            start,
            end,
            matches,
            split: spans,
        });
    }
    batch.spans = total;
    batch.matches = match_count;
    batch.capacity = capacity;
    batch.verified = true;
    Ok(Reference {
        corpus: corpus.to_owned(),
        size,
        pattern_sha256: pattern_sha.to_owned(),
        input_sha256: batch.input_sha256.clone(),
        requests,
    })
}

fn consume(acc: u64, a: usize, b: usize) -> u64 {
    acc.wrapping_add(a as u64).rotate_left(7) ^ (b as u64)
}
fn run(
    batch: &Batch,
    re: &Regex,
    split: &Split,
    engine: &str,
    scratch: &mut Vec<(usize, usize)>,
) -> Result<u64> {
    let mut sum = 0;
    for &(start, end) in &batch.requests {
        // A timed, explicit UTF-8 input contract equivalent to pretok's public API.
        let src = std::str::from_utf8(black_box(&batch.data[start..end]))?;
        match engine {
            "onig-iter" => {
                for (a, b) in re.find_iter(src) {
                    sum = consume(sum, a, b)
                }
            }
            "onig-at" => {
                let mut pos = 0;
                while pos < src.len() {
                    let n = re
                        .match_with_options(src, pos, SearchOptions::SEARCH_OPTION_NONE, None)
                        .ok_or("match-at failed")?;
                    if n == 0 {
                        return Err("empty match".into());
                    }
                    sum = consume(sum, pos, pos + n);
                    pos += n;
                }
            }
            "onig-reused" | "onig-fresh" => {
                let mut fresh = Vec::new();
                let dst = if engine == "onig-reused" {
                    &mut *scratch
                } else {
                    &mut fresh
                };
                dst.clear();
                let mut end = 0;
                for (a, b) in re.find_iter(src) {
                    if a > end {
                        dst.push((end, a))
                    }
                    dst.push((a, b));
                    end = b;
                }
                if end < src.len() {
                    dst.push((end, src.len()))
                }
                black_box(&*dst);
                sum = consume(sum, dst.len(), end);
            }
            "hf-split" => {
                let mut pre = PreTokenizedString::from(src);
                split.pre_tokenize(&mut pre)?;
                let spans = pre.get_splits(OffsetReferential::Original, OffsetType::Byte);
                black_box(&spans);
                sum = consume(sum, spans.len(), src.len());
            }
            _ => return Err(format!("unknown engine {engine}").into()),
        }
    }
    Ok(black_box(sum))
}
fn sample(batch: &Batch, re: &Regex, split: &Split, engine: &str, min_ms: u64) -> Result<Sample> {
    if min_ms == 0 {
        return Err("min_ms must be positive".into());
    }
    if !batch.verified {
        return Err("reference generation must precede measurement".into());
    }
    let mut scratch = Vec::with_capacity(batch.capacity);
    black_box(run(batch, re, split, engine, &mut scratch)?);
    let start = Instant::now();
    let mut iterations = 0;
    let mut checksum = 0;
    loop {
        checksum ^= run(batch, re, split, engine, &mut scratch)?;
        iterations += 1;
        if start.elapsed() >= Duration::from_millis(min_ms) {
            break;
        }
    }
    Ok(Sample {
        iterations,
        elapsed_ns: start.elapsed().as_nanos(),
        checksum,
        bytes: batch.bytes,
        requests: batch.requests.len(),
        spans: if engine == "onig-iter" || engine == "onig-at" {
            batch.matches
        } else {
            batch.spans
        },
    })
}
fn compile_sample(pattern: &str, engine: &str, min_ms: u64) -> Result<Sample> {
    if min_ms == 0 {
        return Err("min_ms must be positive".into());
    }
    let start = Instant::now();
    let mut iterations = 0;
    loop {
        for _ in 0..256 {
            match engine {
                "onig-compile" => {
                    black_box(Regex::new(black_box(pattern))?);
                }
                "hf-compile" => {
                    black_box(Split::new(
                        SplitPattern::Regex(black_box(pattern).to_owned()),
                        SplitDelimiterBehavior::Isolated,
                        false,
                    )?);
                }
                _ => return Err("unknown compile engine".into()),
            }
            iterations += 1;
        }
        if start.elapsed() >= Duration::from_millis(min_ms) {
            break;
        }
    }
    Ok(Sample {
        iterations,
        elapsed_ns: start.elapsed().as_nanos(),
        checksum: 0,
        bytes: pattern.len(),
        requests: 1,
        spans: 0,
    })
}
