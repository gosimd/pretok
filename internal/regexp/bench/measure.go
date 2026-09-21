package bench

import (
	"fmt"
	"math/bits"
	"runtime"
	"time"

	regex "github.com/gosimd/pretok/internal/regexp"
)

// Sample always measures one complete corpus batch per iteration. Compile rows
// use one construction per iteration instead. Allocation values are separate runs.
type Sample struct {
	PatternID      string   `json:"pattern_id,omitempty"`
	Engine         string   `json:"engine"`
	Corpus         string   `json:"corpus"`
	Kind           string   `json:"kind"`
	Size           int      `json:"size"`
	Round          int      `json:"round"`
	Iterations     uint64   `json:"iterations"`
	ElapsedNS      uint64   `json:"elapsed_ns"`
	Checksum       uint64   `json:"checksum"`
	Bytes          int      `json:"bytes"`
	Requests       int      `json:"requests"`
	Scalars        int      `json:"scalars"`
	Spans          int      `json:"spans"`
	BPerBatch      *float64 `json:"bytes_allocated_per_batch"`
	AllocsPerBatch *float64 `json:"allocations_per_batch"`
}

var spanSink []regex.Span
var regexpSink *regex.Regexp
var checksumSink uint64

func consume(acc uint64, a, b int) uint64 { return bits.RotateLeft64(acc+uint64(a), 7) ^ uint64(b) }
func Operation(re *regex.Regexp, b *Batch, engine string) (func() uint64, error) {
	if engine != "pretok-cursor" && engine != "pretok-reused" && engine != "pretok-fresh" {
		return nil, fmt.Errorf("unknown Go engine %q", engine)
	}
	if !b.Verified {
		return nil, fmt.Errorf("batch must pass live reference verification before measurement")
	}
	dst := make([]regex.Span, 0, b.Capacity)
	return func() uint64 {
		var sum uint64
		for _, r := range b.Requests {
			if engine == "pretok-cursor" {
				c, err := re.Cursor(r.Text)
				if err != nil {
					panic(err)
				}
				for s, ok := c.Next(); ok; s, ok = c.Next() {
					sum = consume(sum, s.Start, s.End)
				}
			} else {
				var err error
				if engine == "pretok-reused" {
					dst, err = re.Split(r.Text, dst[:0])
				} else {
					dst, err = re.Split(r.Text, nil)
				}
				if err != nil {
					panic(err)
				}
				sum = consume(sum, len(dst), len(r.Text))
				spanSink = dst
			}
		}
		checksumSink = sum
		return sum
	}, nil
}
func Measure(run func() uint64, minTime time.Duration) Sample {
	run() // warm up outside the timer
	start := time.Now()
	var n, sum uint64
	for {
		sum ^= run()
		n++
		if time.Since(start) >= minTime {
			break
		}
	}
	return Sample{Iterations: n, ElapsedNS: uint64(time.Since(start)), Checksum: sum}
}
func Allocations(run func() uint64) (float64, float64) {
	run()
	runtime.GC()
	var before, after runtime.MemStats
	runtime.ReadMemStats(&before)
	const n = 10
	for range n {
		run()
	}
	runtime.ReadMemStats(&after)
	return float64(after.TotalAlloc-before.TotalAlloc) / n, float64(after.Mallocs-before.Mallocs) / n
}
func MeasureCompile(pattern string, minTime time.Duration) Sample {
	start := time.Now()
	var n uint64
	for {
		for range 256 {
			re, err := regex.Compile(pattern)
			if err != nil {
				panic(err)
			}
			regexpSink = re
			n++
		}
		if time.Since(start) >= minTime {
			break
		}
	}
	return Sample{Iterations: n, ElapsedNS: uint64(time.Since(start)), Bytes: len(pattern), Requests: 1}
}
