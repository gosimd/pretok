# P01 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 36713580 | 1428202 | 8945.8 | 28.56 | 15922582 | 3326208 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-at | 20181383 | 281550 | 4917.5 | 51.96 | 28966052 | 6050973 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 23205503 | 462644 | 5654.4 | 45.19 | 25191223 | 5262416 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 22056745 | 126457 | 5374.5 | 47.54 | 26503230 | 5536492 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 22113140 | 120622 | 5388.2 | 47.42 | 26435639 | 5522373 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 5773520 | 18430 | 1406.8 | 181.62 | 101251055 | 21151221 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 6591304 | 53594 | 1606.1 | 159.08 | 88688824 | 18526986 | 5148576.0 | 25591.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 5812700 | 42230 | 1416.4 | 180.39 | 100568574 | 21008652 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 32220906 | 229419 | 125373.2 | 32.54 | 18142724 | 3686054 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-at | 19389798 | 243336 | 75446.7 | 54.08 | 30148586 | 6125283 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 21153129 | 183910 | 82307.9 | 49.57 | 27635391 | 5614678 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 20642759 | 167840 | 80322.0 | 50.80 | 28318647 | 5753494 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 20862480 | 198324 | 81177.0 | 50.26 | 28020398 | 5692899 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 5692760 | 46813 | 22150.8 | 184.19 | 102687445 | 20862990 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 6065635 | 41786 | 23601.7 | 172.87 | 96374901 | 19580472 | 4719856.0 | 2604.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 5733500 | 50470 | 22309.3 | 182.89 | 101957800 | 20714748 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 31502585 | 214971 | 1853093.2 | 33.29 | 18556414 | 3763247 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-at | 19382612 | 288689 | 1140153.7 | 54.10 | 30159764 | 6116410 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 20359179 | 102846 | 1197598.8 | 51.50 | 28713093 | 5823025 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 20280168 | 188017 | 1192951.1 | 51.70 | 28824958 | 5845711 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 20251328 | 123097 | 1191254.6 | 51.78 | 28866008 | 5854036 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 5682867 | 4362 | 334286.3 | 184.52 | 102866204 | 20861300 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 6036252 | 21402 | 355073.6 | 173.71 | 96844035 | 19640002 | 8081200.0 | 290.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 5730291 | 20252 | 337076.0 | 182.99 | 102014881 | 20688651 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 36012228 | 182810 | 36012227.7 | 29.12 | 16232681 | 3291576 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-at | 19611575 | 246631 | 19611574.5 | 53.47 | 29807653 | 6044237 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 20419907 | 191055 | 20419906.7 | 51.35 | 28627702 | 5804973 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 20473287 | 147851 | 20473287.5 | 51.22 | 28553060 | 5789837 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 20381602 | 174652 | 20381601.7 | 51.45 | 28681504 | 5815882 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 5682809 | 30653 | 5682808.7 | 184.52 | 102867267 | 20858876 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 6134408 | 67291 | 6134407.8 | 170.93 | 95294448 | 19323300 | 11302128.0 | 30.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 5745857 | 46229 | 5745857.2 | 182.49 | 101738518 | 20629994 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 52484150 | 532455 | 12810.4 | 19.98 | 19938153 | 4154016 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-at | 28695862 | 133466 | 7004.1 | 36.54 | 36466477 | 7597611 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 33296013 | 670108 | 8126.9 | 31.49 | 31428298 | 6547931 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 31767809 | 871132 | 7753.9 | 33.01 | 32940170 | 6862922 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 31732850 | 399655 | 7745.4 | 33.04 | 32976458 | 6870483 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 4665894 | 76279 | 1138.9 | 224.73 | 224273656 | 46726313 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 5633720 | 27578 | 1375.1 | 186.12 | 185745297 | 38699119 | 9044080.0 | 29027.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 4698820 | 133320 | 1146.9 | 223.16 | 222702069 | 46398880 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 49549324 | 1203921 | 193552.0 | 21.16 | 21119097 | 4338687 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-at | 29005084 | 903444 | 113301.1 | 36.15 | 36077709 | 7411769 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 31836042 | 816305 | 124359.5 | 32.94 | 32869570 | 6752693 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 30657017 | 746008 | 119754.0 | 34.20 | 34133686 | 7012391 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 30658300 | 747492 | 119759.0 | 34.20 | 34132258 | 7012098 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 4570313 | 59712 | 17852.8 | 229.43 | 228963965 | 47038134 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 5172135 | 185661 | 20203.7 | 202.74 | 202322057 | 41564847 | 9977856.0 | 2929.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 4608051 | 74273 | 18000.2 | 227.55 | 227088839 | 46652910 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 49666822 | 301061 | 3104176.4 | 21.11 | 21069135 | 4324637 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-at | 28521632 | 234780 | 1782602.0 | 36.76 | 36689240 | 7530810 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 30283892 | 286548 | 1892743.3 | 34.62 | 34554244 | 7092582 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 30060757 | 312724 | 1878797.3 | 34.88 | 34810733 | 7145229 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 30226138 | 338426 | 1889133.7 | 34.69 | 34620268 | 7106134 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 4630613 | 15238 | 289413.3 | 226.44 | 225982408 | 46385007 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 5250582 | 28641 | 328161.3 | 199.71 | 199299255 | 40908040 | 14634752.0 | 320.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 4681299 | 20827 | 292581.2 | 223.99 | 223535601 | 45882777 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 54842092 | 1246638 | 54842091.6 | 19.12 | 19080910 | 3916280 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-at | 28861043 | 376163 | 28861042.8 | 36.33 | 36257768 | 7441762 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 30657725 | 331344 | 30657725.5 | 34.20 | 34132897 | 7005640 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 30274724 | 265539 | 30274724.3 | 34.64 | 34564708 | 7094268 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 30662160 | 177218 | 30662160.5 | 34.20 | 34127960 | 7004627 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 4643710 | 32124 | 4643710.1 | 225.81 | 225345033 | 46251165 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 5344151 | 13401 | 5344151.4 | 196.21 | 195809760 | 40189169 | 17962224.0 | 32.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 4712295 | 23576 | 4712295.4 | 222.52 | 222065240 | 45578000 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 22801295 | 78502 | 5543.7 | 45.99 | 16634143 | 2636605 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-at | 12195607 | 94113 | 2965.1 | 85.98 | 31099722 | 4929480 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 14372697 | 96456 | 3494.5 | 72.96 | 26388923 | 4182792 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 13594566 | 33387 | 3305.3 | 77.13 | 27899382 | 4422208 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 13760925 | 43691 | 3345.7 | 76.20 | 27562100 | 4368747 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 5399237 | 15438 | 1312.7 | 194.21 | 70246968 | 11134537 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 5974145 | 10085 | 1452.5 | 175.52 | 63486905 | 10063029 | 2600816.0 | 20602.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 5427803 | 15446 | 1319.7 | 193.19 | 69877259 | 11075936 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 19494051 | 93929 | 75852.3 | 53.79 | 19456192 | 2910939 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-at | 11705309 | 95064 | 45545.9 | 89.58 | 32402391 | 4847886 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 12605235 | 71505 | 49047.6 | 83.19 | 30089086 | 4501780 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 12351270 | 67386 | 48059.4 | 84.90 | 30707773 | 4594345 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 12355653 | 46894 | 48076.5 | 84.87 | 30696880 | 4592716 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 5296264 | 19707 | 20608.0 | 197.98 | 71612738 | 10714344 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 5544515 | 37012 | 21574.0 | 189.12 | 68406345 | 10234619 | 2591472.0 | 2354.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 5312748 | 33461 | 20672.2 | 197.37 | 71390552 | 10681102 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 19211556 | 80053 | 1130091.5 | 54.58 | 19742285 | 2942500 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-at | 11605240 | 41481 | 682661.2 | 90.35 | 32681790 | 4871076 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 12078323 | 38521 | 710489.6 | 86.81 | 31401709 | 4680285 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 12010509 | 33646 | 706500.5 | 87.30 | 31579012 | 4706711 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 12024970 | 54817 | 707351.2 | 87.20 | 31541034 | 4701051 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 5292308 | 25141 | 311312.2 | 198.13 | 71666276 | 10681540 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 5483338 | 31307 | 322549.3 | 191.23 | 69169548 | 10309414 | 3493648.0 | 249.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 5309039 | 21218 | 312296.4 | 197.51 | 71440421 | 10647878 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 21946837 | 691305 | 21946837.0 | 47.78 | 17281761 | 2575041 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-at | 11638589 | 35526 | 11638588.9 | 90.09 | 32588143 | 4855743 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 12096758 | 71242 | 12096757.9 | 86.68 | 31353855 | 4671830 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 12089854 | 152950 | 12089853.7 | 86.73 | 31371761 | 4674498 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 12136565 | 57646 | 12136565.0 | 86.40 | 31251017 | 4656507 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 5296512 | 10432 | 5296511.6 | 197.97 | 71609396 | 10670042 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 5499948 | 23212 | 5499948.5 | 190.65 | 68960646 | 10275369 | 4363504.0 | 26.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 5321969 | 24151 | 5321968.5 | 197.03 | 71266863 | 10619003 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 34206550 | 165859 | 8334.9 | 30.65 | 16805466 | 3099173 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-at | 18558291 | 266815 | 4522.0 | 56.50 | 30975751 | 5712379 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 21633665 | 468319 | 5271.4 | 48.47 | 26572335 | 4900326 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 20461916 | 357468 | 4985.8 | 51.25 | 28093997 | 5180942 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 20536553 | 381618 | 5004.0 | 51.06 | 27991893 | 5162113 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 6036758 | 383814 | 1470.9 | 173.70 | 95226113 | 17561082 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 6797938 | 411047 | 1656.4 | 154.25 | 84563437 | 15594729 | 4675456.0 | 25096.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 6078204 | 378441 | 1481.0 | 172.51 | 94576786 | 17441336 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 29825336 | 231645 | 116051.9 | 35.16 | 19274117 | 3439056 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-at | 17684978 | 140318 | 68813.1 | 59.29 | 32505383 | 5799894 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 19190715 | 155387 | 74672.0 | 54.64 | 29954956 | 5344824 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 18801671 | 75569 | 73158.3 | 55.77 | 30574782 | 5455419 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 18858158 | 38718 | 73378.0 | 55.60 | 30483200 | 5439078 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 5763284 | 4587 | 22425.2 | 181.94 | 99744702 | 17797320 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 6122340 | 21037 | 23822.3 | 171.27 | 93894986 | 16753562 | 4299760.0 | 2574.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 5801239 | 22186 | 22572.9 | 180.75 | 99092108 | 17680878 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 29394407 | 112416 | 1729082.8 | 35.67 | 19556679 | 3482227 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-at | 17602887 | 180090 | 1035464.0 | 59.57 | 32656972 | 5814842 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 18408634 | 117931 | 1082860.8 | 56.96 | 31227575 | 5560326 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 18305746 | 92032 | 1076808.6 | 57.28 | 31403090 | 5591578 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 18358849 | 124820 | 1079932.3 | 57.12 | 31312257 | 5575404 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 5755770 | 9573 | 338574.7 | 182.18 | 99874910 | 17783546 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 6019918 | 20974 | 354112.8 | 174.18 | 95492498 | 17003222 | 6549264.0 | 278.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 5778895 | 5715 | 339935.0 | 181.45 | 99475248 | 17712383 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 33387356 | 311677 | 33387355.5 | 31.41 | 17217806 | 3065352 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-at | 17652627 | 154005 | 17652627.2 | 59.40 | 32564954 | 5797664 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 18469983 | 139397 | 18469982.9 | 56.77 | 31123851 | 5541099 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 18510230 | 130026 | 18510229.9 | 56.65 | 31056178 | 5529051 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 18529775 | 85385 | 18529774.7 | 56.59 | 31023421 | 5523219 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 5736227 | 9396 | 5736226.6 | 182.80 | 100215184 | 17841694 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 6112021 | 32716 | 6112020.8 | 171.56 | 94053508 | 16744707 | 8942832.0 | 29.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 5777095 | 35540 | 5777094.8 | 181.51 | 99506243 | 17715479 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 9368489 | 106278 | 24397.1 | 28.84 | 18549843 | 4161610 | N/A | N/A | 10 |
| synthetic | 0 | onig-at | 5648218 | 45179 | 14708.9 | 47.84 | 30767935 | 6902708 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 6044146 | 64674 | 15740.0 | 44.71 | 28752447 | 6450539 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 5983457 | 36595 | 15581.9 | 45.16 | 29044079 | 6515966 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 5971694 | 40507 | 15551.3 | 45.25 | 29101290 | 6528801 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 1034817 | 4966 | 2694.8 | 261.13 | 167936881 | 37676214 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 1148604 | 2121 | 2991.2 | 235.26 | 151300203 | 33943817 | 2316608.0 | 1050.0 | 10 |
| synthetic | 0 | pretok-reused | 1045763 | 10158 | 2723.3 | 258.40 | 166179067 | 37281853 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 24238768 | 68216 | 5893.2 | 43.26 | 16098302 | 2707687 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-at | 13056941 | 174790 | 3174.6 | 80.31 | 29884719 | 5026522 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 15281221 | 99105 | 3715.3 | 68.62 | 25534805 | 4294879 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 14537809 | 168971 | 3534.6 | 72.13 | 26840565 | 4514504 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 14721281 | 115928 | 3579.2 | 71.23 | 26506049 | 4458240 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 3930929 | 13091 | 955.7 | 266.75 | 99264821 | 16696052 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 4549184 | 22742 | 1106.1 | 230.50 | 85774275 | 14426981 | 2889584.0 | 21367.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 3975009 | 4508 | 966.4 | 263.79 | 98164064 | 16510908 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 20964360 | 174701 | 81573.4 | 50.02 | 18612683 | 2975049 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-at | 12641675 | 142697 | 49189.4 | 82.95 | 30866400 | 4933682 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 13514057 | 73729 | 52583.9 | 77.59 | 28873862 | 4615195 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 13200037 | 85349 | 51362.0 | 79.44 | 29560750 | 4724987 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 13339610 | 110470 | 51905.1 | 78.61 | 29251456 | 4675549 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 3844582 | 5269 | 14959.5 | 272.74 | 101494261 | 16222830 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 4104184 | 14707 | 15969.6 | 255.49 | 95074443 | 15196687 | 2891760.0 | 2394.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 3869073 | 12441 | 15054.8 | 271.01 | 100851796 | 16120139 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 20736965 | 76053 | 1219821.5 | 50.57 | 18816784 | 2997401 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-at | 12710314 | 298860 | 747665.5 | 82.50 | 30699713 | 4890280 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 13233773 | 216936 | 778457.2 | 79.23 | 29485393 | 4696846 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 12998409 | 192245 | 764612.3 | 80.67 | 30019290 | 4781893 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 12998699 | 100972 | 764629.3 | 80.67 | 30018620 | 4781786 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 3846580 | 10293 | 226269.4 | 272.60 | 101441550 | 16159031 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 4046416 | 9228 | 238024.5 | 259.14 | 96431744 | 15360999 | 3764016.0 | 253.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 3868004 | 22985 | 227529.6 | 271.09 | 100879683 | 16069529 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 23524687 | 236833 | 23524686.5 | 44.57 | 16586959 | 2641736 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-at | 12586576 | 158293 | 12586576.0 | 83.31 | 31001521 | 4937483 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 13038259 | 109744 | 13038258.5 | 80.42 | 29927540 | 4766434 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 13010314 | 89992 | 13010313.6 | 80.60 | 29991821 | 4776672 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 13030497 | 94578 | 13030497.3 | 80.47 | 29945365 | 4769273 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 3851306 | 9735 | 3851305.9 | 272.26 | 101317062 | 16136345 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 4077337 | 11784 | 4077336.5 | 257.17 | 95700464 | 15241813 | 5559536.0 | 27.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 3876775 | 14573 | 3876774.5 | 270.48 | 100651455 | 16030336 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.820x |
| ar-wikipedia | 4096 | 3.626x |
| ar-wikipedia | 65536 | 3.569x |
| ar-wikipedia | 1048576 | 3.603x |
| en-wikitext2 | 256 | 6.809x |
| en-wikitext2 | 4096 | 6.708x |
| en-wikitext2 | 65536 | 6.492x |
| en-wikitext2 | 1048576 | 6.520x |
| ja-wikipedia | 256 | 2.518x |
| ja-wikipedia | 4096 | 2.332x |
| ja-wikipedia | 65536 | 2.269x |
| ja-wikipedia | 1048576 | 2.283x |
| ru-wikipedia | 256 | 3.390x |
| ru-wikipedia | 4096 | 3.262x |
| ru-wikipedia | 65536 | 3.180x |
| ru-wikipedia | 1048576 | 3.227x |
| synthetic | 0 | 5.782x |
| zh-wikipedia | 256 | 3.698x |
| zh-wikipedia | 4096 | 3.433x |
| zh-wikipedia | 65536 | 3.379x |
| zh-wikipedia | 1048576 | 3.378x |
| Combined selected natural corpora | 256 | 3.969x |
| Combined selected natural corpora | 4096 | 3.801x |
| Combined selected natural corpora | 65536 | 3.715x |
| Combined selected natural corpora | 1048576 | 3.743x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.804x |
| ar-wikipedia | 4096 | 3.639x |
| ar-wikipedia | 65536 | 3.534x |
| ar-wikipedia | 1048576 | 3.547x |
| en-wikitext2 | 256 | 6.753x |
| en-wikitext2 | 4096 | 6.653x |
| en-wikitext2 | 65536 | 6.457x |
| en-wikitext2 | 1048576 | 6.507x |
| ja-wikipedia | 256 | 2.535x |
| ja-wikipedia | 4096 | 2.326x |
| ja-wikipedia | 65536 | 2.265x |
| ja-wikipedia | 1048576 | 2.280x |
| ru-wikipedia | 256 | 3.379x |
| ru-wikipedia | 4096 | 3.251x |
| ru-wikipedia | 65536 | 3.177x |
| ru-wikipedia | 1048576 | 3.207x |
| synthetic | 0 | 5.710x |
| zh-wikipedia | 256 | 3.703x |
| zh-wikipedia | 4096 | 3.448x |
| zh-wikipedia | 65536 | 3.361x |
| zh-wikipedia | 1048576 | 3.361x |
| Combined selected natural corpora | 256 | 3.957x |
| Combined selected natural corpora | 4096 | 3.794x |
| Combined selected natural corpora | 65536 | 3.700x |
| Combined selected natural corpora | 1048576 | 3.725x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.521x |
| ar-wikipedia | 4096 | 3.487x |
| ar-wikipedia | 65536 | 3.373x |
| ar-wikipedia | 1048576 | 3.329x |
| en-wikitext2 | 256 | 5.910x |
| en-wikitext2 | 4096 | 6.155x |
| en-wikitext2 | 65536 | 5.768x |
| en-wikitext2 | 1048576 | 5.737x |
| ja-wikipedia | 256 | 2.406x |
| ja-wikipedia | 4096 | 2.273x |
| ja-wikipedia | 65536 | 2.203x |
| ja-wikipedia | 1048576 | 2.199x |
| ru-wikipedia | 256 | 3.182x |
| ru-wikipedia | 4096 | 3.135x |
| ru-wikipedia | 65536 | 3.058x |
| ru-wikipedia | 1048576 | 3.022x |
| synthetic | 0 | 5.262x |
| zh-wikipedia | 256 | 3.359x |
| zh-wikipedia | 4096 | 3.293x |
| zh-wikipedia | 65536 | 3.270x |
| zh-wikipedia | 1048576 | 3.198x |
| Combined selected natural corpora | 256 | 3.648x |
| Combined selected natural corpora | 4096 | 3.640x |
| Combined selected natural corpora | 65536 | 3.516x |
| Combined selected natural corpora | 1048576 | 3.485x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 6.316x |
| ar-wikipedia | 4096 | 5.620x |
| ar-wikipedia | 65536 | 5.498x |
| ar-wikipedia | 1048576 | 6.268x |
| en-wikitext2 | 256 | 11.170x |
| en-wikitext2 | 4096 | 10.753x |
| en-wikitext2 | 65536 | 10.610x |
| en-wikitext2 | 1048576 | 11.638x |
| ja-wikipedia | 256 | 4.201x |
| ja-wikipedia | 4096 | 3.669x |
| ja-wikipedia | 65536 | 3.619x |
| ja-wikipedia | 1048576 | 4.124x |
| ru-wikipedia | 256 | 5.628x |
| ru-wikipedia | 4096 | 5.141x |
| ru-wikipedia | 65536 | 5.087x |
| ru-wikipedia | 1048576 | 5.779x |
| synthetic | 0 | 8.959x |
| zh-wikipedia | 256 | 6.098x |
| zh-wikipedia | 4096 | 5.418x |
| zh-wikipedia | 65536 | 5.361x |
| zh-wikipedia | 1048576 | 6.068x |
| Combined selected natural corpora | 256 | 6.557x |
| Combined selected natural corpora | 4096 | 6.004x |
| Combined selected natural corpora | 65536 | 5.933x |
| Combined selected natural corpora | 1048576 | 6.673x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 5.570x |
| ar-wikipedia | 4096 | 5.312x |
| ar-wikipedia | 65536 | 5.219x |
| ar-wikipedia | 1048576 | 5.871x |
| en-wikitext2 | 256 | 9.316x |
| en-wikitext2 | 4096 | 9.580x |
| en-wikitext2 | 65536 | 9.459x |
| en-wikitext2 | 1048576 | 10.262x |
| ja-wikipedia | 256 | 3.817x |
| ja-wikipedia | 4096 | 3.516x |
| ja-wikipedia | 65536 | 3.504x |
| ja-wikipedia | 1048576 | 3.990x |
| ru-wikipedia | 256 | 5.032x |
| ru-wikipedia | 4096 | 4.872x |
| ru-wikipedia | 65536 | 4.883x |
| ru-wikipedia | 1048576 | 5.463x |
| synthetic | 0 | 8.156x |
| zh-wikipedia | 256 | 5.328x |
| zh-wikipedia | 4096 | 5.108x |
| zh-wikipedia | 65536 | 5.125x |
| zh-wikipedia | 1048576 | 5.770x |
| Combined selected natural corpora | 256 | 5.769x |
| Combined selected natural corpora | 4096 | 5.630x |
| Combined selected natural corpora | 65536 | 5.608x |
| Combined selected natural corpora | 1048576 | 6.247x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 95467.8 | 218.3 | 10 |
| onig-compile | 95617.5 | 555.0 | 10 |
| pretok-compile | 11.3 | 0.2 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
