package bench

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

func quantile(values []float64, p float64) float64 {
	copy := append([]float64(nil), values...)
	sort.Float64s(copy)
	if len(copy) == 0 {
		return 0
	}
	at := p * float64(len(copy)-1)
	i := int(at)
	if i == len(copy)-1 {
		return copy[i]
	}
	return copy[i] + (copy[i+1]-copy[i])*(at-float64(i))
}

type Summary struct {
	Sample      Sample
	Median, IQR float64
	Count       int
}

func Summaries(samples []Sample) []Summary {
	groups := map[string][]Sample{}
	for _, s := range samples {
		key := fmt.Sprintf("%s/%09d/%s", s.Corpus, s.Size, s.Engine)
		groups[key] = append(groups[key], s)
	}
	keys := make([]string, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := make([]Summary, 0, len(keys))
	for _, k := range keys {
		ss := groups[k]
		var times []float64
		for _, s := range ss {
			times = append(times, float64(s.ElapsedNS)/float64(s.Iterations))
		}
		out = append(out, Summary{ss[0], quantile(times, .5), quantile(times, .75) - quantile(times, .25), len(ss)})
	}
	return out
}
func WriteReport(destination io.Writer, samples []Sample, environmentFile string) error {
	var buffer bytes.Buffer
	var w io.Writer = &buffer
	fmt.Fprint(w, "# P05 benchmark\n\nNative Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.\n\n")
	fmt.Fprintf(w, "Environment and source fingerprints: [%s](%s). Raw samples: [samples.jsonl](samples.jsonl).\n\n", environmentFile, environmentFile)
	summaries := Summaries(samples)
	short := false
	for _, s := range samples {
		if s.ElapsedNS < 500_000_000 {
			short = true
		}
	}
	for _, s := range summaries {
		if s.Count < 10 {
			short = true
		}
	}
	if short {
		fmt.Fprintln(w, "**Smoke/short run:** fewer than 10 samples or samples shorter than 500 ms. These numbers are not stable performance evidence.")
		fmt.Fprintln(w)
	}
	fmt.Fprint(w, "## Per-workload measurements\n\nOne batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.\n\n")
	fmt.Fprintln(w, "| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |\n|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|")
	index := map[string]Summary{}
	for _, x := range summaries {
		s := x.Sample
		index[fmt.Sprintf("%s/%d/%s", s.Corpus, s.Size, s.Engine)] = x
		if s.Kind == "compile" {
			continue
		}
		bytes, alloc := "N/A", "N/A"
		if s.BPerBatch != nil {
			bytes = fmt.Sprintf("%.1f", *s.BPerBatch)
			alloc = fmt.Sprintf("%.1f", *s.AllocsPerBatch)
		}
		fmt.Fprintf(w, "| %s | %d | %s | %.0f | %.0f | %.1f | %.2f | %.0f | %.0f | %s | %s | %d |\n", s.Corpus, s.Size, s.Engine, x.Median, x.IQR, x.Median/float64(s.Requests), float64(s.Bytes)*1e3/x.Median, float64(s.Scalars)*1e9/x.Median, float64(s.Spans)*1e9/x.Median, bytes, alloc, x.Count)
	}
	pairs := [][3]string{{"Regex iteration", "onig-iter", "pretok-cursor"}, {"Regex span collection, reused", "onig-reused", "pretok-reused"}, {"Regex span collection, fresh", "onig-fresh", "pretok-fresh"}, {"HF stage versus reused destination", "hf-split", "pretok-reused"}, {"HF stage versus fresh destination", "hf-split", "pretok-fresh"}}
	for _, pair := range pairs {
		fmt.Fprintf(w, "\n## %s\n\nSpeedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.\n\n| Corpus | Request bytes | Speedup |\n|---|---:|---:|\n", pair[0])
		totals := map[int][2]float64{}
		for _, x := range summaries {
			s := x.Sample
			if s.Kind == "compile" || s.Engine != pair[1] {
				continue
			}
			other, ok := index[fmt.Sprintf("%s/%d/%s", s.Corpus, s.Size, pair[2])]
			if !ok {
				continue
			}
			fmt.Fprintf(w, "| %s | %d | %.3fx |\n", s.Corpus, s.Size, x.Median/other.Median)
			if s.Kind == "natural" {
				t := totals[s.Size]
				t[0] += x.Median
				t[1] += other.Median
				totals[s.Size] = t
			}
		}
		sizes := make([]int, 0, len(totals))
		for size := range totals {
			sizes = append(sizes, size)
		}
		sort.Ints(sizes)
		for _, size := range sizes {
			t := totals[size]
			fmt.Fprintf(w, "| Combined selected natural corpora | %d | %.3fx |\n", size, t[0]/t[1])
		}
	}
	fmt.Fprintln(w, "\n## Compilation\n\nOne-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.\n\n| Engine | ns/compile | IQR ns | Samples |\n|---|---:|---:|---:|")
	for _, x := range summaries {
		if x.Sample.Kind == "compile" {
			fmt.Fprintf(w, "| %s | %.1f | %.1f | %d |\n", x.Sample.Engine, x.Median, x.IQR, x.Count)
		}
	}
	fmt.Fprint(w, "\nNatural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.\n")
	_, err := io.Copy(destination, &buffer)
	return err
}
func Engines() []string {
	return strings.Fields("pretok-cursor onig-iter pretok-reused onig-reused pretok-fresh onig-fresh hf-split onig-at")
}
