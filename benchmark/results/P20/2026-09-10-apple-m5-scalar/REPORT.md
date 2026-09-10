# P20 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 36457095 | 338120 | 8883.3 | 28.76 | 16034602 | 3337540 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-at | 19845208 | 150583 | 4835.6 | 52.84 | 29456733 | 6131304 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 22185899 | 97213 | 5405.9 | 47.26 | 26348943 | 5484429 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 21186663 | 74477 | 5162.4 | 49.49 | 27591650 | 5743094 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 21225308 | 107439 | 5171.9 | 49.40 | 27541414 | 5732638 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 6213902 | 24814 | 1514.1 | 168.75 | 94075349 | 19581416 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 7181375 | 67102 | 1749.8 | 146.01 | 81401542 | 16943413 | 5045152.0 | 25500.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 6248021 | 23699 | 1522.4 | 167.83 | 93561624 | 19474486 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 32320109 | 142060 | 125759.2 | 32.44 | 18087037 | 3661652 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-at | 19386256 | 166068 | 75432.9 | 54.09 | 30154095 | 6104583 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 20279334 | 128490 | 78907.9 | 51.71 | 28826144 | 5835744 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 19953837 | 83424 | 77641.4 | 52.55 | 29296370 | 5930939 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 20062037 | 29678 | 78062.4 | 52.27 | 29138368 | 5898952 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 6128244 | 39726 | 23845.3 | 171.11 | 95390288 | 19311403 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 6611047 | 32436 | 25723.9 | 158.61 | 88423968 | 17901098 | 4597744.0 | 2595.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 6163532 | 13342 | 23982.6 | 170.13 | 94844157 | 19200841 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 31742060 | 192976 | 1867180.0 | 33.03 | 18416417 | 3721498 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-at | 19216111 | 91885 | 1130359.5 | 54.57 | 30421088 | 6147342 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 19868916 | 116107 | 1168759.8 | 52.77 | 29421585 | 5945367 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 19716969 | 69760 | 1159821.7 | 53.18 | 29648320 | 5991185 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 19790488 | 93847 | 1164146.4 | 52.98 | 29538180 | 5968928 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 6104753 | 14307 | 359103.1 | 171.76 | 95757354 | 19350168 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 6615802 | 28594 | 389164.8 | 158.50 | 88360412 | 17855431 | 8081200.0 | 290.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 6145035 | 26950 | 361472.7 | 170.64 | 95129640 | 19223323 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 35376957 | 655177 | 35376956.9 | 29.64 | 16524174 | 3338699 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-at | 19306089 | 87336 | 19306088.9 | 54.31 | 30279307 | 6117914 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 19802811 | 166797 | 19802810.9 | 52.95 | 29519799 | 5964456 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 19809727 | 177299 | 19809726.8 | 52.93 | 29509493 | 5962374 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 19862637 | 86682 | 19862637.0 | 52.79 | 29430886 | 5946491 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 6119477 | 42814 | 6119477.4 | 171.35 | 95526948 | 19301158 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 6680917 | 34638 | 6680916.7 | 156.95 | 87499220 | 17679161 | 11302128.0 | 30.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 6140028 | 18181 | 6140028.0 | 170.78 | 95207221 | 19236557 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 50638746 | 460285 | 12360.0 | 20.71 | 20664750 | 4123364 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-at | 27622794 | 228607 | 6742.2 | 37.96 | 37883098 | 7559047 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 30060632 | 286355 | 7337.2 | 34.88 | 34810878 | 6946028 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 29119653 | 203345 | 7107.6 | 36.01 | 35935765 | 7170484 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 29270792 | 275151 | 7144.4 | 35.82 | 35750212 | 7133459 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 4688332 | 31638 | 1144.3 | 223.66 | 223200278 | 44536522 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 5910007 | 38008 | 1442.5 | 177.42 | 177061900 | 35330248 | 8546416.0 | 28784.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 4739870 | 7622 | 1156.9 | 221.22 | 220773378 | 44052268 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 47042748 | 221443 | 183760.7 | 22.29 | 22244385 | 4373256 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-at | 26982475 | 132850 | 105400.3 | 38.86 | 38782099 | 7624579 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 28356299 | 211948 | 110766.8 | 36.98 | 36903159 | 7255178 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 27883641 | 117554 | 108920.5 | 37.61 | 37528707 | 7378161 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 28080164 | 397926 | 109688.1 | 37.34 | 37266057 | 7326524 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 4559177 | 23178 | 17809.3 | 229.99 | 229523238 | 45124375 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 5314427 | 25370 | 20759.5 | 197.31 | 196904962 | 38711607 | 8585216.0 | 2861.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 4616862 | 16047 | 18034.6 | 227.12 | 226655483 | 44560573 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 47061831 | 484471 | 2941364.5 | 22.28 | 22235365 | 4367467 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-at | 26865330 | 101888 | 1679083.1 | 39.03 | 38951206 | 7650790 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 27656242 | 122942 | 1728515.1 | 37.91 | 37837281 | 7431993 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 27607361 | 125507 | 1725460.0 | 37.98 | 37904275 | 7445152 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 27687159 | 505462 | 1730447.4 | 37.87 | 37795030 | 7423694 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 4557880 | 20525 | 284867.5 | 230.06 | 229588521 | 45095743 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 5368962 | 21117 | 335560.1 | 195.30 | 194904896 | 38283191 | 14634752.0 | 320.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 4625180 | 12876 | 289073.8 | 226.71 | 226247822 | 44439564 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 52116683 | 756723 | 52116683.4 | 20.12 | 20078734 | 3943593 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-at | 27026519 | 166759 | 27026518.7 | 38.80 | 38718897 | 7604642 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 27878287 | 332789 | 27878287.1 | 37.61 | 37535915 | 7372297 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 27688975 | 187319 | 27688974.8 | 37.87 | 37792551 | 7422702 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 27837697 | 177030 | 27837696.8 | 37.67 | 37590646 | 7383046 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 4575629 | 22607 | 4575628.6 | 229.17 | 228697976 | 44917763 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 5396514 | 21439 | 5396514.3 | 194.31 | 193909797 | 38085139 | 17962224.0 | 32.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 4659456 | 32875 | 4659456.4 | 225.04 | 224583494 | 44109652 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 22835906 | 71135 | 5552.1 | 45.92 | 16608931 | 2337328 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-at | 12812867 | 51022 | 3115.2 | 81.84 | 29601493 | 4165734 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 14347825 | 38230 | 3488.4 | 73.08 | 26434669 | 3720076 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 13846493 | 26313 | 3366.5 | 75.73 | 27391773 | 3854767 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 13838395 | 55622 | 3364.6 | 75.77 | 27407803 | 3857023 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 5342578 | 31013 | 1298.9 | 196.27 | 70991950 | 9990496 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 5996935 | 18126 | 1458.0 | 174.85 | 63245638 | 8900380 | 2307440.0 | 20068.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 5352403 | 5637 | 1301.3 | 195.91 | 70861631 | 9972157 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 19731027 | 63272 | 76774.4 | 53.14 | 19222517 | 2532914 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-at | 12294699 | 94516 | 47839.3 | 85.29 | 30849067 | 4064923 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 12846711 | 103569 | 49987.2 | 81.62 | 29523510 | 3890256 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 12651850 | 116384 | 49229.0 | 82.88 | 29978224 | 3950173 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 12763722 | 133402 | 49664.3 | 82.15 | 29715469 | 3915551 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 5245132 | 39796 | 20409.1 | 199.91 | 72310863 | 9528264 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 5560910 | 20963 | 21637.8 | 188.56 | 68204667 | 8987199 | 2315760.0 | 2312.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 5268201 | 26916 | 20498.8 | 199.04 | 71994215 | 9486540 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 19412704 | 62233 | 1141923.8 | 54.01 | 19537721 | 2563321 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-at | 12307747 | 121100 | 723985.1 | 85.20 | 30816363 | 4043063 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 12525541 | 78101 | 736796.5 | 83.71 | 30280529 | 3972763 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 12505851 | 144896 | 735638.3 | 83.85 | 30328204 | 3979018 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 12478801 | 49066 | 734047.1 | 84.03 | 30393945 | 3987643 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 5244620 | 39809 | 308507.0 | 199.93 | 72317924 | 9488009 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 5447322 | 35279 | 320430.7 | 192.49 | 69626868 | 9134947 | 2912016.0 | 241.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 5242694 | 20275 | 308393.8 | 200.01 | 72344482 | 9491494 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 22305354 | 682052 | 22305354.2 | 47.01 | 17003989 | 2230182 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-at | 12248768 | 64876 | 12248767.8 | 85.61 | 30964747 | 4061225 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 12458052 | 45607 | 12458052.3 | 84.17 | 30444566 | 3993000 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 12470291 | 86613 | 12470291.2 | 84.09 | 30414687 | 3989081 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 12475530 | 54911 | 12475530.0 | 84.05 | 30401915 | 3987406 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 5235927 | 13523 | 5235926.7 | 200.27 | 72437989 | 9500706 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 5502554 | 15990 | 5502553.8 | 190.56 | 68927995 | 9040348 | 4363504.0 | 26.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 5258561 | 34284 | 5258560.5 | 199.40 | 72126202 | 9459813 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 31819900 | 229931 | 7753.4 | 32.95 | 18065959 | 3126628 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-at | 16783693 | 53090 | 4089.6 | 62.48 | 34250924 | 5927718 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 18726545 | 83422 | 4563.0 | 55.99 | 30697441 | 5312726 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 17887425 | 51731 | 4358.5 | 58.62 | 32137493 | 5561952 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 17909985 | 97833 | 4364.0 | 58.55 | 32097012 | 5554946 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 4649814 | 23345 | 1133.0 | 225.51 | 123630095 | 21396338 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 5474485 | 35106 | 1333.9 | 191.54 | 105006589 | 18173216 | 4372352.0 | 24815.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 4679835 | 17553 | 1140.3 | 224.06 | 122837011 | 21259081 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 27808193 | 154884 | 108203.1 | 37.71 | 20672217 | 3453047 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-at | 16220052 | 109958 | 63113.0 | 64.65 | 35441132 | 5920018 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 16960406 | 69478 | 65993.8 | 61.82 | 33894060 | 5661598 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 16709868 | 109900 | 65018.9 | 62.75 | 34402246 | 5746485 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 16826381 | 114384 | 65472.3 | 62.32 | 34164032 | 5706694 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 4533930 | 12576 | 17641.8 | 231.27 | 126790009 | 21178758 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 4990176 | 15181 | 19417.0 | 210.13 | 115197751 | 19242409 | 4244976.0 | 2569.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 4574212 | 29630 | 17798.5 | 229.24 | 125673446 | 20992249 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 27265706 | 181456 | 1603865.1 | 38.46 | 21083518 | 3513865 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-at | 16241562 | 103299 | 955386.0 | 64.56 | 35394195 | 5898940 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 16531714 | 40765 | 972453.8 | 63.43 | 34772982 | 5795406 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 16497998 | 199268 | 970470.5 | 63.56 | 34844046 | 5807250 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 16506903 | 157583 | 970994.3 | 63.52 | 34825249 | 5804117 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 4540738 | 28981 | 267102.2 | 230.93 | 126599900 | 21099653 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 4916699 | 35159 | 289217.6 | 213.27 | 116919300 | 19486245 | 5852944.0 | 273.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 4557663 | 29714 | 268097.8 | 230.07 | 126129776 | 21021300 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 30842435 | 262249 | 30842435.1 | 34.00 | 18638509 | 3105916 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-at | 16213064 | 82001 | 16213063.9 | 64.67 | 35456408 | 5908445 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 16584436 | 179295 | 16584436.1 | 63.23 | 34662439 | 5776138 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 16503966 | 36716 | 16503965.7 | 63.53 | 34831447 | 5804302 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 16617700 | 113106 | 16617700.3 | 63.10 | 34593054 | 5764576 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 4519008 | 10706 | 4519008.4 | 232.04 | 127208658 | 21198013 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 5017099 | 24046 | 5017099.4 | 209.00 | 114579552 | 19093503 | 8942832.0 | 29.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 4569310 | 21236 | 4569309.7 | 229.48 | 125808282 | 20964655 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 77279890 | 694225 | 1731.8 | 21.03 | 10969607 | 2605335 | N/A | N/A | 10 |
| synthetic | 0 | onig-at | 35260686 | 147040 | 790.2 | 46.09 | 24041790 | 5710042 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 46593782 | 528809 | 1044.1 | 34.88 | 18194059 | 4321177 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 44470325 | 356299 | 996.6 | 36.55 | 19062825 | 4527514 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 45500451 | 486266 | 1019.6 | 35.72 | 18631244 | 4425011 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 10515043 | 60182 | 235.6 | 154.56 | 80620691 | 19147807 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 12669710 | 73085 | 283.9 | 128.28 | 66909977 | 15891445 | 8028528.0 | 109699.0 | 10 |
| synthetic | 0 | pretok-reused | 10621130 | 40013 | 238.0 | 153.02 | 79815423 | 18956551 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 22389782 | 77912 | 5443.7 | 46.83 | 17427727 | 2590512 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-at | 11979080 | 88332 | 2912.5 | 87.53 | 32573703 | 4841857 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 13605934 | 73415 | 3308.0 | 77.07 | 28678884 | 4262919 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 12985529 | 48474 | 3157.2 | 80.75 | 30049065 | 4466587 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 13049465 | 26577 | 3172.7 | 80.35 | 29901838 | 4444703 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 4043898 | 12227 | 983.2 | 259.30 | 96491804 | 14342845 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 4700418 | 8757 | 1142.8 | 223.08 | 83014538 | 12339542 | 2524272.0 | 20825.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 4075190 | 15679 | 990.8 | 257.31 | 95750880 | 14232712 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 19472015 | 92312 | 75766.6 | 53.85 | 20039169 | 2809622 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-at | 11558717 | 69939 | 44975.6 | 90.72 | 33758331 | 4733138 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 12040317 | 48832 | 46849.5 | 87.09 | 32408033 | 4543817 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 11934027 | 46767 | 46435.9 | 87.86 | 32696674 | 4584286 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 11920247 | 37197 | 46382.3 | 87.97 | 32734473 | 4589586 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 3966517 | 17514 | 15433.9 | 264.36 | 98374210 | 13792705 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 4258628 | 8076 | 16570.5 | 246.22 | 91626466 | 12846627 | 2430448.0 | 2339.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 3979524 | 20707 | 15484.5 | 263.49 | 98052670 | 13747622 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 19127343 | 105308 | 1125137.8 | 54.82 | 20400272 | 2849115 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-at | 11517317 | 87159 | 677489.2 | 91.04 | 33879678 | 4731657 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 11814621 | 114936 | 694977.7 | 88.75 | 33027128 | 4612590 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 11778815 | 53024 | 692871.5 | 89.02 | 33127526 | 4626611 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 11749442 | 60751 | 691143.7 | 89.24 | 33210342 | 4638178 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 3967865 | 14073 | 233403.8 | 264.27 | 98340794 | 13734338 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 4200525 | 21061 | 247089.7 | 249.63 | 92893871 | 12973617 | 3215152.0 | 246.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 3971420 | 11177 | 233613.0 | 264.03 | 98252754 | 13722042 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 21035805 | 678762 | 21035804.7 | 49.85 | 18549469 | 2590108 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-at | 11583412 | 71494 | 11583411.9 | 90.52 | 33686361 | 4703709 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 11742815 | 55088 | 11742815.4 | 89.29 | 33229084 | 4639858 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 11755740 | 82484 | 11755739.8 | 89.20 | 33192552 | 4634757 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 11764171 | 56408 | 11764171.0 | 89.13 | 33168763 | 4631436 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 3968922 | 8725 | 3968922.0 | 264.20 | 98314606 | 13727909 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 4212013 | 16661 | 4212013.3 | 248.95 | 92640496 | 12935619 | 4363504.0 | 26.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 3968577 | 18316 | 3968577.4 | 264.22 | 98323143 | 13729101 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.410x |
| ar-wikipedia | 4096 | 3.256x |
| ar-wikipedia | 65536 | 3.230x |
| ar-wikipedia | 1048576 | 3.237x |
| en-wikitext2 | 256 | 6.211x |
| en-wikitext2 | 4096 | 6.116x |
| en-wikitext2 | 65536 | 6.057x |
| en-wikitext2 | 1048576 | 6.051x |
| ja-wikipedia | 256 | 2.592x |
| ja-wikipedia | 4096 | 2.412x |
| ja-wikipedia | 65536 | 2.385x |
| ja-wikipedia | 1048576 | 2.382x |
| ru-wikipedia | 256 | 3.847x |
| ru-wikipedia | 4096 | 3.686x |
| ru-wikipedia | 65536 | 3.633x |
| ru-wikipedia | 1048576 | 3.652x |
| synthetic | 0 | 4.229x |
| zh-wikipedia | 256 | 3.211x |
| zh-wikipedia | 4096 | 3.009x |
| zh-wikipedia | 65536 | 2.969x |
| zh-wikipedia | 1048576 | 2.962x |
| Combined selected natural corpora | 256 | 3.810x |
| Combined selected natural corpora | 4096 | 3.648x |
| Combined selected natural corpora | 65536 | 3.609x |
| Combined selected natural corpora | 1048576 | 3.613x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.397x |
| ar-wikipedia | 4096 | 3.255x |
| ar-wikipedia | 65536 | 3.221x |
| ar-wikipedia | 1048576 | 3.235x |
| en-wikitext2 | 256 | 6.175x |
| en-wikitext2 | 4096 | 6.082x |
| en-wikitext2 | 65536 | 5.986x |
| en-wikitext2 | 1048576 | 5.974x |
| ja-wikipedia | 256 | 2.585x |
| ja-wikipedia | 4096 | 2.423x |
| ja-wikipedia | 65536 | 2.380x |
| ja-wikipedia | 1048576 | 2.372x |
| ru-wikipedia | 256 | 3.827x |
| ru-wikipedia | 4096 | 3.679x |
| ru-wikipedia | 65536 | 3.622x |
| ru-wikipedia | 1048576 | 3.637x |
| synthetic | 0 | 4.284x |
| zh-wikipedia | 256 | 3.202x |
| zh-wikipedia | 4096 | 2.995x |
| zh-wikipedia | 65536 | 2.958x |
| zh-wikipedia | 1048576 | 2.964x |
| Combined selected natural corpora | 256 | 3.797x |
| Combined selected natural corpora | 4096 | 3.644x |
| Combined selected natural corpora | 65536 | 3.594x |
| Combined selected natural corpora | 1048576 | 3.601x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.089x |
| ar-wikipedia | 4096 | 3.067x |
| ar-wikipedia | 65536 | 3.003x |
| ar-wikipedia | 1048576 | 2.964x |
| en-wikitext2 | 256 | 5.086x |
| en-wikitext2 | 4096 | 5.336x |
| en-wikitext2 | 65536 | 5.151x |
| en-wikitext2 | 1048576 | 5.166x |
| ja-wikipedia | 256 | 2.393x |
| ja-wikipedia | 4096 | 2.310x |
| ja-wikipedia | 65536 | 2.299x |
| ja-wikipedia | 1048576 | 2.264x |
| ru-wikipedia | 256 | 3.421x |
| ru-wikipedia | 4096 | 3.399x |
| ru-wikipedia | 65536 | 3.362x |
| ru-wikipedia | 1048576 | 3.306x |
| synthetic | 0 | 3.678x |
| zh-wikipedia | 256 | 2.895x |
| zh-wikipedia | 4096 | 2.827x |
| zh-wikipedia | 65536 | 2.813x |
| zh-wikipedia | 1048576 | 2.788x |
| Combined selected natural corpora | 256 | 3.381x |
| Combined selected natural corpora | 4096 | 3.384x |
| Combined selected natural corpora | 65536 | 3.330x |
| Combined selected natural corpora | 1048576 | 3.300x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 5.835x |
| ar-wikipedia | 4096 | 5.244x |
| ar-wikipedia | 65536 | 5.165x |
| ar-wikipedia | 1048576 | 5.762x |
| en-wikitext2 | 256 | 10.684x |
| en-wikitext2 | 4096 | 10.189x |
| en-wikitext2 | 65536 | 10.175x |
| en-wikitext2 | 1048576 | 11.185x |
| ja-wikipedia | 256 | 4.266x |
| ja-wikipedia | 4096 | 3.745x |
| ja-wikipedia | 65536 | 3.703x |
| ja-wikipedia | 1048576 | 4.242x |
| ru-wikipedia | 256 | 6.799x |
| ru-wikipedia | 4096 | 6.079x |
| ru-wikipedia | 65536 | 5.982x |
| ru-wikipedia | 1048576 | 6.750x |
| synthetic | 0 | 7.276x |
| zh-wikipedia | 256 | 5.494x |
| zh-wikipedia | 4096 | 4.893x |
| zh-wikipedia | 65536 | 4.816x |
| zh-wikipedia | 1048576 | 5.301x |
| Combined selected natural corpora | 256 | 6.541x |
| Combined selected natural corpora | 4096 | 5.950x |
| Combined selected natural corpora | 65536 | 5.892x |
| Combined selected natural corpora | 1048576 | 6.573x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 5.077x |
| ar-wikipedia | 4096 | 4.889x |
| ar-wikipedia | 65536 | 4.798x |
| ar-wikipedia | 1048576 | 5.295x |
| en-wikitext2 | 256 | 8.568x |
| en-wikitext2 | 4096 | 8.852x |
| en-wikitext2 | 65536 | 8.766x |
| en-wikitext2 | 1048576 | 9.657x |
| ja-wikipedia | 256 | 3.808x |
| ja-wikipedia | 4096 | 3.548x |
| ja-wikipedia | 65536 | 3.564x |
| ja-wikipedia | 1048576 | 4.054x |
| ru-wikipedia | 256 | 5.812x |
| ru-wikipedia | 4096 | 5.573x |
| ru-wikipedia | 65536 | 5.546x |
| ru-wikipedia | 1048576 | 6.147x |
| synthetic | 0 | 6.100x |
| zh-wikipedia | 256 | 4.763x |
| zh-wikipedia | 4096 | 4.572x |
| zh-wikipedia | 65536 | 4.554x |
| zh-wikipedia | 1048576 | 4.994x |
| Combined selected natural corpora | 256 | 5.609x |
| Combined selected natural corpora | 4096 | 5.475x |
| Combined selected natural corpora | 65536 | 5.447x |
| Combined selected natural corpora | 1048576 | 6.031x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 58887.9 | 184.1 | 10 |
| onig-compile | 58855.4 | 211.8 | 10 |
| pretok-compile | 11.6 | 0.1 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
