# P21 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 9778088 | 28446 | 2382.6 | 107.24 | 59784182 | 419918 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 6202516 | 38148 | 1511.3 | 169.06 | 94248045 | 661989 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 6103735 | 26850 | 1487.3 | 171.79 | 95773326 | 164 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 6098690 | 26734 | 1486.0 | 171.93 | 95852546 | 673259 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 4153290 | 26466 | 1012.0 | 252.47 | 140749877 | 241 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 4204624 | 48405 | 1024.5 | 249.39 | 139031453 | 976544 | 65760.0 | 4106.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 4143570 | 31288 | 1009.6 | 253.06 | 141080044 | 990933 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 8009340 | 29827 | 31164.7 | 130.92 | 72986661 | 32337 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 5604526 | 12756 | 21807.5 | 187.09 | 104304090 | 46213 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 5588690 | 27476 | 21745.9 | 187.62 | 104599646 | 179 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 5575186 | 36685 | 21693.3 | 188.08 | 104853011 | 46456 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 4114089 | 67033 | 16008.1 | 254.87 | 142090984 | 243 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 4139575 | 45950 | 16107.3 | 253.31 | 141216197 | 62567 | 4208.0 | 259.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 4137788 | 62035 | 16100.3 | 253.41 | 141277162 | 62594 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 7902435 | 132325 | 464849.1 | 132.69 | 73974039 | 2404 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 5664527 | 99015 | 333207.5 | 185.11 | 103199254 | 3354 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 5645860 | 197513 | 332109.4 | 185.72 | 103540479 | 177 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 5626978 | 200452 | 330998.7 | 186.35 | 103887922 | 3377 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 4212528 | 83237 | 247795.8 | 248.92 | 138770596 | 237 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 4182537 | 91070 | 246031.6 | 250.70 | 139765639 | 4543 | 368.0 | 19.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 4201242 | 68748 | 247131.9 | 249.59 | 139143392 | 4522 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 9943235 | 88897 | 9943235.3 | 105.46 | 58791227 | 302 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 5644747 | 94771 | 5644746.7 | 185.76 | 103560891 | 531 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 5626007 | 111493 | 5626006.8 | 186.38 | 103905847 | 178 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 5602586 | 56310 | 5602586.1 | 187.16 | 104340208 | 535 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 4181401 | 23375 | 4181401.0 | 250.77 | 139803619 | 239 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 4149933 | 36497 | 4149933.2 | 252.67 | 140863713 | 723 | 112.0 | 3.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 4172242 | 60515 | 4172242.0 | 251.32 | 140110522 | 719 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 10756369 | 51419 | 2625.4 | 97.48 | 97285340 | 382750 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 8142751 | 30489 | 1987.5 | 128.77 | 128511477 | 505603 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 8077444 | 55750 | 1971.6 | 129.82 | 129550513 | 1238 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 8082509 | 27375 | 1972.8 | 129.73 | 129469320 | 509372 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 2274973 | 11308 | 555.3 | 460.92 | 459977669 | 4396 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 2325053 | 91154 | 567.5 | 450.99 | 450070077 | 1770712 | 66608.0 | 4114.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 2281990 | 63615 | 557.0 | 459.50 | 458563424 | 1804127 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 9735394 | 48006 | 38028.9 | 107.71 | 107487896 | 28350 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 7597140 | 10304 | 29676.3 | 138.02 | 137740910 | 36329 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 7593556 | 9439 | 29662.3 | 138.09 | 137805926 | 1317 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 7590893 | 11055 | 29651.9 | 138.14 | 137854268 | 36359 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 2394067 | 15042 | 9351.8 | 437.99 | 437095994 | 4177 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 2397433 | 53517 | 9365.0 | 437.37 | 436482269 | 115123 | 5248.0 | 268.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 2391296 | 5023 | 9341.0 | 438.50 | 437602431 | 115419 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 9616503 | 92646 | 601031.5 | 109.04 | 108816789 | 3744 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 7516831 | 126249 | 469802.0 | 139.50 | 139212513 | 4789 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 7570020 | 113933 | 473126.3 | 138.52 | 138234373 | 1321 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 7509329 | 128574 | 469333.1 | 139.64 | 139351599 | 4794 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 2270831 | 58455 | 141926.9 | 461.76 | 460816745 | 4404 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 2270973 | 68453 | 141935.8 | 461.73 | 460788003 | 15852 | 1248.0 | 21.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 2281949 | 60122 | 142621.8 | 459.51 | 458571673 | 15776 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 11334774 | 75215 | 11334773.6 | 92.51 | 92320944 | 1853 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 7704137 | 87351 | 7704137.3 | 136.11 | 135827928 | 2726 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 7639439 | 87665 | 7639438.8 | 137.26 | 136978256 | 1309 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 7681836 | 86725 | 7681835.7 | 136.50 | 136222258 | 2734 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 2326447 | 45307 | 2326446.8 | 450.72 | 449800532 | 4298 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 2321184 | 54510 | 2321183.7 | 451.74 | 450820408 | 9047 | 1008.0 | 6.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 2323943 | 46786 | 2323942.7 | 451.21 | 450285197 | 9036 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 15593033 | 33424 | 3791.2 | 67.25 | 24323683 | 3980752 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 6945197 | 20039 | 1688.6 | 150.98 | 54610402 | 8937399 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 6304327 | 52773 | 1532.8 | 166.33 | 60161851 | 5168672 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 6353056 | 45452 | 1544.6 | 165.05 | 59700401 | 9770416 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 3136198 | 20187 | 762.5 | 334.35 | 120936235 | 10389968 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 3640596 | 26294 | 885.1 | 288.02 | 104180757 | 17049958 | 2772976.0 | 21650.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 3092866 | 26644 | 752.0 | 339.03 | 122630601 | 20069412 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 12981474 | 62472 | 50511.6 | 80.77 | 29217020 | 4531458 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 5616006 | 27364 | 21852.2 | 186.71 | 67535543 | 10474526 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 5461360 | 28403 | 21250.4 | 192.00 | 69447899 | 5403416 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 5481031 | 12017 | 21327.0 | 191.31 | 69198657 | 10732469 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 3047696 | 11589 | 11858.7 | 344.06 | 124448124 | 9682725 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 3290119 | 10508 | 12802.0 | 318.70 | 115278512 | 17879294 | 2598384.0 | 2368.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 3013811 | 6786 | 11726.9 | 347.92 | 125847327 | 19518480 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 12550701 | 137250 | 738276.5 | 83.55 | 30219826 | 4670337 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 5446072 | 34418 | 320357.2 | 192.54 | 69642852 | 10762986 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 5360824 | 28604 | 315342.6 | 195.60 | 70750317 | 5468376 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 5398947 | 29676 | 317585.1 | 194.22 | 70250740 | 10856933 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 3034830 | 8985 | 178519.4 | 345.51 | 124975697 | 9659519 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 3228844 | 10051 | 189932.0 | 324.75 | 117466181 | 18153864 | 3493648.0 | 249.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 3004998 | 5862 | 176764.6 | 348.94 | 126216400 | 19506171 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 15277985 | 102281 | 15277984.9 | 68.63 | 24825264 | 3835650 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 5404888 | 36222 | 5404888.0 | 194.00 | 70173517 | 10842223 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 5362533 | 19073 | 5362533.0 | 195.54 | 70727770 | 5464022 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 5371849 | 26931 | 5371848.8 | 195.20 | 70605114 | 10908907 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 3054896 | 13521 | 3054896.0 | 343.24 | 124154801 | 9591489 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 3254952 | 36101 | 3254951.7 | 322.15 | 116524002 | 18003647 | 4363504.0 | 26.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 3014934 | 12090 | 3014933.9 | 347.79 | 125800438 | 19436911 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 9482271 | 50864 | 2310.5 | 110.58 | 60624402 | 432808 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 5995726 | 26296 | 1460.9 | 174.89 | 95877802 | 684488 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 5923035 | 27127 | 1443.2 | 177.03 | 97054473 | 0 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 5926946 | 41243 | 1444.2 | 176.92 | 96990418 | 692431 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 4021053 | 36486 | 979.8 | 260.77 | 142961800 | 0 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 4079801 | 66126 | 994.1 | 257.02 | 140903202 | 1005931 | 65664.0 | 4104.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 4026875 | 18073 | 981.2 | 260.39 | 142755102 | 1019152 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 7832894 | 44907 | 30478.2 | 133.87 | 73390114 | 32810 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 5387233 | 37733 | 20962.0 | 194.64 | 106707289 | 47705 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 5382469 | 46108 | 20943.5 | 194.81 | 106801737 | 0 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 5390582 | 53380 | 20975.0 | 194.52 | 106640994 | 47676 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 3941175 | 18865 | 15335.3 | 266.06 | 145859280 | 0 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 3940030 | 16466 | 15330.9 | 266.13 | 145901692 | 65228 | 4112.0 | 257.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 3940983 | 19131 | 15334.6 | 266.07 | 145866408 | 65212 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 7570262 | 40668 | 445309.5 | 138.51 | 75936208 | 2246 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 5371604 | 30609 | 315976.7 | 195.21 | 107017747 | 3165 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 5375172 | 27196 | 316186.6 | 195.08 | 106946716 | 0 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 5361070 | 30337 | 315357.1 | 195.59 | 107228025 | 3171 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 3963935 | 10584 | 233172.6 | 264.53 | 145021820 | 0 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 3955188 | 35316 | 232658.1 | 265.11 | 145342535 | 4298 | 272.0 | 17.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 3965754 | 21458 | 233279.7 | 264.41 | 144955275 | 4287 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 9226251 | 36242 | 9226251.1 | 113.65 | 62306672 | 108 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 5336879 | 16548 | 5336878.5 | 196.48 | 107714087 | 187 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 5342509 | 32329 | 5342509.3 | 196.27 | 107600561 | 0 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 5352726 | 61741 | 5352726.5 | 195.90 | 107395175 | 187 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 3949777 | 22186 | 3949776.7 | 265.48 | 145541644 | 0 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 3958920 | 26948 | 3958920.4 | 264.86 | 145205494 | 253 | 16.0 | 1.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 3960632 | 10999 | 3960631.7 | 264.75 | 145142755 | 252 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 54023369 | 1697344 | 741.4 | 42.35 | 23069942 | 1981383 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 24683157 | 271732 | 338.7 | 92.69 | 50492569 | 4336601 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 23176096 | 478602 | 318.0 | 98.72 | 53775926 | 917972 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 23585422 | 386494 | 323.7 | 97.00 | 52842641 | 4538439 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 7181626 | 34640 | 98.6 | 318.57 | 173542325 | 2962421 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 8812334 | 114890 | 120.9 | 259.62 | 141428598 | 12146726 | 2660448.0 | 96971.0 | 10 |
| synthetic | 0 | pretok-reused | 7303309 | 59705 | 100.2 | 313.26 | 170650861 | 14656507 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 18372318 | 175435 | 4466.9 | 57.07 | 21238638 | 4064974 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 8321644 | 115814 | 2023.3 | 126.01 | 46890135 | 8974549 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 7454074 | 74106 | 1812.3 | 140.67 | 52347613 | 5193133 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 7518769 | 40352 | 1828.0 | 139.46 | 51897192 | 9932876 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 2475128 | 13338 | 601.8 | 423.64 | 157649634 | 15639596 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 3166934 | 22901 | 770.0 | 331.10 | 123211587 | 23582112 | 3344880.0 | 22879.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 2488577 | 9123 | 605.1 | 421.35 | 156797656 | 30010326 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 15462077 | 147343 | 60163.7 | 67.82 | 25236131 | 4630684 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 6850152 | 89641 | 26654.3 | 153.07 | 56962676 | 10452322 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 6605570 | 64631 | 25702.6 | 158.74 | 59071812 | 5434202 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 6685675 | 18637 | 26014.3 | 156.84 | 58364042 | 10709465 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 2417775 | 10090 | 9407.7 | 433.69 | 161389323 | 14846711 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 2776255 | 20857 | 10802.6 | 377.69 | 140550109 | 25790134 | 3372784.0 | 2464.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 2420097 | 19449 | 9416.7 | 433.28 | 161234418 | 29585586 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 15146023 | 153722 | 890942.5 | 69.23 | 25762736 | 4714439 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 6618260 | 48374 | 389309.4 | 158.44 | 58958549 | 10789090 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 6483193 | 88392 | 381364.3 | 161.74 | 60186857 | 5507780 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 6540529 | 82145 | 384737.0 | 160.32 | 59659242 | 10917313 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 2411377 | 11910 | 141845.7 | 434.84 | 161817507 | 14808137 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 2735488 | 15730 | 160911.1 | 383.32 | 142644753 | 26103204 | 4255600.0 | 260.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 2424593 | 8160 | 142623.1 | 432.47 | 160935438 | 29450299 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 19195030 | 296049 | 19195030.1 | 54.63 | 20328335 | 3719452 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 6751404 | 91671 | 6751404.2 | 155.31 | 57795829 | 10574837 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 6762606 | 151874 | 6762606.1 | 155.05 | 57700093 | 5278734 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 6748969 | 93853 | 6748968.9 | 155.37 | 57816684 | 10578653 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 2416963 | 9370 | 2416962.8 | 433.84 | 161443530 | 14769777 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 2739419 | 13488 | 2739419.3 | 382.77 | 142440043 | 26062093 | 5559536.0 | 27.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 2419556 | 14839 | 2419556.4 | 433.37 | 161270474 | 29507475 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.470x |
| ar-wikipedia | 4096 | 1.358x |
| ar-wikipedia | 65536 | 1.340x |
| ar-wikipedia | 1048576 | 1.345x |
| en-wikitext2 | 256 | 3.551x |
| en-wikitext2 | 4096 | 3.172x |
| en-wikitext2 | 65536 | 3.334x |
| en-wikitext2 | 1048576 | 3.284x |
| ja-wikipedia | 256 | 2.010x |
| ja-wikipedia | 4096 | 1.792x |
| ja-wikipedia | 65536 | 1.766x |
| ja-wikipedia | 1048576 | 1.755x |
| ru-wikipedia | 256 | 1.473x |
| ru-wikipedia | 4096 | 1.366x |
| ru-wikipedia | 65536 | 1.356x |
| ru-wikipedia | 1048576 | 1.353x |
| synthetic | 0 | 3.227x |
| zh-wikipedia | 256 | 3.012x |
| zh-wikipedia | 4096 | 2.732x |
| zh-wikipedia | 65536 | 2.689x |
| zh-wikipedia | 1048576 | 2.798x |
| Combined selected natural corpora | 256 | 2.108x |
| Combined selected natural corpora | 4096 | 1.925x |
| Combined selected natural corpora | 65536 | 1.915x |
| Combined selected natural corpora | 1048576 | 1.929x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.472x |
| ar-wikipedia | 4096 | 1.347x |
| ar-wikipedia | 65536 | 1.339x |
| ar-wikipedia | 1048576 | 1.343x |
| en-wikitext2 | 256 | 3.542x |
| en-wikitext2 | 4096 | 3.174x |
| en-wikitext2 | 65536 | 3.291x |
| en-wikitext2 | 1048576 | 3.306x |
| ja-wikipedia | 256 | 2.054x |
| ja-wikipedia | 4096 | 1.819x |
| ja-wikipedia | 65536 | 1.797x |
| ja-wikipedia | 1048576 | 1.782x |
| ru-wikipedia | 256 | 1.472x |
| ru-wikipedia | 4096 | 1.368x |
| ru-wikipedia | 65536 | 1.352x |
| ru-wikipedia | 1048576 | 1.351x |
| synthetic | 0 | 3.229x |
| zh-wikipedia | 256 | 3.021x |
| zh-wikipedia | 4096 | 2.763x |
| zh-wikipedia | 65536 | 2.698x |
| zh-wikipedia | 1048576 | 2.789x |
| Combined selected natural corpora | 256 | 2.119x |
| Combined selected natural corpora | 4096 | 1.932x |
| Combined selected natural corpora | 65536 | 1.917x |
| Combined selected natural corpora | 1048576 | 1.936x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.475x |
| ar-wikipedia | 4096 | 1.354x |
| ar-wikipedia | 65536 | 1.354x |
| ar-wikipedia | 1048576 | 1.360x |
| en-wikitext2 | 256 | 3.502x |
| en-wikitext2 | 4096 | 3.169x |
| en-wikitext2 | 65536 | 3.310x |
| en-wikitext2 | 1048576 | 3.319x |
| ja-wikipedia | 256 | 1.908x |
| ja-wikipedia | 4096 | 1.707x |
| ja-wikipedia | 65536 | 1.687x |
| ja-wikipedia | 1048576 | 1.661x |
| ru-wikipedia | 256 | 1.470x |
| ru-wikipedia | 4096 | 1.367x |
| ru-wikipedia | 65536 | 1.358x |
| ru-wikipedia | 1048576 | 1.348x |
| synthetic | 0 | 2.801x |
| zh-wikipedia | 256 | 2.628x |
| zh-wikipedia | 4096 | 2.467x |
| zh-wikipedia | 65536 | 2.419x |
| zh-wikipedia | 1048576 | 2.465x |
| Combined selected natural corpora | 256 | 2.044x |
| Combined selected natural corpora | 4096 | 1.877x |
| Combined selected natural corpora | 65536 | 1.870x |
| Combined selected natural corpora | 1048576 | 1.878x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.360x |
| ar-wikipedia | 4096 | 1.936x |
| ar-wikipedia | 65536 | 1.881x |
| ar-wikipedia | 1048576 | 2.383x |
| en-wikitext2 | 256 | 4.714x |
| en-wikitext2 | 4096 | 4.071x |
| en-wikitext2 | 65536 | 4.214x |
| en-wikitext2 | 1048576 | 4.877x |
| ja-wikipedia | 256 | 5.042x |
| ja-wikipedia | 4096 | 4.307x |
| ja-wikipedia | 65536 | 4.177x |
| ja-wikipedia | 1048576 | 5.067x |
| ru-wikipedia | 256 | 2.355x |
| ru-wikipedia | 4096 | 1.988x |
| ru-wikipedia | 65536 | 1.909x |
| ru-wikipedia | 1048576 | 2.329x |
| synthetic | 0 | 7.397x |
| zh-wikipedia | 256 | 7.383x |
| zh-wikipedia | 4096 | 6.389x |
| zh-wikipedia | 65536 | 6.247x |
| zh-wikipedia | 1048576 | 7.933x |
| Combined selected natural corpora | 256 | 3.990x |
| Combined selected natural corpora | 4096 | 3.397x |
| Combined selected natural corpora | 65536 | 3.324x |
| Combined selected natural corpora | 1048576 | 4.089x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.326x |
| ar-wikipedia | 4096 | 1.935x |
| ar-wikipedia | 65536 | 1.889x |
| ar-wikipedia | 1048576 | 2.396x |
| en-wikitext2 | 256 | 4.626x |
| en-wikitext2 | 4096 | 4.061x |
| en-wikitext2 | 65536 | 4.235x |
| en-wikitext2 | 1048576 | 4.883x |
| ja-wikipedia | 256 | 4.283x |
| ja-wikipedia | 4096 | 3.946x |
| ja-wikipedia | 65536 | 3.887x |
| ja-wikipedia | 1048576 | 4.694x |
| ru-wikipedia | 256 | 2.324x |
| ru-wikipedia | 4096 | 1.988x |
| ru-wikipedia | 65536 | 1.914x |
| ru-wikipedia | 1048576 | 2.330x |
| synthetic | 0 | 6.130x |
| zh-wikipedia | 256 | 5.801x |
| zh-wikipedia | 4096 | 5.569x |
| zh-wikipedia | 65536 | 5.537x |
| zh-wikipedia | 1048576 | 7.007x |
| Combined selected natural corpora | 256 | 3.674x |
| Combined selected natural corpora | 4096 | 3.265x |
| Combined selected natural corpora | 65536 | 3.224x |
| Combined selected natural corpora | 1048576 | 3.956x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 394.8 | 10.4 | 10 |
| onig-compile | 370.8 | 22.3 | 10 |
| pretok-compile | 10.1 | 0.2 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
