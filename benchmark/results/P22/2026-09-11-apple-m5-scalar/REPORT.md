# P22 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 33270570 | 559846 | 8106.9 | 31.52 | 17570333 | 3518214 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 18917302 | 151859 | 4609.5 | 55.43 | 30901605 | 6187616 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 18082163 | 233593 | 4406.0 | 57.99 | 32328820 | 6303505 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 18076623 | 234549 | 4404.6 | 58.01 | 32338729 | 6475380 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 5670216 | 16663 | 1381.6 | 184.93 | 103095719 | 20101703 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 6653783 | 34063 | 1621.3 | 157.59 | 87856038 | 17591948 | 4765600.0 | 25233.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 5734315 | 17131 | 1397.3 | 182.86 | 101943295 | 20412725 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 29137179 | 375258 | 113374.2 | 35.99 | 20062855 | 3901716 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 17311803 | 199711 | 67361.1 | 60.57 | 33767424 | 6566907 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 17092924 | 273449 | 66509.4 | 61.35 | 34199824 | 6472093 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 17009675 | 150685 | 66185.5 | 61.65 | 34367206 | 6683549 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 5613192 | 34987 | 21841.2 | 186.81 | 104143055 | 19708393 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 6181052 | 43693 | 24050.8 | 169.64 | 94575322 | 18392500 | 4380656.0 | 2579.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 5659267 | 40512 | 22020.5 | 185.28 | 103295182 | 20088291 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 28491617 | 341427 | 1675977.5 | 36.80 | 20517439 | 3982540 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 17013892 | 178505 | 1000817.2 | 61.63 | 34358687 | 6669197 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 16816944 | 190453 | 989232.0 | 62.35 | 34761072 | 6565521 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 16879408 | 207761 | 992906.4 | 62.12 | 34632434 | 6722333 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 5616027 | 26064 | 330354.5 | 186.71 | 104090489 | 19660162 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 6235766 | 48614 | 366809.8 | 168.16 | 93745497 | 18196481 | 8081200.0 | 290.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 5672611 | 21852 | 333683.0 | 184.85 | 103052185 | 20002957 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 31399121 | 651391 | 31399120.6 | 33.40 | 18617560 | 3613254 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 17069438 | 97576 | 17069437.5 | 61.43 | 34246881 | 6646558 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 17045715 | 352633 | 17045714.6 | 61.52 | 34294543 | 6476467 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 17016532 | 391415 | 17016531.9 | 61.62 | 34353357 | 6667222 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 5643945 | 44184 | 5643945.0 | 185.79 | 103575602 | 19560077 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 6216993 | 36052 | 6216993.3 | 168.66 | 94028571 | 18248853 | 8942832.0 | 29.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 5699982 | 21799 | 5699981.5 | 183.96 | 102557350 | 19904100 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 47672038 | 725969 | 11635.8 | 22.00 | 21950750 | 4377304 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 26605925 | 228797 | 6494.0 | 39.41 | 39330975 | 7843178 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 25474698 | 119404 | 6217.9 | 41.16 | 41077504 | 7959623 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 25570151 | 167917 | 6241.2 | 41.01 | 40924162 | 8160883 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 3793620 | 4391 | 926.0 | 276.41 | 275841307 | 53450008 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 5059546 | 41446 | 1234.9 | 207.25 | 206824293 | 41243820 | 8542320.0 | 28782.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 3862732 | 25024 | 942.8 | 271.46 | 270905936 | 54022647 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 44188214 | 894715 | 172610.2 | 23.73 | 23681360 | 4652824 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 24992895 | 336448 | 97628.5 | 41.95 | 41869379 | 8226338 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 24409073 | 272508 | 95347.9 | 42.96 | 42870820 | 8182531 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 24430168 | 198408 | 95430.3 | 42.92 | 42833803 | 8415824 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 3720495 | 15837 | 14533.2 | 281.84 | 281262850 | 53683181 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 4545190 | 30222 | 17754.7 | 230.70 | 230229517 | 45234628 | 8585216.0 | 2861.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 3774952 | 14140 | 14745.9 | 277.77 | 277205351 | 54464263 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 43540104 | 955903 | 2721256.5 | 24.08 | 24033865 | 4717765 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 24254024 | 466002 | 1515876.5 | 43.23 | 43144882 | 8469193 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 24044014 | 219519 | 1502750.9 | 43.61 | 43521727 | 8298947 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 24150064 | 558938 | 1509379.0 | 43.42 | 43330611 | 8505651 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 3709847 | 12657 | 231865.4 | 282.65 | 282070130 | 53786586 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 4692294 | 21436 | 293268.4 | 223.47 | 223011796 | 43776452 | 14634752.0 | 320.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 3777911 | 14807 | 236119.4 | 277.55 | 276988294 | 54371854 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 47025049 | 1315776 | 47025049.3 | 22.30 | 22252757 | 4367842 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 24154976 | 436884 | 24154976.2 | 43.41 | 43321798 | 8503341 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 24070105 | 149923 | 24070105.2 | 43.56 | 43474550 | 8289370 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 23996976 | 411221 | 23996976.2 | 43.70 | 43607036 | 8559328 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 3721932 | 20471 | 3721931.6 | 281.73 | 281154277 | 53608185 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 4633066 | 13817 | 4633066.4 | 226.32 | 225862726 | 44333058 | 17962224.0 | 32.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 3780234 | 26921 | 3780233.9 | 277.38 | 276818058 | 54334734 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 24565297 | 596277 | 5972.6 | 42.69 | 15439667 | 2881626 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 14540787 | 168857 | 3535.3 | 72.11 | 26083870 | 4868237 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 13949405 | 182348 | 3391.5 | 75.17 | 27189690 | 4815259 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 13845140 | 240434 | 3366.2 | 75.74 | 27394451 | 5112841 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 5670286 | 5800 | 1378.6 | 184.92 | 66889047 | 11845964 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 6421470 | 9014 | 1561.3 | 163.29 | 59064359 | 11023644 | 3144048.0 | 22259.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 5693294 | 6837 | 1384.2 | 184.18 | 66618732 | 12433576 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 20968148 | 122130 | 81588.1 | 50.01 | 18088389 | 3222745 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 12917837 | 116756 | 50264.0 | 81.17 | 29360954 | 5231139 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 12834118 | 337937 | 49938.2 | 81.70 | 29552480 | 4985929 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 12742608 | 106385 | 49582.1 | 82.29 | 29764707 | 5303074 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 5575188 | 12206 | 21693.3 | 188.08 | 68029993 | 11477640 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 5990239 | 10777 | 23308.3 | 175.05 | 63316338 | 11280852 | 2990064.0 | 2415.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 5589431 | 12459 | 21748.8 | 187.60 | 67856644 | 12089783 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 20475705 | 231489 | 1204453.2 | 51.21 | 18523416 | 3290143 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 12639759 | 103235 | 743515.2 | 82.96 | 30006902 | 5329849 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 12623555 | 120346 | 742562.1 | 83.06 | 30045419 | 5052776 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 12522904 | 203029 | 736641.4 | 83.73 | 30286905 | 5379583 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 5564160 | 11289 | 327303.6 | 188.45 | 68164821 | 11463365 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 5916332 | 27344 | 348019.5 | 177.23 | 64107292 | 11386786 | 4067088.0 | 256.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 5575684 | 10779 | 327981.4 | 188.06 | 68023938 | 12082463 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 23363370 | 296933 | 23363370.3 | 44.88 | 16233959 | 2882889 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 12543019 | 110133 | 12543019.3 | 83.60 | 30238334 | 5369839 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 12543344 | 106154 | 12543343.7 | 83.60 | 30237551 | 5083971 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 12533999 | 128838 | 12533999.5 | 83.66 | 30260094 | 5373704 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 5570336 | 11090 | 5570336.1 | 188.24 | 68089248 | 11448142 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 5947521 | 5424 | 5947521.3 | 176.30 | 63771104 | 11324718 | 5559536.0 | 27.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 5586868 | 14481 | 5586867.6 | 187.69 | 67887773 | 12055772 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 30759732 | 498912 | 7495.1 | 34.09 | 18688622 | 3258741 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 17338723 | 505972 | 4224.8 | 60.48 | 33154517 | 5781164 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 16584410 | 510228 | 4041.0 | 63.23 | 34662493 | 5833310 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 16621079 | 233670 | 4050.0 | 63.09 | 34586023 | 6030776 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 5662801 | 75486 | 1379.8 | 185.17 | 101514610 | 17083773 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 6538721 | 66835 | 1593.3 | 160.36 | 87915813 | 15329909 | 4403584.0 | 24849.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 5695372 | 68425 | 1387.8 | 184.11 | 100934061 | 17599905 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 26829127 | 287678 | 104393.5 | 39.08 | 21426601 | 3606752 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 15692192 | 195815 | 61059.1 | 66.82 | 36633313 | 6166506 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 15403769 | 352579 | 59936.8 | 68.07 | 37319243 | 6056829 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 15355790 | 218096 | 59750.2 | 68.29 | 37435847 | 6301597 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 5585479 | 57081 | 21733.4 | 187.73 | 102919903 | 16703669 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 6136362 | 63880 | 23876.9 | 170.88 | 93680421 | 15769278 | 4244976.0 | 2569.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 5636530 | 42294 | 21932.0 | 186.03 | 101987744 | 17167654 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 25990871 | 462817 | 1528874.8 | 40.34 | 22117651 | 3714766 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 15318352 | 328213 | 901079.5 | 68.45 | 37527341 | 6302897 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 15207468 | 272757 | 894556.9 | 68.95 | 37800968 | 6120940 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 15215412 | 229557 | 895024.2 | 68.92 | 37781232 | 6345540 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 5585319 | 57282 | 328548.2 | 187.74 | 102922859 | 16665834 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 6027853 | 62245 | 354579.6 | 173.95 | 95366794 | 16017312 | 5852944.0 | 273.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 5634844 | 67963 | 331461.4 | 186.09 | 102018263 | 17134458 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 28766914 | 685449 | 28766914.4 | 36.45 | 19983269 | 3355800 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 15289685 | 205908 | 15289685.0 | 68.58 | 37597701 | 6313799 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 15338417 | 259801 | 15338416.7 | 68.36 | 37478249 | 6067771 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 15419122 | 261203 | 15419122.5 | 68.00 | 37282083 | 6260797 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 5582062 | 68946 | 5582062.0 | 187.85 | 102982911 | 16673050 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 6094982 | 26441 | 6094981.7 | 172.04 | 94316444 | 15838604 | 8942832.0 | 29.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 5622730 | 75699 | 5622729.8 | 186.49 | 102238062 | 17168885 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 143781453 | 2716505 | 1203.1 | 20.31 | 10525551 | 2558411 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 86114083 | 1534738 | 720.6 | 33.91 | 17574117 | 4271682 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 82405434 | 617533 | 689.5 | 35.43 | 18365039 | 3895411 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 83360168 | 1005064 | 697.5 | 35.03 | 18154702 | 4412803 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 23181053 | 92301 | 194.0 | 125.97 | 65285170 | 13847645 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 27574041 | 141522 | 230.7 | 105.90 | 54884194 | 13340519 | 13049504.0 | 262067.0 | 10 |
| synthetic | 0 | pretok-reused | 23559156 | 89550 | 197.1 | 123.94 | 64237402 | 15613972 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 26642718 | 230952 | 6477.7 | 39.36 | 14645765 | 3080954 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 15796973 | 167671 | 3840.7 | 66.38 | 24701124 | 5196249 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 15262948 | 226446 | 3710.9 | 68.70 | 25565376 | 5104715 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 15247523 | 201503 | 3707.2 | 68.77 | 25591239 | 5383497 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 4279693 | 4491 | 1040.5 | 245.01 | 91175461 | 18205277 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 5103208 | 12041 | 1240.8 | 205.47 | 76462292 | 16084980 | 3652592.0 | 23334.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 4314311 | 10594 | 1048.9 | 243.05 | 90443876 | 19026214 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 23375951 | 309272 | 90957.0 | 44.86 | 16692498 | 3381253 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 14396601 | 116981 | 56017.9 | 72.83 | 27103829 | 5490185 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 14218728 | 159527 | 55325.8 | 73.75 | 27442891 | 5267630 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 14255001 | 77382 | 55466.9 | 73.56 | 27373060 | 5544721 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 4187854 | 10562 | 16295.2 | 250.38 | 93174933 | 17884817 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 4678988 | 24542 | 18206.2 | 224.10 | 83394747 | 16892543 | 3727344.0 | 2503.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 4216884 | 15334 | 16408.1 | 248.66 | 92533489 | 18743697 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 23005448 | 309199 | 1353261.6 | 45.58 | 16961330 | 3427232 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 14143668 | 230557 | 831980.5 | 74.14 | 27588530 | 5574579 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 14104263 | 189824 | 829662.5 | 74.34 | 27665607 | 5296484 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 14077106 | 313140 | 828065.1 | 74.49 | 27718979 | 5600938 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 4183348 | 15229 | 246079.3 | 250.65 | 93275281 | 17857226 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 4573180 | 11927 | 269010.6 | 229.29 | 85324218 | 17240739 | 5001072.0 | 267.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 4210393 | 9804 | 247670.2 | 249.04 | 92676142 | 18726280 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 26009405 | 616456 | 26009405.2 | 40.32 | 15002381 | 3030981 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 14115312 | 191162 | 14115312.5 | 74.29 | 27643950 | 5584999 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 14096840 | 146570 | 14096839.7 | 74.38 | 27680176 | 5298563 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 14115227 | 139984 | 14115226.8 | 74.29 | 27644118 | 5585032 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 4178823 | 10559 | 4178822.6 | 250.93 | 93376302 | 17874174 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 4640014 | 14226 | 4640014.5 | 225.99 | 84095212 | 16990033 | 7058672.0 | 28.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 4214803 | 13430 | 4214802.7 | 248.78 | 92579185 | 18704078 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.189x |
| ar-wikipedia | 4096 | 3.045x |
| ar-wikipedia | 65536 | 2.994x |
| ar-wikipedia | 1048576 | 3.020x |
| en-wikitext2 | 256 | 6.715x |
| en-wikitext2 | 4096 | 6.561x |
| en-wikitext2 | 65536 | 6.481x |
| en-wikitext2 | 1048576 | 6.467x |
| ja-wikipedia | 256 | 2.460x |
| ja-wikipedia | 4096 | 2.302x |
| ja-wikipedia | 65536 | 2.269x |
| ja-wikipedia | 1048576 | 2.252x |
| ru-wikipedia | 256 | 2.929x |
| ru-wikipedia | 4096 | 2.758x |
| ru-wikipedia | 65536 | 2.723x |
| ru-wikipedia | 1048576 | 2.748x |
| synthetic | 0 | 3.555x |
| zh-wikipedia | 256 | 3.566x |
| zh-wikipedia | 4096 | 3.395x |
| zh-wikipedia | 65536 | 3.372x |
| zh-wikipedia | 1048576 | 3.373x |
| Combined selected natural corpora | 256 | 3.563x |
| Combined selected natural corpora | 4096 | 3.402x |
| Combined selected natural corpora | 65536 | 3.358x |
| Combined selected natural corpora | 1048576 | 3.365x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.152x |
| ar-wikipedia | 4096 | 3.006x |
| ar-wikipedia | 65536 | 2.976x |
| ar-wikipedia | 1048576 | 2.985x |
| en-wikitext2 | 256 | 6.620x |
| en-wikitext2 | 4096 | 6.472x |
| en-wikitext2 | 65536 | 6.392x |
| en-wikitext2 | 1048576 | 6.348x |
| ja-wikipedia | 256 | 2.432x |
| ja-wikipedia | 4096 | 2.280x |
| ja-wikipedia | 65536 | 2.246x |
| ja-wikipedia | 1048576 | 2.243x |
| ru-wikipedia | 256 | 2.918x |
| ru-wikipedia | 4096 | 2.724x |
| ru-wikipedia | 65536 | 2.700x |
| ru-wikipedia | 1048576 | 2.742x |
| synthetic | 0 | 3.538x |
| zh-wikipedia | 256 | 3.534x |
| zh-wikipedia | 4096 | 3.380x |
| zh-wikipedia | 65536 | 3.343x |
| zh-wikipedia | 1048576 | 3.349x |
| Combined selected natural corpora | 256 | 3.532x |
| Combined selected natural corpora | 4096 | 3.368x |
| Combined selected natural corpora | 65536 | 3.331x |
| Combined selected natural corpora | 1048576 | 3.336x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.843x |
| ar-wikipedia | 4096 | 2.801x |
| ar-wikipedia | 65536 | 2.728x |
| ar-wikipedia | 1048576 | 2.746x |
| en-wikitext2 | 256 | 5.259x |
| en-wikitext2 | 4096 | 5.499x |
| en-wikitext2 | 65536 | 5.169x |
| en-wikitext2 | 1048576 | 5.214x |
| ja-wikipedia | 256 | 2.264x |
| ja-wikipedia | 4096 | 2.156x |
| ja-wikipedia | 65536 | 2.136x |
| ja-wikipedia | 1048576 | 2.109x |
| ru-wikipedia | 256 | 2.652x |
| ru-wikipedia | 4096 | 2.557x |
| ru-wikipedia | 65536 | 2.541x |
| ru-wikipedia | 1048576 | 2.509x |
| synthetic | 0 | 3.123x |
| zh-wikipedia | 256 | 3.095x |
| zh-wikipedia | 4096 | 3.077x |
| zh-wikipedia | 65536 | 3.093x |
| zh-wikipedia | 1048576 | 3.042x |
| Combined selected natural corpora | 256 | 3.130x |
| Combined selected natural corpora | 4096 | 3.099x |
| Combined selected natural corpora | 65536 | 3.038x |
| Combined selected natural corpora | 1048576 | 3.021x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 5.802x |
| ar-wikipedia | 4096 | 5.149x |
| ar-wikipedia | 65536 | 5.023x |
| ar-wikipedia | 1048576 | 5.509x |
| en-wikitext2 | 256 | 12.342x |
| en-wikitext2 | 4096 | 11.706x |
| en-wikitext2 | 65536 | 11.525x |
| en-wikitext2 | 1048576 | 12.440x |
| ja-wikipedia | 256 | 4.315x |
| ja-wikipedia | 4096 | 3.751x |
| ja-wikipedia | 65536 | 3.672x |
| ja-wikipedia | 1048576 | 4.182x |
| ru-wikipedia | 256 | 5.401x |
| ru-wikipedia | 4096 | 4.760x |
| ru-wikipedia | 65536 | 4.613x |
| ru-wikipedia | 1048576 | 5.116x |
| synthetic | 0 | 6.103x |
| zh-wikipedia | 256 | 6.175x |
| zh-wikipedia | 4096 | 5.543x |
| zh-wikipedia | 65536 | 5.464x |
| zh-wikipedia | 1048576 | 6.171x |
| Combined selected natural corpora | 256 | 6.439x |
| Combined selected natural corpora | 4096 | 5.809x |
| Combined selected natural corpora | 65536 | 5.689x |
| Combined selected natural corpora | 1048576 | 6.287x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 5.000x |
| ar-wikipedia | 4096 | 4.714x |
| ar-wikipedia | 65536 | 4.569x |
| ar-wikipedia | 1048576 | 5.051x |
| en-wikitext2 | 256 | 9.422x |
| en-wikitext2 | 4096 | 9.722x |
| en-wikitext2 | 65536 | 9.279x |
| en-wikitext2 | 1048576 | 10.150x |
| ja-wikipedia | 256 | 3.825x |
| ja-wikipedia | 4096 | 3.500x |
| ja-wikipedia | 65536 | 3.461x |
| ja-wikipedia | 1048576 | 3.928x |
| ru-wikipedia | 256 | 4.704x |
| ru-wikipedia | 4096 | 4.372x |
| ru-wikipedia | 65536 | 4.312x |
| ru-wikipedia | 1048576 | 4.720x |
| synthetic | 0 | 5.214x |
| zh-wikipedia | 256 | 5.221x |
| zh-wikipedia | 4096 | 4.996x |
| zh-wikipedia | 65536 | 5.031x |
| zh-wikipedia | 1048576 | 5.605x |
| Combined selected natural corpora | 256 | 5.471x |
| Combined selected natural corpora | 4096 | 5.248x |
| Combined selected natural corpora | 65536 | 5.156x |
| Combined selected natural corpora | 1048576 | 5.686x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 74534.9 | 119.0 | 10 |
| onig-compile | 74630.7 | 477.8 | 10 |
| pretok-compile | 11.3 | 0.1 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
