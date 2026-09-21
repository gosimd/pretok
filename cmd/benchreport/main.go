// benchreport runs offline correctness gates and sequential cross-engine samples.
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"slices"
	"strings"
	"time"

	"github.com/gosimd/pretok/internal/compile"
	regex "github.com/gosimd/pretok/internal/regexp"
	"github.com/gosimd/pretok/internal/regexp/bench"
)

type native struct {
	cmd *exec.Cmd
	in  io.WriteCloser
	out *bufio.Reader
}

func startNative(executable, corpus, fixture, patternID string) (*native, map[string]any, error) {
	cmd := exec.Command(executable, corpus, fixture, patternID)
	cmd.Stderr = os.Stderr
	in, err := cmd.StdinPipe()
	if err != nil {
		return nil, nil, err
	}
	out, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	if err = cmd.Start(); err != nil {
		return nil, nil, err
	}
	n := &native{cmd, in, bufio.NewReader(out)}
	var ready map[string]any
	line, err := n.out.ReadBytes('\n')
	if err == nil {
		err = json.Unmarshal(line, &ready)
	}
	if err != nil || ready["ready"] != true {
		n.close()
		return nil, nil, fmt.Errorf("native correctness gate failed: %s: %v", line, err)
	}
	return n, ready, nil
}
func (n *native) close() {
	n.in.Close()
	if n.cmd.Process != nil {
		n.cmd.Process.Kill()
	}
	n.cmd.Wait()
}
func (n *native) call(command, result any) error {
	if err := json.NewEncoder(n.in).Encode(command); err != nil {
		return err
	}
	line, err := n.out.ReadBytes('\n')
	if err != nil {
		return fmt.Errorf("native runner ended: %w", err)
	}
	var response struct {
		Error string `json:"error"`
	}
	if err = json.Unmarshal(line, &response); err != nil {
		return err
	}
	if response.Error != "" {
		return fmt.Errorf("native: %s", response.Error)
	}
	return json.Unmarshal(line, result)
}
func (n *native) sample(command any) (bench.Sample, error) {
	var s bench.Sample
	err := n.call(command, &s)
	if err == nil && (s.Iterations == 0 || s.ElapsedNS == 0) {
		err = fmt.Errorf("invalid native timing response")
	}
	return s, err
}
func (n *native) reference(b *bench.Batch) (bench.Reference, error) {
	var ref bench.Reference
	err := n.call(map[string]any{"op": "reference", "corpus": b.ID, "size": b.Size}, &ref)
	return ref, err
}
func command(name string, args ...string) string {
	out, err := exec.Command(name, args...).CombinedOutput()
	if err != nil {
		return strings.TrimSpace(string(out)) + " (" + err.Error() + ")"
	}
	return strings.TrimSpace(string(out))
}
func fileSHA(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return "unavailable: " + err.Error()
	}
	return bench.Digest(data)
}
func sourceSHA(root string) string {
	var paths []string
	filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if path != root && (d.Name() == ".git" || d.Name() == "research" || d.Name() == "target" || d.Name() == "results" || d.Name() == "runs" || d.Name() == "bin") {
				return filepath.SkipDir
			}
			return nil
		}
		if slices.Contains([]string{".go", ".rs", ".py", ".toml"}, filepath.Ext(path)) || d.Name() == "Cargo.lock" || d.Name() == "go.mod" {
			paths = append(paths, path)
		}
		return nil
	})
	slices.Sort(paths)
	var contents []byte
	for _, p := range paths {
		relative, _ := filepath.Rel(root, p)
		data, err := os.ReadFile(p)
		if err == nil {
			contents = append(contents, relative...)
			contents = append(contents, 0)
			contents = append(contents, data...)
			contents = append(contents, 0)
		}
	}
	return bench.Digest(contents)
}
func writeJSON(path string, value any) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(data, '\n'), 0644)
}
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "benchreport:", err)
		os.Exit(1)
	}
}
func run() error {
	patternID := flag.String("pattern", "P05", "exact corpus pattern ID: P01–P25")
	root := flag.String("root", ".", "repository root")
	corpus := flag.String("corpus", "benchmark/corpus", "corpus path relative to root")
	reference := flag.String("native", "benchmark/native/target/release/pretok-reference", "native executable path relative to root")
	output := flag.String("output", "benchmark/runs/local", "new result directory relative to root (ignored by Git by default)")
	samples := flag.Int("samples", 10, "samples per engine and workload")
	minTime := flag.Duration("min-time", 500*time.Millisecond, "minimum time per sample")
	filter := flag.String("filter", "", "substring of corpus ID to measure (all corpora are verified)")
	cpuLabel := flag.String("cpu-label", "", "verified CPU description if host discovery is restricted")
	verifyOnly := flag.Bool("verify-only", false, "run all correctness gates without measurements")
	flag.Parse()
	if *samples < 1 || *minTime < time.Millisecond {
		return fmt.Errorf("samples must be positive and min-time at least 1ms")
	}
	absolute, err := filepath.Abs(*root)
	if err != nil {
		return err
	}
	resolve := func(path string) string {
		if filepath.IsAbs(path) {
			return path
		}
		return filepath.Join(absolute, path)
	}
	corpusPath, refPath, outPath := resolve(*corpus), resolve(*reference), resolve(*output)
	if _, err := os.Stat(outPath); err == nil {
		return fmt.Errorf("output directory already exists: %s; choose a new -output", outPath)
	} else if !os.IsNotExist(err) {
		return err
	}
	fixture, expectedPattern := "testdata/qwen/cases.jsonl.gz", compile.QwenPattern
	switch *patternID {
	case "P01":
		fixture, expectedPattern = "testdata/patterns/P01/cases.jsonl.gz", compile.QwenMarksPattern
	case "P02":
		expectedPattern = compile.ASCIITriplesPattern
	case "P03":
		expectedPattern = compile.LineEndWhitespacePattern
	case "P04":
		expectedPattern = compile.CaseWordsPattern
	case "P09":
		expectedPattern = compile.QwenLiteralNewlinesPattern
	case "P25":
		expectedPattern = compile.RightAlignedDecimalsPattern
	case "P14":
		expectedPattern = compile.ConsecutiveDecimalTriplesPattern
	case "P10":
		expectedPattern = compile.InitialDecimalRemainderPattern
	case "P13":
		expectedPattern = compile.QwenPossessivePattern
	case "P16":
		expectedPattern = compile.ASCIIDigitPattern
	case "P07":
		expectedPattern = compile.UnicodeTriplesPattern
	case "P24":
		expectedPattern = compile.QwenNumberTriplesPattern
	case "P20":
		expectedPattern = compile.QwenNumberRunsPattern
	case "P19":
		expectedPattern = compile.CaseWordsSuffixPattern
	case "P06":
		expectedPattern = compile.CaseWordsSuffixTriplesPattern
	case "P11":
		expectedPattern = compile.NewlinePattern
	case "P21":
		expectedPattern = compile.CJKPattern
	case "P17":
		expectedPattern = compile.BroadRangesPattern
	case "P18":
		expectedPattern = compile.PrefixedSymbolsPattern
	case "P08":
		expectedPattern = compile.PrefixedAlphabetPattern
	case "P12":
		expectedPattern = compile.ScriptGroupsPattern
	case "P15":
		expectedPattern = compile.HangulGroupsPattern
	case "P22":
		expectedPattern = compile.PunctuationWordsPattern
	case "P23":
		expectedPattern = compile.ComplementPattern
	case "P05":
	default:
		return fmt.Errorf("unsupported benchmark pattern %q", *patternID)
	}
	runtime.GOMAXPROCS(1)
	fmt.Fprintln(os.Stderr, "Loading pinned corpus inputs...")
	manifest, batches, err := bench.Load(corpusPath, *patternID)
	if err != nil {
		return err
	}
	if manifest.Pattern != expectedPattern {
		return fmt.Errorf("manifest pattern differs from selected %s", *patternID)
	}
	fmt.Fprintln(os.Stderr, "Checking the pinned native reference against regression fixtures...")
	ref, ready, err := startNative(refPath, corpusPath, resolve(fixture), *patternID)
	if err != nil {
		return err
	}
	defer ref.close()
	re, err := regex.Compile(manifest.Pattern)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Generating HF references in memory and verifying Go/Oniguruma on every corpus request...")
	for i := range batches {
		b := &batches[i]
		expected, e := ref.reference(b)
		if e != nil {
			return e
		}
		if e = bench.Verify(re, manifest.PatternSHA, b, expected); e != nil {
			return e
		}
	}
	var selected []bench.Batch
	for _, b := range batches {
		if strings.Contains(b.ID, *filter) {
			selected = append(selected, b)
		}
	}
	if len(selected) == 0 {
		return fmt.Errorf("filter selected no workloads")
	}
	fmt.Fprintf(os.Stderr, "Correctness gate passed: %v original fixtures, %d batches.\n", ready["fixtures"], len(batches))
	if err = os.MkdirAll(outPath, 0755); err != nil {
		return err
	}
	executable, _ := os.Executable()
	gitDiff := command("git", "-C", absolute, "diff", "--binary", "HEAD")
	env := map[string]any{"started_utc": time.Now().UTC().Format(time.RFC3339), "complete": false, "go": runtime.Version(), "rust": command("rustc", "-Vv"), "os": runtime.GOOS, "arch": runtime.GOARCH, "cpu": command("sysctl", "-n", "machdep.cpu.brand_string"), "os_version": command("uname", "-srv"), "workers": 1, "go_allocator": "Go runtime, default GC", "native_allocator": "Rust System + Oniguruma C allocator; counters unavailable", "go_build": command("go", "version", "-m", executable), "native_build": "cargo build --release --locked; thin LTO; 1 codegen unit; default CPU target", "git_commit": command("git", "-C", absolute, "rev-parse", "HEAD"), "git_status": command("git", "-C", absolute, "status", "--short"), "git_diff_sha256": bench.Digest([]byte(gitDiff)), "source_sha256": sourceSHA(absolute), "native_binary_sha256": fileSHA(refPath), "go_binary_sha256": fileSHA(executable), "cargo_lock_sha256": fileSHA(resolve("benchmark/native/Cargo.lock")), "corpus_manifest_sha256": fileSHA(filepath.Join(corpusPath, "manifest.json")), "pattern_sha256": manifest.PatternSHA, "native_gate": ready, "samples": *samples, "min_time_ns": int64(*minTime), "filter": *filter, "selected_batches": len(selected), "pattern_id": *patternID, "backend": "scalar", "utf8": "timed public validation on both sides", "rustflags": os.Getenv("RUSTFLAGS"), "goflags": os.Getenv("GOFLAGS")}
	env["corpus_reference"] = "HF Split generated in memory before timing; exact Go/Oniguruma verification; no stored corpus spans"
	env["input_layout"] = "contiguous corpus storage with request views in both runners"
	if *cpuLabel != "" {
		env["cpu"] = *cpuLabel
		env["cpu_source"] = "explicit -cpu-label"
	}
	env["cc_version"] = command("cc", "--version")
	env["cflags"] = os.Getenv("CFLAGS")
	if info, ok := debug.ReadBuildInfo(); ok {
		env["build_info"] = info
	}
	if err = writeJSON(filepath.Join(outPath, "environment.json"), env); err != nil {
		return err
	}
	if *verifyOnly {
		env["complete"] = true
		env["verification_only"] = true
		return writeJSON(filepath.Join(outPath, "environment.json"), env)
	}
	raw, err := os.Create(filepath.Join(outPath, "samples.jsonl"))
	if err != nil {
		return err
	}
	defer raw.Close()
	var results []bench.Sample
	record := func(s bench.Sample) error {
		s.PatternID = *patternID
		results = append(results, s)
		if err := json.NewEncoder(raw).Encode(s); err != nil {
			return err
		}
		return raw.Sync()
	}
	engines := bench.Engines(manifest.FullCoverage)
	for index := range selected {
		b := &selected[index]
		operations := map[string]func() uint64{}
		allocation := map[string][2]float64{}
		for _, engine := range engines {
			if strings.HasPrefix(engine, "pretok-") {
				f, e := bench.Operation(re, b, engine)
				if e != nil {
					return e
				}
				operations[engine] = f
				bytes, allocs := bench.Allocations(f)
				allocation[engine] = [2]float64{bytes, allocs}
			}
		}
		for round := 0; round < *samples; round++ {
			fmt.Fprintf(os.Stderr, "[%d/%d] %s/%d round %d/%d\n", index+1, len(selected), b.ID, b.Size, round+1, *samples)
			// Each adjacent engine pair is run AB then BA on alternate rounds.
			order := append([]string(nil), engines...)
			if round%2 == 1 {
				slices.Reverse(order)
			}
			for _, engine := range order {
				var s bench.Sample
				if f, ok := operations[engine]; ok {
					s = bench.Measure(f, *minTime)
					a := allocation[engine]
					s.BPerBatch = &a[0]
					s.AllocsPerBatch = &a[1]
				} else {
					s, err = ref.sample(map[string]any{"op": "sample", "corpus": b.ID, "size": b.Size, "engine": engine, "min_ms": minTime.Milliseconds()})
					if err != nil {
						return err
					}
					if s.Bytes != b.Bytes || s.Requests != len(b.Requests) || s.Spans != b.OutputCount(engine) {
						return fmt.Errorf("native workload metadata mismatch")
					}
				}
				s.Engine = engine
				s.Corpus = b.ID
				s.Kind = b.Kind
				s.Size = b.Size
				s.Round = round
				s.Bytes = b.Bytes
				s.Requests = len(b.Requests)
				s.Scalars = b.Scalars
				s.Spans = b.OutputCount(engine)
				if err = record(s); err != nil {
					return err
				}
			}
		}
	}
	for round := 0; round < *samples; round++ {
		order := []string{"pretok-compile", "onig-compile", "hf-compile"}
		if round%2 == 1 {
			slices.Reverse(order)
		}
		for _, engine := range order {
			var s bench.Sample
			if engine == "pretok-compile" {
				s = bench.MeasureCompile(manifest.Pattern, *minTime)
			} else {
				s, err = ref.sample(map[string]any{"op": "compile", "engine": engine, "min_ms": minTime.Milliseconds()})
				if err != nil {
					return err
				}
			}
			s.Engine = engine
			s.Corpus = "compile"
			s.Kind = "compile"
			s.Round = round
			if err = record(s); err != nil {
				return err
			}
		}
	}
	report, err := os.Create(filepath.Join(outPath, "REPORT.md"))
	if err != nil {
		return err
	}
	err = bench.WriteReport(report, results, "environment.json", *patternID)
	closeErr := report.Close()
	if err != nil {
		return err
	}
	if closeErr != nil {
		return closeErr
	}
	env["complete"] = true
	env["finished_utc"] = time.Now().UTC().Format(time.RFC3339)
	env["sample_rows"] = len(results)
	if err = writeJSON(filepath.Join(outPath, "environment.json"), env); err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Report:", filepath.Join(outPath, "REPORT.md"))
	return nil
}
