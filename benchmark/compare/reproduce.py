#!/usr/bin/env python3
"""Rebuild baseline/current sources and run the completed P01 comparison."""
import argparse
import io
import json
import os
from pathlib import Path
import shutil
import subprocess
import tarfile

HERE = Path(__file__).resolve().parent


def snapshot(repo, revision, destination):
    archive = subprocess.check_output(['git', 'archive', revision], cwd=repo)
    destination.mkdir()
    with tarfile.open(fileobj=io.BytesIO(archive)) as tar:
        for item in tar:
            relative = Path(item.name)
            if relative.is_absolute() or '..' in relative.parts:
                raise ValueError('Unsafe archive entry')
            target = destination / relative
            if item.isdir():
                target.mkdir(parents=True, exist_ok=True)
            elif item.isfile():
                target.parent.mkdir(parents=True, exist_ok=True)
                target.write_bytes(tar.extractfile(item).read())
                target.chmod(item.mode)
            else:
                raise ValueError('Unsupported archive entry')


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--repo', type=Path, default=HERE.parents[1])
    parser.add_argument('--result', type=Path, required=True, help='published comparison directory')
    parser.add_argument('--work', type=Path, required=True, help='new temporary directory')
    parser.add_argument('--native', type=Path, help='pinned pretok-reference executable')
    parser.add_argument('--prepare-only', action='store_true')
    args = parser.parse_args()
    work = args.work.resolve()
    work.mkdir(parents=True, exist_ok=False)
    metadata = json.loads((args.result / 'environment.json').read_text())
    snapshot(args.repo, metadata['baseline_commit'], work / 'baseline')
    snapshot(args.repo, metadata['implementation_commit'], work / 'current')
    for name in ('baseline', 'current'):
        source = work / name / 'cmd/p01integration/main.go'
        source.parent.mkdir(parents=True, exist_ok=True)
        shutil.copyfile(HERE / 'main.go', source)
    if args.prepare_only:
        print(work)
        return
    env = os.environ.copy()
    env['GOWORK'] = 'off'
    env['GOTOOLCHAIN'] = 'local'
    env.setdefault('GOPROXY', 'https://proxy.golang.org')
    subprocess.run(['go', 'mod', 'download'], cwd=work / 'current', env=env, check=True)
    for variant in ('old-go', 'new-go', 'old-simd', 'new-simd'):
        source = work / ('baseline' if variant.startswith('old') else 'current')
        env['GOEXPERIMENT'] = 'simd' if variant.endswith('simd') else ''
        subprocess.run(['go', 'build', '-mod=readonly', '-o', str(work / variant),
                        './cmd/p01integration'], cwd=source, env=env, check=True)
    native = args.native.resolve() if args.native else work / 'current/benchmark/native/target/release/pretok-reference'
    if not args.native:
        subprocess.run(['cargo', 'build', '--locked', '--release'],
                       cwd=work / 'current/benchmark/native', check=True)
    subprocess.run(['python3', str(HERE / 'measure.py'), str(work), str(native)], check=True)
    subprocess.run(['python3', str(HERE / 'summarize.py'), str(work / 'results'),
                    '--output', str(work / 'results/summary.csv')], check=True)
    print('Completed:', work / 'results')


if __name__ == '__main__':
    main()
