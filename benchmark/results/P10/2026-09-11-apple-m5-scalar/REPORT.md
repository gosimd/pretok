# P10 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Interpretation of this P10 run

Apple M5 (explicit host label), scalar Go 1.27.0, Rust 1.90.0, HF Tokenizers
0.22.2, onig crate 6.5.1, one worker. Ten samples of at least 500 ms per engine
and workload; 1,570 samples total, including separate compilation timing.

The shared `decimal` workload contains 110 successful subjects, each 4–4,097
Nd scalars, with ASCII, Arabic-Indic, fullwidth, mathematical and mixed-width
digits. Reused Split reaches **491.7 MB/s**, averaging **3.07 us/request**:
**1.50×** Oniguruma reused-span collection and **2.62×** native HF Split throughput.
Its requests vary in byte and scalar length; this average is not a fixed-size latency.

Every natural-text request in this run has zero P10 matches and one preserved
Split gap. Those workloads predominantly measure UTF-8 validation and immediate
rejection, not decimal grouping. At 4 KiB, pretok is about 25% slower than
Oniguruma on Russian and Arabic. Ratios of summed median batch times across the
five 4 KiB languages are **0.96×** for Oniguruma/pretok and
**4.97×** for HF/pretok; a value below one favors the baseline.
The unusually high English throughput is a rejection workload, not general
pre-tokenization throughput.

Cursor and reused Split allocate zero bytes in every measured workload. Fresh
Split allocations, per-request times and IQR are included below. Native allocation
counters are unavailable. Compile medians are 9.98 ns for exact-string
pretok lookup, 2.67 us for Oniguruma and 2.69 us for HF.

