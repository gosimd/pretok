#!/usr/bin/env python3
"""Verify P15's pinned Hangul table by exhaustive HF extraction; stores no outputs.

Run explicitly with tokenizers==0.22.2. Ordinary Go tests remain offline.
"""
import re
from pathlib import Path
import tokenizers
from tokenizers import Regex, pre_tokenizers


def main():
    assert tokenizers.__version__ == '0.22.2', tokenizers.__version__
    text = ''.join(chr(c) for c in range(0x110000) if not 0xD800 <= c <= 0xDFFF)
    split = pre_tokenizers.Split(Regex(r'\p{Hangul}+'), behavior='removed', invert=True)
    actual = [(ord(piece[0]), ord(piece[-1])) for piece, _ in split.pre_tokenize_str(text)]
    path = Path(__file__).resolve().parents[2] / 'internal/scan/grouped_ranges.go'
    section = path.read_text().split('var hangulRanges')[1].split('\n}')[0]
    expected = [(int(lo, 16), int(hi, 16)) for lo, hi in
                re.findall(r'\{0x([0-9A-Fa-f]+), 0x([0-9A-Fa-f]+)\}', section)]
    assert actual == expected, (actual, expected)
    print(f'HF {tokenizers.__version__}: Hangul matches all {len(actual)} pinned ranges over every valid scalar.')


if __name__ == '__main__':
    main()
