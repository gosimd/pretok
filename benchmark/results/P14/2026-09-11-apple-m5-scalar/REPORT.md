# P14 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Interpretation of this P14 run

Apple M5 (explicit host label), scalar Go 1.27.0, Rust 1.90.0, HF Tokenizers
0.22.2 and onig crate 6.5.1; one worker. Ten samples of at least 500 ms per engine
and workload: 1,570 samples total, including separate compilation measurements.

P14 emits uninterrupted Nd triples starting at byte zero. An incomplete triple
or a non-digit ends iteration; later digit runs cannot restart it. Split retains
the remaining suffix as one gap. The plan reuses P10's bounded decimal scanner
and pinned Unicode profile, without additional cursor fields or tables.

The shared `decimal` workload contains 110 successful subjects of 4–4,097 Nd
scalars in five digit layouts: ASCII, Arabic-Indic, fullwidth, mathematical and
mixed-width. It produces 22,100 matches and 22,210 isolated pieces. Reused Split
reaches **421.6 MB/s**, averaging **3.58 us/request**, with throughput ratios
of **6.33×** against Oniguruma reused spans and **11.40×** against native HF Split.
Request sizes vary; this average is not fixed-size latency. P10 and P14 have
different outputs despite using the same decimal input pool.

Natural-text requests mostly exercise early search termination and UTF-8
validation. Their high throughput is not throughput for grouping every character
of a decimal run. At 4 KiB the match/request counts are: en-wikitext2: 3/256, ru-wikipedia: 2/257, ja-wikipedia: 1/257, ar-wikipedia: 0/257, zh-wikipedia: 0/257.
The ratios of summed median batch times across the five 4 KiB languages are
**10.31×** Oniguruma/pretok and **14.38×** HF/pretok.
Individual workloads and variation remain visible below; ratios above one favor
pretok, and ratios below one favor the baseline.

Cursor and reused Split allocate zero bytes across all measured workloads.
Fresh Split allocations, all sizes, ns/request and IQR are reported below.
Native allocation counters are unavailable. Compile medians are
10.13 ns for exact-string pretok lookup,
2.55 us for Oniguruma and 2.58 us for HF.

