# P12 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 21822560 | 137491 | 5317.4 | 48.05 | 26787646 | 188154 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 18348309 | 177950 | 4470.8 | 57.15 | 31859885 | 223781 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 18251011 | 59135 | 4447.1 | 57.45 | 32029732 | 55 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 18300750 | 120156 | 4459.2 | 57.30 | 31942680 | 224362 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 9295967 | 448557 | 2265.1 | 112.80 | 62884795 | 108 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 9278915 | 549552 | 2260.9 | 113.01 | 63000360 | 442509 | 65760.0 | 4106.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 9305292 | 451308 | 2267.4 | 112.69 | 62821782 | 441254 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 20097137 | 142559 | 78199.0 | 52.18 | 29087476 | 12887 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 17683765 | 196303 | 68808.4 | 59.30 | 33057157 | 14646 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 17652759 | 98807 | 68687.8 | 59.40 | 33115219 | 57 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 17729810 | 131042 | 68987.6 | 59.14 | 32971306 | 14608 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 9291489 | 190183 | 36153.7 | 112.85 | 62915106 | 108 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 9297234 | 221299 | 36176.0 | 112.78 | 62876229 | 27858 | 4208.0 | 259.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 9256211 | 157938 | 36016.4 | 113.28 | 63154892 | 27981 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 20050457 | 209513 | 1179438.6 | 52.30 | 29155196 | 948 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 17692765 | 192800 | 1040750.9 | 59.27 | 33040341 | 1074 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 17676212 | 192715 | 1039777.2 | 59.32 | 33071283 | 57 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 17668351 | 123791 | 1039314.7 | 59.35 | 33085997 | 1075 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 9210292 | 161646 | 541781.9 | 113.85 | 63469757 | 109 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 9224885 | 134947 | 542640.3 | 113.67 | 63369352 | 2060 | 368.0 | 19.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 9267552 | 141584 | 545150.1 | 113.14 | 63077608 | 2050 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 21767450 | 341454 | 21767450.0 | 48.17 | 26855465 | 138 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 17608878 | 272927 | 17608877.9 | 59.55 | 33197743 | 170 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 17717269 | 157388 | 17717268.7 | 59.18 | 32994646 | 56 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 17635184 | 116950 | 17635183.9 | 59.46 | 33148222 | 170 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 9127112 | 69759 | 9127112.1 | 114.89 | 64048189 | 110 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 9089460 | 144461 | 9089460.2 | 115.36 | 64313500 | 330 | 112.0 | 3.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 9099725 | 125893 | 9099725.4 | 115.23 | 64240950 | 330 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 21173212 | 303469 | 5168.0 | 49.52 | 49422686 | 197372 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 18161074 | 191561 | 4432.8 | 57.74 | 57619774 | 230108 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 18004392 | 258916 | 4394.5 | 58.24 | 58121207 | 2333 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 18100208 | 760890 | 4417.9 | 57.93 | 57813533 | 230881 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 4445429 | 26389 | 1085.0 | 235.88 | 235396197 | 9448 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 4540973 | 135794 | 1108.4 | 230.91 | 230443364 | 920287 | 70032.0 | 4139.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 4463198 | 83407 | 1089.4 | 234.94 | 234459002 | 936324 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 20029671 | 203467 | 78240.9 | 52.35 | 52244343 | 16875 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 17670331 | 156615 | 69024.7 | 59.34 | 59219999 | 19128 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 17606346 | 117522 | 68774.8 | 59.56 | 59435219 | 2329 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 17568903 | 149799 | 68628.5 | 59.68 | 59561886 | 19239 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 4452222 | 25155 | 17391.5 | 235.52 | 235037043 | 9209 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 4473501 | 19765 | 17474.6 | 234.40 | 233919050 | 75556 | 8256.0 | 279.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 4452606 | 21168 | 17393.0 | 235.50 | 235016771 | 75911 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 19909401 | 82715 | 1244337.6 | 52.67 | 52559943 | 4922 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 17602045 | 114942 | 1100127.8 | 59.57 | 59449739 | 5568 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 17653065 | 169687 | 1103316.5 | 59.40 | 59277923 | 2323 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 17613233 | 147814 | 1100827.0 | 59.53 | 59411978 | 5564 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 4480353 | 21348 | 280022.1 | 234.04 | 233561262 | 9151 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 4469634 | 15069 | 279352.1 | 234.60 | 234121420 | 21926 | 3264.0 | 27.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 4470341 | 10512 | 279396.3 | 234.56 | 234084392 | 21922 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 22483846 | 151312 | 22483846.0 | 46.64 | 46541726 | 3692 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 17573408 | 154784 | 17573408.1 | 59.67 | 59546617 | 4723 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 17679472 | 156019 | 17679472.0 | 59.31 | 59189381 | 2319 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 17683647 | 166574 | 17683646.6 | 59.30 | 59175408 | 4694 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 4478570 | 21270 | 4478569.6 | 234.13 | 233654292 | 9155 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 4472730 | 5888 | 4472729.7 | 234.44 | 233959364 | 18557 | 4080.0 | 8.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 4474706 | 7322 | 4474706.5 | 234.33 | 233856010 | 18549 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 16981513 | 54315 | 4128.7 | 61.75 | 22334876 | 3655269 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 8212744 | 34140 | 1996.8 | 127.68 | 46181884 | 7558010 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 7659291 | 39354 | 1862.2 | 136.90 | 49518941 | 4254310 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 7667956 | 10234 | 1864.3 | 136.75 | 49462987 | 8094987 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 3818167 | 21026 | 928.3 | 274.63 | 99335617 | 8534199 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 4447883 | 41606 | 1081.4 | 235.75 | 85272025 | 13955403 | 2772976.0 | 21650.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 3835336 | 54024 | 932.5 | 273.40 | 98890935 | 16184239 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 14323888 | 56671 | 55735.0 | 73.20 | 26478844 | 4106776 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 6895433 | 31424 | 26830.5 | 152.07 | 55004523 | 8531009 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 6782773 | 45604 | 26392.1 | 154.59 | 55918136 | 4350728 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 6782345 | 19326 | 26390.4 | 154.60 | 55921660 | 8673254 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 3719975 | 21166 | 14474.6 | 281.88 | 101957683 | 7932850 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 4046948 | 29816 | 15746.9 | 259.10 | 93720006 | 14535645 | 2598384.0 | 2368.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 3735455 | 17191 | 14534.8 | 280.71 | 101535157 | 15747747 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 13997198 | 144797 | 823364.6 | 74.91 | 27096851 | 4187695 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 6774501 | 42838 | 398500.1 | 154.78 | 55986408 | 8652445 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 6746082 | 72918 | 396828.4 | 155.43 | 56222259 | 4345485 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 6728146 | 59539 | 395773.3 | 155.85 | 56372139 | 8712058 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 3767888 | 21793 | 221640.5 | 278.29 | 100661166 | 7780221 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 4054912 | 53152 | 238524.2 | 258.59 | 93535936 | 14455554 | 3493648.0 | 249.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 3782867 | 113179 | 222521.6 | 277.19 | 100262585 | 15495127 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 16805153 | 692928 | 16805152.8 | 62.40 | 22569268 | 3487085 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 6716032 | 232481 | 6716031.9 | 156.13 | 56473823 | 8725539 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 6668519 | 35613 | 6668518.6 | 157.24 | 56876200 | 4393929 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 6716651 | 57367 | 6716651.4 | 156.12 | 56468615 | 8724734 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 3731159 | 38918 | 3731158.8 | 281.03 | 101652067 | 7853056 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 4038261 | 85416 | 4038261.2 | 259.66 | 93921613 | 14511444 | 4363504.0 | 26.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 3754909 | 44396 | 3754908.9 | 279.25 | 101009108 | 15606504 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 20742719 | 261728 | 5054.3 | 50.55 | 27713676 | 197853 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 17271631 | 258300 | 4208.5 | 60.71 | 33283306 | 237615 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 17196160 | 375119 | 4190.1 | 60.98 | 33429381 | 0 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 17357124 | 782854 | 4229.3 | 60.41 | 33119370 | 236445 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 9204648 | 178533 | 2242.8 | 113.92 | 62452901 | 0 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 9517496 | 273248 | 2319.1 | 110.17 | 60400022 | 431206 | 65664.0 | 4104.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 9377414 | 265034 | 2284.9 | 111.82 | 61302294 | 437647 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 19077513 | 240859 | 74231.6 | 54.96 | 30132701 | 13471 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 16729478 | 287247 | 65095.2 | 62.68 | 34361921 | 15362 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 16740894 | 157546 | 65139.7 | 62.64 | 34338488 | 0 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 16720753 | 353272 | 65061.3 | 62.71 | 34379852 | 15370 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 9190331 | 233027 | 35760.0 | 114.10 | 62550192 | 0 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 9422325 | 277438 | 36662.7 | 111.29 | 61010100 | 27276 | 4112.0 | 257.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 9381104 | 260409 | 36502.3 | 111.78 | 61278184 | 27395 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 18643597 | 139611 | 1096682.2 | 56.24 | 30834017 | 912 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 16470037 | 299052 | 968825.7 | 63.67 | 34903200 | 1032 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 16568856 | 116429 | 974638.6 | 63.29 | 34695034 | 0 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 16486677 | 247351 | 969804.6 | 63.60 | 34867972 | 1031 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 9206824 | 52479 | 541577.9 | 113.89 | 62438145 | 0 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 9271702 | 132570 | 545394.2 | 113.09 | 62001237 | 1834 | 272.0 | 17.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 9287085 | 226912 | 546299.1 | 112.91 | 61898538 | 1830 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 20351126 | 158437 | 20351125.8 | 51.52 | 28246938 | 49 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 16487839 | 117878 | 16487838.7 | 63.60 | 34865516 | 61 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 16433555 | 359146 | 16433555.1 | 63.81 | 34980684 | 0 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 16527111 | 166750 | 16527110.9 | 63.45 | 34782667 | 61 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 9138154 | 258292 | 9138154.2 | 114.75 | 62907343 | 0 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 9193712 | 143716 | 9193711.7 | 114.05 | 62527194 | 109 | 16.0 | 1.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 9171491 | 217598 | 9171491.3 | 114.33 | 62678684 | 109 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 96356590 | 1036571 | 1001.8 | 27.79 | 14333716 | 1980300 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 56871444 | 447969 | 591.3 | 47.08 | 24285439 | 3355199 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 54753996 | 249215 | 569.3 | 48.90 | 25224607 | 1357581 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 54923300 | 247393 | 571.0 | 48.75 | 25146850 | 3474209 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 17074074 | 237098 | 177.5 | 156.81 | 80891530 | 4353560 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 19440358 | 260028 | 202.1 | 137.72 | 71045399 | 9815406 | 5524112.0 | 161931.0 | 10 |
| synthetic | 0 | pretok-reused | 17193075 | 287536 | 178.7 | 155.72 | 80331645 | 11098364 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 19638424 | 281357 | 4774.7 | 53.39 | 19869364 | 3802902 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 9672643 | 86897 | 2351.7 | 108.41 | 40340889 | 7721054 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 9028910 | 70874 | 2195.2 | 116.14 | 43217064 | 4287339 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 9014758 | 76161 | 2191.8 | 116.32 | 43284912 | 8284527 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 3313827 | 35467 | 805.7 | 316.42 | 117749957 | 11681358 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 3979349 | 51565 | 967.5 | 263.50 | 98056993 | 18767642 | 3344880.0 | 22879.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 3348599 | 33339 | 814.2 | 313.14 | 116527225 | 22302757 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 16728001 | 224703 | 65089.5 | 62.68 | 23326338 | 4280249 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 8318700 | 30467 | 32368.5 | 126.05 | 46906729 | 8607114 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 8130948 | 51212 | 31637.9 | 128.96 | 47989854 | 4414737 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 8143819 | 46688 | 31688.0 | 128.76 | 47914006 | 8791944 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 3240964 | 36274 | 12610.8 | 323.54 | 120397219 | 11075718 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 3611471 | 16466 | 14052.4 | 290.35 | 108045445 | 19825716 | 3372784.0 | 2464.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 3246244 | 21377 | 12631.3 | 323.01 | 120201363 | 22056257 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 16298785 | 370946 | 958752.1 | 64.33 | 23940619 | 4381001 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 8078102 | 100851 | 475182.5 | 129.80 | 48303798 | 8839329 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 8020531 | 51141 | 471795.9 | 130.74 | 48650521 | 4452074 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 8002168 | 64858 | 470715.7 | 131.04 | 48762163 | 8923207 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 3208407 | 21813 | 188729.8 | 326.82 | 121618930 | 11129511 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 3505709 | 39986 | 206218.2 | 299.10 | 111305029 | 20368207 | 4255600.0 | 260.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 3224020 | 19165 | 189648.2 | 325.24 | 121029950 | 22147814 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 20197188 | 563772 | 20197187.5 | 51.92 | 19319670 | 3534898 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 8117162 | 134055 | 8117161.6 | 129.18 | 48071360 | 8795562 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 8049375 | 36432 | 8049375.3 | 130.27 | 48476184 | 4434878 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 8080649 | 55716 | 8080648.9 | 129.76 | 48288573 | 8835305 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 3236084 | 27271 | 3236084.3 | 324.03 | 120578751 | 11031233 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 3563215 | 64587 | 3563215.0 | 294.28 | 109508689 | 20036680 | 5559536.0 | 27.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 3229453 | 20297 | 3229453.2 | 324.69 | 120826336 | 22107458 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.963x |
| ar-wikipedia | 4096 | 1.900x |
| ar-wikipedia | 65536 | 1.919x |
| ar-wikipedia | 1048576 | 1.941x |
| en-wikitext2 | 256 | 4.050x |
| en-wikitext2 | 4096 | 3.955x |
| en-wikitext2 | 65536 | 3.940x |
| en-wikitext2 | 1048576 | 3.948x |
| ja-wikipedia | 256 | 2.006x |
| ja-wikipedia | 4096 | 1.823x |
| ja-wikipedia | 65536 | 1.790x |
| ja-wikipedia | 1048576 | 1.787x |
| ru-wikipedia | 256 | 1.868x |
| ru-wikipedia | 4096 | 1.822x |
| ru-wikipedia | 65536 | 1.800x |
| ru-wikipedia | 1048576 | 1.798x |
| synthetic | 0 | 3.207x |
| zh-wikipedia | 256 | 2.725x |
| zh-wikipedia | 4096 | 2.509x |
| zh-wikipedia | 65536 | 2.500x |
| zh-wikipedia | 1048576 | 2.487x |
| Combined selected natural corpora | 256 | 2.332x |
| Combined selected natural corpora | 4096 | 2.238x |
| Combined selected natural corpora | 65536 | 2.232x |
| Combined selected natural corpora | 1048576 | 2.240x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.967x |
| ar-wikipedia | 4096 | 1.915x |
| ar-wikipedia | 65536 | 1.906x |
| ar-wikipedia | 1048576 | 1.938x |
| en-wikitext2 | 256 | 4.055x |
| en-wikitext2 | 4096 | 3.946x |
| en-wikitext2 | 65536 | 3.940x |
| en-wikitext2 | 1048576 | 3.952x |
| ja-wikipedia | 256 | 1.999x |
| ja-wikipedia | 4096 | 1.816x |
| ja-wikipedia | 65536 | 1.779x |
| ja-wikipedia | 1048576 | 1.789x |
| ru-wikipedia | 256 | 1.851x |
| ru-wikipedia | 4096 | 1.782x |
| ru-wikipedia | 65536 | 1.775x |
| ru-wikipedia | 1048576 | 1.802x |
| synthetic | 0 | 3.195x |
| zh-wikipedia | 256 | 2.692x |
| zh-wikipedia | 4096 | 2.509x |
| zh-wikipedia | 65536 | 2.482x |
| zh-wikipedia | 1048576 | 2.502x |
| Combined selected natural corpora | 256 | 2.322x |
| Combined selected natural corpora | 4096 | 2.226x |
| Combined selected natural corpora | 65536 | 2.214x |
| Combined selected natural corpora | 1048576 | 2.242x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.977x |
| ar-wikipedia | 4096 | 1.902x |
| ar-wikipedia | 65536 | 1.918x |
| ar-wikipedia | 1048576 | 1.937x |
| en-wikitext2 | 256 | 3.999x |
| en-wikitext2 | 4096 | 3.950x |
| en-wikitext2 | 65536 | 3.938x |
| en-wikitext2 | 1048576 | 3.929x |
| ja-wikipedia | 256 | 1.846x |
| ja-wikipedia | 4096 | 1.704x |
| ja-wikipedia | 65536 | 1.671x |
| ja-wikipedia | 1048576 | 1.663x |
| ru-wikipedia | 256 | 1.815x |
| ru-wikipedia | 4096 | 1.776x |
| ru-wikipedia | 65536 | 1.776x |
| ru-wikipedia | 1048576 | 1.793x |
| synthetic | 0 | 2.925x |
| zh-wikipedia | 256 | 2.431x |
| zh-wikipedia | 4096 | 2.303x |
| zh-wikipedia | 65536 | 2.304x |
| zh-wikipedia | 1048576 | 2.278x |
| Combined selected natural corpora | 256 | 2.256x |
| Combined selected natural corpora | 4096 | 2.181x |
| Combined selected natural corpora | 65536 | 2.182x |
| Combined selected natural corpora | 1048576 | 2.191x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.345x |
| ar-wikipedia | 4096 | 2.171x |
| ar-wikipedia | 65536 | 2.164x |
| ar-wikipedia | 1048576 | 2.392x |
| en-wikitext2 | 256 | 4.744x |
| en-wikitext2 | 4096 | 4.498x |
| en-wikitext2 | 65536 | 4.454x |
| en-wikitext2 | 1048576 | 5.025x |
| ja-wikipedia | 256 | 4.428x |
| ja-wikipedia | 4096 | 3.835x |
| ja-wikipedia | 65536 | 3.700x |
| ja-wikipedia | 1048576 | 4.476x |
| ru-wikipedia | 256 | 2.212x |
| ru-wikipedia | 4096 | 2.034x |
| ru-wikipedia | 65536 | 2.007x |
| ru-wikipedia | 1048576 | 2.219x |
| synthetic | 0 | 5.604x |
| zh-wikipedia | 256 | 5.865x |
| zh-wikipedia | 4096 | 5.153x |
| zh-wikipedia | 65536 | 5.055x |
| zh-wikipedia | 1048576 | 6.254x |
| Combined selected natural corpora | 256 | 3.309x |
| Combined selected natural corpora | 4096 | 3.001x |
| Combined selected natural corpora | 65536 | 2.960x |
| Combined selected natural corpora | 1048576 | 3.418x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.352x |
| ar-wikipedia | 4096 | 2.162x |
| ar-wikipedia | 65536 | 2.174x |
| ar-wikipedia | 1048576 | 2.395x |
| en-wikitext2 | 256 | 4.663x |
| en-wikitext2 | 4096 | 4.477x |
| en-wikitext2 | 65536 | 4.454x |
| en-wikitext2 | 1048576 | 5.027x |
| ja-wikipedia | 256 | 3.818x |
| ja-wikipedia | 4096 | 3.539x |
| ja-wikipedia | 65536 | 3.452x |
| ja-wikipedia | 1048576 | 4.161x |
| ru-wikipedia | 256 | 2.179x |
| ru-wikipedia | 4096 | 2.025x |
| ru-wikipedia | 65536 | 2.011x |
| ru-wikipedia | 1048576 | 2.214x |
| synthetic | 0 | 4.957x |
| zh-wikipedia | 256 | 4.935x |
| zh-wikipedia | 4096 | 4.632x |
| zh-wikipedia | 65536 | 4.649x |
| zh-wikipedia | 1048576 | 5.668x |
| Combined selected natural corpora | 256 | 3.159x |
| Combined selected natural corpora | 4096 | 2.926x |
| Combined selected natural corpora | 65536 | 2.912x |
| Combined selected natural corpora | 1048576 | 3.347x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 2407.1 | 35.6 | 10 |
| onig-compile | 2373.5 | 43.1 | 10 |
| pretok-compile | 12.5 | 0.1 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
