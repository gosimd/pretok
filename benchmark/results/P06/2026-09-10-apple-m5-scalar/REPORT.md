# P06 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 36951232 | 206926 | 9003.7 | 28.38 | 15820176 | 3191098 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-at | 21085803 | 164063 | 5137.9 | 49.73 | 27723630 | 5592151 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 23671647 | 173588 | 5767.9 | 44.30 | 24695156 | 4981276 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 22753647 | 132946 | 5544.3 | 46.08 | 25691486 | 5182246 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 22753348 | 147590 | 5544.2 | 46.08 | 25691824 | 5182314 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 7096397 | 29974 | 1729.1 | 147.76 | 82376313 | 16616179 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 7991755 | 22078 | 1947.3 | 131.21 | 73147266 | 14754582 | 4850592.0 | 25313.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 7102571 | 31897 | 1730.6 | 147.63 | 82304700 | 16601734 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 32837227 | 119067 | 127771.3 | 31.93 | 17802204 | 3488723 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-at | 20464619 | 89515 | 79628.9 | 51.24 | 28565154 | 5597954 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 21795938 | 89021 | 84809.1 | 48.11 | 26820364 | 5256025 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 21366090 | 25489 | 83136.5 | 49.08 | 27359942 | 5361767 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 21541802 | 75445 | 83820.2 | 48.68 | 27136773 | 5318032 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 6947856 | 6541 | 27034.5 | 150.92 | 84137460 | 16488539 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 7444806 | 33660 | 28968.1 | 140.85 | 78521182 | 15387909 | 4462064.0 | 2585.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 6986468 | 22506 | 27184.7 | 150.09 | 83672467 | 16397413 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 31949612 | 168699 | 1879388.9 | 32.82 | 18296779 | 3578885 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-at | 20369490 | 86082 | 1198205.3 | 51.48 | 28698558 | 5613494 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 21322716 | 147333 | 1254277.4 | 49.18 | 27415597 | 5362544 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 21210888 | 26656 | 1247699.3 | 49.44 | 27560138 | 5390816 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 21286748 | 87927 | 1252161.7 | 49.26 | 27461921 | 5371605 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 6944305 | 21164 | 408488.5 | 151.00 | 84180493 | 16465867 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 7493081 | 23056 | 440769.5 | 139.94 | 78015300 | 15259944 | 8081200.0 | 290.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 6973933 | 15924 | 410231.3 | 150.36 | 83822860 | 16395913 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 36083702 | 409469 | 36083702.4 | 29.06 | 16200527 | 3168439 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-at | 20438824 | 55255 | 20438824.2 | 51.30 | 28601205 | 5593717 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 21375402 | 127551 | 21375401.9 | 49.06 | 27348024 | 5348625 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 21285296 | 220006 | 21285296.0 | 49.26 | 27463795 | 5371267 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 21418043 | 182363 | 21418042.5 | 48.96 | 27293577 | 5337976 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 6927606 | 12144 | 6927605.6 | 151.36 | 84383412 | 16503393 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 7452934 | 22902 | 7452934.4 | 140.69 | 78435548 | 15340132 | 8942832.0 | 29.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 6974536 | 9121 | 6974535.9 | 150.34 | 83815613 | 16392345 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 52113271 | 671723 | 12719.9 | 20.12 | 20080048 | 4049602 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-at | 28323517 | 172136 | 6913.2 | 37.02 | 36945870 | 7450981 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 30996288 | 151990 | 7565.6 | 33.83 | 33760075 | 6808493 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 29656828 | 152898 | 7238.7 | 35.36 | 35284859 | 7116000 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 29727876 | 165987 | 7256.0 | 35.27 | 35200530 | 7098993 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 4821488 | 34395 | 1176.8 | 217.48 | 217036093 | 43770301 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 6084842 | 27981 | 1485.2 | 172.33 | 171974381 | 34682575 | 8595568.0 | 28808.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 4883610 | 28872 | 1192.0 | 214.71 | 214275317 | 43213528 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 48000432 | 123545 | 187501.7 | 21.85 | 21800575 | 4333044 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-at | 27795542 | 102303 | 108576.3 | 37.72 | 37647656 | 7482783 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 29124057 | 198071 | 113765.8 | 36.00 | 35930331 | 7141450 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 28549917 | 191170 | 111523.1 | 36.73 | 36652892 | 7285065 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 28866249 | 129480 | 112758.8 | 36.33 | 36251229 | 7205231 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 4727772 | 14299 | 18467.9 | 221.79 | 221338287 | 43992813 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 5522027 | 25748 | 21570.4 | 189.89 | 189502348 | 37665158 | 9015296.0 | 2882.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 4785477 | 11282 | 18693.3 | 219.12 | 218669329 | 43462336 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 47651223 | 488793 | 2978201.5 | 22.01 | 21960339 | 4360832 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-at | 27583906 | 267972 | 1723994.1 | 38.01 | 37936506 | 7533342 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 28330839 | 276711 | 1770677.4 | 37.01 | 36936322 | 7334728 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 28328128 | 122388 | 1770508.0 | 37.02 | 36939856 | 7335430 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 28217446 | 118709 | 1763590.4 | 37.16 | 37084753 | 7364203 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 4761857 | 31199 | 297616.1 | 220.20 | 219753975 | 43638228 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 5609393 | 23981 | 350587.0 | 186.93 | 186550858 | 37044831 | 14634752.0 | 320.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 4817585 | 12930 | 301099.1 | 217.66 | 217211937 | 43133436 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 52627069 | 476002 | 52627068.8 | 19.92 | 19884007 | 3948253 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-at | 27840273 | 211141 | 27840273.1 | 37.66 | 37587167 | 7463468 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 28562575 | 186754 | 28562575.2 | 36.71 | 36636647 | 7274729 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 28457674 | 157326 | 28457673.6 | 36.85 | 36771699 | 7301546 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 28562976 | 91016 | 28562975.7 | 36.71 | 36636134 | 7274627 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 4757090 | 51559 | 4757090.4 | 220.42 | 219974170 | 43679011 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 5591515 | 16694 | 5591514.6 | 187.53 | 187147326 | 37160772 | 17962224.0 | 32.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 4846704 | 17144 | 4846703.7 | 216.35 | 215906946 | 42871405 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 24133949 | 90610 | 5867.7 | 43.45 | 15715621 | 2282138 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-at | 14415179 | 87329 | 3504.8 | 72.74 | 26311155 | 3820764 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 15952945 | 61906 | 3878.7 | 65.73 | 23774920 | 3452466 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 15313123 | 62018 | 3723.1 | 68.48 | 24768298 | 3596719 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 15331513 | 54428 | 3727.6 | 68.39 | 24738590 | 3592405 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 6343509 | 31669 | 1542.3 | 165.30 | 59790250 | 8682418 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 6992597 | 15930 | 1700.1 | 149.96 | 54240221 | 7876473 | 2385136.0 | 20228.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 6334855 | 26919 | 1540.2 | 165.52 | 59871930 | 8694279 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 21293694 | 133804 | 82854.8 | 49.24 | 17811846 | 2427949 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-at | 13851072 | 55886 | 53895.2 | 75.70 | 27382718 | 3732563 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 14430964 | 123860 | 56151.6 | 72.66 | 26282374 | 3582574 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 14196028 | 70347 | 55237.5 | 73.86 | 26717332 | 3641864 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 14251626 | 99948 | 55453.8 | 73.58 | 26613103 | 3627656 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 6190458 | 20333 | 24087.4 | 169.39 | 61268487 | 8351563 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 6521856 | 19273 | 25376.9 | 160.78 | 58155229 | 7927192 | 2401776.0 | 2326.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 6199951 | 7742 | 24124.3 | 169.13 | 61174678 | 8338776 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 21027937 | 139676 | 1236937.4 | 49.87 | 18036958 | 2448267 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-at | 13858328 | 79945 | 815195.7 | 75.66 | 27368382 | 3714878 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 14078939 | 24294 | 828172.9 | 74.48 | 26939530 | 3656668 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 14024676 | 46386 | 824980.9 | 74.77 | 27043762 | 3670816 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 14008339 | 62427 | 824019.9 | 74.85 | 27075301 | 3675097 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 6202204 | 16647 | 364835.5 | 169.06 | 61152458 | 8300598 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 6419708 | 10860 | 377629.9 | 163.34 | 59080571 | 8019368 | 3108624.0 | 244.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 6185036 | 6123 | 363825.6 | 169.53 | 61322198 | 8323638 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 23954381 | 254928 | 23954381.0 | 43.77 | 15833429 | 2148501 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-at | 13789753 | 37247 | 13789752.8 | 76.04 | 27504481 | 3732192 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 14037207 | 123383 | 14037207.2 | 74.70 | 27019620 | 3666399 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 14003747 | 45890 | 14003746.5 | 74.88 | 27084181 | 3675159 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 14018547 | 119044 | 14018547.4 | 74.80 | 27055585 | 3671279 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 6173094 | 14645 | 6173093.8 | 169.86 | 61440829 | 8337149 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 6475299 | 34890 | 6475298.9 | 161.93 | 58573358 | 7948050 | 4363504.0 | 26.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 6191330 | 22679 | 6191330.2 | 169.36 | 61259856 | 8312592 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 33981978 | 227519 | 8280.2 | 30.86 | 16916526 | 2976519 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-at | 18669695 | 98369 | 4549.1 | 56.16 | 30790915 | 5417764 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 20783875 | 178621 | 5064.3 | 50.45 | 27658798 | 4866657 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 19806417 | 64294 | 4826.1 | 52.94 | 29023775 | 5106830 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 19912502 | 57442 | 4852.0 | 52.66 | 28869149 | 5079623 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 7314572 | 22167 | 1782.3 | 143.35 | 78590656 | 13828287 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 8177020 | 27215 | 1992.5 | 128.23 | 70301524 | 12369787 | 4424064.0 | 24864.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 7352424 | 24434 | 1791.5 | 142.62 | 78186049 | 13757095 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 29865143 | 161289 | 116206.8 | 35.11 | 19248426 | 3271372 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-at | 18075585 | 39015 | 70333.0 | 58.01 | 31802954 | 5405081 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 18973046 | 46690 | 73825.1 | 55.27 | 30298614 | 5149410 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 18569110 | 101993 | 72253.3 | 56.47 | 30957704 | 5261426 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 18709247 | 70362 | 72798.6 | 56.05 | 30725823 | 5222017 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 7173800 | 14405 | 27913.6 | 146.17 | 80132845 | 13619002 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 7646960 | 30674 | 29754.7 | 137.12 | 75174583 | 12776320 | 4258544.0 | 2570.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 7209140 | 18102 | 28051.1 | 145.45 | 79740026 | 13552241 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 29064323 | 216696 | 1709666.1 | 36.08 | 19778785 | 3354147 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-at | 17987258 | 84537 | 1058074.0 | 58.30 | 31959123 | 5419725 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 18475297 | 32217 | 1086782.2 | 56.76 | 31114899 | 5276559 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 18469262 | 276108 | 1086427.2 | 56.77 | 31125066 | 5278283 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 18439156 | 101444 | 1084656.2 | 56.87 | 31175885 | 5286902 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 7172584 | 20642 | 421916.7 | 146.19 | 80146427 | 13591475 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 7558683 | 15530 | 444628.4 | 138.72 | 76052533 | 12897220 | 5852944.0 | 273.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 7202126 | 17583 | 423654.4 | 145.59 | 79817686 | 13535726 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 33197206 | 488626 | 33197205.8 | 31.59 | 17316427 | 2936151 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-at | 18069275 | 155613 | 18069274.5 | 58.03 | 31814061 | 5394350 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 18590621 | 147854 | 18590621.1 | 56.40 | 30921882 | 5243074 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 18467484 | 101075 | 18467484.4 | 56.78 | 31128062 | 5278033 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 18509624 | 73299 | 18509623.9 | 56.65 | 31057195 | 5266017 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 7151683 | 11790 | 7151682.7 | 146.62 | 80380663 | 13629240 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 7672402 | 32920 | 7672401.8 | 136.67 | 74925299 | 12704236 | 8942832.0 | 29.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 7180509 | 31863 | 7180508.6 | 146.03 | 80057978 | 13574526 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 101257725 | 910404 | 2177.8 | 21.08 | 11290684 | 2884896 | N/A | N/A | 10 |
| synthetic | 0 | onig-at | 51937950 | 273595 | 1117.0 | 41.11 | 22012209 | 5624365 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 62571001 | 212850 | 1345.7 | 34.12 | 18271547 | 4668584 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 60785697 | 374587 | 1307.3 | 35.12 | 18808191 | 4805703 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 61179982 | 270351 | 1315.8 | 34.90 | 18686979 | 4774732 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 16182256 | 70325 | 348.0 | 131.93 | 70649543 | 18051747 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 18895052 | 38211 | 406.4 | 112.99 | 60506265 | 15460026 | 13309360.0 | 118536.0 | 10 |
| synthetic | 0 | pretok-reused | 16332638 | 51163 | 351.3 | 130.72 | 69999039 | 17885536 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 23973772 | 152501 | 5828.8 | 43.74 | 16276246 | 2514748 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-at | 13618548 | 78343 | 3311.1 | 77.00 | 28652319 | 4426904 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 15270080 | 79938 | 3712.6 | 68.67 | 25553436 | 3948113 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 14551895 | 39266 | 3538.0 | 72.06 | 26814584 | 4142966 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 14611505 | 99485 | 3552.5 | 71.76 | 26705188 | 4126064 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 5231646 | 7664 | 1272.0 | 200.43 | 74585124 | 11523714 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 5936622 | 67146 | 1443.4 | 176.63 | 65728119 | 10155270 | 2632176.0 | 21018.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 5277731 | 17000 | 1283.2 | 198.68 | 73933856 | 11423091 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 20947205 | 54633 | 81506.6 | 50.06 | 18627927 | 2722177 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-at | 13169559 | 107108 | 51243.4 | 79.62 | 29629161 | 4329834 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 13673175 | 73584 | 53203.0 | 76.69 | 28537848 | 4170355 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 13472655 | 79433 | 52422.8 | 77.83 | 28962591 | 4232425 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 13478960 | 42880 | 52447.3 | 77.79 | 28949044 | 4230445 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 5154358 | 21594 | 20055.9 | 203.43 | 75703513 | 11062872 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 5459372 | 18721 | 21242.7 | 192.07 | 71473974 | 10444791 | 2620144.0 | 2361.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 5169471 | 19165 | 20114.7 | 202.84 | 75482198 | 11030530 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 20635327 | 72335 | 1213842.7 | 50.81 | 18909466 | 2752997 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-at | 13084144 | 83670 | 769655.5 | 80.14 | 29822586 | 4341820 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 13413145 | 85011 | 789008.5 | 78.18 | 29091090 | 4235323 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 13286770 | 19728 | 781574.7 | 78.92 | 29367784 | 4275606 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 13391573 | 148160 | 787739.6 | 78.30 | 29137951 | 4242145 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 5131959 | 11235 | 301879.9 | 204.32 | 76033935 | 11069653 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 5414991 | 22018 | 318528.9 | 193.64 | 72059765 | 10491060 | 3436336.0 | 249.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 5159814 | 8956 | 303518.5 | 203.22 | 75623462 | 11009893 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 23758260 | 278526 | 23758260.4 | 44.14 | 16423888 | 2390663 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-at | 13089950 | 73232 | 13089949.8 | 80.11 | 29809358 | 4339054 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 13362497 | 55876 | 13362497.3 | 78.47 | 29201353 | 4250553 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 13339407 | 47717 | 13339406.8 | 78.61 | 29251900 | 4257910 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 13363253 | 43084 | 13363253.3 | 78.47 | 29199701 | 4250312 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 5145107 | 11314 | 5145106.7 | 203.80 | 75839632 | 11039227 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 5403035 | 9643 | 5403035.2 | 194.07 | 72219223 | 10512240 | 4363504.0 | 26.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 5158928 | 7997 | 5158928.0 | 203.25 | 75636449 | 11009652 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.206x |
| ar-wikipedia | 4096 | 3.075x |
| ar-wikipedia | 65536 | 3.054x |
| ar-wikipedia | 1048576 | 3.073x |
| en-wikitext2 | 256 | 6.151x |
| en-wikitext2 | 4096 | 6.039x |
| en-wikitext2 | 65536 | 5.949x |
| en-wikitext2 | 1048576 | 5.982x |
| ja-wikipedia | 256 | 2.414x |
| ja-wikipedia | 4096 | 2.293x |
| ja-wikipedia | 65536 | 2.261x |
| ja-wikipedia | 1048576 | 2.269x |
| ru-wikipedia | 256 | 2.708x |
| ru-wikipedia | 4096 | 2.588x |
| ru-wikipedia | 65536 | 2.575x |
| ru-wikipedia | 1048576 | 2.582x |
| synthetic | 0 | 3.756x |
| zh-wikipedia | 256 | 2.782x |
| zh-wikipedia | 4096 | 2.614x |
| zh-wikipedia | 65536 | 2.589x |
| zh-wikipedia | 1048576 | 2.593x |
| Combined selected natural corpora | 256 | 3.314x |
| Combined selected natural corpora | 4096 | 3.185x |
| Combined selected natural corpora | 65536 | 3.155x |
| Combined selected natural corpora | 1048576 | 3.169x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.204x |
| ar-wikipedia | 4096 | 3.083x |
| ar-wikipedia | 65536 | 3.052x |
| ar-wikipedia | 1048576 | 3.071x |
| en-wikitext2 | 256 | 6.087x |
| en-wikitext2 | 4096 | 6.032x |
| en-wikitext2 | 65536 | 5.857x |
| en-wikitext2 | 1048576 | 5.893x |
| ja-wikipedia | 256 | 2.420x |
| ja-wikipedia | 4096 | 2.299x |
| ja-wikipedia | 65536 | 2.265x |
| ja-wikipedia | 1048576 | 2.264x |
| ru-wikipedia | 256 | 2.708x |
| ru-wikipedia | 4096 | 2.595x |
| ru-wikipedia | 65536 | 2.560x |
| ru-wikipedia | 1048576 | 2.578x |
| synthetic | 0 | 3.746x |
| zh-wikipedia | 256 | 2.769x |
| zh-wikipedia | 4096 | 2.607x |
| zh-wikipedia | 65536 | 2.595x |
| zh-wikipedia | 1048576 | 2.590x |
| Combined selected natural corpora | 256 | 3.306x |
| Combined selected natural corpora | 4096 | 3.191x |
| Combined selected natural corpora | 65536 | 3.143x |
| Combined selected natural corpora | 1048576 | 3.159x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.962x |
| ar-wikipedia | 4096 | 2.928x |
| ar-wikipedia | 65536 | 2.846x |
| ar-wikipedia | 1048576 | 2.868x |
| en-wikitext2 | 256 | 5.094x |
| en-wikitext2 | 4096 | 5.274x |
| en-wikitext2 | 65536 | 5.051x |
| en-wikitext2 | 1048576 | 5.108x |
| ja-wikipedia | 256 | 2.281x |
| ja-wikipedia | 4096 | 2.213x |
| ja-wikipedia | 65536 | 2.193x |
| ja-wikipedia | 1048576 | 2.168x |
| ru-wikipedia | 256 | 2.542x |
| ru-wikipedia | 4096 | 2.481x |
| ru-wikipedia | 65536 | 2.444x |
| ru-wikipedia | 1048576 | 2.423x |
| synthetic | 0 | 3.312x |
| zh-wikipedia | 256 | 2.572x |
| zh-wikipedia | 4096 | 2.505x |
| zh-wikipedia | 65536 | 2.477x |
| zh-wikipedia | 1048576 | 2.473x |
| Combined selected natural corpora | 256 | 3.032x |
| Combined selected natural corpora | 4096 | 3.007x |
| Combined selected natural corpora | 65536 | 2.943x |
| Combined selected natural corpora | 1048576 | 2.943x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 5.203x |
| ar-wikipedia | 4096 | 4.700x |
| ar-wikipedia | 65536 | 4.581x |
| ar-wikipedia | 1048576 | 5.174x |
| en-wikitext2 | 256 | 10.671x |
| en-wikitext2 | 4096 | 10.030x |
| en-wikitext2 | 65536 | 9.891x |
| en-wikitext2 | 1048576 | 10.858x |
| ja-wikipedia | 256 | 3.810x |
| ja-wikipedia | 4096 | 3.434x |
| ja-wikipedia | 65536 | 3.400x |
| ja-wikipedia | 1048576 | 3.869x |
| ru-wikipedia | 256 | 4.622x |
| ru-wikipedia | 4096 | 4.143x |
| ru-wikipedia | 65536 | 4.036x |
| ru-wikipedia | 1048576 | 4.623x |
| synthetic | 0 | 6.200x |
| zh-wikipedia | 256 | 4.542x |
| zh-wikipedia | 4096 | 4.052x |
| zh-wikipedia | 65536 | 3.999x |
| zh-wikipedia | 1048576 | 4.605x |
| Combined selected natural corpora | 256 | 5.530x |
| Combined selected natural corpora | 4096 | 5.039x |
| Combined selected natural corpora | 65536 | 4.955x |
| Combined selected natural corpora | 1048576 | 5.588x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 4.624x |
| ar-wikipedia | 4096 | 4.411x |
| ar-wikipedia | 65536 | 4.264x |
| ar-wikipedia | 1048576 | 4.842x |
| en-wikitext2 | 256 | 8.564x |
| en-wikitext2 | 4096 | 8.693x |
| en-wikitext2 | 65536 | 8.495x |
| en-wikitext2 | 1048576 | 9.412x |
| ja-wikipedia | 256 | 3.451x |
| ja-wikipedia | 4096 | 3.265x |
| ja-wikipedia | 65536 | 3.276x |
| ja-wikipedia | 1048576 | 3.699x |
| ru-wikipedia | 256 | 4.156x |
| ru-wikipedia | 4096 | 3.905x |
| ru-wikipedia | 65536 | 3.845x |
| ru-wikipedia | 1048576 | 4.327x |
| synthetic | 0 | 5.359x |
| zh-wikipedia | 256 | 4.038x |
| zh-wikipedia | 4096 | 3.837x |
| zh-wikipedia | 65536 | 3.811x |
| zh-wikipedia | 1048576 | 4.397x |
| Combined selected natural corpora | 256 | 4.865x |
| Combined selected natural corpora | 4096 | 4.692x |
| Combined selected natural corpora | 65536 | 4.626x |
| Combined selected natural corpora | 1048576 | 5.204x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 313657.7 | 155.2 | 10 |
| onig-compile | 314287.9 | 941.8 | 10 |
| pretok-compile | 13.4 | 0.1 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
