# P05 shared-template regression check

Before/after Go microbenchmarks on Apple M5, Go 1.27.0, one worker. The exact same current benchmark harness was compiled against the baseline library (commit `5e798dbb3e007ec6b42a833c54b0d296fad3c4d6`) and the P01/P05 shared-template library. Ten paired rounds, 500 ms per benchmark, alternating AB/BA order. These are small API workloads, not the full multilingual cross-engine suite.

[Environment and binary hashes](environment.json) · [Raw samples](samples.jsonl)

Positive changes mean the new library is slower. Values are medians; no statistical significance claim is made.

| Benchmark | Before ns/op | After ns/op | Time change |
|---|---:|---:|---:|
| BenchmarkQwenSplit/multilingual | 19416.00 | 20006.00 | +3.04% |
| BenchmarkQwenSplit/short | 134.00 | 134.95 | +0.71% |
| BenchmarkQwenSplitFresh/multilingual | 21424.50 | 21691.00 | +1.24% |
| BenchmarkQwenSplitFresh/short | 208.85 | 210.85 | +0.96% |

The shared implementation removes duplicate scanner logic and supports P01 with one preselected word-class table. Retain the measured tradeoff above when evaluating the refactor; the earlier P05 cross-engine report describes the earlier implementation. Both variants preserve the same allocation counts on these workloads.
