# P18 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 15722811 | 187608 | 3831.1 | 66.69 | 37180057 | 1258426 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 10246623 | 47351 | 2496.7 | 102.33 | 57050502 | 1930978 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 10075640 | 108609 | 2455.1 | 104.07 | 58018645 | 803621 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 10069666 | 93180 | 2453.6 | 104.13 | 58053066 | 1964911 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 5332868 | 31657 | 1299.4 | 196.63 | 109617378 | 1518320 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 5604290 | 20884 | 1365.6 | 187.10 | 104308493 | 3530510 | 828064.0 | 13279.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 5336293 | 24005 | 1300.3 | 196.50 | 109547012 | 3707817 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 13528733 | 86047 | 52641.0 | 77.51 | 43209887 | 1182446 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 9589397 | 73635 | 37312.8 | 109.35 | 60960559 | 1668197 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 9548605 | 100082 | 37154.1 | 109.81 | 61220985 | 846721 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 9545601 | 118437 | 37142.4 | 109.85 | 61240253 | 1675851 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 5245136 | 50639 | 20409.1 | 199.91 | 111450879 | 1541428 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 5386073 | 18263 | 20957.5 | 194.68 | 108534539 | 2970067 | 709136.0 | 1850.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 5262101 | 25556 | 20475.1 | 199.27 | 111091556 | 3040040 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 13218542 | 145112 | 777561.3 | 79.33 | 44223863 | 1192189 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 9507392 | 81058 | 559258.3 | 110.29 | 61486371 | 1657552 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 9441300 | 55624 | 555370.6 | 111.06 | 61916793 | 856238 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 9499910 | 87267 | 558818.2 | 110.38 | 61534796 | 1658858 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 5228266 | 83892 | 307545.1 | 200.56 | 111810492 | 1546211 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 5349327 | 93532 | 314666.3 | 196.02 | 109280108 | 2945978 | 760336.0 | 190.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 5221190 | 88844 | 307128.8 | 200.83 | 111962022 | 3018277 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 15419254 | 157413 | 15419254.4 | 68.00 | 37912015 | 1021061 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 9579521 | 77681 | 9579521.2 | 109.46 | 61023405 | 1643506 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 9575268 | 160293 | 9575268.5 | 109.51 | 61050507 | 844258 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 9568358 | 144657 | 9568357.7 | 109.59 | 61094601 | 1645423 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 5273461 | 101223 | 5273461.4 | 198.84 | 110852239 | 1532959 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 5388048 | 17859 | 5388048.2 | 194.61 | 108494761 | 2922023 | 1209584.0 | 21.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 5273879 | 50483 | 5273878.5 | 198.82 | 110843471 | 2985279 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 43619462 | 488279 | 10646.7 | 24.04 | 23990140 | 4602647 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 23740284 | 273709 | 5794.6 | 44.17 | 44078537 | 8456723 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 22525496 | 422365 | 5498.0 | 46.55 | 46455669 | 8504718 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 22699336 | 231996 | 5540.5 | 46.19 | 46099895 | 8844532 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 3525256 | 9669 | 860.4 | 297.45 | 296839964 | 54342997 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 4714320 | 102832 | 1150.7 | 222.42 | 221969848 | 42586201 | 8358000.0 | 28692.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 3578689 | 23778 | 873.5 | 293.01 | 292407931 | 56100155 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 40353532 | 552737 | 157631.0 | 25.98 | 25931733 | 4899100 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 22387426 | 415712 | 87450.9 | 46.84 | 46742176 | 8830671 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 21781458 | 380123 | 85083.8 | 48.14 | 48042559 | 8690649 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 21981036 | 608592 | 85863.4 | 47.70 | 47606354 | 8993934 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 3442676 | 14351 | 13448.0 | 304.58 | 303960329 | 54984840 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 4163400 | 93002 | 16263.3 | 251.86 | 251341932 | 47484268 | 7827456.0 | 2824.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 3485894 | 47762 | 13616.8 | 300.81 | 300191829 | 56713136 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 38945521 | 1456812 | 2434095.1 | 26.92 | 26869252 | 5071238 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 21744570 | 453874 | 1359035.6 | 48.22 | 48124061 | 9082819 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 21384499 | 322802 | 1336531.2 | 49.03 | 48934370 | 8845473 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 21440815 | 312720 | 1340050.9 | 48.91 | 48805840 | 9211497 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 3444262 | 10928 | 215266.4 | 304.44 | 303820346 | 54919160 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 4316069 | 39047 | 269754.3 | 242.95 | 242451404 | 45759694 | 14634752.0 | 320.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 3531488 | 49171 | 220718.0 | 296.92 | 296316165 | 55926000 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 45825864 | 483813 | 45825863.6 | 22.88 | 22835074 | 4309531 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 22046871 | 404170 | 22046871.4 | 47.56 | 47464195 | 8957643 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 21824211 | 85628 | 21824211.1 | 48.05 | 47948446 | 8666751 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 21928413 | 153137 | 21928413.0 | 47.82 | 47720599 | 9006032 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 3496018 | 13576 | 3496017.6 | 299.93 | 299322574 | 54102988 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 4314226 | 42032 | 4314226.3 | 243.05 | 242554963 | 45775995 | 17962224.0 | 32.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 3542896 | 11291 | 3542896.1 | 295.97 | 295362032 | 55741967 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 16632283 | 147771 | 4043.8 | 63.04 | 22803845 | 3026945 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 8903461 | 49290 | 2164.7 | 117.77 | 42599165 | 5654543 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 8345903 | 122501 | 2029.2 | 125.64 | 45445051 | 2903940 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 8390401 | 19588 | 2040.0 | 124.97 | 45204037 | 6000309 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 3359635 | 37985 | 816.8 | 312.11 | 112893209 | 7213878 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 3937369 | 31352 | 957.3 | 266.31 | 96328295 | 12786459 | 2235248.0 | 20363.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 3378686 | 22050 | 821.5 | 310.35 | 112256649 | 14900762 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 14274487 | 49402 | 55542.7 | 73.46 | 26570483 | 3288595 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 7890595 | 27426 | 30702.7 | 132.89 | 48067353 | 5949235 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 7768260 | 38341 | 30226.7 | 134.98 | 48824321 | 3112023 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 7786628 | 38143 | 30298.2 | 134.66 | 48709149 | 6028669 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 3280357 | 38315 | 12764.0 | 319.65 | 115621566 | 7369625 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 3555440 | 52637 | 13834.4 | 294.92 | 106675980 | 13203149 | 2183408.0 | 2310.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 3289611 | 34503 | 12800.0 | 318.75 | 115296323 | 14270078 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 13802282 | 211550 | 811898.9 | 75.97 | 27479514 | 3385527 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 7693190 | 53182 | 452540.6 | 136.30 | 49300746 | 6073943 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 7627041 | 76863 | 448649.5 | 137.48 | 49728326 | 3168988 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 7660123 | 47579 | 450595.5 | 136.89 | 49513563 | 6100163 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 3270303 | 20441 | 192370.7 | 320.64 | 115977037 | 7390753 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 3482357 | 36836 | 204844.5 | 301.11 | 108914750 | 13418499 | 2838288.0 | 241.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 3269034 | 22262 | 192296.1 | 320.76 | 116022026 | 14294129 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 17112111 | 616956 | 17112110.8 | 61.28 | 22164419 | 2729821 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 7697317 | 137235 | 7697316.9 | 136.23 | 49274313 | 6068738 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 7662805 | 109015 | 7662805.2 | 136.84 | 49496234 | 3154197 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 7696492 | 161720 | 7696491.7 | 136.24 | 49279596 | 6069389 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 3271595 | 46853 | 3271594.9 | 320.51 | 115931223 | 7387834 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 3491237 | 41684 | 3491237.0 | 300.34 | 108637713 | 13380071 | 3413232.0 | 25.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 3290589 | 66900 | 3290588.7 | 318.66 | 115262051 | 14195940 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 17318818 | 151281 | 4220.0 | 60.55 | 33192624 | 1878881 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 10704283 | 122766 | 2608.3 | 97.96 | 53703456 | 3039905 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 10431873 | 165758 | 2541.9 | 100.52 | 55105827 | 1377030 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 10439109 | 77301 | 2543.6 | 100.45 | 55067631 | 3117124 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 5196818 | 51460 | 1266.3 | 201.77 | 110617110 | 2764191 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 5632603 | 100956 | 1372.5 | 186.16 | 102058857 | 5777081 | 1432128.0 | 17605.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 5210947 | 64535 | 1269.7 | 201.23 | 110317187 | 6244547 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 14759164 | 169743 | 57428.7 | 71.05 | 38949158 | 1951940 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 9849556 | 83852 | 38325.1 | 106.46 | 58363750 | 2924904 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 9774569 | 42933 | 38033.3 | 107.28 | 58811496 | 1468505 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 9785870 | 83398 | 38077.3 | 107.15 | 58743577 | 2943939 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 5083257 | 28801 | 19779.2 | 206.28 | 113088313 | 2823780 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 5285313 | 52124 | 20565.4 | 198.39 | 108764993 | 5450765 | 1325072.0 | 2114.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 5098074 | 31630 | 19836.9 | 205.68 | 112759636 | 5650957 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 14600807 | 220096 | 858871.0 | 71.82 | 39371592 | 1957289 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 9883210 | 52221 | 581365.3 | 106.10 | 58165007 | 2891571 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 9785754 | 190989 | 575632.6 | 107.15 | 58744272 | 1466724 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 9879144 | 120470 | 581126.1 | 106.14 | 58188947 | 2892761 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 5130429 | 102377 | 301789.9 | 204.38 | 112048529 | 2797622 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 5317529 | 74667 | 312795.8 | 197.19 | 108106042 | 5374301 | 1552144.0 | 216.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 5162126 | 123802 | 303654.4 | 203.13 | 111360521 | 5536092 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 16723916 | 125341 | 16723916.0 | 62.70 | 34373349 | 1707854 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 9825638 | 71825 | 9825638.1 | 106.72 | 58505819 | 2906885 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 9807024 | 62989 | 9807023.6 | 106.92 | 58616867 | 1463543 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 9814268 | 63218 | 9814267.6 | 106.84 | 58573602 | 2910253 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 5141916 | 47656 | 5141915.6 | 203.93 | 111798218 | 2791372 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 5327659 | 37516 | 5327659.1 | 196.82 | 107900484 | 5361079 | 2061552.0 | 23.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 5163263 | 53118 | 5163263.1 | 203.08 | 111335988 | 5531773 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 75014804 | 2233445 | 1029.4 | 30.50 | 16614267 | 2812698 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 37153201 | 462694 | 509.8 | 61.58 | 33545320 | 5679026 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 35351394 | 397075 | 485.1 | 64.72 | 35255073 | 3096455 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 35635103 | 373326 | 489.0 | 64.20 | 34974391 | 5920959 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 9343671 | 74785 | 128.2 | 244.86 | 133386113 | 11715309 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 11953813 | 147359 | 164.0 | 191.39 | 104260954 | 17650769 | 7184448.0 | 155542.0 | 10 |
| synthetic | 0 | pretok-reused | 9528712 | 22475 | 130.8 | 240.10 | 130795853 | 22142972 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 19355401 | 279909 | 4705.9 | 54.17 | 20159902 | 3501142 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 10063420 | 120121 | 2446.7 | 104.20 | 38774393 | 6733894 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 9390380 | 98194 | 2283.1 | 111.66 | 41553485 | 3587714 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 9397626 | 98911 | 2284.9 | 111.58 | 41521445 | 7210970 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 3555590 | 29782 | 864.5 | 294.91 | 109743536 | 9475221 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 4221304 | 43528 | 1026.3 | 248.40 | 92436612 | 16053335 | 3023280.0 | 22248.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 3551169 | 42863 | 863.4 | 295.28 | 109880167 | 19082732 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 16726653 | 223971 | 65084.3 | 62.69 | 23328217 | 3858811 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 8961480 | 110664 | 34869.6 | 117.01 | 43542248 | 7202493 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 8771071 | 81577 | 34128.7 | 119.55 | 44487495 | 3825986 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 8809744 | 91208 | 34279.2 | 119.02 | 44292208 | 7326547 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 3451186 | 29974 | 13428.7 | 303.83 | 113063465 | 9723615 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 3787607 | 37148 | 14737.8 | 276.84 | 103020989 | 17041104 | 2896112.0 | 2405.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 3472792 | 24734 | 13512.8 | 301.94 | 112360018 | 18585909 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 16380171 | 133765 | 963539.5 | 64.01 | 23821668 | 3928042 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 8726659 | 53121 | 513332.9 | 120.16 | 44713904 | 7373039 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 8654608 | 144896 | 509094.6 | 121.16 | 45086155 | 3877241 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 8696985 | 34528 | 511587.4 | 120.57 | 44866467 | 7398196 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 3451722 | 15038 | 203042.5 | 303.78 | 113045894 | 9721524 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 3756523 | 22672 | 220972.0 | 279.13 | 103873436 | 17128071 | 3903344.0 | 256.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 3463580 | 40199 | 203740.0 | 302.74 | 112658852 | 18576730 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 19452487 | 379912 | 19452487.2 | 53.90 | 20059286 | 3307084 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 8684614 | 60953 | 8684613.9 | 120.74 | 44930380 | 7407468 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 8649797 | 84315 | 8649796.7 | 121.23 | 45111234 | 3879398 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 8653891 | 60801 | 8653891.2 | 121.17 | 45089890 | 7433766 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 3439236 | 17093 | 3439236.0 | 304.89 | 113456302 | 9756818 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 3801143 | 36556 | 3801143.0 | 275.86 | 102654123 | 16924120 | 5559536.0 | 27.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 3443322 | 37930 | 3443321.8 | 304.52 | 113321678 | 18682831 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.889x |
| ar-wikipedia | 4096 | 1.820x |
| ar-wikipedia | 65536 | 1.806x |
| ar-wikipedia | 1048576 | 1.816x |
| en-wikitext2 | 256 | 6.390x |
| en-wikitext2 | 4096 | 6.327x |
| en-wikitext2 | 65536 | 6.209x |
| en-wikitext2 | 1048576 | 6.243x |
| ja-wikipedia | 256 | 2.484x |
| ja-wikipedia | 4096 | 2.368x |
| ja-wikipedia | 65536 | 2.332x |
| ja-wikipedia | 1048576 | 2.342x |
| ru-wikipedia | 256 | 2.007x |
| ru-wikipedia | 4096 | 1.923x |
| ru-wikipedia | 65536 | 1.907x |
| ru-wikipedia | 1048576 | 1.907x |
| synthetic | 0 | 3.783x |
| zh-wikipedia | 256 | 2.641x |
| zh-wikipedia | 4096 | 2.541x |
| zh-wikipedia | 65536 | 2.507x |
| zh-wikipedia | 1048576 | 2.515x |
| Combined selected natural corpora | 256 | 2.898x |
| Combined selected natural corpora | 4096 | 2.812x |
| Combined selected natural corpora | 65536 | 2.772x |
| Combined selected natural corpora | 1048576 | 2.789x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.887x |
| ar-wikipedia | 4096 | 1.814x |
| ar-wikipedia | 65536 | 1.819x |
| ar-wikipedia | 1048576 | 1.814x |
| en-wikitext2 | 256 | 6.343x |
| en-wikitext2 | 4096 | 6.306x |
| en-wikitext2 | 65536 | 6.071x |
| en-wikitext2 | 1048576 | 6.189x |
| ja-wikipedia | 256 | 2.483x |
| ja-wikipedia | 4096 | 2.367x |
| ja-wikipedia | 65536 | 2.343x |
| ja-wikipedia | 1048576 | 2.339x |
| ru-wikipedia | 256 | 2.003x |
| ru-wikipedia | 4096 | 1.920x |
| ru-wikipedia | 65536 | 1.914x |
| ru-wikipedia | 1048576 | 1.901x |
| synthetic | 0 | 3.740x |
| zh-wikipedia | 256 | 2.646x |
| zh-wikipedia | 4096 | 2.537x |
| zh-wikipedia | 65536 | 2.511x |
| zh-wikipedia | 1048576 | 2.513x |
| Combined selected natural corpora | 256 | 2.897x |
| Combined selected natural corpora | 4096 | 2.810x |
| Combined selected natural corpora | 65536 | 2.769x |
| Combined selected natural corpora | 1048576 | 2.784x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.828x |
| ar-wikipedia | 4096 | 1.780x |
| ar-wikipedia | 65536 | 1.777x |
| ar-wikipedia | 1048576 | 1.778x |
| en-wikitext2 | 256 | 5.036x |
| en-wikitext2 | 4096 | 5.377x |
| en-wikitext2 | 65536 | 5.038x |
| en-wikitext2 | 1048576 | 5.110x |
| ja-wikipedia | 256 | 2.261x |
| ja-wikipedia | 4096 | 2.219x |
| ja-wikipedia | 65536 | 2.209x |
| ja-wikipedia | 1048576 | 2.205x |
| ru-wikipedia | 256 | 1.900x |
| ru-wikipedia | 4096 | 1.864x |
| ru-wikipedia | 65536 | 1.859x |
| ru-wikipedia | 1048576 | 1.844x |
| synthetic | 0 | 3.108x |
| zh-wikipedia | 256 | 2.384x |
| zh-wikipedia | 4096 | 2.366x |
| zh-wikipedia | 65536 | 2.323x |
| zh-wikipedia | 1048576 | 2.285x |
| Combined selected natural corpora | 256 | 2.640x |
| Combined selected natural corpora | 4096 | 2.646x |
| Combined selected natural corpora | 65536 | 2.590x |
| Combined selected natural corpora | 1048576 | 2.591x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.946x |
| ar-wikipedia | 4096 | 2.571x |
| ar-wikipedia | 65536 | 2.532x |
| ar-wikipedia | 1048576 | 2.924x |
| en-wikitext2 | 256 | 12.189x |
| en-wikitext2 | 4096 | 11.576x |
| en-wikitext2 | 65536 | 11.028x |
| en-wikitext2 | 1048576 | 12.935x |
| ja-wikipedia | 256 | 4.923x |
| ja-wikipedia | 4096 | 4.339x |
| ja-wikipedia | 65536 | 4.222x |
| ja-wikipedia | 1048576 | 5.200x |
| ru-wikipedia | 256 | 3.324x |
| ru-wikipedia | 4096 | 2.895x |
| ru-wikipedia | 65536 | 2.828x |
| ru-wikipedia | 1048576 | 3.239x |
| synthetic | 0 | 7.873x |
| zh-wikipedia | 256 | 5.450x |
| zh-wikipedia | 4096 | 4.816x |
| zh-wikipedia | 65536 | 4.729x |
| zh-wikipedia | 1048576 | 5.649x |
| Combined selected natural corpora | 256 | 5.350x |
| Combined selected natural corpora | 4096 | 4.835x |
| Combined selected natural corpora | 65536 | 4.695x |
| Combined selected natural corpora | 1048576 | 5.529x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.805x |
| ar-wikipedia | 4096 | 2.512x |
| ar-wikipedia | 65536 | 2.471x |
| ar-wikipedia | 1048576 | 2.862x |
| en-wikitext2 | 256 | 9.253x |
| en-wikitext2 | 4096 | 9.692x |
| en-wikitext2 | 65536 | 9.023x |
| en-wikitext2 | 1048576 | 10.622x |
| ja-wikipedia | 256 | 4.224x |
| ja-wikipedia | 4096 | 4.015x |
| ja-wikipedia | 65536 | 3.963x |
| ja-wikipedia | 1048576 | 4.901x |
| ru-wikipedia | 256 | 3.075x |
| ru-wikipedia | 4096 | 2.792x |
| ru-wikipedia | 65536 | 2.746x |
| ru-wikipedia | 1048576 | 3.139x |
| synthetic | 0 | 6.275x |
| zh-wikipedia | 256 | 4.585x |
| zh-wikipedia | 4096 | 4.416x |
| zh-wikipedia | 65536 | 4.360x |
| zh-wikipedia | 1048576 | 5.118x |
| Combined selected natural corpora | 256 | 4.672x |
| Combined selected natural corpora | 4096 | 4.493x |
| Combined selected natural corpora | 65536 | 4.363x |
| Combined selected natural corpora | 1048576 | 5.131x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 860.4 | 12.7 | 10 |
| onig-compile | 839.4 | 11.3 | 10 |
| pretok-compile | 10.5 | 0.1 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
