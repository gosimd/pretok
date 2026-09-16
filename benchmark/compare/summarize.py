#!/usr/bin/env python3
"""Summarize the complete four-build P01 comparison."""
import argparse
import collections
import csv
import json
from pathlib import Path
import statistics

parser = argparse.ArgumentParser(description=__doc__)
parser.add_argument('result', type=Path)
parser.add_argument('--output', type=Path, required=True, help='generated CSV outside published results')
args = parser.parse_args()
root = args.result
completion = root / 'complete.json'
metadata = json.loads((completion if completion.exists() else root / 'environment.json').read_text())
assert metadata.get('completion', metadata)['complete']
rows = [json.loads(line) for line in (root / 'samples.jsonl').read_text().splitlines()]
assert len(rows) == 600
by_case = collections.defaultdict(list)
for row in rows:
    by_case[row['corpus'], row['size'], row['engine'], row['build']].append(row)
assert len(by_case) == 120 and all(len(group) == 5 for group in by_case.values())
summary = []
for (corpus, size, engine, build), group in sorted(by_case.items()):
    ns = [r['elapsed_ns'] / r['iterations'] / r['requests'] for r in group]
    mb = [r['bytes'] * r['iterations'] / r['elapsed_ns'] * 1000 for r in group]
    baseline = by_case[corpus, size, engine, build.replace('new', 'old')]
    base_ns = statistics.median(r['elapsed_ns'] / r['iterations'] / r['requests'] for r in baseline)
    allocation = next(r for r in group if r['bytes_allocated_per_batch'] is not None)
    summary.append(dict(
        corpus=corpus, size=size, engine=engine, build=build,
        ns_request=statistics.median(ns), mb_s=statistics.median(mb),
        speedup=base_ns / statistics.median(ns),
        bytes_request=allocation['bytes_allocated_per_batch'] / allocation['requests'],
        allocs_request=allocation['allocations_per_batch'] / allocation['requests'],
        min_ns=min(ns), max_ns=max(ns),
    ))
args.output.parent.mkdir(parents=True, exist_ok=True)
with args.output.open('w', newline='') as output:
    writer = csv.DictWriter(output, fieldnames=list(summary[0]))
    writer.writeheader()
    writer.writerows(summary)
