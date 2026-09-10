# P09 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 38340838 | 205917 | 9342.3 | 27.35 | 15246798 | 3313412 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-at | 21375193 | 190967 | 5208.4 | 49.06 | 27348291 | 5943291 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 24223043 | 131725 | 5902.3 | 43.29 | 24133013 | 5244552 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 22873428 | 583035 | 5573.4 | 45.84 | 25556948 | 5553999 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 23153443 | 382265 | 5641.7 | 45.29 | 25247865 | 5486830 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 6407984 | 95074 | 1561.4 | 163.64 | 91226044 | 19825113 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 7411560 | 91584 | 1805.9 | 141.48 | 78873408 | 17140656 | 5425056.0 | 25855.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 6485750 | 179814 | 1580.3 | 161.67 | 90132209 | 19587402 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 33975974 | 185385 | 132202.2 | 30.86 | 17205541 | 3641514 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-at | 20767734 | 225156 | 80808.3 | 50.49 | 28148232 | 5957511 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 21925155 | 196309 | 85311.9 | 47.83 | 26662297 | 5643016 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 21620065 | 252302 | 84124.8 | 48.50 | 27038540 | 5722647 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 21684174 | 199463 | 84374.2 | 48.36 | 26958602 | 5705728 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 6247112 | 25131 | 24307.8 | 167.85 | 93575237 | 19804991 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 6824505 | 14501 | 26554.5 | 153.65 | 85658223 | 18129373 | 5045488.0 | 2628.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 6297998 | 20563 | 24505.8 | 166.49 | 92819183 | 19644974 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 33889301 | 1440778 | 1993488.3 | 30.94 | 17249544 | 3644454 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-at | 20470779 | 590588 | 1204163.5 | 51.22 | 28556558 | 6033381 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 21173232 | 808874 | 1245484.2 | 49.52 | 27609153 | 5833214 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 21165968 | 897949 | 1245056.9 | 49.54 | 27618628 | 5835216 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 21052139 | 919199 | 1238361.1 | 49.81 | 27767962 | 5866767 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 6266801 | 197198 | 368635.4 | 167.32 | 93281242 | 19708300 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 6878303 | 252743 | 404606.1 | 152.45 | 84988261 | 17956174 | 8081200.0 | 290.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 6255219 | 148719 | 367954.0 | 167.63 | 93453967 | 19744793 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 35938929 | 319772 | 35938928.6 | 29.18 | 16265788 | 3436190 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-at | 20068442 | 493917 | 20068441.7 | 52.25 | 29129068 | 6153592 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 20650394 | 76796 | 20650394.2 | 50.78 | 28308176 | 5980176 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 20562446 | 223161 | 20562445.9 | 50.99 | 28429254 | 6005754 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 20704888 | 128864 | 20704888.3 | 50.64 | 28233671 | 5964437 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 6196013 | 22755 | 6196012.6 | 169.23 | 94346968 | 19931044 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 6778493 | 69643 | 6778493.0 | 154.69 | 86239671 | 18218356 | 11302128.0 | 30.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 6232924 | 11243 | 6232923.9 | 168.23 | 93788246 | 19813013 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 53186235 | 563506 | 12981.8 | 19.72 | 19674959 | 4099501 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-at | 28855025 | 205634 | 7043.0 | 36.34 | 36265329 | 7556292 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 31863464 | 291664 | 7777.3 | 32.91 | 32841282 | 6842853 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 30693038 | 273468 | 7491.6 | 34.16 | 34093627 | 7103793 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 30917504 | 242094 | 7546.4 | 33.92 | 33846103 | 7052219 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 4788924 | 43427 | 1168.9 | 218.96 | 218511915 | 45529432 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 6186636 | 31739 | 1510.0 | 169.49 | 169144749 | 35243224 | 9046128.0 | 29028.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 4837556 | 20648 | 1180.8 | 216.76 | 216315218 | 45071725 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 49722345 | 367588 | 194227.9 | 21.09 | 21045609 | 4323951 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-at | 28302358 | 434765 | 110556.1 | 37.05 | 36973492 | 7596434 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 30269216 | 353928 | 118239.1 | 34.64 | 34570998 | 7102827 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 29636548 | 143936 | 115767.8 | 35.38 | 35309005 | 7254455 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 29735129 | 314899 | 116152.8 | 35.26 | 35191945 | 7230404 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 4688489 | 14026 | 18314.4 | 223.65 | 223192826 | 45856356 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 5619406 | 88739 | 21950.8 | 186.60 | 186218436 | 38259738 | 9977856.0 | 2929.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 4750314 | 43589 | 18555.9 | 220.74 | 220287961 | 45259534 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 49854738 | 1498543 | 3115921.1 | 21.03 | 20989720 | 4308698 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-at | 28475956 | 852798 | 1779747.3 | 36.82 | 36748090 | 7543522 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 29530325 | 659657 | 1845645.3 | 35.51 | 35436014 | 7274183 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 29319052 | 603213 | 1832440.8 | 35.76 | 35691365 | 7326601 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 29672096 | 329205 | 1854506.0 | 35.34 | 35266704 | 7239428 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 4712183 | 14754 | 294511.4 | 222.52 | 222070543 | 45585880 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 5672069 | 44854 | 354504.3 | 184.87 | 184489483 | 37871369 | 14634752.0 | 320.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 4793020 | 28244 | 299563.7 | 218.77 | 218325197 | 44817048 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 55454177 | 2932208 | 55454177.1 | 18.91 | 18870301 | 3873378 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-at | 29216106 | 466515 | 29216106.5 | 35.89 | 35817127 | 7351938 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 30366542 | 732907 | 30366541.7 | 34.53 | 34460197 | 7073410 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 30206388 | 1761658 | 30206388.5 | 34.71 | 34642903 | 7110913 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 30532795 | 364259 | 30532795.3 | 34.34 | 34272558 | 7034895 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 4795975 | 142652 | 4795974.7 | 218.64 | 218190684 | 44786517 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 5623648 | 95132 | 5623647.6 | 186.46 | 186077983 | 38194961 | 17962224.0 | 32.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 4803281 | 108926 | 4803281.0 | 218.30 | 217858795 | 44718392 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 26771880 | 196712 | 6509.1 | 39.17 | 14167103 | 2245565 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-at | 15414069 | 150194 | 3747.6 | 68.03 | 24606092 | 3900203 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 17554322 | 185949 | 4268.0 | 59.73 | 21606075 | 3424684 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 16610402 | 178777 | 4038.5 | 63.13 | 22833885 | 3619298 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 16793856 | 143654 | 4083.1 | 62.44 | 22584450 | 3579761 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 5766005 | 33328 | 1401.9 | 181.85 | 65778643 | 10426283 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 6614653 | 43230 | 1608.2 | 158.52 | 57339367 | 9088610 | 2600816.0 | 20602.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 5804051 | 63103 | 1411.1 | 180.66 | 65347458 | 10357937 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 23660378 | 46843 | 92063.7 | 44.32 | 16030175 | 2398356 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-at | 15022810 | 80658 | 58454.5 | 69.80 | 25246941 | 3777323 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 15689763 | 72090 | 61049.7 | 66.83 | 24173724 | 3616753 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 15411251 | 101383 | 59966.0 | 68.04 | 24610592 | 3682115 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 15469573 | 176679 | 60192.9 | 67.78 | 24517806 | 3668233 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 5774927 | 22614 | 22470.5 | 181.57 | 65677024 | 9826272 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 6113631 | 62766 | 23788.4 | 171.51 | 62038420 | 9281882 | 2591472.0 | 2354.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 5768056 | 62405 | 22443.8 | 181.79 | 65755259 | 9837977 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 23934981 | 137752 | 1407940.1 | 43.81 | 15846263 | 2361815 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-at | 14799061 | 209971 | 870533.0 | 70.85 | 25628654 | 3819837 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 15388641 | 235334 | 905214.2 | 68.14 | 24646751 | 3673489 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 15255585 | 180746 | 897387.4 | 68.73 | 24861714 | 3705528 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 15414918 | 174087 | 906759.9 | 68.02 | 24604737 | 3667227 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 5737906 | 75934 | 337523.9 | 182.75 | 66100769 | 9852026 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 6126718 | 122843 | 360395.2 | 171.15 | 61905899 | 9226799 | 3493648.0 | 249.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 5823181 | 105711 | 342540.1 | 180.07 | 65132787 | 9707753 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 23594948 | 1409496 | 23594947.9 | 44.44 | 16074628 | 2395174 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-at | 13477200 | 1077096 | 13477200.1 | 77.80 | 28142344 | 4193304 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 13830637 | 934723 | 13830637.4 | 75.82 | 27423176 | 4086146 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 14237679 | 1045136 | 14237678.9 | 73.65 | 26639174 | 3969327 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 14451607 | 798778 | 14451606.5 | 72.56 | 26244833 | 3910569 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 5413907 | 310133 | 5413906.6 | 193.68 | 70056621 | 10438673 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 5679520 | 198018 | 5679520.0 | 184.62 | 66780291 | 9950489 | 4363504.0 | 26.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 5425172 | 228417 | 5425172.5 | 193.28 | 69911141 | 10416996 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 33497640 | 877867 | 8162.2 | 31.30 | 17161119 | 3167746 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-at | 17785006 | 91150 | 4333.6 | 58.96 | 32322563 | 5966374 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 19850589 | 183100 | 4836.9 | 52.82 | 28959191 | 5345534 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 19015126 | 230886 | 4633.3 | 55.14 | 30231564 | 5580400 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 19069007 | 74056 | 4646.4 | 54.99 | 30146142 | 5564632 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 4800624 | 27092 | 1169.7 | 218.42 | 119746298 | 22103791 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 5781455 | 60377 | 1408.7 | 181.37 | 99431203 | 18353858 | 4684160.0 | 25105.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 4840739 | 27148 | 1179.5 | 216.61 | 118753981 | 21920621 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 29424865 | 168466 | 114493.6 | 35.64 | 19536436 | 3489362 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-at | 17206076 | 160156 | 66949.7 | 60.94 | 33410116 | 5967310 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 18131987 | 189052 | 70552.5 | 57.83 | 31704027 | 5662590 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 17798207 | 137735 | 69253.7 | 58.91 | 32298591 | 5768783 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 17758532 | 89789 | 69099.3 | 59.05 | 32370751 | 5781672 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 4711153 | 18748 | 18331.3 | 222.57 | 122020442 | 21793815 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 5213213 | 51853 | 20284.9 | 201.14 | 110269223 | 19694954 | 4299760.0 | 2574.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 4755941 | 26789 | 18505.6 | 220.48 | 120871357 | 21588579 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 29329872 | 759722 | 1725286.6 | 35.75 | 19599710 | 3493401 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-at | 17340648 | 392515 | 1020038.1 | 60.47 | 33150837 | 5908718 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 17855615 | 594430 | 1050330.3 | 58.73 | 32194746 | 5738307 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 17735667 | 455073 | 1043274.5 | 59.12 | 32412483 | 5777116 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 17772554 | 542651 | 1045444.4 | 59.00 | 32345210 | 5765125 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 4660588 | 12009 | 274152.2 | 224.99 | 123344309 | 21984565 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 5190574 | 20456 | 305327.9 | 202.02 | 110750172 | 19739820 | 6549264.0 | 278.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 4736923 | 24941 | 278642.5 | 221.36 | 121356619 | 21630285 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 34821975 | 1509692 | 34821975.0 | 30.11 | 16508455 | 2942022 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-at | 18730671 | 1030788 | 18730671.3 | 55.98 | 30690678 | 5469478 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 19381285 | 1238233 | 19381285.3 | 54.10 | 29660417 | 5285872 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 19118448 | 1407594 | 19118447.5 | 54.85 | 30068184 | 5358542 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 19343153 | 1271805 | 19343152.7 | 54.21 | 29718889 | 5296293 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 4765760 | 149618 | 4765759.7 | 220.02 | 120622322 | 21496468 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 5312890 | 187478 | 5312889.7 | 197.36 | 108200440 | 19282727 | 8942832.0 | 29.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 4822328 | 146943 | 4822327.5 | 217.44 | 119207374 | 21244306 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 80287429 | 1649693 | 1844.5 | 13.08 | 7695651 | 2937122 | N/A | N/A | 10 |
| synthetic | 0 | onig-at | 36924115 | 1964776 | 848.3 | 28.45 | 16733346 | 6386450 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 47198932 | 648085 | 1084.4 | 22.26 | 13090635 | 4996172 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 45302852 | 381484 | 1040.8 | 23.19 | 13638523 | 5205279 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 46352655 | 779554 | 1064.9 | 22.66 | 13329636 | 5087389 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 9563855 | 150529 | 219.7 | 109.84 | 64604072 | 24656793 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 12122197 | 302179 | 278.5 | 86.66 | 50969639 | 19453075 | 10138304.0 | 117364.0 | 10 |
| synthetic | 0 | pretok-reused | 9719214 | 103663 | 223.3 | 108.08 | 63571394 | 24262661 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 23834536 | 183559 | 5794.9 | 43.99 | 16371328 | 2753609 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-at | 12941430 | 77971 | 3146.5 | 81.02 | 30151459 | 5071387 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 14621899 | 188539 | 3555.0 | 71.71 | 26686205 | 4488541 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 14004513 | 71147 | 3404.9 | 74.87 | 27862661 | 4686418 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 14049566 | 92393 | 3415.9 | 74.63 | 27773313 | 4671390 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 4112180 | 5952 | 999.8 | 254.99 | 94889560 | 15960146 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 4848264 | 16729 | 1178.8 | 216.28 | 80483038 | 13537011 | 2889584.0 | 21367.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 4140865 | 9754 | 1006.8 | 253.23 | 94232236 | 15849586 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 20899253 | 160843 | 81320.0 | 50.17 | 18670668 | 2984317 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-at | 12507420 | 114376 | 48667.0 | 83.84 | 31197722 | 4986640 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 12979670 | 54050 | 50504.6 | 80.79 | 30062629 | 4805207 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 12874767 | 77906 | 50096.4 | 81.44 | 30307577 | 4844359 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 12833931 | 36707 | 49937.5 | 81.70 | 30404014 | 4859774 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 4036292 | 17335 | 15705.4 | 259.79 | 96673630 | 15452301 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 4383655 | 22693 | 17057.0 | 239.20 | 89013163 | 14227853 | 2891760.0 | 2394.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 4050384 | 19603 | 15760.3 | 258.88 | 96337276 | 15398538 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 20494344 | 186475 | 1205549.7 | 51.16 | 19039546 | 3032886 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-at | 12355758 | 71382 | 726809.3 | 84.87 | 31580660 | 5030610 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 12641024 | 73282 | 743589.6 | 82.95 | 30867990 | 4917086 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 12742421 | 81886 | 749554.2 | 82.29 | 30622359 | 4877958 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 12617584 | 108452 | 742210.8 | 83.10 | 30925334 | 4926220 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 4019734 | 13771 | 236454.9 | 260.86 | 97071850 | 15462964 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 4323582 | 27999 | 254328.4 | 242.52 | 90249935 | 14376274 | 3764016.0 | 253.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 4047075 | 17269 | 238063.2 | 259.09 | 96416055 | 15358500 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 22682598 | 607753 | 22682598.1 | 46.23 | 17202747 | 2739810 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-at | 12358735 | 109475 | 12358734.7 | 84.84 | 31573054 | 5028508 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 12751047 | 121971 | 12751046.9 | 82.23 | 30601644 | 4873796 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 12679130 | 146258 | 12679129.7 | 82.70 | 30775220 | 4901441 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 12682720 | 167716 | 12682719.8 | 82.68 | 30766508 | 4900053 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 4016532 | 15280 | 4016531.5 | 261.06 | 97149244 | 15472554 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 4361573 | 32436 | 4361573.4 | 240.41 | 89463817 | 14248528 | 5559536.0 | 27.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 4048636 | 25843 | 4048635.8 | 258.99 | 96378885 | 15349862 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.570x |
| ar-wikipedia | 4096 | 3.461x |
| ar-wikipedia | 65536 | 3.377x |
| ar-wikipedia | 1048576 | 3.319x |
| en-wikitext2 | 256 | 6.409x |
| en-wikitext2 | 4096 | 6.321x |
| en-wikitext2 | 65536 | 6.222x |
| en-wikitext2 | 1048576 | 6.298x |
| ja-wikipedia | 256 | 2.881x |
| ja-wikipedia | 4096 | 2.669x |
| ja-wikipedia | 65536 | 2.659x |
| ja-wikipedia | 1048576 | 2.630x |
| ru-wikipedia | 256 | 3.961x |
| ru-wikipedia | 4096 | 3.778x |
| ru-wikipedia | 65536 | 3.805x |
| ru-wikipedia | 1048576 | 4.012x |
| synthetic | 0 | 4.737x |
| zh-wikipedia | 256 | 3.406x |
| zh-wikipedia | 4096 | 3.190x |
| zh-wikipedia | 65536 | 3.170x |
| zh-wikipedia | 1048576 | 3.157x |
| Combined selected natural corpora | 256 | 3.988x |
| Combined selected natural corpora | 4096 | 3.824x |
| Combined selected natural corpora | 65536 | 3.789x |
| Combined selected natural corpora | 1048576 | 3.843x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.570x |
| ar-wikipedia | 4096 | 3.443x |
| ar-wikipedia | 65536 | 3.366x |
| ar-wikipedia | 1048576 | 3.322x |
| en-wikitext2 | 256 | 6.391x |
| en-wikitext2 | 4096 | 6.260x |
| en-wikitext2 | 65536 | 6.191x |
| en-wikitext2 | 1048576 | 6.357x |
| ja-wikipedia | 256 | 2.893x |
| ja-wikipedia | 4096 | 2.682x |
| ja-wikipedia | 65536 | 2.647x |
| ja-wikipedia | 1048576 | 2.664x |
| ru-wikipedia | 256 | 3.939x |
| ru-wikipedia | 4096 | 3.734x |
| ru-wikipedia | 65536 | 3.752x |
| ru-wikipedia | 1048576 | 4.011x |
| synthetic | 0 | 4.769x |
| zh-wikipedia | 256 | 3.393x |
| zh-wikipedia | 4096 | 3.169x |
| zh-wikipedia | 65536 | 3.118x |
| zh-wikipedia | 1048576 | 3.133x |
| Combined selected natural corpora | 256 | 3.983x |
| Combined selected natural corpora | 4096 | 3.804x |
| Combined selected natural corpora | 65536 | 3.763x |
| Combined selected natural corpora | 1048576 | 3.857x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.268x |
| ar-wikipedia | 4096 | 3.213x |
| ar-wikipedia | 65536 | 3.078x |
| ar-wikipedia | 1048576 | 3.046x |
| en-wikitext2 | 256 | 5.150x |
| en-wikitext2 | 4096 | 5.387x |
| en-wikitext2 | 65536 | 5.206x |
| en-wikitext2 | 1048576 | 5.400x |
| ja-wikipedia | 256 | 2.654x |
| ja-wikipedia | 4096 | 2.566x |
| ja-wikipedia | 65536 | 2.512x |
| ja-wikipedia | 1048576 | 2.435x |
| ru-wikipedia | 256 | 3.433x |
| ru-wikipedia | 4096 | 3.478x |
| ru-wikipedia | 65536 | 3.440x |
| ru-wikipedia | 1048576 | 3.648x |
| synthetic | 0 | 3.894x |
| zh-wikipedia | 256 | 3.016x |
| zh-wikipedia | 4096 | 2.961x |
| zh-wikipedia | 65536 | 2.924x |
| zh-wikipedia | 1048576 | 2.923x |
| Combined selected natural corpora | 256 | 3.505x |
| Combined selected natural corpora | 4096 | 3.516x |
| Combined selected natural corpora | 65536 | 3.426x |
| Combined selected natural corpora | 1048576 | 3.494x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 5.912x |
| ar-wikipedia | 4096 | 5.395x |
| ar-wikipedia | 65536 | 5.418x |
| ar-wikipedia | 1048576 | 5.766x |
| en-wikitext2 | 256 | 10.994x |
| en-wikitext2 | 4096 | 10.467x |
| en-wikitext2 | 65536 | 10.402x |
| en-wikitext2 | 1048576 | 11.545x |
| ja-wikipedia | 256 | 4.613x |
| ja-wikipedia | 4096 | 4.102x |
| ja-wikipedia | 65536 | 4.110x |
| ja-wikipedia | 1048576 | 4.349x |
| ru-wikipedia | 256 | 6.920x |
| ru-wikipedia | 4096 | 6.187x |
| ru-wikipedia | 65536 | 6.192x |
| ru-wikipedia | 1048576 | 7.221x |
| synthetic | 0 | 8.261x |
| zh-wikipedia | 256 | 5.756x |
| zh-wikipedia | 4096 | 5.160x |
| zh-wikipedia | 65536 | 5.064x |
| zh-wikipedia | 1048576 | 5.603x |
| Combined selected natural corpora | 256 | 6.727x |
| Combined selected natural corpora | 4096 | 6.154x |
| Combined selected natural corpora | 65536 | 6.139x |
| Combined selected natural corpora | 1048576 | 6.809x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 5.173x |
| ar-wikipedia | 4096 | 4.979x |
| ar-wikipedia | 65536 | 4.927x |
| ar-wikipedia | 1048576 | 5.302x |
| en-wikitext2 | 256 | 8.597x |
| en-wikitext2 | 4096 | 8.848x |
| en-wikitext2 | 65536 | 8.790x |
| en-wikitext2 | 1048576 | 9.861x |
| ja-wikipedia | 256 | 4.047x |
| ja-wikipedia | 4096 | 3.870x |
| ja-wikipedia | 65536 | 3.907x |
| ja-wikipedia | 1048576 | 4.154x |
| ru-wikipedia | 256 | 5.794x |
| ru-wikipedia | 4096 | 5.644x |
| ru-wikipedia | 65536 | 5.651x |
| ru-wikipedia | 1048576 | 6.554x |
| synthetic | 0 | 6.623x |
| zh-wikipedia | 256 | 4.916x |
| zh-wikipedia | 4096 | 4.768x |
| zh-wikipedia | 65536 | 4.740x |
| zh-wikipedia | 1048576 | 5.201x |
| Combined selected natural corpora | 256 | 5.694x |
| Combined selected natural corpora | 4096 | 5.601x |
| Combined selected natural corpora | 65536 | 5.587x |
| Combined selected natural corpora | 1048576 | 6.215x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 60941.7 | 768.3 | 10 |
| onig-compile | 61106.8 | 313.2 | 10 |
| pretok-compile | 12.7 | 0.3 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
