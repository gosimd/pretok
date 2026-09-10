# P23 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 28416931 | 214455 | 6924.2 | 36.90 | 20571363 | 3950075 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 14969319 | 56560 | 3647.5 | 70.05 | 39051544 | 7498605 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 14211472 | 116891 | 3462.8 | 73.78 | 41134023 | 7005890 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 14267151 | 121779 | 3476.4 | 73.50 | 40973492 | 7867653 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 5301360 | 43435 | 1291.8 | 197.79 | 110268882 | 18780842 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 6172430 | 37894 | 1504.0 | 169.88 | 94707427 | 18185543 | 4446624.0 | 24922.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 5355572 | 42671 | 1305.0 | 195.79 | 109152682 | 20959294 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 24720425 | 192920 | 96188.4 | 42.42 | 23647450 | 4406154 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 13582035 | 62445 | 52848.4 | 77.20 | 43040310 | 8019564 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 13390781 | 78983 | 52104.2 | 78.31 | 43655035 | 7231318 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 13494234 | 91930 | 52506.7 | 77.71 | 43320354 | 8071744 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 5274089 | 56757 | 20521.7 | 198.82 | 110839042 | 18360137 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 5732590 | 50765 | 22305.8 | 182.91 | 101973982 | 19000488 | 4231408.0 | 2568.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 5323652 | 72744 | 20714.6 | 196.97 | 109807139 | 20460015 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 24166547 | 515710 | 1421561.6 | 43.39 | 24189431 | 4498284 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 13195988 | 97247 | 776234.6 | 79.46 | 44299448 | 8237958 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 13144890 | 167894 | 773228.8 | 79.77 | 44471654 | 7353352 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 13038252 | 53202 | 766956.0 | 80.42 | 44835383 | 8337621 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 5250829 | 40505 | 308872.3 | 199.70 | 111330035 | 18408331 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 5763023 | 56095 | 339001.4 | 181.95 | 101435479 | 18863017 | 7663408.0 | 287.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 5305108 | 30178 | 312065.2 | 197.65 | 110190973 | 20491195 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 28018856 | 481270 | 28018856.5 | 37.42 | 20863628 | 3879245 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 13240386 | 122145 | 13240386.0 | 79.20 | 44150903 | 8209126 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 13180212 | 75005 | 13180211.6 | 79.56 | 44352474 | 7332735 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 13213117 | 109818 | 13213116.8 | 79.36 | 44242022 | 8226068 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 5225912 | 55214 | 5225912.1 | 200.65 | 111860856 | 18493805 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 5811677 | 26942 | 5811677.4 | 180.43 | 100586278 | 18702346 | 8942832.0 | 29.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 5292051 | 48985 | 5292051.1 | 198.14 | 110462841 | 20538728 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 41802915 | 446976 | 10203.3 | 25.08 | 25032632 | 4796388 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 22206722 | 306675 | 5420.2 | 47.22 | 47122534 | 9028933 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 21088282 | 143168 | 5147.2 | 49.72 | 49621728 | 8494149 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 21199326 | 124716 | 5174.4 | 49.46 | 49361806 | 9457990 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 3406538 | 17807 | 831.5 | 307.81 | 307184853 | 52583291 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 4695872 | 45268 | 1146.2 | 223.30 | 222841869 | 42697710 | 8419440.0 | 28722.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 3462568 | 14761 | 845.1 | 302.83 | 302214178 | 57905874 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 38405156 | 745114 | 150020.1 | 27.30 | 27247305 | 5141757 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 20704822 | 341129 | 80878.2 | 50.64 | 50540740 | 9537392 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 20275868 | 104595 | 79202.6 | 51.72 | 51609972 | 8720416 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 20426788 | 198975 | 79792.1 | 51.33 | 51228662 | 9667208 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 3339269 | 11848 | 13044.0 | 314.01 | 313373113 | 52949918 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 4106852 | 46201 | 16042.4 | 255.32 | 254802719 | 48083060 | 7847936.0 | 2825.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 3401607 | 13573 | 13287.5 | 308.26 | 307630231 | 58051982 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 37726222 | 336434 | 2357888.9 | 27.79 | 27737657 | 5229334 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 20190657 | 184872 | 1261916.0 | 51.93 | 51827784 | 9771005 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 20058631 | 204602 | 1253664.4 | 52.28 | 52168915 | 8807929 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 20128366 | 254897 | 1258022.9 | 52.09 | 51988175 | 9801243 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 3328033 | 14065 | 208002.1 | 315.07 | 314431079 | 53086914 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 4328756 | 75970 | 270547.3 | 242.23 | 241740798 | 45574984 | 14634752.0 | 320.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 3407751 | 15605 | 212984.5 | 307.70 | 307075521 | 57892429 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 44065375 | 1063612 | 44065375.0 | 23.80 | 23747375 | 4476735 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 20367104 | 130594 | 20367104.2 | 51.48 | 51378782 | 9685668 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 20377725 | 280288 | 20377725.0 | 51.46 | 51352003 | 8669466 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 20377629 | 343445 | 20377629.2 | 51.46 | 51352245 | 9680665 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 3354636 | 24461 | 3354635.6 | 312.58 | 311937611 | 52662651 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 4306729 | 17889 | 4306729.0 | 243.47 | 242977211 | 45804832 | 17962224.0 | 32.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 3440008 | 38352 | 3440008.3 | 304.82 | 304196071 | 57345502 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 14531780 | 146775 | 3533.1 | 72.16 | 26100037 | 3327879 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 6705313 | 50483 | 1630.3 | 156.38 | 56564098 | 7212191 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 6232032 | 43811 | 1515.2 | 168.26 | 60859763 | 4411884 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 6271176 | 22380 | 1524.7 | 167.21 | 60479889 | 7711473 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 3440838 | 7782 | 836.6 | 304.74 | 110228947 | 7990785 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 3969012 | 25346 | 965.0 | 264.19 | 95560295 | 12184391 | 2120944.0 | 20069.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 3457312 | 9355 | 840.6 | 303.29 | 109703726 | 13987746 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 11653734 | 199575 | 45345.3 | 89.98 | 32545792 | 3855245 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 5379205 | 48430 | 20930.8 | 194.93 | 70508558 | 8352163 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 5282447 | 33566 | 20554.3 | 198.50 | 71800053 | 4563036 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 5245862 | 55398 | 20411.9 | 199.89 | 72300803 | 8564465 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 3333138 | 62530 | 12969.4 | 314.59 | 113790680 | 7231625 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 3573669 | 41327 | 13905.3 | 293.42 | 106131824 | 12571954 | 2147824.0 | 2294.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 3344867 | 52781 | 13015.0 | 313.49 | 113391637 | 13431922 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 11170731 | 264285 | 657101.8 | 93.87 | 33953017 | 4002514 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 5090791 | 61167 | 299458.3 | 205.97 | 74503150 | 8782721 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 5056286 | 44799 | 297428.6 | 207.38 | 75011579 | 4725603 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 5047309 | 111632 | 296900.5 | 207.75 | 75144994 | 8858384 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 3346150 | 236804 | 196832.4 | 313.37 | 113348170 | 7140743 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 3476906 | 30756 | 204523.9 | 301.58 | 109085506 | 12859423 | 2436880.0 | 234.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 3311941 | 137201 | 194820.0 | 316.60 | 114518954 | 13499939 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 14186992 | 420635 | 14186992.5 | 73.91 | 26734348 | 3150492 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 5167627 | 145647 | 5167626.6 | 202.91 | 73395395 | 8649232 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 5098253 | 38552 | 5098253.4 | 205.67 | 74394106 | 4683957 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 5122347 | 38554 | 5122347.1 | 204.71 | 74044181 | 8725687 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 3277987 | 29367 | 3277987.2 | 319.88 | 115705150 | 7284958 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 3491053 | 13578 | 3491053.4 | 300.36 | 108643426 | 12803012 | 3413232.0 | 25.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 3285318 | 18300 | 3285317.8 | 319.17 | 115446974 | 13604772 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 25846476 | 317619 | 6297.9 | 40.57 | 22241214 | 3649356 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 12893305 | 278081 | 3141.6 | 81.33 | 44585698 | 7315657 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 12236140 | 119669 | 2981.5 | 85.69 | 46980256 | 6438060 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 12264152 | 187441 | 2988.3 | 85.50 | 46872950 | 7690951 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 5177744 | 23681 | 1261.6 | 202.52 | 111024613 | 15214542 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 5991296 | 89386 | 1459.9 | 175.02 | 95948691 | 15743339 | 4201344.0 | 24621.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 5233005 | 29237 | 1275.1 | 200.38 | 109852174 | 18024633 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 22371034 | 639170 | 87046.8 | 46.87 | 25696488 | 4063156 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 11651881 | 153804 | 45338.1 | 89.99 | 49335984 | 7801058 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 11276059 | 152755 | 43875.7 | 92.99 | 50980312 | 6725488 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 11378132 | 66328 | 44272.9 | 92.16 | 50522970 | 7988746 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 5070929 | 40850 | 19731.2 | 206.78 | 113363250 | 14955248 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 5547993 | 49061 | 21587.5 | 189.00 | 103615305 | 16383762 | 4204272.0 | 2566.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 5148883 | 19310 | 20034.6 | 203.65 | 111646932 | 17653732 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 21672810 | 326426 | 1274871.2 | 48.38 | 26524341 | 4183998 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 11297044 | 327353 | 664532.0 | 92.82 | 50885614 | 8026790 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 11176237 | 185307 | 657425.7 | 93.82 | 51435651 | 6768736 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 11258163 | 143617 | 662244.9 | 93.14 | 51061350 | 8054511 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 5078132 | 15518 | 298713.6 | 206.49 | 113202464 | 14897015 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 5505659 | 138134 | 323862.3 | 190.45 | 104412031 | 16470146 | 5852944.0 | 273.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 5131970 | 38512 | 301880.6 | 204.32 | 112014883 | 17669434 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 25901211 | 493039 | 25901211.5 | 40.48 | 22194213 | 3500415 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 11428893 | 252618 | 11428893.5 | 91.75 | 50298570 | 7932964 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 11377069 | 96327 | 11377069.1 | 92.17 | 50527688 | 6648373 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 11381098 | 120628 | 11381098.0 | 92.13 | 50509801 | 7966279 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 5118085 | 26423 | 5118084.8 | 204.88 | 112318772 | 14778770 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 5629365 | 64791 | 5629365.4 | 186.27 | 102117549 | 16105723 | 7058672.0 | 28.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 5160659 | 14258 | 5160658.9 | 203.19 | 111392171 | 17568493 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 66800159 | 240005 | 916.7 | 34.25 | 18657381 | 2846176 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 32326234 | 200281 | 443.6 | 70.77 | 38554320 | 5881446 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 30494317 | 285901 | 418.5 | 75.03 | 40870434 | 4030030 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 30885440 | 215893 | 423.8 | 74.08 | 40352865 | 6155813 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 8514605 | 14155 | 116.8 | 268.70 | 146373915 | 14433201 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 10906435 | 50584 | 149.7 | 209.77 | 114273456 | 17432369 | 6457376.0 | 140412.0 | 10 |
| synthetic | 0 | pretok-reused | 8693656 | 15897 | 119.3 | 263.16 | 143359251 | 21869396 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 16289008 | 213145 | 3960.4 | 64.37 | 23954988 | 3854133 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 7557329 | 72292 | 1837.4 | 138.75 | 51632396 | 8307168 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 6943695 | 30012 | 1688.2 | 151.01 | 56195296 | 4960039 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 7000358 | 53374 | 1702.0 | 149.79 | 55740438 | 8968113 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 3478353 | 5927 | 845.7 | 301.46 | 112180395 | 9901526 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 4065211 | 9569 | 988.4 | 257.94 | 95985926 | 15443235 | 2776560.0 | 21763.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 3500705 | 3665 | 851.1 | 299.53 | 111464134 | 17933533 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 13692633 | 235431 | 53278.7 | 76.58 | 28497295 | 4347082 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 6288841 | 74650 | 24470.2 | 166.74 | 62046890 | 9464861 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 6067079 | 78971 | 23607.3 | 172.83 | 64314804 | 5148606 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 6115195 | 38992 | 23794.5 | 171.47 | 63808754 | 9733622 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 3424235 | 13327 | 13323.9 | 306.22 | 113953342 | 9122330 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 3701078 | 14674 | 14401.1 | 283.32 | 105429558 | 16082612 | 2604016.0 | 2370.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 3426230 | 39985 | 13331.6 | 306.04 | 113886969 | 17372737 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 13162052 | 57546 | 774238.3 | 79.67 | 29646062 | 4506744 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 6029977 | 39961 | 354704.5 | 173.89 | 64710533 | 9837186 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 5976655 | 62213 | 351568.0 | 175.44 | 65287855 | 5192369 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 5975342 | 50645 | 351490.7 | 175.48 | 65302201 | 9927130 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 3415538 | 6701 | 200914.0 | 307.00 | 114243491 | 9085830 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 3657692 | 4167 | 215158.3 | 286.68 | 106680121 | 16217332 | 3575664.0 | 252.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 3422143 | 7842 | 201302.5 | 306.41 | 114022987 | 17333582 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 16182081 | 290122 | 16182081.3 | 64.80 | 24113276 | 3664980 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 5978764 | 57343 | 5978764.1 | 175.38 | 65264826 | 9919609 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 5969278 | 42930 | 5969277.8 | 175.66 | 65368544 | 5196944 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 5979159 | 37452 | 5979159.0 | 175.37 | 65260516 | 9918954 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 3413914 | 5991 | 3413914.0 | 307.15 | 114297842 | 9086931 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 3664317 | 11152 | 3664317.4 | 286.16 | 106487228 | 16185006 | 4363504.0 | 26.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 3418968 | 7137 | 3418967.7 | 306.69 | 114128894 | 17346464 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.681x |
| ar-wikipedia | 4096 | 2.539x |
| ar-wikipedia | 65536 | 2.503x |
| ar-wikipedia | 1048576 | 2.522x |
| en-wikitext2 | 256 | 6.191x |
| en-wikitext2 | 4096 | 6.072x |
| en-wikitext2 | 65536 | 6.027x |
| en-wikitext2 | 1048576 | 6.074x |
| ja-wikipedia | 256 | 1.811x |
| ja-wikipedia | 4096 | 1.585x |
| ja-wikipedia | 65536 | 1.511x |
| ja-wikipedia | 1048576 | 1.555x |
| ru-wikipedia | 256 | 2.363x |
| ru-wikipedia | 4096 | 2.224x |
| ru-wikipedia | 65536 | 2.201x |
| ru-wikipedia | 1048576 | 2.223x |
| synthetic | 0 | 3.581x |
| zh-wikipedia | 256 | 1.996x |
| zh-wikipedia | 4096 | 1.772x |
| zh-wikipedia | 65536 | 1.750x |
| zh-wikipedia | 1048576 | 1.749x |
| Combined selected natural corpora | 256 | 2.918x |
| Combined selected natural corpora | 4096 | 2.754x |
| Combined selected natural corpora | 65536 | 2.714x |
| Combined selected natural corpora | 1048576 | 2.746x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.664x |
| ar-wikipedia | 4096 | 2.535x |
| ar-wikipedia | 65536 | 2.458x |
| ar-wikipedia | 1048576 | 2.497x |
| en-wikitext2 | 256 | 6.122x |
| en-wikitext2 | 4096 | 6.005x |
| en-wikitext2 | 65536 | 5.907x |
| en-wikitext2 | 1048576 | 5.924x |
| ja-wikipedia | 256 | 1.814x |
| ja-wikipedia | 4096 | 1.568x |
| ja-wikipedia | 65536 | 1.524x |
| ja-wikipedia | 1048576 | 1.559x |
| ru-wikipedia | 256 | 2.344x |
| ru-wikipedia | 4096 | 2.210x |
| ru-wikipedia | 65536 | 2.194x |
| ru-wikipedia | 1048576 | 2.205x |
| synthetic | 0 | 3.553x |
| zh-wikipedia | 256 | 2.000x |
| zh-wikipedia | 4096 | 1.785x |
| zh-wikipedia | 65536 | 1.746x |
| zh-wikipedia | 1048576 | 1.749x |
| Combined selected natural corpora | 256 | 2.904x |
| Combined selected natural corpora | 4096 | 2.744x |
| Combined selected natural corpora | 65536 | 2.694x |
| Combined selected natural corpora | 1048576 | 2.722x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.425x |
| ar-wikipedia | 4096 | 2.369x |
| ar-wikipedia | 65536 | 2.290x |
| ar-wikipedia | 1048576 | 2.278x |
| en-wikitext2 | 256 | 4.729x |
| en-wikitext2 | 4096 | 5.042x |
| en-wikitext2 | 65536 | 4.664x |
| en-wikitext2 | 1048576 | 4.729x |
| ja-wikipedia | 256 | 1.689x |
| ja-wikipedia | 4096 | 1.505x |
| ja-wikipedia | 65536 | 1.464x |
| ja-wikipedia | 1048576 | 1.480x |
| ru-wikipedia | 256 | 2.152x |
| ru-wikipedia | 4096 | 2.100x |
| ru-wikipedia | 65536 | 2.052x |
| ru-wikipedia | 1048576 | 2.030x |
| synthetic | 0 | 2.964x |
| zh-wikipedia | 256 | 1.859x |
| zh-wikipedia | 4096 | 1.699x |
| zh-wikipedia | 65536 | 1.649x |
| zh-wikipedia | 1048576 | 1.632x |
| Combined selected natural corpora | 256 | 2.584x |
| Combined selected natural corpora | 4096 | 2.542x |
| Combined selected natural corpora | 65536 | 2.455x |
| Combined selected natural corpora | 1048576 | 2.453x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 5.306x |
| ar-wikipedia | 4096 | 4.644x |
| ar-wikipedia | 65536 | 4.555x |
| ar-wikipedia | 1048576 | 5.295x |
| en-wikitext2 | 256 | 12.073x |
| en-wikitext2 | 4096 | 11.290x |
| en-wikitext2 | 65536 | 11.071x |
| en-wikitext2 | 1048576 | 12.810x |
| ja-wikipedia | 256 | 4.203x |
| ja-wikipedia | 4096 | 3.484x |
| ja-wikipedia | 65536 | 3.373x |
| ja-wikipedia | 1048576 | 4.318x |
| ru-wikipedia | 256 | 4.939x |
| ru-wikipedia | 4096 | 4.345x |
| ru-wikipedia | 65536 | 4.223x |
| ru-wikipedia | 1048576 | 5.019x |
| synthetic | 0 | 7.684x |
| zh-wikipedia | 256 | 4.653x |
| zh-wikipedia | 4096 | 3.996x |
| zh-wikipedia | 65536 | 3.846x |
| zh-wikipedia | 1048576 | 4.733x |
| Combined selected natural corpora | 256 | 6.040x |
| Combined selected natural corpora | 4096 | 5.369x |
| Combined selected natural corpora | 65536 | 5.243x |
| Combined selected natural corpora | 1048576 | 6.232x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 4.604x |
| ar-wikipedia | 4096 | 4.312x |
| ar-wikipedia | 65536 | 4.193x |
| ar-wikipedia | 1048576 | 4.821x |
| en-wikitext2 | 256 | 8.902x |
| en-wikitext2 | 4096 | 9.351x |
| en-wikitext2 | 65536 | 8.715x |
| en-wikitext2 | 1048576 | 10.232x |
| ja-wikipedia | 256 | 3.661x |
| ja-wikipedia | 4096 | 3.261x |
| ja-wikipedia | 65536 | 3.213x |
| ja-wikipedia | 1048576 | 4.064x |
| ru-wikipedia | 256 | 4.314x |
| ru-wikipedia | 4096 | 4.032x |
| ru-wikipedia | 65536 | 3.936x |
| ru-wikipedia | 1048576 | 4.601x |
| synthetic | 0 | 6.125x |
| zh-wikipedia | 256 | 4.007x |
| zh-wikipedia | 4096 | 3.700x |
| zh-wikipedia | 65536 | 3.598x |
| zh-wikipedia | 1048576 | 4.416x |
| Combined selected natural corpora | 256 | 5.097x |
| Combined selected natural corpora | 4096 | 4.891x |
| Combined selected natural corpora | 65536 | 4.747x |
| Combined selected natural corpora | 1048576 | 5.604x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 1016.7 | 6.5 | 10 |
| onig-compile | 992.7 | 4.7 | 10 |
| pretok-compile | 10.6 | 0.4 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