All 162,826 shared requests passed exact live HF/Oniguruma comparison before
timing. All 24 supported patterns passed the expanded corpus gate. No new
reference output arrays are stored. These comparisons measure regex boundaries
and native HF Split, not a complete tokenizer or BPE.

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 10333247 | 39737 | 2517.8 | 101.48 | 56572247 | 398423 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 6864634 | 28066 | 1672.7 | 152.75 | 85157485 | 599741 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 6768782 | 38060 | 1649.3 | 154.91 | 86363402 | 1921 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 6794563 | 18986 | 1655.6 | 154.33 | 86035702 | 605926 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 1095976 | 2873 | 267.1 | 956.75 | 533383109 | 11862 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 1134175 | 4055 | 276.4 | 924.53 | 515418662 | 3629951 | 66080.0 | 4117.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 1103647 | 1491 | 268.9 | 950.10 | 529675864 | 3730361 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 8641888 | 31105 | 33626.0 | 121.34 | 67644365 | 29739 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 6217655 | 14223 | 24193.2 | 168.64 | 94018571 | 41334 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 6208004 | 38271 | 24155.7 | 168.91 | 94164729 | 0 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 6218435 | 15011 | 24196.2 | 168.62 | 94006773 | 41329 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 1037629 | 2233 | 4037.5 | 1010.55 | 563375676 | 0 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 1039366 | 3128 | 4044.2 | 1008.86 | 562434321 | 247266 | 4112.0 | 257.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 1039513 | 2565 | 4044.8 | 1008.72 | 562354637 | 247231 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 8484035 | 56436 | 499060.9 | 123.59 | 68902943 | 2004 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 6117635 | 33687 | 359860.9 | 171.40 | 95555715 | 2779 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 6144377 | 16195 | 361433.9 | 170.66 | 95139836 | 0 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 6133101 | 22020 | 360770.7 | 170.97 | 95314750 | 2772 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 1035090 | 2225 | 60887.7 | 1013.03 | 564757587 | 0 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 1036058 | 2572 | 60944.6 | 1012.08 | 564229846 | 16408 | 272.0 | 17.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 1035093 | 2537 | 60887.8 | 1013.03 | 564755846 | 16424 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 10142536 | 112398 | 10142536.3 | 103.38 | 57635978 | 99 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 6119160 | 16638 | 6119159.5 | 171.36 | 95531910 | 163 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 6124281 | 22631 | 6124281.2 | 171.22 | 95452017 | 0 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 6118720 | 14691 | 6118719.5 | 171.37 | 95538780 | 163 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 1013862 | 2311 | 1013862.2 | 1034.24 | 576582289 | 0 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 1016748 | 2329 | 1016748.2 | 1031.30 | 574945706 | 984 | 16.0 | 1.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 1015310 | 2849 | 1015310.4 | 1032.76 | 575759893 | 985 | 0.0 | 0.0 | 10 |
| decimal | 0 | hf-split | 4491938 | 81314 | 40835.8 | 36.99 | 14796510 | 4944414 | N/A | N/A | 10 |
| decimal | 0 | onig-fresh | 2557606 | 42608 | 23251.0 | 64.96 | 25987195 | 8683903 | N/A | N/A | 10 |
| decimal | 0 | onig-iter | 2501896 | 47732 | 22744.5 | 66.41 | 26565856 | 8833302 | N/A | N/A | 10 |
| decimal | 0 | onig-reused | 2493213 | 42278 | 22665.6 | 66.64 | 26658376 | 8908185 | N/A | N/A | 10 |
| decimal | 0 | pretok-cursor | 390163 | 3669 | 3546.9 | 425.83 | 170351973 | 56643024 | 0.0 | 0.0 | 10 |
| decimal | 0 | pretok-fresh | 504350 | 8934 | 4585.0 | 329.42 | 131783576 | 44036910 | 1198240.0 | 760.0 | 10 |
| decimal | 0 | pretok-reused | 394092 | 6304 | 3582.7 | 421.59 | 168653563 | 56357416 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 12412659 | 41847 | 3029.7 | 84.48 | 84304019 | 330952 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 9646707 | 28677 | 2354.6 | 108.70 | 108476087 | 425845 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 9565094 | 24325 | 2334.7 | 109.63 | 109401648 | 1150 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 9596645 | 9801 | 2342.4 | 109.26 | 109041962 | 428066 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 66637 | 289 | 16.3 | 15735.60 | 15703498797 | 165073 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 119344 | 2363 | 29.1 | 8786.16 | 8768239300 | 34421496 | 65904.0 | 4108.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 73984 | 238 | 18.1 | 14172.94 | 14144025583 | 55525232 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 11424441 | 53854 | 44626.7 | 91.78 | 91596339 | 22671 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 9139377 | 17972 | 35700.7 | 114.73 | 114497626 | 28339 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 9124900 | 32294 | 35644.1 | 114.91 | 114679284 | 329 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 9135932 | 36338 | 35687.2 | 114.77 | 114540807 | 28350 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 21239 | 46 | 83.0 | 49371.40 | 49270687501 | 141253 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 24620 | 119 | 96.2 | 42590.44 | 42503555685 | 10519908 | 4192.0 | 259.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 21738 | 114 | 84.9 | 48237.72 | 48139315870 | 11914795 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 11193338 | 20734 | 699583.7 | 93.68 | 93487480 | 1429 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 9101439 | 25120 | 568839.9 | 115.21 | 114974895 | 1758 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 9088377 | 20124 | 568023.6 | 115.38 | 115140140 | 0 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 9096184 | 14565 | 568511.5 | 115.28 | 115041322 | 1759 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 18788 | 64 | 1174.2 | 55811.19 | 55697339887 | 0 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 18975 | 71 | 1185.9 | 55261.39 | 55148663766 | 843222 | 256.0 | 16.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 18808 | 55 | 1175.5 | 55750.22 | 55636496496 | 850681 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 12359404 | 122152 | 12359404.0 | 84.84 | 84667271 | 81 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 8987388 | 28406 | 8987388.4 | 116.67 | 116433935 | 111 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 8963769 | 40964 | 8963769.0 | 116.98 | 116740737 | 0 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 8977552 | 34402 | 8977552.1 | 116.80 | 116561507 | 111 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 18658 | 92 | 18658.5 | 56198.35 | 56083712075 | 0 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 18702 | 66 | 18702.4 | 56066.48 | 55952107494 | 53469 | 16.0 | 1.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 18676 | 79 | 18676.2 | 56145.18 | 56030644874 | 53544 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 7404215 | 11710 | 1800.2 | 141.62 | 51224874 | 556980 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 4575937 | 14804 | 1112.6 | 229.15 | 82885751 | 901236 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 4514016 | 14513 | 1097.5 | 232.29 | 84022735 | 2437 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 4514891 | 9958 | 1097.7 | 232.25 | 84006462 | 913422 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 476274 | 2611 | 115.8 | 2201.62 | 796348539 | 23096 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 525525 | 1513 | 127.8 | 1995.29 | 721716300 | 7847390 | 66160.0 | 4124.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 484587 | 1069 | 117.8 | 2163.85 | 782687655 | 8510346 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 6218101 | 92570 | 24194.9 | 168.63 | 60996114 | 41492 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 3980448 | 35641 | 15488.1 | 263.43 | 95285759 | 64817 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 3971540 | 22390 | 15453.5 | 264.02 | 95499475 | 252 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 3968891 | 46674 | 15443.2 | 264.20 | 95563218 | 65006 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 424388 | 8170 | 1651.3 | 2470.79 | 893710848 | 2356 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 428532 | 4869 | 1667.4 | 2446.90 | 885068745 | 602056 | 4144.0 | 258.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 425539 | 1882 | 1655.8 | 2464.11 | 891293262 | 606290 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 6221301 | 194734 | 365958.9 | 168.55 | 60964736 | 2733 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 3945796 | 136979 | 232105.7 | 265.74 | 96122550 | 4308 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 3978841 | 88494 | 234049.4 | 263.54 | 95324250 | 0 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 3931427 | 24887 | 231260.4 | 266.72 | 96473886 | 4324 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 427477 | 4885 | 25145.7 | 2452.94 | 887251723 | 0 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 424957 | 7792 | 24997.5 | 2467.49 | 892514452 | 40004 | 272.0 | 17.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 426765 | 6578 | 25103.8 | 2457.03 | 888733176 | 39835 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 7852548 | 69618 | 7852547.5 | 133.53 | 48300249 | 127 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 3900686 | 42951 | 3900686.1 | 268.82 | 97234178 | 256 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 3891855 | 45907 | 3891855.1 | 269.43 | 97454809 | 0 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 3904078 | 40227 | 3904078.2 | 268.58 | 97149694 | 256 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 416468 | 1195 | 416468.1 | 2517.78 | 910706007 | 0 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 419499 | 11649 | 419499.3 | 2499.59 | 904125564 | 2384 | 16.0 | 1.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 417627 | 1720 | 417626.6 | 2510.80 | 908179795 | 2394 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 9994211 | 16836 | 2435.2 | 104.92 | 57518999 | 411938 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 6636794 | 6085 | 1617.2 | 157.99 | 86616672 | 620330 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 6557544 | 13662 | 1597.8 | 159.90 | 87663466 | 1982 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 6572288 | 25234 | 1601.4 | 159.54 | 87466798 | 626418 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 1009445 | 939 | 246.0 | 1038.76 | 569478481 | 12878 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 1039577 | 2451 | 253.3 | 1008.66 | 552972043 | 3960265 | 66080.0 | 4117.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 1015533 | 1289 | 247.4 | 1032.54 | 566064226 | 4054028 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 8406594 | 35388 | 32710.5 | 124.73 | 68381680 | 30809 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 6014489 | 38733 | 23402.7 | 174.34 | 95578700 | 43063 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 6016622 | 33435 | 23411.0 | 174.28 | 95544813 | 332 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 6001714 | 46922 | 23353.0 | 174.71 | 95782146 | 43154 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 952983 | 2043 | 3708.1 | 1100.31 | 603218540 | 2099 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 951376 | 4661 | 3701.9 | 1102.17 | 604237750 | 272237 | 4176.0 | 259.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 953475 | 2513 | 3710.0 | 1099.74 | 602907018 | 271638 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 8172501 | 9399 | 480735.3 | 128.31 | 70340404 | 2080 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 5908283 | 7156 | 347546.0 | 177.48 | 97296802 | 2877 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 5895416 | 13996 | 346789.2 | 177.86 | 97509143 | 0 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 5895834 | 17506 | 346813.8 | 177.85 | 97502239 | 2883 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 948363 | 1196 | 55786.0 | 1105.67 | 606157360 | 0 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 947695 | 2020 | 55746.8 | 1106.45 | 606584323 | 17938 | 272.0 | 17.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 949099 | 1358 | 55829.3 | 1104.81 | 605687186 | 17912 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 9396186 | 79139 | 9396186.3 | 111.60 | 61179821 | 106 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 5888878 | 26004 | 5888877.7 | 178.06 | 97617412 | 170 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 5876893 | 35990 | 5876893.2 | 178.42 | 97816479 | 0 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 5888977 | 28590 | 5888977.5 | 178.06 | 97615758 | 170 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 929690 | 3170 | 929690.4 | 1127.88 | 618331623 | 0 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 930053 | 1912 | 930053.1 | 1127.44 | 618090497 | 1075 | 16.0 | 1.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 933372 | 2719 | 933372.0 | 1123.43 | 615892687 | 1071 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 106067912 | 1331360 | 753.3 | 39.38 | 19194382 | 2626544 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 53233935 | 744648 | 378.0 | 78.47 | 38244552 | 5233353 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 50619390 | 1023817 | 359.5 | 82.53 | 40219924 | 2725043 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 50472519 | 952297 | 358.4 | 82.77 | 40336961 | 5519677 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 5755513 | 49513 | 40.9 | 725.81 | 353731794 | 23966586 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 8552299 | 129801 | 60.7 | 488.45 | 238053873 | 32575099 | 9722992.0 | 149650.0 | 10 |
| synthetic | 0 | pretok-reused | 5997360 | 84019 | 42.6 | 696.54 | 339467387 | 46452442 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 7664263 | 18163 | 1863.4 | 136.81 | 50912004 | 537690 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 4797456 | 20492 | 1166.4 | 218.57 | 81335404 | 858997 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 4715224 | 10901 | 1146.4 | 222.38 | 82753870 | 1697 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 4731774 | 13678 | 1150.4 | 221.60 | 82464424 | 870921 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 472439 | 2743 | 114.9 | 2219.49 | 825932232 | 16933 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 525679 | 4473 | 127.8 | 1994.71 | 742284391 | 7839391 | 66064.0 | 4121.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 480898 | 2386 | 116.9 | 2180.45 | 811404090 | 8569376 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 6385776 | 20111 | 24847.4 | 164.20 | 61105021 | 40246 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 4134010 | 17676 | 16085.6 | 253.65 | 94388502 | 62167 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 4123141 | 38148 | 16043.4 | 254.31 | 94637315 | 0 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 4128428 | 25416 | 16063.9 | 253.99 | 94516128 | 62251 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 414894 | 1331 | 1614.4 | 2527.33 | 940488753 | 0 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 419589 | 4446 | 1632.6 | 2499.05 | 929964937 | 612504 | 4112.0 | 257.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 416822 | 2025 | 1621.9 | 2515.64 | 936137814 | 616570 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 6357045 | 34653 | 373943.8 | 164.95 | 61381196 | 2674 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 4070521 | 19598 | 239442.4 | 257.60 | 95860706 | 4176 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 4069363 | 22686 | 239374.3 | 257.68 | 95887977 | 0 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 4082788 | 16434 | 240164.0 | 256.83 | 95572680 | 4164 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 412776 | 2355 | 24280.9 | 2540.30 | 945314661 | 0 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 414130 | 973 | 24360.6 | 2531.99 | 942223021 | 41050 | 272.0 | 17.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 414954 | 2305 | 24409.0 | 2526.97 | 940353228 | 40968 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 8286436 | 55852 | 8286436.5 | 126.54 | 47089361 | 121 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 4091430 | 91033 | 4091429.5 | 256.29 | 95370822 | 244 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 4075654 | 25667 | 4075654.3 | 257.28 | 95739965 | 0 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 4075175 | 18820 | 4075174.6 | 257.31 | 95751234 | 245 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 411228 | 3946 | 411227.6 | 2549.86 | 948873569 | 0 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 413483 | 3088 | 413482.9 | 2535.95 | 943697987 | 2418 | 16.0 | 1.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 412438 | 4829 | 412437.9 | 2542.38 | 946089055 | 2425 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 6.176x |
| ar-wikipedia | 4096 | 5.983x |
| ar-wikipedia | 65536 | 5.936x |
| ar-wikipedia | 1048576 | 6.041x |
| decimal | 0 | 6.412x |
| en-wikitext2 | 256 | 143.540x |
| en-wikitext2 | 4096 | 429.639x |
| en-wikitext2 | 65536 | 483.735x |
| en-wikitext2 | 1048576 | 480.413x |
| ja-wikipedia | 256 | 9.478x |
| ja-wikipedia | 4096 | 9.358x |
| ja-wikipedia | 65536 | 9.308x |
| ja-wikipedia | 1048576 | 9.345x |
| ru-wikipedia | 256 | 6.496x |
| ru-wikipedia | 4096 | 6.313x |
| ru-wikipedia | 65536 | 6.216x |
| ru-wikipedia | 1048576 | 6.321x |
| synthetic | 0 | 8.795x |
| zh-wikipedia | 256 | 9.981x |
| zh-wikipedia | 4096 | 9.938x |
| zh-wikipedia | 65536 | 9.859x |
| zh-wikipedia | 1048576 | 9.911x |
| Combined selected natural corpora | 256 | 10.293x |
| Combined selected natural corpora | 4096 | 10.327x |
| Combined selected natural corpora | 65536 | 10.264x |
| Combined selected natural corpora | 1048576 | 10.370x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 6.156x |
| ar-wikipedia | 4096 | 5.982x |
| ar-wikipedia | 65536 | 5.925x |
| ar-wikipedia | 1048576 | 6.026x |
| decimal | 0 | 6.326x |
| en-wikitext2 | 256 | 129.712x |
| en-wikitext2 | 4096 | 420.281x |
| en-wikitext2 | 65536 | 483.622x |
| en-wikitext2 | 1048576 | 480.696x |
| ja-wikipedia | 256 | 9.317x |
| ja-wikipedia | 4096 | 9.327x |
| ja-wikipedia | 65536 | 9.212x |
| ja-wikipedia | 1048576 | 9.348x |
| ru-wikipedia | 256 | 6.472x |
| ru-wikipedia | 4096 | 6.295x |
| ru-wikipedia | 65536 | 6.212x |
| ru-wikipedia | 1048576 | 6.309x |
| synthetic | 0 | 8.416x |
| zh-wikipedia | 256 | 9.839x |
| zh-wikipedia | 4096 | 9.905x |
| zh-wikipedia | 65536 | 9.839x |
| zh-wikipedia | 1048576 | 9.881x |
| Combined selected natural corpora | 256 | 10.197x |
| Combined selected natural corpora | 4096 | 10.309x |
| Combined selected natural corpora | 65536 | 10.243x |
| Combined selected natural corpora | 1048576 | 10.354x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 6.053x |
| ar-wikipedia | 4096 | 5.982x |
| ar-wikipedia | 65536 | 5.905x |
| ar-wikipedia | 1048576 | 6.018x |
| decimal | 0 | 5.071x |
| en-wikitext2 | 256 | 80.831x |
| en-wikitext2 | 4096 | 371.218x |
| en-wikitext2 | 65536 | 479.658x |
| en-wikitext2 | 1048576 | 480.548x |
| ja-wikipedia | 256 | 8.707x |
| ja-wikipedia | 4096 | 9.289x |
| ja-wikipedia | 65536 | 9.285x |
| ja-wikipedia | 1048576 | 9.298x |
| ru-wikipedia | 256 | 6.384x |
| ru-wikipedia | 4096 | 6.322x |
| ru-wikipedia | 65536 | 6.234x |
| ru-wikipedia | 1048576 | 6.332x |
| synthetic | 0 | 6.225x |
| zh-wikipedia | 256 | 9.126x |
| zh-wikipedia | 4096 | 9.853x |
| zh-wikipedia | 65536 | 9.829x |
| zh-wikipedia | 1048576 | 9.895x |
| Combined selected natural corpora | 256 | 9.724x |
| Combined selected natural corpora | 4096 | 10.297x |
| Combined selected natural corpora | 65536 | 10.255x |
| Combined selected natural corpora | 1048576 | 10.358x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 9.363x |
| ar-wikipedia | 4096 | 8.313x |
| ar-wikipedia | 65536 | 8.196x |
| ar-wikipedia | 1048576 | 9.990x |
| decimal | 0 | 11.398x |
| en-wikitext2 | 256 | 167.774x |
| en-wikitext2 | 4096 | 525.559x |
| en-wikitext2 | 65536 | 595.122x |
| en-wikitext2 | 1048576 | 661.775x |
| ja-wikipedia | 256 | 15.279x |
| ja-wikipedia | 4096 | 14.612x |
| ja-wikipedia | 65536 | 14.578x |
| ja-wikipedia | 1048576 | 18.803x |
| ru-wikipedia | 256 | 9.841x |
| ru-wikipedia | 4096 | 8.817x |
| ru-wikipedia | 65536 | 8.611x |
| ru-wikipedia | 1048576 | 10.067x |
| synthetic | 0 | 17.686x |
| zh-wikipedia | 256 | 15.937x |
| zh-wikipedia | 4096 | 15.320x |
| zh-wikipedia | 65536 | 15.320x |
| zh-wikipedia | 1048576 | 20.091x |
| Combined selected natural corpora | 256 | 15.136x |
| Combined selected natural corpora | 4096 | 14.377x |
| Combined selected natural corpora | 65536 | 14.212x |
| Combined selected natural corpora | 1048576 | 17.172x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 9.111x |
| ar-wikipedia | 4096 | 8.315x |
| ar-wikipedia | 65536 | 8.189x |
| ar-wikipedia | 1048576 | 9.975x |
| decimal | 0 | 8.906x |
| en-wikitext2 | 256 | 104.007x |
| en-wikitext2 | 4096 | 464.031x |
| en-wikitext2 | 65536 | 589.904x |
| en-wikitext2 | 1048576 | 660.847x |
| ja-wikipedia | 256 | 14.089x |
| ja-wikipedia | 4096 | 14.510x |
| ja-wikipedia | 65536 | 14.640x |
| ja-wikipedia | 1048576 | 18.719x |
| ru-wikipedia | 256 | 9.614x |
| ru-wikipedia | 4096 | 8.836x |
| ru-wikipedia | 65536 | 8.624x |
| ru-wikipedia | 1048576 | 10.103x |
| synthetic | 0 | 12.402x |
| zh-wikipedia | 256 | 14.580x |
| zh-wikipedia | 4096 | 15.219x |
| zh-wikipedia | 65536 | 15.350x |
| zh-wikipedia | 1048576 | 20.041x |
| Combined selected natural corpora | 256 | 14.296x |
| Combined selected natural corpora | 4096 | 14.345x |
| Combined selected natural corpora | 65536 | 14.226x |
| Combined selected natural corpora | 1048576 | 17.165x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 2575.3 | 53.4 | 10 |
| onig-compile | 2548.9 | 45.0 | 10 |
| pretok-compile | 10.1 | 0.3 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
