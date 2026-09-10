# P07 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 12189416 | 50149 | 2970.1 | 86.02 | 47957588 | 930561 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 7709119 | 31839 | 1878.4 | 136.02 | 75829029 | 1471374 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 7469843 | 14784 | 1820.1 | 140.37 | 78258000 | 565206 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 7478707 | 33321 | 1822.3 | 140.21 | 78165249 | 1516706 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 4401465 | 13649 | 1072.5 | 238.23 | 132813747 | 959226 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 4554349 | 16470 | 1109.7 | 230.24 | 128355341 | 2490587 | 388032.0 | 7808.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 4418400 | 3418 | 1076.6 | 237.32 | 132304690 | 2567219 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 9954486 | 24566 | 38733.4 | 105.34 | 58724782 | 757146 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 6917520 | 30631 | 26916.4 | 151.58 | 84506442 | 1089552 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 6857775 | 15422 | 26683.9 | 152.90 | 85242665 | 614777 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 6847773 | 39993 | 26645.0 | 153.13 | 85367165 | 1100650 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 4331040 | 8040 | 16852.3 | 242.11 | 134973348 | 973438 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 4415162 | 10779 | 17179.6 | 237.49 | 132401702 | 1707072 | 345360.0 | 1435.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 4339546 | 7456 | 16885.4 | 241.63 | 134708798 | 1736818 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 9680180 | 57530 | 569422.4 | 108.32 | 60388855 | 754015 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 6833591 | 17987 | 401975.9 | 153.44 | 85544340 | 1068106 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 6824777 | 26484 | 401457.5 | 153.64 | 85654816 | 617749 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 6812225 | 15038 | 400719.1 | 153.93 | 85812639 | 1071456 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 4318344 | 12997 | 254020.3 | 242.82 | 135370164 | 976300 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 4378512 | 7073 | 257559.5 | 239.48 | 133509970 | 1667005 | 322832.0 | 164.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 4327339 | 8908 | 254549.4 | 242.31 | 135088790 | 1686718 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 10891018 | 50311 | 10891018.1 | 96.28 | 53674963 | 668716 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 6843464 | 31129 | 6843464.0 | 153.22 | 85420922 | 1064227 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 6852351 | 68844 | 6852350.9 | 153.02 | 85310138 | 615263 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 6845827 | 38450 | 6845827.4 | 153.17 | 85391431 | 1063860 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 4312286 | 14819 | 4312286.1 | 243.16 | 135560347 | 977672 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 4370341 | 9851 | 4370341.5 | 239.93 | 133759570 | 1666460 | 505072.0 | 18.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 4327999 | 10277 | 4327998.9 | 242.28 | 135068194 | 1682764 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 13379036 | 28818 | 3265.6 | 78.37 | 78214677 | 1352788 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 9123632 | 25184 | 2226.9 | 114.93 | 114695220 | 1983749 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 8865629 | 19577 | 2163.9 | 118.27 | 118033019 | 913077 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 8896188 | 13749 | 2171.4 | 117.87 | 117627570 | 2034467 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 2511928 | 6875 | 613.1 | 417.44 | 416587180 | 3222624 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 2720353 | 5377 | 664.0 | 385.46 | 384669574 | 6653181 | 688336.0 | 10614.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 2518995 | 8445 | 614.8 | 416.27 | 415418490 | 7185009 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 11972014 | 30764 | 46765.7 | 87.59 | 87406932 | 1195872 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 8455495 | 25557 | 33029.3 | 124.01 | 123758222 | 1693218 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 8378687 | 22507 | 32729.2 | 125.15 | 124892722 | 964829 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 8386691 | 20883 | 32760.5 | 125.03 | 124773525 | 1707110 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 2450522 | 5204 | 9572.4 | 427.90 | 427026151 | 3298889 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 2563436 | 15674 | 10013.4 | 409.05 | 408216534 | 5585082 | 642624.0 | 1725.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 2455278 | 4240 | 9590.9 | 427.07 | 426198921 | 5831111 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 11801670 | 21987 | 737604.4 | 88.85 | 88668552 | 1193221 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 8358041 | 8266 | 522377.6 | 125.46 | 125201224 | 1684845 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 8326856 | 23533 | 520428.5 | 125.93 | 125670118 | 970714 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 8331899 | 16999 | 520743.7 | 125.85 | 125594050 | 1690131 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 2457779 | 8907 | 153611.2 | 426.64 | 425765329 | 3288742 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 2538913 | 8028 | 158682.0 | 413.00 | 412159513 | 5546469 | 678400.0 | 185.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 2459193 | 9149 | 153699.6 | 426.39 | 425520501 | 5726269 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 12989659 | 246170 | 12989658.6 | 80.72 | 80559238 | 1082938 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 8402435 | 24896 | 8402435.4 | 124.79 | 124539725 | 1674157 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 8378106 | 19818 | 8378105.6 | 125.16 | 124901386 | 964777 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 8389913 | 27559 | 8389912.8 | 124.98 | 124725610 | 1676656 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 2472053 | 8730 | 2472053.4 | 424.17 | 423306800 | 3269751 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 2542794 | 10123 | 2542793.8 | 412.37 | 411530423 | 5532104 | 914672.0 | 20.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 2470553 | 14099 | 2470552.9 | 424.43 | 423563899 | 5693867 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 9875880 | 24506 | 2401.1 | 106.18 | 38404678 | 1575556 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 5677572 | 18590 | 1380.4 | 184.69 | 66803208 | 2740608 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 5442632 | 21121 | 1323.3 | 192.66 | 69686878 | 1209893 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 5471039 | 12566 | 1330.2 | 191.66 | 69325044 | 2844067 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 3775506 | 14336 | 917.9 | 277.73 | 100458047 | 1744137 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 3970134 | 6668 | 965.3 | 264.12 | 95533309 | 3919264 | 575408.0 | 9520.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 3785295 | 8961 | 920.3 | 277.01 | 100198273 | 4110644 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 8209664 | 71258 | 31944.2 | 127.72 | 46199213 | 1438183 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 4975573 | 18764 | 19360.2 | 210.74 | 76228406 | 2372993 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 4885269 | 18119 | 19008.8 | 214.64 | 77637482 | 1346088 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 4877447 | 15425 | 18978.4 | 214.98 | 77761990 | 2420734 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 3687624 | 8970 | 14348.7 | 284.35 | 102852141 | 1783262 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 3790469 | 15852 | 14748.9 | 276.63 | 100061507 | 3114918 | 529264.0 | 1563.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 3691605 | 17556 | 14364.2 | 284.04 | 102741221 | 3198338 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 8053736 | 24400 | 473749.2 | 130.20 | 47093673 | 1436725 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 4860093 | 8810 | 285887.8 | 215.75 | 78039665 | 2380819 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 4842164 | 23224 | 284833.2 | 216.55 | 78328617 | 1358071 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 4832022 | 16614 | 284236.6 | 217.01 | 78493023 | 2394650 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 3684647 | 7064 | 216743.9 | 284.58 | 102935241 | 1784703 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 3755353 | 7444 | 220903.1 | 279.22 | 100997169 | 3081202 | 576528.0 | 178.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 3691325 | 12463 | 217136.8 | 284.06 | 102749007 | 3134647 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 9186287 | 43634 | 9186287.1 | 114.15 | 41287627 | 1257853 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 4852053 | 16101 | 4852053.1 | 216.11 | 78168972 | 2381466 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 4840844 | 9808 | 4840843.8 | 216.61 | 78349978 | 1358441 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 4839736 | 12361 | 4839735.8 | 216.66 | 78367915 | 2387527 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 3685895 | 11590 | 3685894.9 | 284.48 | 102900383 | 1784099 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 3763432 | 4849 | 3763432.5 | 278.62 | 100780338 | 3070335 | 914672.0 | 20.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 3698208 | 9404 | 3698207.6 | 283.54 | 102557791 | 3124487 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 12277773 | 17303 | 2991.7 | 85.40 | 46820949 | 1038706 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 7837606 | 37267 | 1909.7 | 133.79 | 73345989 | 1627155 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 7601756 | 33396 | 1852.3 | 137.94 | 75621609 | 690235 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 7640741 | 20301 | 1861.8 | 137.23 | 75235770 | 1669079 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 4263544 | 14589 | 1038.9 | 245.94 | 134830784 | 1230666 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 4431538 | 26398 | 1079.8 | 236.62 | 129719530 | 2877782 | 437888.0 | 8526.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 4278996 | 8174 | 1042.6 | 245.05 | 134343901 | 2980372 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 10272067 | 25038 | 39969.1 | 102.08 | 55963127 | 871879 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 7136727 | 27688 | 27769.4 | 146.93 | 80549108 | 1254917 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 7055263 | 27002 | 27452.4 | 148.62 | 81479178 | 742708 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 7034569 | 13200 | 27371.9 | 149.06 | 81718871 | 1273141 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 4185726 | 17275 | 16286.9 | 250.51 | 137337458 | 1251874 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 4269101 | 16974 | 16611.3 | 245.62 | 134655294 | 2097866 | 406224.0 | 1508.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 4192536 | 4673 | 16313.4 | 250.11 | 137114400 | 2136177 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 10035650 | 63446 | 590332.4 | 104.49 | 57281491 | 869102 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 7031016 | 62003 | 413589.2 | 149.14 | 81760165 | 1240504 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 7001873 | 13105 | 411874.9 | 149.76 | 82100458 | 748228 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 6994889 | 22678 | 411464.1 | 149.91 | 82182435 | 1246910 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 4179468 | 6694 | 245851.1 | 250.89 | 137543102 | 1253509 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 4241836 | 7420 | 249519.7 | 247.20 | 135520810 | 2056185 | 385296.0 | 169.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 4186194 | 7000 | 246246.7 | 250.48 | 137322114 | 2083516 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 11254126 | 47942 | 11254126.4 | 93.17 | 51079664 | 773583 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 7025677 | 24262 | 7025676.8 | 149.25 | 81822295 | 1239169 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 7046890 | 48559 | 7046889.6 | 148.80 | 81575991 | 743449 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 7038008 | 12128 | 7038008.4 | 148.99 | 81678931 | 1236998 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 4177583 | 18506 | 4177582.8 | 251.00 | 137605172 | 1254074 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 4232156 | 8218 | 4232156.0 | 247.76 | 135830768 | 2057108 | 685296.0 | 19.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 4177436 | 7740 | 4177435.8 | 251.01 | 137610015 | 2084054 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 53916077 | 536109 | 1208.2 | 30.14 | 15723139 | 3372760 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 27103338 | 162240 | 607.4 | 59.96 | 31277698 | 6709358 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 25469156 | 145688 | 570.8 | 63.81 | 33284573 | 4659008 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 25694180 | 125616 | 575.8 | 63.25 | 32993074 | 7077323 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 6139440 | 28672 | 137.6 | 264.72 | 138079362 | 19327658 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 7828200 | 26802 | 175.4 | 207.61 | 108291823 | 23229607 | 8079472.0 | 78094.0 | 10 |
| synthetic | 0 | pretok-reused | 6247160 | 21428 | 140.0 | 260.16 | 135698456 | 29108586 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 9949971 | 26923 | 2419.2 | 105.38 | 39216496 | 1565331 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 5779108 | 12528 | 1405.1 | 181.44 | 67519587 | 2695052 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 5540207 | 11888 | 1347.0 | 189.27 | 70431126 | 1206273 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 5562967 | 11474 | 1352.5 | 188.49 | 70142963 | 2799765 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 2976336 | 10356 | 723.6 | 352.30 | 131101801 | 2245378 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 3158047 | 7810 | 767.8 | 332.03 | 123558329 | 4931846 | 574960.0 | 9016.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 2993530 | 6649 | 727.8 | 350.28 | 130348766 | 5202887 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 8307689 | 24034 | 32325.6 | 126.22 | 46968899 | 1422778 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 5065347 | 20919 | 19709.5 | 207.01 | 77033814 | 2333503 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 4974851 | 25403 | 19357.4 | 210.77 | 78435109 | 1340945 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 4971823 | 33568 | 19345.6 | 210.90 | 78482882 | 2377398 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 2912371 | 13463 | 11332.2 | 360.04 | 133981191 | 2290573 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 3002837 | 13509 | 11684.2 | 349.19 | 129944774 | 3936277 | 540496.0 | 1496.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 2920913 | 9499 | 11365.4 | 358.99 | 133589388 | 4046680 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 8145804 | 33020 | 479164.9 | 128.73 | 47902330 | 1421591 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 4961202 | 17126 | 291835.4 | 211.35 | 78650907 | 2334112 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 4934955 | 23161 | 290291.5 | 212.48 | 79069217 | 1351785 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 4922974 | 21044 | 289586.7 | 213.00 | 79261634 | 2352236 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 2910851 | 10792 | 171226.5 | 360.23 | 134051181 | 2291770 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 2971413 | 9132 | 174789.0 | 352.89 | 131319001 | 3897136 | 554768.0 | 176.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 2909376 | 7442 | 171139.8 | 360.41 | 134119145 | 3980235 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 9257727 | 23857 | 9257726.8 | 113.26 | 42148900 | 1249227 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 4944921 | 32368 | 4944920.5 | 212.05 | 78909862 | 2338764 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 4935784 | 18676 | 4935784.1 | 212.44 | 79055929 | 1351558 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 4930521 | 17702 | 4930520.8 | 212.67 | 79140321 | 2345594 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 2905387 | 20872 | 2905387.2 | 360.91 | 134303271 | 2296080 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 2985700 | 12256 | 2985699.7 | 351.20 | 130690640 | 3873464 | 914672.0 | 20.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 2913658 | 7685 | 2913658.4 | 359.88 | 133922012 | 3969237 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.697x |
| ar-wikipedia | 4096 | 1.583x |
| ar-wikipedia | 65536 | 1.580x |
| ar-wikipedia | 1048576 | 1.589x |
| en-wikitext2 | 256 | 3.529x |
| en-wikitext2 | 4096 | 3.419x |
| en-wikitext2 | 65536 | 3.388x |
| en-wikitext2 | 1048576 | 3.389x |
| ja-wikipedia | 256 | 1.442x |
| ja-wikipedia | 4096 | 1.325x |
| ja-wikipedia | 65536 | 1.314x |
| ja-wikipedia | 1048576 | 1.313x |
| ru-wikipedia | 256 | 1.783x |
| ru-wikipedia | 4096 | 1.686x |
| ru-wikipedia | 65536 | 1.675x |
| ru-wikipedia | 1048576 | 1.687x |
| synthetic | 0 | 4.148x |
| zh-wikipedia | 256 | 1.861x |
| zh-wikipedia | 4096 | 1.708x |
| zh-wikipedia | 65536 | 1.695x |
| zh-wikipedia | 1048576 | 1.699x |
| Combined selected natural corpora | 256 | 1.948x |
| Combined selected natural corpora | 4096 | 1.830x |
| Combined selected natural corpora | 65536 | 1.819x |
| Combined selected natural corpora | 1048576 | 1.826x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.693x |
| ar-wikipedia | 4096 | 1.578x |
| ar-wikipedia | 65536 | 1.574x |
| ar-wikipedia | 1048576 | 1.582x |
| en-wikitext2 | 256 | 3.532x |
| en-wikitext2 | 4096 | 3.416x |
| en-wikitext2 | 65536 | 3.388x |
| en-wikitext2 | 1048576 | 3.396x |
| ja-wikipedia | 256 | 1.445x |
| ja-wikipedia | 4096 | 1.321x |
| ja-wikipedia | 65536 | 1.309x |
| ja-wikipedia | 1048576 | 1.309x |
| ru-wikipedia | 256 | 1.786x |
| ru-wikipedia | 4096 | 1.678x |
| ru-wikipedia | 65536 | 1.671x |
| ru-wikipedia | 1048576 | 1.685x |
| synthetic | 0 | 4.113x |
| zh-wikipedia | 256 | 1.858x |
| zh-wikipedia | 4096 | 1.702x |
| zh-wikipedia | 65536 | 1.692x |
| zh-wikipedia | 1048576 | 1.692x |
| Combined selected natural corpora | 256 | 1.948x |
| Combined selected natural corpora | 4096 | 1.825x |
| Combined selected natural corpora | 65536 | 1.815x |
| Combined selected natural corpora | 1048576 | 1.822x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.693x |
| ar-wikipedia | 4096 | 1.567x |
| ar-wikipedia | 65536 | 1.561x |
| ar-wikipedia | 1048576 | 1.566x |
| en-wikitext2 | 256 | 3.354x |
| en-wikitext2 | 4096 | 3.299x |
| en-wikitext2 | 65536 | 3.292x |
| en-wikitext2 | 1048576 | 3.304x |
| ja-wikipedia | 256 | 1.430x |
| ja-wikipedia | 4096 | 1.313x |
| ja-wikipedia | 65536 | 1.294x |
| ja-wikipedia | 1048576 | 1.289x |
| ru-wikipedia | 256 | 1.769x |
| ru-wikipedia | 4096 | 1.672x |
| ru-wikipedia | 65536 | 1.658x |
| ru-wikipedia | 1048576 | 1.660x |
| synthetic | 0 | 3.462x |
| zh-wikipedia | 256 | 1.830x |
| zh-wikipedia | 4096 | 1.687x |
| zh-wikipedia | 65536 | 1.670x |
| zh-wikipedia | 1048576 | 1.656x |
| Combined selected natural corpora | 256 | 1.918x |
| Combined selected natural corpora | 4096 | 1.804x |
| Combined selected natural corpora | 65536 | 1.792x |
| Combined selected natural corpora | 1048576 | 1.792x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.759x |
| ar-wikipedia | 4096 | 2.294x |
| ar-wikipedia | 65536 | 2.237x |
| ar-wikipedia | 1048576 | 2.516x |
| en-wikitext2 | 256 | 5.311x |
| en-wikitext2 | 4096 | 4.876x |
| en-wikitext2 | 65536 | 4.799x |
| en-wikitext2 | 1048576 | 5.258x |
| ja-wikipedia | 256 | 2.609x |
| ja-wikipedia | 4096 | 2.224x |
| ja-wikipedia | 65536 | 2.182x |
| ja-wikipedia | 1048576 | 2.484x |
| ru-wikipedia | 256 | 2.869x |
| ru-wikipedia | 4096 | 2.450x |
| ru-wikipedia | 65536 | 2.397x |
| ru-wikipedia | 1048576 | 2.694x |
| synthetic | 0 | 8.630x |
| zh-wikipedia | 256 | 3.324x |
| zh-wikipedia | 4096 | 2.844x |
| zh-wikipedia | 65536 | 2.800x |
| zh-wikipedia | 1048576 | 3.177x |
| Combined selected natural corpora | 256 | 3.205x |
| Combined selected natural corpora | 4096 | 2.768x |
| Combined selected natural corpora | 65536 | 2.715x |
| Combined selected natural corpora | 1048576 | 3.046x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.676x |
| ar-wikipedia | 4096 | 2.255x |
| ar-wikipedia | 65536 | 2.211x |
| ar-wikipedia | 1048576 | 2.492x |
| en-wikitext2 | 256 | 4.918x |
| en-wikitext2 | 4096 | 4.670x |
| en-wikitext2 | 65536 | 4.648x |
| en-wikitext2 | 1048576 | 5.108x |
| ja-wikipedia | 256 | 2.488x |
| ja-wikipedia | 4096 | 2.166x |
| ja-wikipedia | 65536 | 2.145x |
| ja-wikipedia | 1048576 | 2.441x |
| ru-wikipedia | 256 | 2.771x |
| ru-wikipedia | 4096 | 2.406x |
| ru-wikipedia | 65536 | 2.366x |
| ru-wikipedia | 1048576 | 2.659x |
| synthetic | 0 | 6.887x |
| zh-wikipedia | 256 | 3.151x |
| zh-wikipedia | 4096 | 2.767x |
| zh-wikipedia | 65536 | 2.741x |
| zh-wikipedia | 1048576 | 3.101x |
| Combined selected natural corpora | 256 | 3.062x |
| Combined selected natural corpora | 4096 | 2.700x |
| Combined selected natural corpora | 65536 | 2.668x |
| Combined selected natural corpora | 1048576 | 2.994x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 2155.4 | 3.3 | 10 |
| onig-compile | 2149.4 | 3.6 | 10 |
| pretok-compile | 9.9 | 0.1 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
