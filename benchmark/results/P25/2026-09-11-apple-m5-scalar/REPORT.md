# P25 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Interpretation of this P25 run

Apple M5 (explicit host label), scalar Go 1.27.0, Rust 1.90.0, HF Tokenizers
0.22.2, onig crate 6.5.1, one worker. Ten repetitions of at least 500 ms per
engine/workload: 1,570 samples, including separate compilation timing.

P25 groups Nd runs from the right in blocks of 510 scalars and resumes search
after gaps. Cursor-owned value state caches each run's byte end. Long runs are
scanned at most twice overall; the compiled plan remains immutable. Short runs
return the already known endpoint directly. Public signatures are unchanged.

The shared decimal workload has 110 subjects in five digit layouts, with selected
lengths from 4 to 4,097 scalars. It produces 220 matches and
220 isolated pieces. Reused Split reaches **285.7 MB/s**,
averaging **5.29 us/request**, with ratios of **236.78×** against Oniguruma
reused spans and **237.57×** against native HF Split. Sizes vary, so the mean per
request is not fixed-size latency. P10/P14/P25 have different output boundaries
on the same decimal input pool.

Unlike P10/P14, P25 searches through natural text for subsequent digit runs.
Natural rows below use 4 KiB requests. Baseline/pretok ratios above one favor
pretok; ratios below one favor the baseline.

| Workload | pretok MB/s | Oniguruma / pretok | HF Split / pretok |
| --- | ---: | ---: | ---: |
| English, 4 KiB | 467.6 | 5.03× | 6.54× |
| Russian, 4 KiB | 266.5 | 2.20× | 3.00× |
| Japanese, 4 KiB | 432.9 | 2.58× | 3.83× |
| Arabic, 4 KiB | 253.2 | 2.11× | 2.84× |
| Chinese, 4 KiB | 409.4 | 2.57× | 3.76× |
| Decimal subjects | 285.7 | 236.78× | 237.57× |

Ratios of summed median batch times across the five 4 KiB languages are
2.71× Oniguruma/pretok and 3.74× HF/pretok.
Cursor and reused Split allocate zero bytes in every measured workload. The full
report includes fresh allocations, all sizes, ns/request and IQR. Native allocation
counters are unavailable. Compile medians: 10.79 ns for exact-string
pretok lookup, 4.73 us for Oniguruma and 4.75 us for HF.

All 163,066 shared requests passed live HF/Oniguruma comparison before timing;
all 25 supported patterns passed the expanded gate. Expected spans are generated
in memory, not stored. These measurements cover regex boundaries and native HF
Split, not a complete tokenizer or BPE. Compatibility remains exact-string-only.

## Existing-pattern check for cursor state

The new internal cache adds one integer (8 bytes on this host) to each cursor.
An uncontended paired A/B used three samples per version, 300 ms per sample,
`-test.cpu=1`, existing Go reused-Split benchmarks, and order before/after,
after/before, before/after. Compilation and setup are outside timing. Medians:

| Existing workload | Before ns/op | After ns/op | Change |
| --- | ---: | ---: | ---: |
| BenchmarkP14Split/short | 10.22 | 11.78 | +15.26% |
| BenchmarkP14Split/multilingual | 1327.00 | 1339.00 | +0.90% |
| BenchmarkQwenSplit/short | 146.00 | 145.00 | -0.68% |
| BenchmarkQwenSplit/multilingual | 20598.00 | 20502.00 | -0.47% |
| BenchmarkP01Split/multilingual | 21023.00 | 20824.00 | -0.95% |
| BenchmarkP01Split/short | 146.90 | 148.20 | +0.88% |

The short P14 rejection costs 1.56 ns more (15.3%). The other selected medians
changed by less than 1%; these small differences are not claimed as improvements.
All paired reused-Split samples reported zero allocations. This is a scoped
regression check, not a fresh full benchmark of every previously supported pattern.

