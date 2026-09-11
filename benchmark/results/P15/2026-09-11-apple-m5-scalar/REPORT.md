# P15 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 17408017 | 450974 | 4241.7 | 60.24 | 33580793 | 236558 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 13861652 | 318066 | 3377.6 | 75.65 | 42172102 | 297079 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 13887686 | 281936 | 3383.9 | 75.50 | 42093046 | 504 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 13853841 | 251070 | 3375.7 | 75.69 | 42195878 | 297246 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 7729886 | 136770 | 1883.5 | 135.65 | 75625305 | 906 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 7795324 | 173845 | 1899.4 | 134.51 | 74990468 | 528265 | 66400.0 | 4116.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 7771022 | 90007 | 1893.5 | 134.93 | 75224982 | 529917 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 15469479 | 182091 | 60192.5 | 67.78 | 37788927 | 17518 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 13005083 | 156408 | 50603.4 | 80.63 | 44949733 | 20838 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 13111455 | 243779 | 51017.3 | 79.97 | 44585061 | 534 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 13092919 | 243708 | 50945.2 | 80.09 | 44648178 | 20698 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 7522299 | 267719 | 29269.6 | 139.40 | 77712284 | 931 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 7631647 | 176133 | 29695.1 | 137.40 | 76598796 | 35510 | 4848.0 | 269.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 7541854 | 243951 | 29345.7 | 139.03 | 77510780 | 35933 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 15307775 | 277725 | 900457.4 | 68.50 | 38188110 | 2025 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 12962723 | 306694 | 762513.1 | 80.89 | 45096619 | 2391 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 13054250 | 59287 | 767897.1 | 80.32 | 44780436 | 536 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 12950174 | 291923 | 761774.9 | 80.97 | 45140321 | 2394 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 7603053 | 188373 | 447238.4 | 137.92 | 76886873 | 921 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 7576862 | 140548 | 445697.8 | 138.39 | 77152652 | 4091 | 1040.0 | 28.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 7644766 | 204319 | 449692.1 | 137.16 | 76467346 | 4055 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 17283088 | 172703 | 17283088.4 | 60.67 | 33823527 | 868 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 13139705 | 241513 | 13139705.4 | 79.80 | 44489201 | 1142 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 13179209 | 117928 | 13179208.5 | 79.56 | 44355850 | 531 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 13099174 | 179789 | 13099174.2 | 80.05 | 44626859 | 1145 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 7434177 | 238541 | 7434177.4 | 141.05 | 78633448 | 942 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 7575605 | 208521 | 7575604.5 | 138.41 | 77165459 | 1980 | 496.0 | 5.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 7606267 | 288563 | 7606266.5 | 137.86 | 76854394 | 1972 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 16773976 | 263202 | 4094.2 | 62.51 | 62384552 | 262669 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 13912034 | 74848 | 3395.7 | 75.37 | 75218119 | 316704 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 13905564 | 113261 | 3394.1 | 75.41 | 75253113 | 11147 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 13932095 | 137547 | 3400.6 | 75.26 | 75109810 | 316248 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 3180051 | 97558 | 776.2 | 329.74 | 329062955 | 48741 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 3269440 | 58925 | 798.0 | 320.72 | 320066097 | 1347631 | 81424.0 | 4364.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 3231107 | 52070 | 788.7 | 324.53 | 323863268 | 1363619 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 16091274 | 215833 | 62856.5 | 65.16 | 65031333 | 35174 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 13761367 | 344856 | 53755.3 | 76.20 | 76041649 | 41130 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 13567120 | 431861 | 52996.6 | 77.29 | 77130369 | 11425 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 13758356 | 373150 | 53743.6 | 76.21 | 76058286 | 41139 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 3195987 | 54911 | 12484.3 | 328.09 | 327422181 | 48498 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 3125636 | 103702 | 12209.5 | 335.48 | 334791675 | 181083 | 19232.0 | 458.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 3142944 | 124106 | 12277.1 | 333.63 | 332947999 | 180086 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 16161461 | 258047 | 1010091.3 | 64.88 | 64748910 | 20171 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 13949421 | 69501 | 871838.8 | 75.17 | 75016517 | 23370 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 13755082 | 307029 | 859692.6 | 76.23 | 76076393 | 11269 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 13896943 | 328576 | 868559.0 | 75.45 | 75299796 | 23458 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 3169528 | 81289 | 198095.5 | 330.83 | 330155418 | 48903 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 3137165 | 150744 | 196072.8 | 334.24 | 333561394 | 103915 | 15904.0 | 86.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 3135382 | 132921 | 195961.4 | 334.43 | 333751022 | 103975 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 17491419 | 325399 | 17491418.8 | 59.95 | 59825736 | 17780 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 13824823 | 344629 | 13824823.2 | 75.85 | 75692614 | 22496 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 13666564 | 394574 | 13666564.2 | 76.73 | 76569135 | 11342 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 13685617 | 381963 | 13685616.5 | 76.62 | 76462540 | 22725 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 3115954 | 52872 | 3115953.8 | 336.52 | 335832001 | 49744 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 3079161 | 58221 | 3079160.6 | 340.54 | 339844889 | 101002 | 16368.0 | 10.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 3089308 | 102733 | 3089307.9 | 339.42 | 338728617 | 100670 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 19299983 | 135455 | 4692.4 | 54.33 | 19651831 | 3488241 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 10090465 | 44124 | 2453.3 | 103.92 | 37587961 | 6671942 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 9617323 | 80384 | 2338.3 | 109.03 | 39437169 | 5379667 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 9636681 | 39617 | 2343.0 | 108.81 | 39357947 | 6986119 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 3809299 | 20846 | 926.2 | 275.27 | 99566869 | 13582026 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 4492677 | 44542 | 1092.3 | 233.40 | 84421818 | 14985051 | 2970736.0 | 22015.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 3864119 | 46402 | 939.5 | 271.36 | 98154337 | 17422602 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 16578591 | 45817 | 64508.1 | 63.25 | 22877699 | 3866794 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 8834571 | 35374 | 34375.8 | 118.69 | 42931342 | 7256266 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 8696231 | 46029 | 33837.5 | 120.58 | 43614299 | 5594378 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 8705680 | 57236 | 33874.2 | 120.45 | 43566961 | 7363699 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 3702256 | 24481 | 14405.7 | 283.23 | 102445634 | 13140635 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 4038643 | 28095 | 15714.6 | 259.64 | 93912737 | 15873154 | 2897392.0 | 2405.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 3736450 | 20184 | 14538.7 | 280.63 | 101508122 | 17156928 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 16103839 | 156329 | 947284.6 | 65.11 | 23552149 | 3967998 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 8640589 | 107578 | 508269.9 | 121.35 | 43895155 | 7395329 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 8606864 | 64869 | 506286.1 | 121.83 | 44067155 | 5629809 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 8606094 | 82072 | 506240.8 | 121.84 | 44071098 | 7424971 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 3680132 | 30080 | 216478.3 | 284.93 | 103061524 | 13166648 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 4028840 | 42602 | 236990.6 | 260.27 | 94141233 | 15860643 | 4067088.0 | 256.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 3724787 | 30884 | 219105.1 | 281.51 | 101825966 | 17155345 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 19528439 | 309714 | 19528439.1 | 53.69 | 19421931 | 3271383 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 8598732 | 133308 | 8598732.0 | 121.95 | 44108829 | 7429584 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 8662757 | 98541 | 8662757.2 | 121.04 | 43782827 | 5591869 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 8638577 | 121130 | 8638577.2 | 121.38 | 43905378 | 7395315 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 3686684 | 16532 | 3686684.3 | 284.42 | 102878351 | 13139449 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 4069363 | 34546 | 4069363.5 | 257.68 | 93203766 | 15699015 | 5559536.0 | 27.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 3742620 | 28516 | 3742620.2 | 280.17 | 101340767 | 17069592 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 17525989 | 395937 | 4270.5 | 59.83 | 32800260 | 379494 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 13777399 | 211925 | 3357.1 | 76.11 | 41724640 | 482747 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 13758939 | 208557 | 3352.6 | 76.21 | 41780621 | 93176 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 13736297 | 86198 | 3347.1 | 76.34 | 41849489 | 484192 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 7608636 | 193777 | 1854.0 | 137.81 | 75553223 | 168493 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 7618208 | 295449 | 1856.3 | 137.64 | 75458301 | 873040 | 192384.0 | 6182.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 7535451 | 235137 | 1836.1 | 139.15 | 76287006 | 882628 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 15581189 | 242119 | 60627.2 | 67.30 | 36894297 | 180987 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 13036329 | 173769 | 50725.0 | 80.43 | 44096541 | 216319 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 12954463 | 101211 | 50406.5 | 80.94 | 44375210 | 98962 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 12988372 | 125989 | 50538.4 | 80.73 | 44259358 | 217117 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 7520197 | 203797 | 29261.5 | 139.43 | 76441752 | 170474 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 7503106 | 138398 | 29195.0 | 139.75 | 76615870 | 375844 | 125936.0 | 1142.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 7443581 | 118782 | 28963.3 | 140.87 | 77228556 | 378850 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 15389585 | 125322 | 905269.7 | 68.14 | 37353638 | 167711 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 12815714 | 105776 | 753865.5 | 81.82 | 44855637 | 201393 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 12927376 | 88610 | 760433.9 | 81.11 | 44468189 | 99169 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 12837168 | 147653 | 755127.5 | 81.68 | 44780672 | 201057 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 7462316 | 274010 | 438959.7 | 140.52 | 77034666 | 171797 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 7455434 | 186983 | 438554.9 | 140.65 | 77105774 | 346190 | 106256.0 | 139.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 7431832 | 217565 | 437166.6 | 141.09 | 77350641 | 347290 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 16783908 | 310187 | 16783908.3 | 62.48 | 34250485 | 152825 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 12775391 | 242145 | 12775391.2 | 82.08 | 44997213 | 200777 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 12808916 | 139370 | 12808916.1 | 81.86 | 44879441 | 100087 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 12788350 | 130978 | 12788350.0 | 81.99 | 44951616 | 200573 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 7524919 | 149287 | 7524919.5 | 139.35 | 76393774 | 170367 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 7576082 | 119957 | 7576082.4 | 138.41 | 75877871 | 338565 | 177392.0 | 15.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 7509882 | 204878 | 7509882.5 | 139.63 | 76546737 | 341550 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 88068326 | 952182 | 915.6 | 30.40 | 15682687 | 2025972 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 47153070 | 267221 | 490.2 | 56.78 | 29290733 | 3783932 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 44954856 | 191635 | 467.4 | 59.56 | 30722999 | 1478928 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 45243370 | 274484 | 470.4 | 59.18 | 30527081 | 3943650 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 13958843 | 116853 | 145.1 | 191.80 | 98944302 | 4762931 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 16206335 | 189065 | 168.5 | 165.20 | 85222721 | 11009522 | 4895312.0 | 160531.0 | 10 |
| synthetic | 0 | pretok-reused | 14009422 | 221623 | 145.6 | 191.11 | 98587076 | 12736000 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 22586245 | 158176 | 5491.4 | 46.43 | 17276134 | 3550745 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 12049277 | 91999 | 2929.6 | 87.02 | 32383935 | 6655835 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 11569672 | 206019 | 2813.0 | 90.63 | 33726366 | 5749515 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 11630429 | 63015 | 2827.7 | 90.16 | 33550182 | 6895533 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 3301392 | 49087 | 802.7 | 317.62 | 118193477 | 20149077 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 4038378 | 29339 | 981.9 | 259.65 | 96623689 | 19858962 | 3538992.0 | 23206.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 3338143 | 71893 | 811.6 | 314.12 | 116892229 | 24024733 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 19766448 | 401947 | 76912.2 | 53.05 | 19740674 | 3903079 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 10783798 | 142943 | 41960.3 | 97.24 | 36184191 | 7154251 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 10657606 | 285949 | 41469.3 | 98.39 | 36612631 | 5976014 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 10617850 | 206378 | 41314.6 | 98.76 | 36749719 | 7266066 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 3240101 | 78908 | 12607.4 | 323.62 | 120429257 | 19656792 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 3671030 | 80686 | 14284.2 | 285.63 | 106292503 | 21015898 | 3673072.0 | 2500.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 3282840 | 46736 | 12773.7 | 319.41 | 118861405 | 23500991 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 20080795 | 195209 | 1181223.2 | 52.22 | 19431651 | 3832318 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 10955076 | 147541 | 644416.2 | 95.72 | 35618466 | 7024689 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 10942411 | 48664 | 643671.2 | 95.83 | 35659692 | 5803200 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 10909550 | 62287 | 641738.3 | 96.12 | 35767102 | 7054003 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 3285549 | 29193 | 193267.6 | 319.15 | 118763409 | 19327364 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 3733461 | 29667 | 219615.3 | 280.86 | 104515090 | 20612510 | 4894576.0 | 266.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 3335907 | 31571 | 196229.8 | 314.33 | 116970597 | 23068990 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 23565869 | 1061204 | 23565869.3 | 44.50 | 16557972 | 3265146 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 10609275 | 103520 | 10609274.9 | 98.84 | 36779422 | 7252711 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 10659816 | 89747 | 10659816.5 | 98.37 | 36605039 | 5956106 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 10609880 | 116970 | 10609879.7 | 98.83 | 36777326 | 7252297 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 3227999 | 17814 | 3227998.8 | 324.84 | 120880776 | 19668843 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 3658008 | 43718 | 3658008.2 | 286.65 | 106670892 | 21034945 | 7058672.0 | 28.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 3256289 | 44904 | 3256288.8 | 322.02 | 119830587 | 23629968 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.797x |
| ar-wikipedia | 4096 | 1.743x |
| ar-wikipedia | 65536 | 1.717x |
| ar-wikipedia | 1048576 | 1.773x |
| en-wikitext2 | 256 | 4.373x |
| en-wikitext2 | 4096 | 4.245x |
| en-wikitext2 | 65536 | 4.340x |
| en-wikitext2 | 1048576 | 4.386x |
| ja-wikipedia | 256 | 2.525x |
| ja-wikipedia | 4096 | 2.349x |
| ja-wikipedia | 65536 | 2.339x |
| ja-wikipedia | 1048576 | 2.350x |
| ru-wikipedia | 256 | 1.808x |
| ru-wikipedia | 4096 | 1.723x |
| ru-wikipedia | 65536 | 1.732x |
| ru-wikipedia | 1048576 | 1.702x |
| synthetic | 0 | 3.221x |
| zh-wikipedia | 256 | 3.504x |
| zh-wikipedia | 4096 | 3.289x |
| zh-wikipedia | 65536 | 3.330x |
| zh-wikipedia | 1048576 | 3.302x |
| Combined selected natural corpora | 256 | 2.448x |
| Combined selected natural corpora | 4096 | 2.343x |
| Combined selected natural corpora | 65536 | 2.353x |
| Combined selected natural corpora | 1048576 | 2.360x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.783x |
| ar-wikipedia | 4096 | 1.736x |
| ar-wikipedia | 65536 | 1.694x |
| ar-wikipedia | 1048576 | 1.722x |
| en-wikitext2 | 256 | 4.312x |
| en-wikitext2 | 4096 | 4.378x |
| en-wikitext2 | 65536 | 4.432x |
| en-wikitext2 | 1048576 | 4.430x |
| ja-wikipedia | 256 | 2.494x |
| ja-wikipedia | 4096 | 2.330x |
| ja-wikipedia | 65536 | 2.310x |
| ja-wikipedia | 1048576 | 2.308x |
| ru-wikipedia | 256 | 1.823x |
| ru-wikipedia | 4096 | 1.745x |
| ru-wikipedia | 65536 | 1.727x |
| ru-wikipedia | 1048576 | 1.703x |
| synthetic | 0 | 3.229x |
| zh-wikipedia | 256 | 3.484x |
| zh-wikipedia | 4096 | 3.234x |
| zh-wikipedia | 65536 | 3.270x |
| zh-wikipedia | 1048576 | 3.258x |
| Combined selected natural corpora | 256 | 2.439x |
| Combined selected natural corpora | 4096 | 2.353x |
| Combined selected natural corpora | 65536 | 2.342x |
| Combined selected natural corpora | 1048576 | 2.334x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.778x |
| ar-wikipedia | 4096 | 1.704x |
| ar-wikipedia | 65536 | 1.711x |
| ar-wikipedia | 1048576 | 1.734x |
| en-wikitext2 | 256 | 4.255x |
| en-wikitext2 | 4096 | 4.403x |
| en-wikitext2 | 65536 | 4.447x |
| en-wikitext2 | 1048576 | 4.490x |
| ja-wikipedia | 256 | 2.246x |
| ja-wikipedia | 4096 | 2.188x |
| ja-wikipedia | 65536 | 2.145x |
| ja-wikipedia | 1048576 | 2.113x |
| ru-wikipedia | 256 | 1.808x |
| ru-wikipedia | 4096 | 1.737x |
| ru-wikipedia | 65536 | 1.719x |
| ru-wikipedia | 1048576 | 1.686x |
| synthetic | 0 | 2.910x |
| zh-wikipedia | 256 | 2.984x |
| zh-wikipedia | 4096 | 2.938x |
| zh-wikipedia | 65536 | 2.934x |
| zh-wikipedia | 1048576 | 2.900x |
| Combined selected natural corpora | 256 | 2.340x |
| Combined selected natural corpora | 4096 | 2.288x |
| Combined selected natural corpora | 65536 | 2.288x |
| Combined selected natural corpora | 1048576 | 2.271x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.240x |
| ar-wikipedia | 4096 | 2.051x |
| ar-wikipedia | 65536 | 2.002x |
| ar-wikipedia | 1048576 | 2.272x |
| en-wikitext2 | 256 | 5.191x |
| en-wikitext2 | 4096 | 5.120x |
| en-wikitext2 | 65536 | 5.155x |
| en-wikitext2 | 1048576 | 5.662x |
| ja-wikipedia | 256 | 4.995x |
| ja-wikipedia | 4096 | 4.437x |
| ja-wikipedia | 65536 | 4.323x |
| ja-wikipedia | 1048576 | 5.218x |
| ru-wikipedia | 256 | 2.326x |
| ru-wikipedia | 4096 | 2.093x |
| ru-wikipedia | 65536 | 2.071x |
| ru-wikipedia | 1048576 | 2.235x |
| synthetic | 0 | 6.286x |
| zh-wikipedia | 256 | 6.766x |
| zh-wikipedia | 4096 | 6.021x |
| zh-wikipedia | 65536 | 6.020x |
| zh-wikipedia | 1048576 | 7.237x |
| Combined selected natural corpora | 256 | 3.636x |
| Combined selected natural corpora | 4096 | 3.320x |
| Combined selected natural corpora | 65536 | 3.286x |
| Combined selected natural corpora | 1048576 | 3.755x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.233x |
| ar-wikipedia | 4096 | 2.027x |
| ar-wikipedia | 65536 | 2.020x |
| ar-wikipedia | 1048576 | 2.281x |
| en-wikitext2 | 256 | 5.131x |
| en-wikitext2 | 4096 | 5.148x |
| en-wikitext2 | 65536 | 5.152x |
| en-wikitext2 | 1048576 | 5.681x |
| ja-wikipedia | 256 | 4.296x |
| ja-wikipedia | 4096 | 4.105x |
| ja-wikipedia | 65536 | 3.997x |
| ja-wikipedia | 1048576 | 4.799x |
| ru-wikipedia | 256 | 2.301x |
| ru-wikipedia | 4096 | 2.077x |
| ru-wikipedia | 65536 | 2.064x |
| ru-wikipedia | 1048576 | 2.215x |
| synthetic | 0 | 5.434x |
| zh-wikipedia | 256 | 5.593x |
| zh-wikipedia | 4096 | 5.384x |
| zh-wikipedia | 65536 | 5.379x |
| zh-wikipedia | 1048576 | 6.442x |
| Combined selected natural corpora | 256 | 3.439x |
| Combined selected natural corpora | 4096 | 3.215x |
| Combined selected natural corpora | 65536 | 3.202x |
| Combined selected natural corpora | 1048576 | 3.646x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 1611.4 | 11.1 | 10 |
| onig-compile | 1584.1 | 12.2 | 10 |
| pretok-compile | 11.0 | 0.3 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
