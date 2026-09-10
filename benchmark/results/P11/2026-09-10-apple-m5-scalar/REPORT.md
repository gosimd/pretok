# P11 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 7530678 | 20737 | 1835.0 | 139.24 | 77625811 | 1692278 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 3129053 | 9743 | 762.4 | 335.11 | 186821684 | 4072797 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 2984201 | 16213 | 727.1 | 351.38 | 195889938 | 1748206 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 3001622 | 15397 | 731.4 | 349.34 | 194753069 | 4245705 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 1875178 | 8610 | 456.9 | 559.19 | 311743781 | 2782136 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 2011608 | 4962 | 490.2 | 521.26 | 290600875 | 6335231 | 425024.0 | 8613.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 1882870 | 12543 | 458.8 | 556.90 | 310470230 | 6768392 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 5828259 | 17741 | 22678.1 | 179.91 | 100300107 | 1533219 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 2725189 | 22906 | 10603.8 | 384.77 | 214508069 | 3279039 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 2669494 | 30124 | 10387.1 | 392.80 | 218983483 | 1954303 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 2682234 | 13254 | 10436.7 | 390.93 | 217943312 | 3331551 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 1811733 | 1470 | 7049.5 | 578.77 | 322660595 | 2879563 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 1868483 | 4447 | 7270.4 | 561.19 | 312860706 | 4782489 | 408720.0 | 1605.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 1812861 | 1908 | 7053.9 | 578.41 | 322459839 | 4929224 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 5603835 | 19899 | 329637.4 | 187.12 | 104316955 | 1552508 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 2672477 | 10599 | 157204.6 | 392.36 | 218738984 | 3255406 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 2643579 | 8410 | 155504.7 | 396.65 | 221130119 | 1973461 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 2665577 | 9347 | 156798.7 | 393.38 | 219305201 | 3263833 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 1808524 | 3236 | 106383.8 | 579.80 | 323233121 | 2884672 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 1843484 | 6698 | 108440.2 | 568.80 | 317103364 | 4719325 | 423440.0 | 172.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 1812912 | 6149 | 106641.9 | 578.39 | 322450844 | 4798909 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 7083065 | 41236 | 7083065.4 | 148.04 | 82531357 | 1226023 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 2686646 | 8067 | 2686646.4 | 390.29 | 217585389 | 3232282 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 2666644 | 15033 | 2666644.2 | 393.22 | 219217474 | 1956392 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 2675918 | 19921 | 2675917.8 | 391.86 | 218457758 | 3245242 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 1786952 | 18108 | 1786951.7 | 586.80 | 327135318 | 2919497 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 1829612 | 13765 | 1829611.8 | 573.11 | 319507673 | 4746362 | 505072.0 | 18.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 1797495 | 4421 | 1797495.1 | 583.35 | 325216469 | 4831168 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 6386987 | 90193 | 1558.9 | 164.17 | 163838903 | 2021141 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 2519692 | 31614 | 615.0 | 416.15 | 415303502 | 5123245 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 2337131 | 25109 | 570.4 | 448.66 | 447744306 | 2838095 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 2377758 | 53466 | 580.4 | 440.99 | 440093964 | 5429064 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 498514 | 3720 | 121.7 | 2103.40 | 2099110710 | 13305532 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 654672 | 27510 | 159.8 | 1601.68 | 1598413617 | 19718264 | 552784.0 | 8985.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 507636 | 20336 | 123.9 | 2065.61 | 2061393922 | 25429657 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 4966132 | 16920 | 19399.0 | 211.15 | 210714714 | 1832211 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 2258240 | 7825 | 8821.2 | 464.33 | 463386155 | 4029245 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 2168622 | 12483 | 8471.2 | 483.52 | 482535488 | 3058625 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 2187869 | 11304 | 8546.4 | 479.27 | 478290560 | 4158842 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 432094 | 1314 | 1687.9 | 2426.73 | 2421783480 | 15350843 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 494770 | 3153 | 1932.7 | 2119.32 | 2114996212 | 18390358 | 430592.0 | 1665.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 434376 | 2603 | 1696.8 | 2413.98 | 2409059898 | 20947306 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 5048346 | 49880 | 315521.6 | 207.71 | 207283136 | 1755030 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 2185383 | 17760 | 136586.5 | 479.81 | 478834515 | 4054209 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 2162019 | 9742 | 135126.2 | 485.00 | 484009152 | 3067966 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 2170116 | 7447 | 135632.3 | 483.19 | 482203159 | 4082730 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 431045 | 24221 | 26940.3 | 2432.64 | 2427676704 | 15388198 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 480105 | 22839 | 30006.5 | 2184.06 | 2179601903 | 18454310 | 411392.0 | 170.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 435490 | 24510 | 27218.1 | 2407.81 | 2402896142 | 20344904 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 6503044 | 141703 | 6503043.6 | 161.24 | 160914961 | 1360132 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 2181797 | 10358 | 2181797.5 | 480.60 | 479621512 | 4053997 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 2168503 | 18253 | 2168502.6 | 483.55 | 482562014 | 3058793 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 2170264 | 17418 | 2170264.4 | 483.16 | 482170276 | 4075540 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 425414 | 2479 | 425414.3 | 2464.83 | 2459806612 | 15591858 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 470272 | 1900 | 470271.6 | 2229.72 | 2225175695 | 18808279 | 685296.0 | 19.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 427612 | 3379 | 427612.2 | 2452.17 | 2447163365 | 20684628 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 6903051 | 43189 | 1678.3 | 151.90 | 54943818 | 2561911 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 2635498 | 9147 | 640.8 | 397.87 | 143912085 | 6710307 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 2458961 | 6832 | 597.9 | 426.43 | 154244009 | 3283094 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 2485677 | 10407 | 604.3 | 421.85 | 152586198 | 7114762 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 903907 | 2926 | 219.8 | 1160.05 | 419600915 | 8931233 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 1081229 | 3965 | 262.9 | 969.80 | 350785891 | 16356382 | 651568.0 | 10484.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 906213 | 1902 | 220.3 | 1157.10 | 418533100 | 19515287 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 5532792 | 11490 | 21528.4 | 189.52 | 68551286 | 2511752 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 2254746 | 11663 | 8773.3 | 465.05 | 168214093 | 6163445 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 2191563 | 7999 | 8527.5 | 478.46 | 173063677 | 3683672 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 2209204 | 3908 | 8596.1 | 474.64 | 171681737 | 6290501 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 835104 | 15780 | 3249.4 | 1255.62 | 454170811 | 9667056 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 913326 | 5122 | 3553.8 | 1148.08 | 415273513 | 15215819 | 648048.0 | 1767.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 837893 | 12252 | 3260.3 | 1251.44 | 452659422 | 16585657 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 5423521 | 17159 | 319030.6 | 193.34 | 69932435 | 2519397 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 2209965 | 14903 | 129998.0 | 474.48 | 171622611 | 6182903 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 2179827 | 9293 | 128225.1 | 481.04 | 173995438 | 3703504 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 2197120 | 15089 | 129242.4 | 477.25 | 172625985 | 6219050 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 826702 | 3424 | 48629.5 | 1268.38 | 458786710 | 9765306 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 879603 | 4290 | 51741.3 | 1192.10 | 431194569 | 15534282 | 659216.0 | 185.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 826390 | 3846 | 48611.2 | 1268.86 | 458959947 | 16534562 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 6863600 | 18182 | 6863600.2 | 152.77 | 55259629 | 1988607 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 2212579 | 16675 | 2212579.2 | 473.92 | 171419855 | 6168819 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 2199880 | 17066 | 2199879.6 | 476.65 | 172409438 | 3669746 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 2213993 | 18403 | 2213993.3 | 473.61 | 171310367 | 6164879 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 823227 | 4856 | 823227.5 | 1273.74 | 460723197 | 9806524 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 879463 | 6754 | 879463.0 | 1192.29 | 431263176 | 15519698 | 914672.0 | 20.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 826722 | 8883 | 826722.3 | 1268.35 | 458775565 | 16509776 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 6999731 | 85652 | 1705.6 | 149.80 | 82125586 | 1428483 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 2904813 | 14864 | 707.8 | 360.98 | 197898086 | 3442218 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 2793789 | 12813 | 680.7 | 375.32 | 205762465 | 1291078 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 2814050 | 28858 | 685.7 | 372.62 | 204280995 | 3553241 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 1779255 | 28415 | 433.5 | 589.33 | 323088613 | 2027253 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 1879614 | 28854 | 458.0 | 557.87 | 305837796 | 5319709 | 306016.0 | 7606.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 1780136 | 31294 | 433.8 | 589.04 | 322928614 | 5616985 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 5393565 | 20456 | 20986.6 | 194.41 | 106582020 | 1144141 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 2527610 | 18327 | 9835.1 | 414.85 | 227431086 | 2441437 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 2488158 | 23748 | 9681.5 | 421.43 | 231037187 | 1449667 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 2497336 | 16300 | 9717.3 | 419.88 | 230188057 | 2471033 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 1699409 | 10264 | 6612.5 | 617.02 | 338268865 | 2122503 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 1731375 | 8754 | 6736.9 | 605.63 | 332023322 | 3564218 | 264144.0 | 1453.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 1698921 | 6171 | 6610.6 | 617.20 | 338365930 | 3632305 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 5147660 | 27575 | 302803.5 | 203.70 | 111673458 | 1152174 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 2453129 | 29857 | 144301.7 | 427.44 | 234336233 | 2417729 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 2430586 | 27482 | 142975.6 | 431.41 | 236509655 | 1484004 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 2443509 | 22824 | 143735.8 | 429.13 | 235258847 | 2427248 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 1678970 | 15343 | 98762.9 | 624.53 | 342386671 | 2148341 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 1703107 | 12067 | 100182.7 | 615.68 | 337534367 | 3482460 | 278032.0 | 161.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 1679309 | 17065 | 98782.9 | 624.41 | 342317669 | 3531811 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 6560552 | 28638 | 6560552.2 | 159.83 | 87623264 | 901753 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 2460031 | 34954 | 2460030.9 | 426.24 | 233678777 | 2404848 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 2454320 | 38813 | 2454319.6 | 427.24 | 234222548 | 1469654 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 2456562 | 45601 | 2456561.7 | 426.85 | 234008779 | 2408244 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 1653656 | 4399 | 1653656.4 | 634.09 | 347627848 | 2181227 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 1676101 | 1848 | 1676101.2 | 625.60 | 342972734 | 3529620 | 365808.0 | 17.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 1654647 | 6306 | 1654647.3 | 633.72 | 347419663 | 3575384 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 47828167 | 292528 | 656.3 | 47.84 | 26058201 | 3924758 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 16235163 | 125748 | 222.8 | 140.92 | 76766462 | 11562188 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 14741157 | 144428 | 202.3 | 155.20 | 84546689 | 5139556 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 14867276 | 189475 | 204.0 | 153.89 | 83829477 | 12625984 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 2861921 | 33434 | 39.3 | 799.41 | 435482378 | 26472782 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 5073171 | 54246 | 69.6 | 450.97 | 245668041 | 37001315 | 7069856.0 | 122057.0 | 10 |
| synthetic | 0 | pretok-reused | 3011101 | 44199 | 41.3 | 759.81 | 413907083 | 62340654 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 6954617 | 21278 | 1690.9 | 150.77 | 56107045 | 2558301 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 2745429 | 11179 | 667.5 | 381.93 | 142128257 | 6480591 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 2570377 | 17515 | 624.9 | 407.95 | 151807665 | 3241547 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 2595118 | 11739 | 631.0 | 404.06 | 150360421 | 6855951 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 908016 | 5646 | 220.8 | 1154.80 | 429731203 | 9176045 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 1086124 | 8383 | 264.1 | 965.43 | 359261774 | 16381180 | 648752.0 | 10417.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 911835 | 7923 | 221.7 | 1149.96 | 427931801 | 19512312 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 5588949 | 7823 | 21746.9 | 187.62 | 69816882 | 2511027 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 2356749 | 9300 | 9170.2 | 444.92 | 165568358 | 5954814 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 2280535 | 7098 | 8873.7 | 459.79 | 171101529 | 3653529 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 2299713 | 4932 | 8948.3 | 455.96 | 169674643 | 6102500 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 824473 | 32103 | 3208.1 | 1271.81 | 473275550 | 10105847 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 902205 | 31310 | 3510.5 | 1162.23 | 432499214 | 15555221 | 624368.0 | 1769.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 823604 | 29116 | 3204.7 | 1273.15 | 473775211 | 17039749 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 5497960 | 31017 | 323409.4 | 190.72 | 70972321 | 2509658 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 2303733 | 18982 | 135513.7 | 455.16 | 169378600 | 5989410 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 2275469 | 16354 | 133851.1 | 460.82 | 171482469 | 3661663 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 2289571 | 9172 | 134680.6 | 457.98 | 170426260 | 6026457 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 824043 | 24283 | 48473.1 | 1272.47 | 473522669 | 10111124 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 873572 | 28996 | 51386.6 | 1200.33 | 446675467 | 15794927 | 637456.0 | 184.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 825165 | 25947 | 48539.1 | 1270.74 | 472878751 | 16721504 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 6948686 | 55046 | 6948685.9 | 150.90 | 56154934 | 1983397 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 2305982 | 4301 | 2305981.8 | 454.72 | 169213394 | 5976630 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 2299812 | 11109 | 2299811.7 | 455.94 | 169667366 | 3622905 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 2304706 | 5555 | 2304705.6 | 454.97 | 169307088 | 5979939 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 841919 | 26249 | 841918.8 | 1245.46 | 463468695 | 9896442 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 899292 | 28993 | 899291.9 | 1166.00 | 433900278 | 15325391 | 914672.0 | 20.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 845476 | 28308 | 845475.6 | 1240.22 | 461518946 | 16300885 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.591x |
| ar-wikipedia | 4096 | 1.473x |
| ar-wikipedia | 65536 | 1.462x |
| ar-wikipedia | 1048576 | 1.492x |
| en-wikitext2 | 256 | 4.688x |
| en-wikitext2 | 4096 | 5.019x |
| en-wikitext2 | 65536 | 5.016x |
| en-wikitext2 | 1048576 | 5.097x |
| ja-wikipedia | 256 | 2.720x |
| ja-wikipedia | 4096 | 2.624x |
| ja-wikipedia | 65536 | 2.637x |
| ja-wikipedia | 1048576 | 2.672x |
| ru-wikipedia | 256 | 1.570x |
| ru-wikipedia | 4096 | 1.464x |
| ru-wikipedia | 65536 | 1.448x |
| ru-wikipedia | 1048576 | 1.484x |
| synthetic | 0 | 5.151x |
| zh-wikipedia | 256 | 2.831x |
| zh-wikipedia | 4096 | 2.766x |
| zh-wikipedia | 65536 | 2.761x |
| zh-wikipedia | 1048576 | 2.732x |
| Combined selected natural corpora | 256 | 2.204x |
| Combined selected natural corpora | 4096 | 2.106x |
| Combined selected natural corpora | 65536 | 2.099x |
| Combined selected natural corpora | 1048576 | 2.131x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.594x |
| ar-wikipedia | 4096 | 1.480x |
| ar-wikipedia | 65536 | 1.470x |
| ar-wikipedia | 1048576 | 1.489x |
| en-wikitext2 | 256 | 4.684x |
| en-wikitext2 | 4096 | 5.037x |
| en-wikitext2 | 65536 | 4.983x |
| en-wikitext2 | 1048576 | 5.075x |
| ja-wikipedia | 256 | 2.743x |
| ja-wikipedia | 4096 | 2.637x |
| ja-wikipedia | 65536 | 2.659x |
| ja-wikipedia | 1048576 | 2.678x |
| ru-wikipedia | 256 | 1.581x |
| ru-wikipedia | 4096 | 1.470x |
| ru-wikipedia | 65536 | 1.455x |
| ru-wikipedia | 1048576 | 1.485x |
| synthetic | 0 | 4.937x |
| zh-wikipedia | 256 | 2.846x |
| zh-wikipedia | 4096 | 2.792x |
| zh-wikipedia | 65536 | 2.775x |
| zh-wikipedia | 1048576 | 2.726x |
| Combined selected natural corpora | 256 | 2.217x |
| Combined selected natural corpora | 4096 | 2.118x |
| Combined selected natural corpora | 65536 | 2.109x |
| Combined selected natural corpora | 1048576 | 2.129x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.555x |
| ar-wikipedia | 4096 | 1.459x |
| ar-wikipedia | 65536 | 1.450x |
| ar-wikipedia | 1048576 | 1.468x |
| en-wikitext2 | 256 | 3.849x |
| en-wikitext2 | 4096 | 4.564x |
| en-wikitext2 | 65536 | 4.552x |
| en-wikitext2 | 1048576 | 4.639x |
| ja-wikipedia | 256 | 2.438x |
| ja-wikipedia | 4096 | 2.469x |
| ja-wikipedia | 65536 | 2.512x |
| ja-wikipedia | 1048576 | 2.516x |
| ru-wikipedia | 256 | 1.545x |
| ru-wikipedia | 4096 | 1.460x |
| ru-wikipedia | 65536 | 1.440x |
| ru-wikipedia | 1048576 | 1.468x |
| synthetic | 0 | 3.200x |
| zh-wikipedia | 256 | 2.528x |
| zh-wikipedia | 4096 | 2.612x |
| zh-wikipedia | 65536 | 2.637x |
| zh-wikipedia | 1048576 | 2.564x |
| Combined selected natural corpora | 256 | 2.076x |
| Combined selected natural corpora | 4096 | 2.051x |
| Combined selected natural corpora | 65536 | 2.046x |
| Combined selected natural corpora | 1048576 | 2.059x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 4.000x |
| ar-wikipedia | 4096 | 3.215x |
| ar-wikipedia | 65536 | 3.091x |
| ar-wikipedia | 1048576 | 3.941x |
| en-wikitext2 | 256 | 12.582x |
| en-wikitext2 | 4096 | 11.433x |
| en-wikitext2 | 65536 | 11.592x |
| en-wikitext2 | 1048576 | 15.208x |
| ja-wikipedia | 256 | 7.617x |
| ja-wikipedia | 4096 | 6.603x |
| ja-wikipedia | 65536 | 6.563x |
| ja-wikipedia | 1048576 | 8.302x |
| ru-wikipedia | 256 | 3.932x |
| ru-wikipedia | 4096 | 3.175x |
| ru-wikipedia | 65536 | 3.065x |
| ru-wikipedia | 1048576 | 3.965x |
| synthetic | 0 | 15.884x |
| zh-wikipedia | 256 | 7.627x |
| zh-wikipedia | 4096 | 6.786x |
| zh-wikipedia | 65536 | 6.663x |
| zh-wikipedia | 1048576 | 8.219x |
| Combined selected natural corpora | 256 | 5.807x |
| Combined selected natural corpora | 4096 | 4.870x |
| Combined selected natural corpora | 65536 | 4.789x |
| Combined selected natural corpora | 1048576 | 6.117x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.744x |
| ar-wikipedia | 4096 | 3.119x |
| ar-wikipedia | 65536 | 3.040x |
| ar-wikipedia | 1048576 | 3.871x |
| en-wikitext2 | 256 | 9.756x |
| en-wikitext2 | 4096 | 10.037x |
| en-wikitext2 | 65536 | 10.515x |
| en-wikitext2 | 1048576 | 13.828x |
| ja-wikipedia | 256 | 6.384x |
| ja-wikipedia | 4096 | 6.058x |
| ja-wikipedia | 65536 | 6.166x |
| ja-wikipedia | 1048576 | 7.804x |
| ru-wikipedia | 256 | 3.724x |
| ru-wikipedia | 4096 | 3.115x |
| ru-wikipedia | 65536 | 3.023x |
| ru-wikipedia | 1048576 | 3.914x |
| synthetic | 0 | 9.428x |
| zh-wikipedia | 256 | 6.403x |
| zh-wikipedia | 4096 | 6.195x |
| zh-wikipedia | 65536 | 6.294x |
| zh-wikipedia | 1048576 | 7.727x |
| Combined selected natural corpora | 256 | 5.180x |
| Combined selected natural corpora | 4096 | 4.621x |
| Combined selected natural corpora | 65536 | 4.623x |
| Combined selected natural corpora | 1048576 | 5.901x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 314.7 | 2.3 | 10 |
| onig-compile | 300.0 | 3.1 | 10 |
| pretok-compile | 9.8 | 0.1 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
