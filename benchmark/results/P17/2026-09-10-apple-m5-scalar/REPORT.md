# P17 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 9799661 | 23179 | 2387.8 | 107.00 | 59652571 | 430729 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 7226576 | 35495 | 1760.9 | 145.10 | 80892391 | 584094 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 7150454 | 70580 | 1742.3 | 146.64 | 81753553 | 8531 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 7148852 | 41811 | 1741.9 | 146.68 | 81771866 | 590444 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 4215935 | 35112 | 1027.3 | 248.72 | 138658434 | 14469 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 4231021 | 27468 | 1031.0 | 247.83 | 138164050 | 997632 | 71360.0 | 4198.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 4203268 | 21995 | 1024.2 | 249.47 | 139076308 | 1004219 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 7987527 | 26674 | 31079.9 | 131.28 | 73185980 | 46948 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 5571201 | 42119 | 21677.8 | 188.21 | 104927997 | 67310 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 5552523 | 60320 | 21605.1 | 188.85 | 105280971 | 10626 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 5578439 | 27621 | 21706.0 | 187.97 | 104791862 | 67223 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 4114835 | 26177 | 16011.0 | 254.83 | 142065245 | 14338 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 4133947 | 66071 | 16085.4 | 253.65 | 141408433 | 90712 | 9648.0 | 328.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 4123962 | 36368 | 16046.5 | 254.26 | 141750830 | 90932 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 7888891 | 172498 | 464052.4 | 132.92 | 74101034 | 17113 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 6100651 | 1132952 | 358861.8 | 171.88 | 95821737 | 22129 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 6549309 | 993036 | 385253.5 | 160.10 | 89257509 | 9009 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 6141126 | 1072091 | 361242.7 | 170.75 | 95190204 | 21983 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 4164541 | 105041 | 244973.0 | 251.79 | 140369601 | 14167 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 4132209 | 39794 | 243071.1 | 253.76 | 141467925 | 32670 | 5616.0 | 62.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 4146285 | 93353 | 243899.1 | 252.90 | 140987663 | 32559 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 9513551 | 70284 | 9513551.1 | 110.22 | 61446561 | 12508 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 6552267 | 1010145 | 6552266.9 | 160.03 | 89217214 | 18162 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 6598376 | 1110429 | 6598375.7 | 158.91 | 88593774 | 8942 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 6598003 | 1072884 | 6598002.7 | 158.92 | 88598782 | 18036 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 4140339 | 20314 | 4140339.0 | 253.26 | 141190129 | 14250 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 4181521 | 57988 | 4181521.0 | 250.76 | 139799609 | 28459 | 4080.0 | 8.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 4143525 | 48462 | 4143524.6 | 253.06 | 141081580 | 28720 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 10988438 | 49360 | 2682.1 | 95.43 | 95230737 | 490334 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 8165473 | 41212 | 1993.0 | 128.42 | 128153878 | 659852 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 8105946 | 42739 | 1978.5 | 129.36 | 129094995 | 80435 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 8116566 | 52289 | 1981.1 | 129.19 | 128926078 | 663828 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 2277331 | 13771 | 555.9 | 460.44 | 459501426 | 286300 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 2341609 | 42624 | 571.5 | 447.80 | 446888102 | 2300982 | 130096.0 | 5144.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 2283274 | 8873 | 557.3 | 459.24 | 458305419 | 2359769 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 9959932 | 31502 | 38906.0 | 105.28 | 105064671 | 156427 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 8141209 | 13873 | 31801.6 | 128.80 | 128535828 | 191372 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 8139798 | 45109 | 31796.1 | 128.82 | 128558104 | 79977 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 8136107 | 37296 | 31781.7 | 128.88 | 128616433 | 191492 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 2219957 | 62745 | 8671.7 | 472.34 | 471377217 | 293249 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 2250072 | 6847 | 8789.3 | 466.02 | 465068262 | 692422 | 68096.0 | 818.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 2219334 | 43539 | 8669.3 | 472.47 | 471509554 | 702013 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 9950942 | 72045 | 621933.9 | 105.37 | 105159592 | 132450 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 8067390 | 36783 | 504211.9 | 129.98 | 129711956 | 163374 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 8066915 | 21486 | 504182.2 | 129.98 | 129719597 | 80700 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 8070135 | 32854 | 504383.4 | 129.93 | 129667847 | 163318 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 2266610 | 41673 | 141663.1 | 462.62 | 461674884 | 287213 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 2282557 | 60964 | 142659.8 | 459.39 | 458449361 | 577422 | 57088.0 | 121.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 2266142 | 41602 | 141633.9 | 462.71 | 461770314 | 581605 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 11292191 | 258020 | 11292190.7 | 92.86 | 92669086 | 115389 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 7965877 | 392627 | 7965876.8 | 131.63 | 131364949 | 163573 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 8027372 | 388376 | 8027371.9 | 130.63 | 130358605 | 81098 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 7948226 | 384901 | 7948225.7 | 131.93 | 131656680 | 163936 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 2213229 | 4035 | 2213228.5 | 473.78 | 472810191 | 294140 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 2224927 | 22689 | 2224927.2 | 471.29 | 470324143 | 585637 | 79088.0 | 13.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 2214198 | 33836 | 2214197.9 | 473.57 | 472603191 | 588475 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 11302367 | 104750 | 2748.0 | 92.77 | 33557573 | 3090857 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 5018401 | 23263 | 1220.1 | 208.95 | 75577864 | 6961182 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 4686899 | 20664 | 1139.5 | 223.72 | 80923440 | 4105273 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 4694699 | 35054 | 1141.4 | 223.35 | 80788992 | 7441159 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 2132893 | 3104 | 518.6 | 491.62 | 177824197 | 9021080 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 2500821 | 7940 | 608.0 | 419.29 | 151662179 | 13969011 | 1519824.0 | 16215.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 2150741 | 6384 | 522.9 | 487.54 | 176348489 | 16242771 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 9121482 | 114504 | 35492.1 | 114.96 | 41580962 | 3437051 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 3918531 | 55117 | 15247.2 | 267.59 | 96791376 | 8000702 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 3846629 | 39934 | 14967.4 | 272.60 | 98600619 | 4103333 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 3855472 | 40324 | 15001.8 | 271.97 | 98374459 | 8131559 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 2084086 | 21920 | 8109.3 | 503.13 | 181988664 | 7573584 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 2276118 | 46001 | 8856.5 | 460.69 | 166634588 | 13773890 | 1399920.0 | 2097.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 2104008 | 28314 | 8186.8 | 498.37 | 180265440 | 14900606 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 8912225 | 143994 | 524248.5 | 117.66 | 42557273 | 3492618 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 3867892 | 30266 | 227523.1 | 271.10 | 98058569 | 8047535 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 3799997 | 74578 | 223529.2 | 275.94 | 99810609 | 4097635 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 3809859 | 63919 | 224109.3 | 275.23 | 99552251 | 8170120 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 2094345 | 12040 | 123196.7 | 500.67 | 181097221 | 7434784 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 2242221 | 37705 | 131895.3 | 467.65 | 169153728 | 13882219 | 1761040.0 | 220.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 2099811 | 13337 | 123518.3 | 499.37 | 180625804 | 14823717 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 10730554 | 144341 | 10730554.1 | 97.72 | 35345798 | 2899291 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 3788607 | 58622 | 3788606.9 | 276.77 | 100110678 | 8211726 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 3789644 | 46212 | 3789644.3 | 276.69 | 100083272 | 4104871 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 3796367 | 61360 | 3796367.0 | 276.20 | 99906041 | 8194940 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 2074552 | 10828 | 2074551.7 | 505.45 | 182825037 | 7498487 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 2234286 | 32705 | 2234286.0 | 469.31 | 169754453 | 13924359 | 2659568.0 | 24.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 2080099 | 22002 | 2080099.1 | 504.10 | 182337469 | 14956499 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 10120671 | 87403 | 2466.1 | 103.61 | 56800286 | 667248 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 6434639 | 198206 | 1567.9 | 162.96 | 89337873 | 1049476 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 6322347 | 246451 | 1540.5 | 165.85 | 90924622 | 210839 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 6333724 | 210179 | 1543.3 | 165.55 | 90761299 | 1066197 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 4083599 | 20257 | 995.0 | 256.78 | 140772156 | 326428 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 4180818 | 47575 | 1018.7 | 250.81 | 137498677 | 1615234 | 197376.0 | 6247.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 4069314 | 50742 | 991.5 | 257.68 | 141266322 | 1659494 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 8333890 | 63640 | 32427.6 | 125.82 | 68978234 | 350617 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 5605063 | 216912 | 21809.6 | 187.08 | 102560307 | 521314 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 5563252 | 222338 | 21646.9 | 188.48 | 103331112 | 239608 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 5573558 | 204460 | 21687.0 | 188.13 | 103140039 | 524261 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 4022398 | 56415 | 15651.4 | 260.68 | 142913985 | 331394 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 4069508 | 40717 | 15834.7 | 257.67 | 141259570 | 718023 | 129840.0 | 1156.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 4016527 | 44283 | 15628.5 | 261.07 | 143122914 | 727494 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 8083501 | 174544 | 475500.1 | 129.72 | 71114852 | 331911 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 5515079 | 220428 | 324416.4 | 190.13 | 104233684 | 486484 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 5525830 | 111651 | 325048.8 | 189.76 | 104030889 | 241231 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 5506158 | 13550 | 323891.6 | 190.44 | 104402571 | 487273 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 4031213 | 94922 | 237130.2 | 260.11 | 142601504 | 330670 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 4052968 | 52082 | 238409.9 | 258.72 | 141836060 | 661984 | 114448.0 | 141.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 4013166 | 105172 | 236068.6 | 261.28 | 143242767 | 668549 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 9793486 | 125777 | 9793485.8 | 107.07 | 58697895 | 272324 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 5744574 | 51530 | 5744574.3 | 182.53 | 100069555 | 464264 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 5744254 | 51463 | 5744254.3 | 182.54 | 100075131 | 232058 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 5742767 | 36789 | 5742766.8 | 182.59 | 100101052 | 464410 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 3985763 | 19947 | 3985763.4 | 263.08 | 144227578 | 334440 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 4012516 | 56963 | 4012515.8 | 261.33 | 143265976 | 664670 | 177392.0 | 15.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 3989734 | 37238 | 3989734.1 | 262.82 | 144084037 | 668466 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 63134888 | 1699301 | 866.4 | 36.24 | 19740528 | 3137014 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 28907078 | 256706 | 396.7 | 79.15 | 43114562 | 6851436 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 27088243 | 91376 | 371.7 | 84.46 | 46009480 | 2704199 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 27269583 | 60011 | 374.2 | 83.90 | 45703522 | 7262854 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 7007038 | 52041 | 96.2 | 326.51 | 177866306 | 10454060 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 9467566 | 97160 | 129.9 | 241.65 | 131640587 | 20919315 | 7614112.0 | 128141.0 | 10 |
| synthetic | 0 | pretok-reused | 7167694 | 110549 | 98.4 | 319.19 | 173879624 | 27631619 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 14600540 | 96919 | 3549.9 | 71.82 | 26725244 | 3728013 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 6516937 | 28578 | 1584.5 | 160.90 | 59875217 | 8352237 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 6005183 | 86473 | 1460.0 | 174.61 | 64977708 | 4790862 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 6042432 | 31854 | 1469.1 | 173.54 | 64577144 | 9008128 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 2309354 | 14326 | 561.5 | 454.06 | 168966280 | 12458028 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 2835678 | 6957 | 689.4 | 369.78 | 137604835 | 19195057 | 2404080.0 | 20718.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 2327499 | 14201 | 565.9 | 450.52 | 167649028 | 23386043 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 12125222 | 183893 | 47179.9 | 86.48 | 32181101 | 4213201 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 5288190 | 67760 | 20576.6 | 198.29 | 73787633 | 9660395 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 5170049 | 67892 | 20116.9 | 202.82 | 75473750 | 4961075 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 5169167 | 68085 | 20113.5 | 202.85 | 75486632 | 9882830 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 2263828 | 7508 | 8808.7 | 463.19 | 172364246 | 11329925 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 2525564 | 49834 | 9827.1 | 415.18 | 154501333 | 20227561 | 2343408.0 | 2333.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 2267885 | 25842 | 8824.5 | 462.36 | 172055927 | 22525837 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 11794879 | 102860 | 693816.4 | 88.90 | 33082407 | 4313312 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 5103407 | 41101 | 300200.4 | 205.47 | 76459312 | 9968830 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 5035050 | 30280 | 296179.4 | 208.25 | 77497344 | 5053177 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 5068428 | 49378 | 298142.8 | 206.88 | 76986986 | 10037629 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 2264375 | 16041 | 133198.5 | 463.07 | 172322609 | 11236213 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 2509925 | 18508 | 147642.6 | 417.77 | 155464008 | 20269530 | 3051376.0 | 245.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 2280048 | 7836 | 134120.5 | 459.89 | 171138077 | 22313128 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 14267020 | 998078 | 14267019.6 | 73.50 | 27350001 | 3565216 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 5037580 | 39025 | 5037579.6 | 208.15 | 77458429 | 10097111 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 5020729 | 37543 | 5020729.4 | 208.85 | 77718389 | 5065599 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 5024596 | 51506 | 5024596.2 | 208.69 | 77658578 | 10123201 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 2264891 | 14110 | 2264891.2 | 462.97 | 172283330 | 11229237 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 2536277 | 27748 | 2536277.3 | 413.43 | 153848711 | 20054983 | 4363504.0 | 26.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 2256079 | 11609 | 2256079.4 | 464.78 | 172956236 | 22545749 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.696x |
| ar-wikipedia | 4096 | 1.349x |
| ar-wikipedia | 65536 | 1.573x |
| ar-wikipedia | 1048576 | 1.594x |
| en-wikitext2 | 256 | 3.559x |
| en-wikitext2 | 4096 | 3.667x |
| en-wikitext2 | 65536 | 3.559x |
| en-wikitext2 | 1048576 | 3.627x |
| ja-wikipedia | 256 | 2.197x |
| ja-wikipedia | 4096 | 1.846x |
| ja-wikipedia | 65536 | 1.814x |
| ja-wikipedia | 1048576 | 1.827x |
| ru-wikipedia | 256 | 1.548x |
| ru-wikipedia | 4096 | 1.383x |
| ru-wikipedia | 65536 | 1.371x |
| ru-wikipedia | 1048576 | 1.441x |
| synthetic | 0 | 3.866x |
| zh-wikipedia | 256 | 2.600x |
| zh-wikipedia | 4096 | 2.284x |
| zh-wikipedia | 65536 | 2.224x |
| zh-wikipedia | 1048576 | 2.217x |
| Combined selected natural corpora | 256 | 2.149x |
| Combined selected natural corpora | 4096 | 1.923x |
| Combined selected natural corpora | 65536 | 1.955x |
| Combined selected natural corpora | 1048576 | 1.988x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.701x |
| ar-wikipedia | 4096 | 1.353x |
| ar-wikipedia | 65536 | 1.481x |
| ar-wikipedia | 1048576 | 1.592x |
| en-wikitext2 | 256 | 3.555x |
| en-wikitext2 | 4096 | 3.666x |
| en-wikitext2 | 65536 | 3.561x |
| en-wikitext2 | 1048576 | 3.590x |
| ja-wikipedia | 256 | 2.183x |
| ja-wikipedia | 4096 | 1.832x |
| ja-wikipedia | 65536 | 1.814x |
| ja-wikipedia | 1048576 | 1.825x |
| ru-wikipedia | 256 | 1.556x |
| ru-wikipedia | 4096 | 1.388x |
| ru-wikipedia | 65536 | 1.372x |
| ru-wikipedia | 1048576 | 1.439x |
| synthetic | 0 | 3.805x |
| zh-wikipedia | 256 | 2.596x |
| zh-wikipedia | 4096 | 2.279x |
| zh-wikipedia | 65536 | 2.223x |
| zh-wikipedia | 1048576 | 2.227x |
| Combined selected natural corpora | 256 | 2.151x |
| Combined selected natural corpora | 4096 | 1.922x |
| Combined selected natural corpora | 65536 | 1.931x |
| Combined selected natural corpora | 1048576 | 1.982x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.708x |
| ar-wikipedia | 4096 | 1.348x |
| ar-wikipedia | 65536 | 1.476x |
| ar-wikipedia | 1048576 | 1.567x |
| en-wikitext2 | 256 | 3.487x |
| en-wikitext2 | 4096 | 3.618x |
| en-wikitext2 | 65536 | 3.534x |
| en-wikitext2 | 1048576 | 3.580x |
| ja-wikipedia | 256 | 2.007x |
| ja-wikipedia | 4096 | 1.722x |
| ja-wikipedia | 65536 | 1.725x |
| ja-wikipedia | 1048576 | 1.696x |
| ru-wikipedia | 256 | 1.539x |
| ru-wikipedia | 4096 | 1.377x |
| ru-wikipedia | 65536 | 1.361x |
| ru-wikipedia | 1048576 | 1.432x |
| synthetic | 0 | 3.053x |
| zh-wikipedia | 256 | 2.298x |
| zh-wikipedia | 4096 | 2.094x |
| zh-wikipedia | 65536 | 2.033x |
| zh-wikipedia | 1048576 | 1.986x |
| Combined selected natural corpora | 256 | 2.073x |
| Combined selected natural corpora | 4096 | 1.870x |
| Combined selected natural corpora | 65536 | 1.883x |
| Combined selected natural corpora | 1048576 | 1.915x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.331x |
| ar-wikipedia | 4096 | 1.937x |
| ar-wikipedia | 65536 | 1.903x |
| ar-wikipedia | 1048576 | 2.296x |
| en-wikitext2 | 256 | 4.813x |
| en-wikitext2 | 4096 | 4.488x |
| en-wikitext2 | 65536 | 4.391x |
| en-wikitext2 | 1048576 | 5.100x |
| ja-wikipedia | 256 | 5.255x |
| ja-wikipedia | 4096 | 4.335x |
| ja-wikipedia | 65536 | 4.244x |
| ja-wikipedia | 1048576 | 5.159x |
| ru-wikipedia | 256 | 2.487x |
| ru-wikipedia | 4096 | 2.075x |
| ru-wikipedia | 65536 | 2.014x |
| ru-wikipedia | 1048576 | 2.455x |
| synthetic | 0 | 8.808x |
| zh-wikipedia | 256 | 6.273x |
| zh-wikipedia | 4096 | 5.346x |
| zh-wikipedia | 65536 | 5.173x |
| zh-wikipedia | 1048576 | 6.324x |
| Combined selected natural corpora | 256 | 3.779x |
| Combined selected natural corpora | 4096 | 3.226x |
| Combined selected natural corpora | 65536 | 3.150x |
| Combined selected natural corpora | 1048576 | 3.786x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.316x |
| ar-wikipedia | 4096 | 1.932x |
| ar-wikipedia | 65536 | 1.909x |
| ar-wikipedia | 1048576 | 2.275x |
| en-wikitext2 | 256 | 4.693x |
| en-wikitext2 | 4096 | 4.426x |
| en-wikitext2 | 65536 | 4.360x |
| en-wikitext2 | 1048576 | 5.075x |
| ja-wikipedia | 256 | 4.519x |
| ja-wikipedia | 4096 | 4.007x |
| ja-wikipedia | 65536 | 3.975x |
| ja-wikipedia | 1048576 | 4.803x |
| ru-wikipedia | 256 | 2.421x |
| ru-wikipedia | 4096 | 2.048x |
| ru-wikipedia | 65536 | 1.994x |
| ru-wikipedia | 1048576 | 2.441x |
| synthetic | 0 | 6.669x |
| zh-wikipedia | 256 | 5.149x |
| zh-wikipedia | 4096 | 4.801x |
| zh-wikipedia | 65536 | 4.699x |
| zh-wikipedia | 1048576 | 5.625x |
| Combined selected natural corpora | 256 | 3.531x |
| Combined selected natural corpora | 4096 | 3.116x |
| Combined selected natural corpora | 65536 | 3.064x |
| Combined selected natural corpora | 1048576 | 3.660x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 354.7 | 1.6 | 10 |
| onig-compile | 329.6 | 2.0 | 10 |
| pretok-compile | 10.0 | 0.1 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