Baseline commit: `b970be86e6a093878df34ddc9f3229dd8cf7429d`.
Before test binary SHA-256: `78c5d3efcd4501b601b0addbb9699aac03b521266775f0c691fe3d4346c004d4`.
After test binary SHA-256: `41fa1493061e750af1367c94f36c961351ef13cc62302743316e436e26fc324e`.

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 13941387 | 130037 | 3397.0 | 75.21 | 41930906 | 730845 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 9544047 | 63980 | 2325.5 | 109.87 | 61250221 | 1067576 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 9344264 | 78316 | 2276.9 | 112.22 | 62559767 | 329186 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 9363409 | 65721 | 2281.5 | 111.99 | 62431856 | 1088172 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 4224595 | 46629 | 1029.4 | 248.21 | 138374226 | 728117 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 4341686 | 29599 | 1057.9 | 241.51 | 134642381 | 2346784 | 365792.0 | 7732.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 4227319 | 19012 | 1030.0 | 248.05 | 138285052 | 2410275 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 11769528 | 47227 | 45795.8 | 89.09 | 49668518 | 541398 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 8758457 | 68322 | 34079.6 | 119.72 | 66744063 | 727525 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 8741988 | 74210 | 34015.5 | 119.95 | 66869803 | 349920 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 8738522 | 88497 | 34002.0 | 119.99 | 66896322 | 729185 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 4122820 | 24902 | 16042.1 | 254.33 | 141790073 | 741968 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 4181464 | 17881 | 16270.3 | 250.77 | 139801524 | 1523868 | 289552.0 | 1381.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 4140701 | 36713 | 16111.7 | 253.24 | 141177788 | 1538870 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 11764320 | 85427 | 692018.8 | 89.13 | 49690504 | 521322 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 8663916 | 67229 | 509642.1 | 121.03 | 67472374 | 707878 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 8726325 | 124080 | 513313.3 | 120.16 | 66989823 | 350434 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 8639852 | 58140 | 508226.6 | 121.37 | 67660305 | 709850 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 4141291 | 14415 | 243605.4 | 253.20 | 141157674 | 738417 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 4205797 | 34654 | 247399.8 | 249.32 | 138992684 | 1458225 | 309264.0 | 163.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 4135784 | 26758 | 243281.4 | 253.54 | 141345634 | 1482911 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 13484645 | 154973 | 13484645.3 | 77.76 | 43351159 | 453627 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 8662054 | 62639 | 8662053.5 | 121.05 | 67486884 | 706184 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 8746872 | 86800 | 8746872.5 | 119.88 | 66832459 | 349611 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 8660475 | 30304 | 8660474.9 | 121.08 | 67499186 | 706312 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 4117090 | 57465 | 4117089.8 | 254.69 | 141987429 | 742758 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 4166368 | 44335 | 4166368.5 | 251.68 | 140308041 | 1468185 | 365808.0 | 17.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 4134389 | 32691 | 4134388.8 | 253.62 | 141393331 | 1479542 | 0.0 | 0.0 | 10 |
| decimal | 0 | hf-split | 138135026 | 2802266 | 1255773.0 | 1.20 | 481160 | 1593 | N/A | N/A | 10 |
| decimal | 0 | onig-fresh | 137798854 | 2246755 | 1252716.9 | 1.21 | 482333 | 1597 | N/A | N/A | 10 |
| decimal | 0 | onig-iter | 138037078 | 2146138 | 1254882.5 | 1.20 | 481501 | 1594 | N/A | N/A | 10 |
| decimal | 0 | onig-reused | 137677964 | 1418849 | 1251617.9 | 1.21 | 482757 | 1598 | N/A | N/A | 10 |
| decimal | 0 | pretok-cursor | 579309 | 1985 | 5266.4 | 286.80 | 114731549 | 379763 | 0.0 | 0.0 | 10 |
| decimal | 0 | pretok-fresh | 585662 | 7234 | 5324.2 | 283.69 | 113486925 | 375643 | 7840.0 | 180.0 | 10 |
| decimal | 0 | pretok-reused | 581456 | 6233 | 5286.0 | 285.74 | 114307865 | 378360 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 16247588 | 24031 | 3965.7 | 64.54 | 64405683 | 985377 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 12080388 | 20504 | 2948.6 | 86.80 | 86622798 | 1325289 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 11796422 | 39851 | 2879.3 | 88.89 | 88708003 | 509561 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 11808721 | 33050 | 2882.3 | 88.80 | 88615609 | 1355778 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 2307742 | 8924 | 563.3 | 454.37 | 453446188 | 2604710 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 2516603 | 6673 | 614.3 | 416.66 | 415813246 | 6361750 | 639024.0 | 10442.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 2313994 | 7491 | 564.8 | 453.15 | 452221091 | 6918773 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 14663324 | 50341 | 57278.6 | 71.51 | 71364243 | 832485 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 11366356 | 29115 | 44399.8 | 92.25 | 92064426 | 1073959 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 11277298 | 21864 | 44051.9 | 92.98 | 92791463 | 530180 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 11267730 | 32047 | 44014.6 | 93.06 | 92870262 | 1083359 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 2246964 | 10227 | 8777.2 | 466.66 | 465711446 | 2660923 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 2345877 | 10862 | 9163.6 | 446.99 | 446074993 | 5203598 | 572736.0 | 1674.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 2242266 | 9588 | 8758.9 | 467.64 | 466687313 | 5444047 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 14761828 | 183564 | 922614.3 | 71.03 | 70888034 | 811011 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 11294821 | 55273 | 705926.3 | 92.84 | 92647504 | 1059955 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 11274193 | 128133 | 704637.1 | 93.01 | 92817020 | 530238 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 11248201 | 80797 | 703012.6 | 93.22 | 93031496 | 1064348 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 2250507 | 19924 | 140656.7 | 465.93 | 464978296 | 2656290 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 2325489 | 28345 | 145343.1 | 450.91 | 449985732 | 5148164 | 540672.0 | 178.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 2249380 | 20181 | 140586.2 | 466.16 | 465211313 | 5322356 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 16628728 | 171958 | 16628728.5 | 63.06 | 62929466 | 719057 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 11301584 | 98496 | 11301583.8 | 92.78 | 92592067 | 1057993 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 11320749 | 65258 | 11320748.6 | 92.62 | 92435318 | 528057 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 11289858 | 85072 | 11289858.3 | 92.88 | 92688231 | 1059092 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 2270015 | 14041 | 2270015.0 | 461.92 | 460982418 | 2633463 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 2348065 | 15282 | 2348065.5 | 446.57 | 445659207 | 5092277 | 914672.0 | 20.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 2271169 | 12844 | 2271169.3 | 461.69 | 460748125 | 5264689 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 11073361 | 119913 | 2692.3 | 94.69 | 34251570 | 1260773 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 6999014 | 35182 | 1701.7 | 149.82 | 54190487 | 1994709 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 6854012 | 25220 | 1666.4 | 152.99 | 55336930 | 728770 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 6852296 | 42606 | 1666.0 | 153.03 | 55350793 | 2037419 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 2458469 | 18326 | 597.7 | 426.52 | 154274879 | 2031752 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 2661672 | 25583 | 647.1 | 393.95 | 142496886 | 5245199 | 529200.0 | 9352.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 2481394 | 19434 | 603.3 | 422.57 | 152849545 | 5626272 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 9284767 | 103986 | 36127.5 | 112.93 | 40849704 | 1097604 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 6323849 | 20830 | 24606.4 | 165.81 | 59976132 | 1611519 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 6227541 | 30450 | 24231.7 | 168.38 | 60903650 | 797907 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 6259235 | 53395 | 24355.0 | 167.52 | 60595267 | 1628154 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 2406915 | 35144 | 9365.4 | 435.65 | 157579331 | 2064469 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 2494409 | 43296 | 9705.9 | 420.37 | 152052068 | 4085537 | 466064.0 | 1522.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 2422418 | 20453 | 9425.8 | 432.86 | 156570826 | 4206953 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 9537724 | 137242 | 561042.6 | 109.94 | 39766299 | 1043750 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 6280183 | 48370 | 369422.5 | 166.97 | 60393146 | 1585145 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 6209344 | 46540 | 365255.5 | 168.87 | 61082138 | 800246 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 6217990 | 47419 | 365764.1 | 168.64 | 60997198 | 1601000 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 2412066 | 24609 | 141886.2 | 434.72 | 157242788 | 2060060 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 2449136 | 14291 | 144066.8 | 428.14 | 154862775 | 4064699 | 442896.0 | 170.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 2401775 | 13117 | 141280.9 | 436.58 | 157916564 | 4144852 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 11341726 | 124857 | 11341725.9 | 92.45 | 33441118 | 876322 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 6245830 | 41449 | 6245830.0 | 167.88 | 60725316 | 1591302 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 6227176 | 14923 | 6227175.7 | 168.39 | 60907227 | 797954 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 6220956 | 25766 | 6220955.8 | 168.56 | 60968124 | 1597664 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 2406918 | 21785 | 2406918.3 | 435.65 | 157579094 | 2064466 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 2472125 | 34833 | 2472125.0 | 424.16 | 153422660 | 4020428 | 685296.0 | 19.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 2425706 | 37198 | 2425706.1 | 432.28 | 156358595 | 4097364 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 13952865 | 176433 | 3399.8 | 75.15 | 41199925 | 785788 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 9485179 | 43499 | 2311.2 | 110.55 | 60605813 | 1155909 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 9286238 | 90802 | 2262.7 | 112.92 | 61904185 | 374425 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 9302481 | 67091 | 2266.7 | 112.72 | 61796097 | 1178610 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 4042819 | 26259 | 985.1 | 259.37 | 142192129 | 860044 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 4196169 | 35600 | 1022.5 | 249.89 | 136995671 | 2612859 | 405312.0 | 8406.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 4048810 | 31835 | 986.6 | 258.98 | 141981719 | 2707956 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 11822163 | 72605 | 46000.6 | 88.70 | 48625367 | 604458 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 8680328 | 90084 | 33775.6 | 120.80 | 66225260 | 823241 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 8621526 | 61803 | 33546.8 | 121.62 | 66676942 | 400045 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 8648842 | 60252 | 33653.1 | 121.24 | 66466358 | 826238 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 3935303 | 38717 | 15312.5 | 266.45 | 146076933 | 876426 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 4009716 | 33280 | 15602.0 | 261.51 | 143366002 | 1782171 | 328912.0 | 1443.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 3934491 | 28252 | 15309.3 | 266.51 | 146107074 | 1816245 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 11710106 | 108547 | 688829.7 | 89.54 | 49090676 | 590174 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 8632213 | 47168 | 507777.3 | 121.47 | 66594392 | 800606 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 8575112 | 29707 | 504418.3 | 122.28 | 67037845 | 401977 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 8590018 | 67807 | 505295.2 | 122.07 | 66921516 | 804539 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 3918381 | 32919 | 230493.0 | 267.60 | 146707797 | 879700 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 3975820 | 41160 | 233871.8 | 263.74 | 144588288 | 1738258 | 303888.0 | 163.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 3934254 | 38345 | 231426.7 | 266.52 | 146115886 | 1756623 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 13417798 | 195470 | 13417797.7 | 78.15 | 42842873 | 513870 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 8586814 | 89136 | 8586813.6 | 122.11 | 66946487 | 802975 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 8612071 | 51911 | 8612071.3 | 121.76 | 66750144 | 400252 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 8563461 | 33746 | 8563461.2 | 122.45 | 67129049 | 805165 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 3910114 | 49596 | 3910114.0 | 268.17 | 147017965 | 881560 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 3977465 | 21392 | 3977464.6 | 263.63 | 144528501 | 1733516 | 505072.0 | 18.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 3934381 | 40848 | 3934381.1 | 266.52 | 146111163 | 1752499 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 2171544000 | 23710178 | 15395.3 | 2.66 | 1231645 | 135667 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 2084474271 | 28823959 | 14778.1 | 2.77 | 1283092 | 141334 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 2094798834 | 45933698 | 14851.3 | 2.76 | 1276768 | 48830 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 2086368728 | 25135646 | 14791.5 | 2.77 | 1281927 | 141206 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 20932432 | 104204 | 148.4 | 275.83 | 127771678 | 4886628 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 24623897 | 191813 | 174.6 | 234.48 | 108616927 | 11964272 | 9749040.0 | 229230.0 | 10 |
| synthetic | 0 | pretok-reused | 21268534 | 283202 | 150.8 | 271.47 | 125752533 | 13851778 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 11302358 | 112609 | 2748.0 | 92.77 | 34524034 | 1221780 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 7353260 | 73788 | 1787.8 | 142.60 | 53065306 | 1877943 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 7176661 | 43001 | 1744.9 | 146.11 | 54371106 | 685556 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 7165630 | 46638 | 1742.2 | 146.33 | 54454807 | 1927116 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 2629608 | 13379 | 639.3 | 398.76 | 148388281 | 1871001 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 2803221 | 15755 | 681.6 | 374.06 | 139198074 | 4926118 | 530672.0 | 8872.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 2643317 | 15603 | 642.7 | 396.69 | 147618717 | 5224119 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 9623495 | 58247 | 37445.5 | 108.96 | 40546911 | 1042657 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 6656319 | 36040 | 25900.1 | 157.53 | 58621439 | 1507440 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 6592049 | 31540 | 25650.0 | 159.07 | 59192978 | 741651 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 6578748 | 55745 | 25598.2 | 159.39 | 59312657 | 1525214 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 2564476 | 15410 | 9978.5 | 408.88 | 152157011 | 1906432 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 2640447 | 31605 | 10274.1 | 397.12 | 147779155 | 3800114 | 454224.0 | 1444.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 2561491 | 7127 | 9966.9 | 409.36 | 152334301 | 3917249 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 9793201 | 157410 | 576070.7 | 107.07 | 39844275 | 1000082 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 6594560 | 67442 | 387915.3 | 159.01 | 59170437 | 1485164 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 6552762 | 43685 | 385456.6 | 160.02 | 59547868 | 746098 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 6574339 | 62463 | 386725.8 | 159.49 | 59352433 | 1489732 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 2565986 | 15445 | 150940.4 | 408.64 | 152067451 | 1905310 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 2609291 | 18360 | 153487.7 | 401.86 | 149543706 | 3753510 | 483856.0 | 172.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 2559986 | 13097 | 150587.4 | 409.60 | 152423901 | 3825803 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 11536747 | 96797 | 11536746.7 | 90.89 | 33822620 | 847639 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 6567980 | 35239 | 6567980.0 | 159.65 | 59409895 | 1488890 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 6561901 | 56872 | 6561901.2 | 159.80 | 59464930 | 745058 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 6573546 | 14972 | 6573545.7 | 159.51 | 59359593 | 1487629 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 2555424 | 26904 | 2555424.1 | 410.33 | 152695985 | 1913185 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 2633516 | 19633 | 2633516.1 | 398.17 | 148168070 | 3713287 | 685296.0 | 19.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 2577363 | 9641 | 2577362.8 | 406.84 | 151396223 | 3794188 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.212x |
| ar-wikipedia | 4096 | 2.120x |
| ar-wikipedia | 65536 | 2.107x |
| ar-wikipedia | 1048576 | 2.125x |
| decimal | 0 | 238.279x |
| en-wikitext2 | 256 | 5.112x |
| en-wikitext2 | 4096 | 5.019x |
| en-wikitext2 | 65536 | 5.010x |
| en-wikitext2 | 1048576 | 4.987x |
| ja-wikipedia | 256 | 2.788x |
| ja-wikipedia | 4096 | 2.587x |
| ja-wikipedia | 65536 | 2.574x |
| ja-wikipedia | 1048576 | 2.587x |
| ru-wikipedia | 256 | 2.297x |
| ru-wikipedia | 4096 | 2.191x |
| ru-wikipedia | 65536 | 2.188x |
| ru-wikipedia | 1048576 | 2.203x |
| synthetic | 0 | 100.074x |
| zh-wikipedia | 256 | 2.729x |
| zh-wikipedia | 4096 | 2.571x |
| zh-wikipedia | 65536 | 2.554x |
| zh-wikipedia | 1048576 | 2.568x |
| Combined selected natural corpora | 256 | 2.838x |
| Combined selected natural corpora | 4096 | 2.714x |
| Combined selected natural corpora | 65536 | 2.704x |
| Combined selected natural corpora | 1048576 | 2.718x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.215x |
| ar-wikipedia | 4096 | 2.110x |
| ar-wikipedia | 65536 | 2.089x |
| ar-wikipedia | 1048576 | 2.095x |
| decimal | 0 | 236.781x |
| en-wikitext2 | 256 | 5.103x |
| en-wikitext2 | 4096 | 5.025x |
| en-wikitext2 | 65536 | 5.001x |
| en-wikitext2 | 1048576 | 4.971x |
| ja-wikipedia | 256 | 2.761x |
| ja-wikipedia | 4096 | 2.584x |
| ja-wikipedia | 65536 | 2.589x |
| ja-wikipedia | 1048576 | 2.565x |
| ru-wikipedia | 256 | 2.298x |
| ru-wikipedia | 4096 | 2.198x |
| ru-wikipedia | 65536 | 2.183x |
| ru-wikipedia | 1048576 | 2.177x |
| synthetic | 0 | 98.096x |
| zh-wikipedia | 256 | 2.711x |
| zh-wikipedia | 4096 | 2.568x |
| zh-wikipedia | 65536 | 2.568x |
| zh-wikipedia | 1048576 | 2.550x |
| Combined selected natural corpora | 256 | 2.831x |
| Combined selected natural corpora | 4096 | 2.712x |
| Combined selected natural corpora | 65536 | 2.701x |
| Combined selected natural corpora | 1048576 | 2.692x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.198x |
| ar-wikipedia | 4096 | 2.095x |
| ar-wikipedia | 65536 | 2.060x |
| ar-wikipedia | 1048576 | 2.079x |
| decimal | 0 | 235.287x |
| en-wikitext2 | 256 | 4.800x |
| en-wikitext2 | 4096 | 4.845x |
| en-wikitext2 | 65536 | 4.857x |
| en-wikitext2 | 1048576 | 4.813x |
| ja-wikipedia | 256 | 2.630x |
| ja-wikipedia | 4096 | 2.535x |
| ja-wikipedia | 65536 | 2.564x |
| ja-wikipedia | 1048576 | 2.527x |
| ru-wikipedia | 256 | 2.260x |
| ru-wikipedia | 4096 | 2.165x |
| ru-wikipedia | 65536 | 2.171x |
| ru-wikipedia | 1048576 | 2.159x |
| synthetic | 0 | 84.652x |
| zh-wikipedia | 256 | 2.623x |
| zh-wikipedia | 4096 | 2.521x |
| zh-wikipedia | 65536 | 2.527x |
| zh-wikipedia | 1048576 | 2.494x |
| Combined selected natural corpora | 256 | 2.752x |
| Combined selected natural corpora | 4096 | 2.666x |
| Combined selected natural corpora | 65536 | 2.664x |
| Combined selected natural corpora | 1048576 | 2.652x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.298x |
| ar-wikipedia | 4096 | 2.842x |
| ar-wikipedia | 65536 | 2.845x |
| ar-wikipedia | 1048576 | 3.262x |
| decimal | 0 | 237.567x |
| en-wikitext2 | 256 | 7.021x |
| en-wikitext2 | 4096 | 6.540x |
| en-wikitext2 | 65536 | 6.563x |
| en-wikitext2 | 1048576 | 7.322x |
| ja-wikipedia | 256 | 4.463x |
| ja-wikipedia | 4096 | 3.833x |
| ja-wikipedia | 65536 | 3.971x |
| ja-wikipedia | 1048576 | 4.676x |
| ru-wikipedia | 256 | 3.446x |
| ru-wikipedia | 4096 | 3.005x |
| ru-wikipedia | 65536 | 2.976x |
| ru-wikipedia | 1048576 | 3.410x |
| synthetic | 0 | 102.101x |
| zh-wikipedia | 256 | 4.276x |
| zh-wikipedia | 4096 | 3.757x |
| zh-wikipedia | 65536 | 3.825x |
| zh-wikipedia | 1048576 | 4.476x |
| Combined selected natural corpora | 256 | 4.233x |
| Combined selected natural corpora | 4096 | 3.736x |
| Combined selected natural corpora | 65536 | 3.767x |
| Combined selected natural corpora | 1048576 | 4.328x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.211x |
| ar-wikipedia | 4096 | 2.815x |
| ar-wikipedia | 65536 | 2.797x |
| ar-wikipedia | 1048576 | 3.237x |
| decimal | 0 | 235.861x |
| en-wikitext2 | 256 | 6.456x |
| en-wikitext2 | 4096 | 6.251x |
| en-wikitext2 | 65536 | 6.348x |
| en-wikitext2 | 1048576 | 7.082x |
| ja-wikipedia | 256 | 4.160x |
| ja-wikipedia | 4096 | 3.722x |
| ja-wikipedia | 65536 | 3.894x |
| ja-wikipedia | 1048576 | 4.588x |
| ru-wikipedia | 256 | 3.325x |
| ru-wikipedia | 4096 | 2.948x |
| ru-wikipedia | 65536 | 2.945x |
| ru-wikipedia | 1048576 | 3.373x |
| synthetic | 0 | 88.188x |
| zh-wikipedia | 256 | 4.032x |
| zh-wikipedia | 4096 | 3.645x |
| zh-wikipedia | 65536 | 3.753x |
| zh-wikipedia | 1048576 | 4.381x |
| Combined selected natural corpora | 256 | 4.027x |
| Combined selected natural corpora | 4096 | 3.647x |
| Combined selected natural corpora | 65536 | 3.698x |
| Combined selected natural corpora | 1048576 | 4.258x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 4754.7 | 50.6 | 10 |
| onig-compile | 4731.6 | 32.7 | 10 |
| pretok-compile | 10.8 | 0.1 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