All 162,794 shared requests passed exact live HF/Oniguruma comparison before
timing. All 23 supported patterns also passed the expanded corpus gate. Results
measure regex boundaries and native HF Split, not a complete tokenizer or BPE.

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 4875926 | 14852 | 1188.1 | 215.05 | 119890048 | 841686 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 1450705 | 19757 | 353.5 | 722.80 | 402959336 | 2828970 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 1370431 | 11996 | 333.9 | 765.14 | 426562933 | 0 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 1372737 | 12038 | 334.5 | 763.86 | 425846325 | 2989648 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 1082581 | 1424 | 263.8 | 968.59 | 539982695 | 0 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 1121549 | 4570 | 273.3 | 934.94 | 521221115 | 3659225 | 65664.0 | 4104.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 1089889 | 2252 | 265.6 | 962.09 | 536362037 | 3765522 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 3170042 | 13556 | 12334.8 | 330.78 | 184406063 | 81071 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 827140 | 13671 | 3218.4 | 1267.71 | 706742711 | 310709 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 818389 | 11063 | 3184.4 | 1281.27 | 714299870 | 0 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 816448 | 14827 | 3176.8 | 1284.31 | 715997444 | 314778 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 1018029 | 2232 | 3961.2 | 1030.01 | 574222103 | 0 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 1018733 | 2783 | 3963.9 | 1029.29 | 573825392 | 252274 | 4112.0 | 257.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 1016478 | 3098 | 3955.2 | 1031.58 | 575098373 | 252834 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 3041636 | 28152 | 178919.8 | 344.74 | 192190948 | 5589 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 781096 | 26019 | 45946.8 | 1342.44 | 748403748 | 21764 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 779739 | 22067 | 45867.0 | 1344.78 | 749706442 | 0 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 770388 | 22842 | 45317.0 | 1361.10 | 758805531 | 22067 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 1013123 | 3891 | 59595.5 | 1034.99 | 577003148 | 0 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 1014181 | 2923 | 59657.7 | 1033.91 | 576400841 | 16762 | 272.0 | 17.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 1013040 | 2800 | 59590.6 | 1035.08 | 577050396 | 16781 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 4363405 | 16599 | 4363405.4 | 240.31 | 133972194 | 229 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 772047 | 9394 | 772047.5 | 1358.18 | 757174929 | 1295 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 766865 | 1834 | 766864.6 | 1367.35 | 762292319 | 0 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 766803 | 3388 | 766802.9 | 1367.46 | 762353659 | 1304 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 1002192 | 7568 | 1002192.3 | 1046.28 | 583296263 | 0 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 1005947 | 2283 | 1005947.2 | 1042.38 | 581118946 | 994 | 16.0 | 1.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 1006376 | 6455 | 1006376.1 | 1041.93 | 580871292 | 994 | 0.0 | 0.0 | 10 |
| decimal | 0 | hf-split | 884953 | 5701 | 8045.0 | 187.74 | 75105650 | 248601 | N/A | N/A | 10 |
| decimal | 0 | onig-fresh | 506260 | 5899 | 4602.4 | 328.18 | 131286199 | 434559 | N/A | N/A | 10 |
| decimal | 0 | onig-iter | 508892 | 7845 | 4626.3 | 326.48 | 130607334 | 216156 | N/A | N/A | 10 |
| decimal | 0 | onig-reused | 508270 | 3793 | 4620.6 | 326.88 | 130767197 | 432841 | N/A | N/A | 10 |
| decimal | 0 | pretok-cursor | 338191 | 4476 | 3074.5 | 491.27 | 196530719 | 325260 | 0.0 | 0.0 | 10 |
| decimal | 0 | pretok-fresh | 339318 | 1649 | 3084.7 | 489.64 | 195877918 | 648358 | 5280.0 | 220.0 | 10 |
| decimal | 0 | pretok-reused | 337926 | 2449 | 3072.1 | 491.66 | 196684852 | 651029 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 3404133 | 6355 | 830.9 | 308.03 | 307401945 | 1203537 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 660072 | 2202 | 161.1 | 1588.58 | 1585337440 | 6206898 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 594348 | 3997 | 145.1 | 1764.25 | 1760647311 | 0 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 600279 | 6536 | 146.5 | 1746.81 | 1743250901 | 6825159 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 67472 | 361 | 16.5 | 15540.84 | 15509135862 | 0 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 120642 | 1835 | 29.4 | 8691.67 | 8673935837 | 33960110 | 65552.0 | 4097.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 74324 | 651 | 18.1 | 14108.16 | 14079380905 | 55123456 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 2294283 | 8649 | 8962.0 | 457.04 | 456106290 | 111582 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 66330 | 423 | 259.1 | 15808.51 | 15776262186 | 3859500 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 62324 | 107 | 243.5 | 16824.55 | 16790233141 | 0 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 62626 | 230 | 244.6 | 16743.38 | 16709224041 | 4087739 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 21626 | 107 | 84.5 | 48485.79 | 48386884869 | 0 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 25141 | 226 | 98.2 | 41707.27 | 41622195390 | 10182440 | 4096.0 | 256.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 22164 | 176 | 86.6 | 47309.72 | 47213213335 | 11550225 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 2256570 | 11279 | 141035.6 | 464.68 | 463729001 | 7090 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 28895 | 134 | 1806.0 | 36288.58 | 36214557999 | 553720 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 28677 | 77 | 1792.3 | 36564.96 | 36490374117 | 0 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 28689 | 121 | 1793.1 | 36549.78 | 36475220509 | 557705 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 19090 | 98 | 1193.1 | 54929.21 | 54817160358 | 0 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 19334 | 58 | 1208.4 | 54234.56 | 54123926422 | 827554 | 256.0 | 16.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 19154 | 43 | 1197.2 | 54743.11 | 54631441224 | 835314 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 3794747 | 173390 | 3794747.2 | 276.32 | 275759345 | 264 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 26042 | 319 | 26042.0 | 40264.82 | 40182682671 | 38400 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 26206 | 233 | 26206.3 | 40012.35 | 39930728711 | 0 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 26024 | 263 | 26023.9 | 40292.82 | 40210628350 | 38426 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 18902 | 78 | 18901.9 | 55474.75 | 55361582220 | 0 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 18876 | 243 | 18876.4 | 55549.60 | 55436284683 | 52976 | 16.0 | 1.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 18914 | 199 | 18914.4 | 55437.94 | 55324852501 | 52870 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 4051107 | 48067 | 985.0 | 258.84 | 93623785 | 1015278 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 1156427 | 13979 | 281.2 | 906.74 | 327975763 | 3556645 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 1088899 | 23296 | 264.7 | 962.97 | 348315029 | 0 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 1091090 | 20610 | 265.3 | 961.03 | 347615598 | 3769624 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 477912 | 1853 | 116.2 | 2194.08 | 793619188 | 0 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 526888 | 1286 | 128.1 | 1990.13 | 719849771 | 7806217 | 65808.0 | 4113.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 485505 | 1861 | 118.0 | 2159.76 | 781207761 | 8471598 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 2721675 | 36499 | 10590.2 | 385.27 | 139355346 | 94427 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 532893 | 20026 | 2073.5 | 1967.70 | 711737786 | 482273 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 530035 | 12688 | 2062.4 | 1978.31 | 715575342 | 0 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 529645 | 4896 | 2060.9 | 1979.77 | 716102560 | 485231 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 423265 | 4000 | 1646.9 | 2477.35 | 896082335 | 0 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 426117 | 2339 | 1658.0 | 2460.77 | 890084069 | 603121 | 4112.0 | 257.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 424845 | 3255 | 1653.1 | 2468.14 | 892749082 | 604926 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 2657829 | 21196 | 156342.9 | 394.52 | 142702945 | 6396 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 474504 | 18157 | 27912.0 | 2209.83 | 799318715 | 35827 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 475077 | 27935 | 27945.7 | 2207.17 | 798355344 | 0 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 478170 | 18845 | 28127.7 | 2192.89 | 793190459 | 35552 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 420613 | 1716 | 24741.9 | 2492.97 | 901731511 | 0 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 422017 | 1028 | 24824.5 | 2484.68 | 898732501 | 40283 | 272.0 | 17.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 421594 | 483 | 24799.6 | 2487.17 | 899634353 | 40323 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 4016832 | 87927 | 4016832.5 | 261.05 | 94422658 | 249 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 472650 | 29832 | 472650.4 | 2218.50 | 802453554 | 2116 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 474285 | 24802 | 474285.2 | 2210.85 | 799687567 | 0 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 472524 | 24232 | 472523.5 | 2219.10 | 802668993 | 2116 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 419013 | 2915 | 419013.5 | 2502.49 | 905173777 | 0 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 421833 | 2657 | 421833.4 | 2485.76 | 899122664 | 2371 | 16.0 | 1.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 420543 | 1620 | 420543.2 | 2493.38 | 901881093 | 2378 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 4735658 | 26600 | 1153.9 | 221.42 | 121389054 | 866617 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 1388791 | 12434 | 338.4 | 755.03 | 413926223 | 2955088 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 1309070 | 22433 | 319.0 | 801.01 | 439133989 | 0 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 1312028 | 21001 | 319.7 | 799.20 | 438143761 | 3127981 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 1012533 | 2430 | 246.7 | 1035.60 | 567741224 | 0 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 1045324 | 4686 | 254.7 | 1003.11 | 549931695 | 3926054 | 65664.0 | 4104.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 1019113 | 2223 | 248.3 | 1028.91 | 564075659 | 4027030 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 3134407 | 14553 | 12196.1 | 334.54 | 183402177 | 81993 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 771755 | 12368 | 3002.9 | 1358.69 | 744869825 | 333007 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 757313 | 9365 | 2946.7 | 1384.60 | 759074680 | 0 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 758454 | 13318 | 2951.2 | 1382.52 | 757932371 | 338847 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 948702 | 5177 | 3691.4 | 1105.27 | 605940378 | 0 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 945271 | 5194 | 3678.1 | 1109.28 | 608139615 | 271880 | 4112.0 | 257.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 949597 | 3823 | 3694.9 | 1104.23 | 605369555 | 270641 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 2964841 | 8402 | 174402.4 | 353.67 | 193891360 | 5734 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 706254 | 8777 | 41544.4 | 1484.70 | 813951747 | 24071 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 704337 | 3694 | 41431.6 | 1488.74 | 816167107 | 0 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 704248 | 7241 | 41426.3 | 1488.93 | 816271101 | 24139 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 930020 | 3665 | 54707.0 | 1127.48 | 618112783 | 0 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 929139 | 4754 | 54655.2 | 1128.54 | 618698651 | 18297 | 272.0 | 17.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 930038 | 3861 | 54708.1 | 1127.45 | 618100429 | 18279 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 4569061 | 170854 | 4569061.2 | 229.49 | 125815124 | 219 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 703034 | 13100 | 703034.4 | 1491.50 | 817679745 | 1422 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 698153 | 8381 | 698153.1 | 1501.93 | 823396752 | 0 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 699353 | 8695 | 699353.0 | 1499.35 | 821984062 | 1430 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 916560 | 1345 | 916560.2 | 1144.03 | 627189539 | 0 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 916778 | 3133 | 916778.3 | 1143.76 | 627040332 | 1091 | 16.0 | 1.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 916216 | 1619 | 916216.2 | 1144.46 | 627425060 | 1091 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 71986539 | 124054 | 511.3 | 58.02 | 28277342 | 1987566 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 27557058 | 98643 | 195.7 | 151.57 | 73868117 | 5192064 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 25219541 | 114615 | 179.1 | 165.62 | 80714714 | 91159 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 25320269 | 82601 | 179.9 | 164.96 | 80393618 | 5650730 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 5476870 | 42359 | 38.9 | 762.65 | 371669918 | 419765 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 7415946 | 42238 | 52.7 | 563.23 | 274487972 | 19293290 | 2326032.0 | 143078.0 | 10 |
| synthetic | 0 | pretok-reused | 5669213 | 14400 | 40.3 | 736.77 | 359060066 | 25237718 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 4056269 | 52055 | 986.2 | 258.51 | 96197512 | 1013986 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 1174731 | 39489 | 285.6 | 892.61 | 332163636 | 3501226 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 1107678 | 52937 | 269.3 | 946.64 | 352271298 | 0 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 1111395 | 37646 | 270.2 | 943.48 | 351092960 | 3700754 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 471275 | 4912 | 114.6 | 2224.97 | 827972818 | 0 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 520958 | 4206 | 126.7 | 2012.78 | 749010341 | 7895069 | 65808.0 | 4113.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 480681 | 2134 | 116.9 | 2181.43 | 811770687 | 8556605 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 2742756 | 13457 | 10672.2 | 382.31 | 142266778 | 93701 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 545877 | 11502 | 2124.0 | 1920.90 | 714818656 | 470802 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 544129 | 8933 | 2117.2 | 1927.07 | 717115375 | 0 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 542739 | 10327 | 2111.8 | 1932.00 | 718951214 | 473524 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 418476 | 1501 | 1628.3 | 2505.70 | 932438039 | 0 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 421983 | 769 | 1642.0 | 2484.87 | 924687872 | 609029 | 4112.0 | 257.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 419109 | 658 | 1630.8 | 2501.91 | 931029819 | 613206 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 2718322 | 15792 | 159901.3 | 385.74 | 143545553 | 6254 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 532092 | 27727 | 31299.5 | 1970.66 | 733337080 | 31949 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 530515 | 26031 | 31206.7 | 1976.52 | 735517727 | 0 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 530312 | 25549 | 31194.9 | 1977.28 | 735798287 | 32057 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 413935 | 1900 | 24349.1 | 2533.18 | 942667171 | 0 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 415352 | 1519 | 24432.5 | 2524.54 | 939450870 | 40929 | 272.0 | 17.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 415318 | 742 | 24430.5 | 2524.75 | 939527320 | 40932 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 4036814 | 21292 | 4036813.8 | 259.75 | 96661133 | 248 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 509992 | 24743 | 509992.1 | 2056.06 | 765115759 | 1961 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 502997 | 22860 | 502997.5 | 2084.65 | 775755353 | 0 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 503414 | 21690 | 503413.6 | 2082.93 | 775114140 | 1986 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 410224 | 3716 | 410223.9 | 2556.10 | 951195102 | 0 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 412692 | 1173 | 412692.1 | 2540.81 | 945506305 | 2423 | 16.0 | 1.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 412099 | 2448 | 412099.1 | 2544.47 | 946867014 | 2427 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.266x |
| ar-wikipedia | 4096 | 0.804x |
| ar-wikipedia | 65536 | 0.770x |
| ar-wikipedia | 1048576 | 0.765x |
| decimal | 0 | 1.505x |
| en-wikitext2 | 256 | 8.809x |
| en-wikitext2 | 4096 | 2.882x |
| en-wikitext2 | 65536 | 1.502x |
| en-wikitext2 | 1048576 | 1.386x |
| ja-wikipedia | 256 | 2.278x |
| ja-wikipedia | 4096 | 1.252x |
| ja-wikipedia | 65536 | 1.129x |
| ja-wikipedia | 1048576 | 1.132x |
| ru-wikipedia | 256 | 1.293x |
| ru-wikipedia | 4096 | 0.798x |
| ru-wikipedia | 65536 | 0.757x |
| ru-wikipedia | 1048576 | 0.762x |
| synthetic | 0 | 4.605x |
| zh-wikipedia | 256 | 2.350x |
| zh-wikipedia | 4096 | 1.300x |
| zh-wikipedia | 65536 | 1.282x |
| zh-wikipedia | 1048576 | 1.226x |
| Combined selected natural corpora | 256 | 1.758x |
| Combined selected natural corpora | 4096 | 0.958x |
| Combined selected natural corpora | 65536 | 0.900x |
| Combined selected natural corpora | 1048576 | 0.892x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.260x |
| ar-wikipedia | 4096 | 0.803x |
| ar-wikipedia | 65536 | 0.760x |
| ar-wikipedia | 1048576 | 0.762x |
| decimal | 0 | 1.504x |
| en-wikitext2 | 256 | 8.077x |
| en-wikitext2 | 4096 | 2.826x |
| en-wikitext2 | 65536 | 1.498x |
| en-wikitext2 | 1048576 | 1.376x |
| ja-wikipedia | 256 | 2.247x |
| ja-wikipedia | 4096 | 1.247x |
| ja-wikipedia | 65536 | 1.134x |
| ja-wikipedia | 1048576 | 1.124x |
| ru-wikipedia | 256 | 1.287x |
| ru-wikipedia | 4096 | 0.799x |
| ru-wikipedia | 65536 | 0.757x |
| ru-wikipedia | 1048576 | 0.763x |
| synthetic | 0 | 4.466x |
| zh-wikipedia | 256 | 2.312x |
| zh-wikipedia | 4096 | 1.295x |
| zh-wikipedia | 65536 | 1.277x |
| zh-wikipedia | 1048576 | 1.222x |
| Combined selected natural corpora | 256 | 1.742x |
| Combined selected natural corpora | 4096 | 0.957x |
| Combined selected natural corpora | 65536 | 0.897x |
| Combined selected natural corpora | 1048576 | 0.890x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.293x |
| ar-wikipedia | 4096 | 0.812x |
| ar-wikipedia | 65536 | 0.770x |
| ar-wikipedia | 1048576 | 0.767x |
| decimal | 0 | 1.492x |
| en-wikitext2 | 256 | 5.471x |
| en-wikitext2 | 4096 | 2.638x |
| en-wikitext2 | 65536 | 1.495x |
| en-wikitext2 | 1048576 | 1.380x |
| ja-wikipedia | 256 | 2.195x |
| ja-wikipedia | 4096 | 1.251x |
| ja-wikipedia | 65536 | 1.124x |
| ja-wikipedia | 1048576 | 1.120x |
| ru-wikipedia | 256 | 1.329x |
| ru-wikipedia | 4096 | 0.816x |
| ru-wikipedia | 65536 | 0.760x |
| ru-wikipedia | 1048576 | 0.767x |
| synthetic | 0 | 3.716x |
| zh-wikipedia | 256 | 2.255x |
| zh-wikipedia | 4096 | 1.294x |
| zh-wikipedia | 65536 | 1.281x |
| zh-wikipedia | 1048576 | 1.236x |
| Combined selected natural corpora | 256 | 1.748x |
| Combined selected natural corpora | 4096 | 0.967x |
| Combined selected natural corpora | 65536 | 0.901x |
| Combined selected natural corpora | 1048576 | 0.895x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 4.474x |
| ar-wikipedia | 4096 | 3.119x |
| ar-wikipedia | 65536 | 3.002x |
| ar-wikipedia | 1048576 | 4.336x |
| decimal | 0 | 2.619x |
| en-wikitext2 | 256 | 45.801x |
| en-wikitext2 | 4096 | 103.514x |
| en-wikitext2 | 65536 | 117.809x |
| en-wikitext2 | 1048576 | 200.627x |
| ja-wikipedia | 256 | 8.344x |
| ja-wikipedia | 4096 | 6.406x |
| ja-wikipedia | 65536 | 6.304x |
| ja-wikipedia | 1048576 | 9.552x |
| ru-wikipedia | 256 | 4.647x |
| ru-wikipedia | 4096 | 3.301x |
| ru-wikipedia | 65536 | 3.188x |
| ru-wikipedia | 1048576 | 4.987x |
| synthetic | 0 | 12.698x |
| zh-wikipedia | 256 | 8.439x |
| zh-wikipedia | 4096 | 6.544x |
| zh-wikipedia | 65536 | 6.545x |
| zh-wikipedia | 1048576 | 9.796x |
| Combined selected natural corpora | 256 | 6.707x |
| Combined selected natural corpora | 4096 | 4.965x |
| Combined selected natural corpora | 65536 | 4.873x |
| Combined selected natural corpora | 1048576 | 7.491x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 4.347x |
| ar-wikipedia | 4096 | 3.112x |
| ar-wikipedia | 65536 | 2.999x |
| ar-wikipedia | 1048576 | 4.338x |
| decimal | 0 | 2.608x |
| en-wikitext2 | 256 | 28.217x |
| en-wikitext2 | 4096 | 91.255x |
| en-wikitext2 | 65536 | 116.715x |
| en-wikitext2 | 1048576 | 201.031x |
| ja-wikipedia | 256 | 7.689x |
| ja-wikipedia | 4096 | 6.387x |
| ja-wikipedia | 65536 | 6.298x |
| ja-wikipedia | 1048576 | 9.522x |
| ru-wikipedia | 256 | 4.530x |
| ru-wikipedia | 4096 | 3.316x |
| ru-wikipedia | 65536 | 3.191x |
| ru-wikipedia | 1048576 | 4.984x |
| synthetic | 0 | 9.707x |
| zh-wikipedia | 256 | 7.786x |
| zh-wikipedia | 4096 | 6.500x |
| zh-wikipedia | 65536 | 6.545x |
| zh-wikipedia | 1048576 | 9.782x |
| Combined selected natural corpora | 256 | 6.333x |
| Combined selected natural corpora | 4096 | 4.957x |
| Combined selected natural corpora | 65536 | 4.871x |
| Combined selected natural corpora | 1048576 | 7.486x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 2690.7 | 12.6 | 10 |
| onig-compile | 2673.3 | 9.3 | 10 |
| pretok-compile | 10.0 | 0.1 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
