package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gosimd/pretok"
	"github.com/gosimd/pretok/internal/bench"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"
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

func must(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	runtime.GOMAXPROCS(1)
	m, batches, err := bench.Load("benchmark/corpus", "P01")
	must(err)
	re, err := pretok.Compile(m.Pattern)
	must(err)
	ref, ready, err := startNative(os.Args[1], "benchmark/corpus", "testdata/patterns/P01/cases.jsonl.gz", "P01")
	must(err)
	total := 0
	for i := range batches {
		oracle, err := ref.reference(&batches[i])
		must(err)
		must(bench.Verify(re, m.PatternSHA, &batches[i], oracle))
		total += len(batches[i].Requests)
	}
	ref.close()
	out := json.NewEncoder(os.Stdout)
	must(out.Encode(map[string]any{"ready": true, "oracle": ready, "requests": total, "go": runtime.Version(), "pattern_sha256": m.PatternSHA}))
	type op struct {
		run           func() uint64
		b             *bench.Batch
		bytes, allocs float64
	}
	ops := map[string]op{}
	in := json.NewDecoder(os.Stdin)
	for {
		var cmd struct {
			Corpus string `json:"corpus"`
			Size   int    `json:"size"`
			Engine string `json:"engine"`
			Round  int    `json:"round"`
		}
		if err := in.Decode(&cmd); err == io.EOF {
			return
		} else {
			must(err)
		}
		key := fmt.Sprintf("%s/%d/%s", cmd.Corpus, cmd.Size, cmd.Engine)
		o, ok := ops[key]
		if !ok {
			for i := range batches {
				b := &batches[i]
				if b.ID == cmd.Corpus && b.Size == cmd.Size {
					o.b = b
					break
				}
			}
			if o.b == nil {
				panic("unknown batch")
			}
			o.run, err = bench.Operation(re, o.b, cmd.Engine)
			must(err)
			o.bytes, o.allocs = bench.Allocations(o.run)
			ops[key] = o
		}
		s := bench.Measure(o.run, 150*time.Millisecond)
		s.PatternID = "P01"
		s.Engine = cmd.Engine
		s.Corpus = cmd.Corpus
		s.Size = cmd.Size
		s.Round = cmd.Round
		s.Kind = o.b.Kind
		s.Bytes = o.b.Bytes
		s.Requests = len(o.b.Requests)
		s.Scalars = o.b.Scalars
		s.Spans = o.b.Spans
		if cmd.Round == 0 {
			s.BPerBatch = &o.bytes
			s.AllocsPerBatch = &o.allocs
		}
		must(out.Encode(s))
	}
}
