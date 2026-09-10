# P16 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 8123320 | 15525 | 1979.4 | 129.08 | 71962573 | 1913257 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 3595370 | 37231 | 876.1 | 291.65 | 162591074 | 4322782 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 3378193 | 8859 | 823.1 | 310.40 | 173043678 | 2496009 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 3408520 | 31622 | 830.5 | 307.63 | 171504048 | 4559750 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 1509972 | 1678 | 367.9 | 694.43 | 387143033 | 5584211 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 1656151 | 3473 | 403.5 | 633.14 | 352972037 | 9384410 | 590848.0 | 8645.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 1515598 | 6327 | 369.3 | 691.86 | 385705942 | 10254701 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 6404590 | 13916 | 24920.6 | 163.72 | 91274388 | 1833373 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 3137402 | 16046 | 12207.8 | 334.22 | 186324532 | 3742587 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 3054218 | 15541 | 11884.1 | 343.32 | 191399256 | 2760772 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 3070549 | 19312 | 11947.7 | 341.49 | 190381259 | 3824072 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 1445447 | 7222 | 5624.3 | 725.43 | 404424973 | 5833488 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 1506933 | 4285 | 5863.6 | 695.83 | 387923656 | 7791985 | 528944.0 | 1571.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 1446192 | 5556 | 5627.2 | 725.06 | 404216741 | 8119254 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 6147042 | 9935 | 361590.7 | 170.58 | 95098583 | 1871469 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 3065321 | 9941 | 180313.0 | 342.08 | 190705956 | 3752951 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 3022147 | 25135 | 177773.3 | 346.96 | 193430392 | 2790070 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 3046464 | 6478 | 179203.8 | 344.19 | 191886404 | 3776181 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 1435350 | 8882 | 84432.4 | 730.54 | 407269872 | 5874523 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 1482377 | 2058 | 87198.7 | 707.36 | 394349725 | 7760508 | 552976.0 | 179.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 1435323 | 5757 | 84430.8 | 730.55 | 407277646 | 8014920 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 7312710 | 41730 | 7312709.5 | 143.39 | 79939590 | 1570964 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 3062243 | 31711 | 3062242.9 | 342.42 | 190897660 | 3751499 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 3035070 | 24264 | 3035070.2 | 345.49 | 192606748 | 2778189 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 3051913 | 14318 | 3051913.4 | 343.58 | 191543773 | 3764196 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 1416518 | 6071 | 1416518.3 | 740.25 | 412684410 | 5952624 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 1474836 | 4055 | 1474835.7 | 710.98 | 396366195 | 7789342 | 914672.0 | 20.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 1419973 | 8477 | 1419972.7 | 738.45 | 411680457 | 8090297 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 8437110 | 64806 | 2059.3 | 124.28 | 124027895 | 2991901 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 3332424 | 8950 | 813.4 | 314.66 | 314016805 | 7574967 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 3013717 | 30723 | 735.6 | 347.93 | 347224664 | 5058537 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 3071079 | 18316 | 749.6 | 341.44 | 340739243 | 8219588 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 534461 | 2495 | 130.5 | 1961.93 | 1957931394 | 28524091 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 767595 | 6751 | 187.4 | 1366.05 | 1363267946 | 32885852 | 1021104.0 | 11879.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 545363 | 1251 | 133.1 | 1922.71 | 1918791453 | 46286640 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 6633393 | 12230 | 25911.7 | 158.08 | 157752900 | 3236956 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 2852204 | 13804 | 11141.4 | 367.64 | 366887126 | 7528213 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 2745877 | 17688 | 10726.1 | 381.87 | 381093912 | 5551960 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 2792073 | 24900 | 10906.5 | 375.55 | 374788520 | 7690343 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 476175 | 3002 | 1860.1 | 2202.08 | 2197587928 | 32015523 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 573417 | 1193 | 2239.9 | 1828.65 | 1824916028 | 37445730 | 964032.0 | 1875.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 479725 | 1379 | 1873.9 | 2185.78 | 2181325098 | 44758942 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 6638877 | 37225 | 414929.8 | 157.94 | 157622600 | 3199035 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 2840503 | 9129 | 177531.5 | 369.15 | 368398431 | 7476844 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 2767087 | 6432 | 172942.9 | 378.95 | 378172784 | 5509404 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 2820549 | 11561 | 176284.3 | 371.76 | 371004718 | 7529740 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 476620 | 1879 | 29788.8 | 2200.02 | 2195535207 | 31985618 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 556415 | 1950 | 34775.9 | 1884.52 | 1880676896 | 38169346 | 1093376.0 | 201.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 480352 | 1506 | 30022.0 | 2182.93 | 2178481588 | 44213452 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 8124167 | 176363 | 8124167.0 | 129.07 | 128805452 | 2612329 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 2843206 | 47492 | 2843206.3 | 368.80 | 368048212 | 7464460 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 2805391 | 37985 | 2805391.4 | 373.77 | 373009269 | 5434179 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 2888694 | 117309 | 2888694.4 | 362.99 | 362252585 | 7346918 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 475214 | 2760 | 475213.9 | 2206.53 | 2202033672 | 32080291 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 566696 | 2740 | 566696.4 | 1850.33 | 1846556502 | 37450385 | 1586416.0 | 22.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 481478 | 1796 | 481478.3 | 2177.83 | 2173383439 | 44078828 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 7720851 | 57312 | 1877.2 | 135.81 | 49124119 | 2681052 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 3203695 | 25653 | 778.9 | 327.30 | 118388281 | 6461288 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 2912810 | 18979 | 708.2 | 359.99 | 130211024 | 4029099 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 2996392 | 22598 | 728.5 | 349.95 | 126578915 | 6908309 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 936297 | 4598 | 227.6 | 1119.92 | 405085157 | 12534485 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 1143150 | 5414 | 277.9 | 917.27 | 331785029 | 18107863 | 815376.0 | 10314.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 944623 | 3976 | 229.7 | 1110.05 | 401514563 | 21913498 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 6264229 | 33807 | 24374.4 | 167.39 | 60546955 | 2706798 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 2750856 | 14639 | 10703.7 | 381.18 | 137877106 | 6163901 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 2654637 | 19341 | 10329.3 | 395.00 | 142874501 | 4420943 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 2692397 | 23971 | 10476.3 | 389.46 | 140870735 | 6297733 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 878129 | 2841 | 3416.8 | 1194.10 | 431918487 | 13364784 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 975115 | 3767 | 3794.2 | 1075.33 | 388959199 | 17388716 | 763664.0 | 1664.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 882662 | 7612 | 3434.5 | 1187.97 | 429700184 | 19210073 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 6189529 | 43871 | 364089.9 | 169.41 | 61277688 | 2701337 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 2698407 | 13992 | 158729.8 | 388.59 | 140556989 | 6196248 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 2633421 | 11532 | 154907.1 | 398.18 | 144025606 | 4456561 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 2680411 | 14742 | 157671.3 | 391.20 | 141500665 | 6237848 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 872334 | 4691 | 51313.8 | 1202.03 | 434787546 | 13453561 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 948375 | 3342 | 55786.8 | 1105.65 | 399926163 | 17630156 | 861712.0 | 189.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 878319 | 3358 | 51665.8 | 1193.84 | 431824903 | 19036365 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 7497401 | 185645 | 7497401.0 | 139.86 | 50588197 | 2227972 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 2677578 | 18873 | 2677578.2 | 391.61 | 141650390 | 6238473 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 2621675 | 33793 | 2621675.0 | 399.96 | 144670871 | 4476527 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 2670845 | 18629 | 2670845.2 | 392.60 | 142007483 | 6254200 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 865820 | 1584 | 865820.2 | 1211.08 | 438058630 | 13554778 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 943373 | 15018 | 943373.3 | 1111.52 | 402046550 | 17706669 | 1209584.0 | 21.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 872158 | 3410 | 872158.0 | 1202.28 | 434875322 | 19152493 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 8426411 | 41286 | 2053.2 | 124.44 | 68220856 | 2087128 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 3714331 | 29609 | 905.1 | 282.31 | 154767324 | 4734904 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 3483458 | 11600 | 848.8 | 301.02 | 165024830 | 2899418 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 3532839 | 30030 | 860.8 | 296.81 | 162718120 | 4978149 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 1446979 | 6090 | 352.6 | 724.66 | 397280739 | 6980058 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 1614021 | 8329 | 393.3 | 649.67 | 356164576 | 10896391 | 685184.0 | 9541.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 1454883 | 3434 | 354.5 | 720.73 | 395122611 | 12088261 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 6766330 | 33942 | 26328.1 | 154.97 | 84958460 | 2039067 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 3305696 | 21018 | 12862.6 | 317.20 | 173898916 | 4173705 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 3210768 | 18457 | 12493.3 | 326.58 | 179040325 | 3145665 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 3247000 | 12824 | 12634.2 | 322.94 | 177042509 | 4249153 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 1391865 | 3817 | 5415.8 | 753.36 | 413011904 | 7256449 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 1471427 | 10454 | 5725.4 | 712.62 | 390680024 | 9376614 | 613648.0 | 1654.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 1390879 | 1761 | 5412.0 | 753.89 | 413304877 | 9919628 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 6564235 | 69089 | 386131.4 | 159.74 | 87574110 | 2066349 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 3226051 | 58983 | 189767.7 | 325.03 | 178192142 | 4204521 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 3183923 | 53952 | 187289.6 | 329.33 | 180549929 | 3172188 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 3208880 | 16755 | 188757.7 | 326.77 | 179145671 | 4227020 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 1384502 | 4352 | 81441.3 | 757.37 | 415208479 | 7295041 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 1436661 | 10363 | 84509.4 | 729.87 | 400134190 | 9441340 | 658192.0 | 183.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 1385250 | 5230 | 81485.3 | 756.96 | 414984288 | 9791734 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 8112939 | 284856 | 8112938.8 | 129.25 | 70856814 | 1669925 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 3229228 | 74161 | 3229228.2 | 324.71 | 178016841 | 4195430 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 3176298 | 25422 | 3176297.6 | 330.12 | 180983356 | 3179803 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 3209786 | 60303 | 3209786.2 | 326.68 | 179095107 | 4220842 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 1367642 | 7408 | 1367641.9 | 766.70 | 420327132 | 7384974 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 1427080 | 15310 | 1427080.2 | 734.77 | 402820400 | 9493510 | 914672.0 | 20.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 1365013 | 6635 | 1365012.9 | 768.18 | 421136684 | 9925181 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 39475673 | 266484 | 884.6 | 41.17 | 21474745 | 4163450 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 14564748 | 193560 | 326.4 | 111.59 | 58204236 | 11284439 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 13143013 | 127701 | 294.5 | 123.66 | 64500432 | 6312556 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 13487217 | 141646 | 302.2 | 120.50 | 62854330 | 12185983 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 2216968 | 9640 | 49.7 | 733.09 | 382382683 | 37423191 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 3751006 | 47130 | 84.1 | 433.28 | 226000729 | 43816250 | 7486832.0 | 74391.0 | 10 |
| synthetic | 0 | pretok-reused | 2320412 | 7599 | 52.0 | 700.41 | 365336038 | 70830105 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 7728401 | 54730 | 1879.0 | 135.68 | 50489485 | 2772889 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 3297867 | 35262 | 801.8 | 317.96 | 118319819 | 6498140 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 3032928 | 28885 | 737.4 | 345.73 | 128655529 | 4136267 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 3087217 | 7133 | 750.6 | 339.65 | 126393127 | 6941527 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 921341 | 2233 | 224.0 | 1138.10 | 423516565 | 13616029 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 1110298 | 4361 | 269.9 | 944.41 | 351439876 | 19301124 | 845520.0 | 9905.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 929368 | 6009 | 226.0 | 1128.27 | 419858570 | 23058688 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 6332388 | 30072 | 24639.6 | 165.59 | 61620193 | 2792943 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 2848433 | 21697 | 11083.4 | 368.12 | 136988647 | 6209028 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 2737466 | 20963 | 10651.6 | 383.05 | 142541663 | 4582705 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 2777471 | 16580 | 10807.3 | 377.53 | 140488584 | 6367663 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 858667 | 2520 | 3341.1 | 1221.17 | 454428935 | 14609859 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 948193 | 4580 | 3689.5 | 1105.87 | 411522766 | 18652321 | 802384.0 | 1627.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 860689 | 1772 | 3349.0 | 1218.30 | 453361105 | 20548649 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 6246402 | 65604 | 367435.4 | 167.87 | 62468441 | 2792968 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 2790906 | 12949 | 164170.9 | 375.71 | 139812321 | 6251017 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 2719451 | 8965 | 159967.7 | 385.58 | 143485967 | 4613064 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 2762757 | 22487 | 162515.1 | 379.54 | 141236804 | 6314706 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 852118 | 1713 | 50124.6 | 1230.55 | 457921354 | 14722140 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 926650 | 4440 | 54508.8 | 1131.57 | 421089755 | 18826949 | 891664.0 | 190.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 856583 | 3812 | 50387.2 | 1224.14 | 455534481 | 20366974 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 7373663 | 42405 | 7373663.0 | 142.21 | 52918475 | 2363954 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 2772445 | 12224 | 2772445.4 | 378.21 | 140743257 | 6287229 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 2731457 | 12067 | 2731456.5 | 383.89 | 142855285 | 4592788 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 2769433 | 14007 | 2769433.5 | 378.62 | 140896326 | 6294067 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 849565 | 1975 | 849565.2 | 1234.25 | 459297323 | 14766378 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 924421 | 2664 | 924421.0 | 1134.30 | 422105313 | 18856128 | 1209584.0 | 21.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 854669 | 1966 | 854669.0 | 1226.88 | 456554538 | 20395031 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.237x |
| ar-wikipedia | 4096 | 2.113x |
| ar-wikipedia | 65536 | 2.106x |
| ar-wikipedia | 1048576 | 2.143x |
| en-wikitext2 | 256 | 5.639x |
| en-wikitext2 | 4096 | 5.767x |
| en-wikitext2 | 65536 | 5.806x |
| en-wikitext2 | 1048576 | 5.903x |
| ja-wikipedia | 256 | 3.111x |
| ja-wikipedia | 4096 | 3.023x |
| ja-wikipedia | 65536 | 3.019x |
| ja-wikipedia | 1048576 | 3.028x |
| ru-wikipedia | 256 | 2.407x |
| ru-wikipedia | 4096 | 2.307x |
| ru-wikipedia | 65536 | 2.300x |
| ru-wikipedia | 1048576 | 2.322x |
| synthetic | 0 | 5.928x |
| zh-wikipedia | 256 | 3.292x |
| zh-wikipedia | 4096 | 3.188x |
| zh-wikipedia | 65536 | 3.191x |
| zh-wikipedia | 1048576 | 3.215x |
| Combined selected natural corpora | 256 | 2.958x |
| Combined selected natural corpora | 4096 | 2.852x |
| Combined selected natural corpora | 65536 | 2.853x |
| Combined selected natural corpora | 1048576 | 2.889x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.249x |
| ar-wikipedia | 4096 | 2.123x |
| ar-wikipedia | 65536 | 2.122x |
| ar-wikipedia | 1048576 | 2.149x |
| en-wikitext2 | 256 | 5.631x |
| en-wikitext2 | 4096 | 5.820x |
| en-wikitext2 | 65536 | 5.872x |
| en-wikitext2 | 1048576 | 6.000x |
| ja-wikipedia | 256 | 3.172x |
| ja-wikipedia | 4096 | 3.050x |
| ja-wikipedia | 65536 | 3.052x |
| ja-wikipedia | 1048576 | 3.062x |
| ru-wikipedia | 256 | 2.428x |
| ru-wikipedia | 4096 | 2.334x |
| ru-wikipedia | 65536 | 2.316x |
| ru-wikipedia | 1048576 | 2.351x |
| synthetic | 0 | 5.812x |
| zh-wikipedia | 256 | 3.322x |
| zh-wikipedia | 4096 | 3.227x |
| zh-wikipedia | 65536 | 3.225x |
| zh-wikipedia | 1048576 | 3.240x |
| Combined selected natural corpora | 256 | 2.986x |
| Combined selected natural corpora | 4096 | 2.881x |
| Combined selected natural corpora | 65536 | 2.883x |
| Combined selected natural corpora | 1048576 | 2.922x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.171x |
| ar-wikipedia | 4096 | 2.082x |
| ar-wikipedia | 65536 | 2.068x |
| ar-wikipedia | 1048576 | 2.076x |
| en-wikitext2 | 256 | 4.341x |
| en-wikitext2 | 4096 | 4.974x |
| en-wikitext2 | 65536 | 5.105x |
| en-wikitext2 | 1048576 | 5.017x |
| ja-wikipedia | 256 | 2.803x |
| ja-wikipedia | 4096 | 2.821x |
| ja-wikipedia | 65536 | 2.845x |
| ja-wikipedia | 1048576 | 2.838x |
| ru-wikipedia | 256 | 2.301x |
| ru-wikipedia | 4096 | 2.247x |
| ru-wikipedia | 65536 | 2.246x |
| ru-wikipedia | 1048576 | 2.263x |
| synthetic | 0 | 3.883x |
| zh-wikipedia | 256 | 2.970x |
| zh-wikipedia | 4096 | 3.004x |
| zh-wikipedia | 65536 | 3.012x |
| zh-wikipedia | 1048576 | 2.999x |
| Combined selected natural corpora | 256 | 2.725x |
| Combined selected natural corpora | 4096 | 2.720x |
| Combined selected natural corpora | 65536 | 2.733x |
| Combined selected natural corpora | 1048576 | 2.733x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 5.360x |
| ar-wikipedia | 4096 | 4.429x |
| ar-wikipedia | 65536 | 4.283x |
| ar-wikipedia | 1048576 | 5.150x |
| en-wikitext2 | 256 | 15.471x |
| en-wikitext2 | 4096 | 13.827x |
| en-wikitext2 | 65536 | 13.821x |
| en-wikitext2 | 1048576 | 16.873x |
| ja-wikipedia | 256 | 8.173x |
| ja-wikipedia | 4096 | 7.097x |
| ja-wikipedia | 65536 | 7.047x |
| ja-wikipedia | 1048576 | 8.596x |
| ru-wikipedia | 256 | 5.792x |
| ru-wikipedia | 4096 | 4.865x |
| ru-wikipedia | 65536 | 4.739x |
| ru-wikipedia | 1048576 | 5.943x |
| synthetic | 0 | 17.012x |
| zh-wikipedia | 256 | 8.316x |
| zh-wikipedia | 4096 | 7.357x |
| zh-wikipedia | 65536 | 7.292x |
| zh-wikipedia | 1048576 | 8.628x |
| Combined selected natural corpora | 256 | 7.502x |
| Combined selected natural corpora | 4096 | 6.403x |
| Combined selected natural corpora | 65536 | 6.312x |
| Combined selected natural corpora | 1048576 | 7.695x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 4.905x |
| ar-wikipedia | 4096 | 4.250x |
| ar-wikipedia | 65536 | 4.147x |
| ar-wikipedia | 1048576 | 4.958x |
| en-wikitext2 | 256 | 10.992x |
| en-wikitext2 | 4096 | 11.568x |
| en-wikitext2 | 65536 | 11.932x |
| en-wikitext2 | 1048576 | 14.336x |
| ja-wikipedia | 256 | 6.754x |
| ja-wikipedia | 4096 | 6.424x |
| ja-wikipedia | 65536 | 6.526x |
| ja-wikipedia | 1048576 | 7.947x |
| ru-wikipedia | 256 | 5.221x |
| ru-wikipedia | 4096 | 4.598x |
| ru-wikipedia | 65536 | 4.569x |
| ru-wikipedia | 1048576 | 5.685x |
| synthetic | 0 | 10.524x |
| zh-wikipedia | 256 | 6.961x |
| zh-wikipedia | 4096 | 6.678x |
| zh-wikipedia | 65536 | 6.741x |
| zh-wikipedia | 1048576 | 7.977x |
| Combined selected natural corpora | 256 | 6.427x |
| Combined selected natural corpora | 4096 | 5.918x |
| Combined selected natural corpora | 65536 | 5.941x |
| Combined selected natural corpora | 1048576 | 7.200x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 335.1 | 2.0 | 10 |
| onig-compile | 322.3 | 1.3 | 10 |
| pretok-compile | 9.9 | 0.1 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
