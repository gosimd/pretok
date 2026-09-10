# P24 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 37687061 | 119998 | 9183.0 | 27.82 | 15511292 | 3258811 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-at | 20795019 | 111989 | 5067.0 | 50.42 | 28111299 | 5905982 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 23436690 | 194221 | 5710.7 | 44.74 | 24942728 | 5240288 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 22160416 | 257771 | 5399.7 | 47.32 | 26379243 | 5542089 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 22335099 | 168373 | 5442.3 | 46.95 | 26172931 | 5498744 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 6246410 | 8880 | 1522.0 | 167.87 | 93585752 | 19661693 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 7241220 | 41216 | 1764.4 | 144.81 | 80728804 | 16960541 | 5136288.0 | 25586.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 6295231 | 16823 | 1533.9 | 166.57 | 92859976 | 19509213 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 33323739 | 221594 | 129664.4 | 31.47 | 17542299 | 3585852 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-at | 20242266 | 98737 | 78763.7 | 51.80 | 28878931 | 5903193 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 21248720 | 215179 | 82679.8 | 49.35 | 27511068 | 5623586 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 20898800 | 148363 | 81318.3 | 50.17 | 27971701 | 5717744 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 20926047 | 101464 | 81424.3 | 50.11 | 27935281 | 5710300 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 6156736 | 14869 | 23956.2 | 170.31 | 94948857 | 19408662 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 6651153 | 31425 | 25880.0 | 157.65 | 87890780 | 17965908 | 4692720.0 | 2602.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 6179212 | 25185 | 24043.6 | 169.69 | 94603484 | 19338064 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 32537924 | 810563 | 1913995.6 | 32.23 | 17965958 | 3665815 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-at | 19571612 | 363220 | 1151271.3 | 53.58 | 29868515 | 6094439 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 20453690 | 529549 | 1203158.2 | 51.27 | 28580418 | 5831613 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 20301342 | 659240 | 1194196.6 | 51.65 | 28794895 | 5875375 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 20402633 | 602866 | 1200154.9 | 51.39 | 28651939 | 5846206 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 6135814 | 24873 | 360930.2 | 170.89 | 95272611 | 19439638 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 6707950 | 32490 | 394585.3 | 156.32 | 87146591 | 17781587 | 8081200.0 | 290.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 6178217 | 21248 | 363424.5 | 169.72 | 94618721 | 19306217 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 35808734 | 1006793 | 35808733.6 | 29.28 | 16324928 | 3330556 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-at | 19478665 | 171111 | 19478665.1 | 53.83 | 30011040 | 6122750 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 20153838 | 200850 | 20153838.3 | 52.03 | 29005641 | 5917632 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 20041639 | 110781 | 20041639.2 | 52.32 | 29168023 | 5950761 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 20304745 | 242167 | 20304745.0 | 51.64 | 28790069 | 5873652 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 6127913 | 19574 | 6127913.4 | 171.11 | 95395441 | 19462253 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 6742447 | 19166 | 6742447.2 | 155.52 | 86700716 | 17688385 | 11302128.0 | 30.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 6170089 | 18309 | 6170089.2 | 169.95 | 94743363 | 19329218 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 51678396 | 912358 | 12613.7 | 20.29 | 20249023 | 4080641 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-at | 27681986 | 178531 | 6756.6 | 37.88 | 37802093 | 7617987 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 30831929 | 439436 | 7525.5 | 34.01 | 33940043 | 6839695 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 29619167 | 219849 | 7229.5 | 35.40 | 35329725 | 7119748 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 29676854 | 139831 | 7243.6 | 35.33 | 35261049 | 7105908 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 4692705 | 5444 | 1145.4 | 223.45 | 222992272 | 44938045 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 5954043 | 42473 | 1453.3 | 176.11 | 175752332 | 35418116 | 8591472.0 | 28806.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 4753205 | 14970 | 1160.2 | 220.60 | 220153981 | 44366065 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 47889152 | 384499 | 187067.0 | 21.90 | 21851233 | 4339814 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-at | 27351442 | 356675 | 106841.6 | 38.34 | 38258934 | 7598503 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 28886580 | 162983 | 112838.2 | 36.30 | 36225715 | 7194690 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 28379436 | 158760 | 110857.2 | 36.95 | 36873072 | 7323260 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 28676597 | 241691 | 112018.0 | 36.57 | 36490975 | 7247373 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 4569683 | 12052 | 17850.3 | 229.46 | 228995546 | 45480181 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 5376336 | 21297 | 21001.3 | 195.04 | 194637587 | 38656441 | 8953856.0 | 2879.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 4630239 | 14377 | 18086.9 | 226.46 | 226000661 | 44885375 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 48698839 | 964469 | 3043677.4 | 21.53 | 21487925 | 4263777 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-at | 27689312 | 641900 | 1730582.0 | 37.87 | 37792090 | 7498958 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 28514074 | 522858 | 1782129.6 | 36.77 | 36698965 | 7282053 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 28514686 | 755672 | 1782167.9 | 36.77 | 36698177 | 7281897 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 28292438 | 963691 | 1768277.3 | 37.06 | 36986456 | 7339099 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 4588617 | 10602 | 286788.6 | 228.52 | 228050635 | 45251326 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 5443588 | 32868 | 340224.2 | 192.63 | 192232959 | 38144144 | 14634752.0 | 320.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 4652969 | 13968 | 290810.6 | 225.36 | 224896613 | 44625484 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 53606062 | 1863288 | 53606062.5 | 19.56 | 19520870 | 3873200 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-at | 28143956 | 312793 | 28143956.0 | 37.26 | 37181589 | 7377321 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 29478341 | 499330 | 29478340.7 | 35.57 | 35498504 | 7043375 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 29062448 | 468913 | 29062447.9 | 36.08 | 36006499 | 7144168 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 29268074 | 357268 | 29268074.1 | 35.83 | 35753531 | 7093975 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 4616501 | 25890 | 4616500.6 | 227.14 | 226673209 | 44974975 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 5437111 | 28493 | 5437110.7 | 192.86 | 192461964 | 38187010 | 17962224.0 | 32.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 4679389 | 13441 | 4679389.4 | 224.08 | 223626826 | 44370533 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 23401236 | 155732 | 5689.6 | 44.81 | 16207691 | 2348423 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-at | 13436323 | 123470 | 3266.8 | 78.04 | 28227960 | 4090107 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 15107856 | 121397 | 3673.2 | 69.41 | 25104820 | 3637578 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 14456886 | 86836 | 3514.9 | 72.53 | 26235248 | 3801372 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 14449836 | 112607 | 3513.2 | 72.57 | 26248048 | 3803226 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 5365439 | 37749 | 1304.5 | 195.43 | 70689465 | 10242592 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 6033655 | 17633 | 1467.0 | 173.79 | 62860740 | 9108244 | 2381424.0 | 20218.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 5382973 | 23161 | 1308.8 | 194.79 | 70459204 | 10209228 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 20195787 | 93554 | 78582.8 | 51.92 | 18780155 | 2553750 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-at | 12823989 | 70215 | 49898.8 | 81.77 | 29575822 | 4021760 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 13382821 | 90475 | 52073.2 | 78.35 | 28340810 | 3853821 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 13153376 | 61941 | 51180.4 | 79.72 | 28835184 | 3921047 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 13226606 | 132477 | 51465.4 | 79.28 | 28675534 | 3899337 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 5251524 | 17474 | 20433.9 | 199.67 | 72222840 | 9820958 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 5584396 | 38708 | 21729.2 | 187.77 | 67917822 | 9235556 | 2401776.0 | 2326.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 5269368 | 6475 | 20503.4 | 198.99 | 71978272 | 9787701 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 20089530 | 42732 | 1181737.1 | 52.20 | 18879486 | 2556506 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-at | 12799826 | 152467 | 752930.9 | 81.92 | 29631653 | 4012476 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 13069191 | 127167 | 768776.0 | 80.23 | 29020924 | 3929776 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 13035545 | 85415 | 766796.8 | 80.44 | 29095830 | 3939920 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 13077326 | 109779 | 769254.5 | 80.18 | 29002870 | 3927332 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 5244425 | 11472 | 308495.6 | 199.94 | 72320611 | 9793067 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 5540768 | 18671 | 325927.5 | 189.25 | 68452609 | 9269293 | 3051280.0 | 243.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 5281884 | 14683 | 310699.1 | 198.52 | 71807708 | 9723613 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 21843084 | 180592 | 21843084.2 | 48.00 | 17363848 | 2350538 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-at | 12811495 | 159372 | 12811495.4 | 81.85 | 29604663 | 4007573 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 13045250 | 86881 | 13045250.0 | 80.38 | 29074184 | 3935762 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 13019006 | 67240 | 13019005.9 | 80.54 | 29132793 | 3943696 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 13091814 | 70232 | 13091813.6 | 80.09 | 28970776 | 3921764 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 5234165 | 32223 | 5234164.7 | 200.33 | 72462374 | 9809206 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 5556962 | 21525 | 5556961.8 | 188.70 | 68253123 | 9239401 | 4363504.0 | 26.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 5271172 | 12763 | 5271171.7 | 198.93 | 71953642 | 9740339 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 32922721 | 332458 | 8022.1 | 31.85 | 17460798 | 3075080 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-at | 17542065 | 87471 | 4274.4 | 59.77 | 32770202 | 5771270 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 19811953 | 125723 | 4827.5 | 52.93 | 29015666 | 5110047 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 18748758 | 98819 | 4568.4 | 55.93 | 30661071 | 5399824 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 18963586 | 166287 | 4620.8 | 55.29 | 30313729 | 5338653 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 4674115 | 76929 | 1138.9 | 224.34 | 122987355 | 21659717 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 5531414 | 45525 | 1347.8 | 189.57 | 103925869 | 18302734 | 4433792.0 | 24874.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 4697032 | 45101 | 1144.5 | 223.24 | 122387283 | 21554036 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 28700084 | 135064 | 111673.5 | 36.54 | 20029802 | 3407481 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-at | 16986272 | 173428 | 66094.4 | 61.73 | 33842447 | 5757296 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 17892723 | 244926 | 69621.5 | 58.60 | 32127977 | 5465630 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 17584643 | 79086 | 68422.7 | 59.63 | 32690854 | 5561387 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 17709126 | 240609 | 68907.1 | 59.21 | 32461060 | 5522294 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 4591947 | 72135 | 17867.5 | 228.35 | 125188087 | 21297069 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 5064393 | 19465 | 19705.8 | 207.05 | 113509554 | 19310310 | 4258544.0 | 2570.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 4617847 | 35217 | 17968.3 | 227.07 | 124485942 | 21177619 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 28458603 | 77409 | 1674035.5 | 36.85 | 20199762 | 3428875 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-at | 16848192 | 142711 | 991070.1 | 62.24 | 34119803 | 5791779 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 17438414 | 246048 | 1025789.0 | 60.13 | 32964982 | 5595750 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 17312590 | 115328 | 1018387.6 | 60.57 | 33204564 | 5636418 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 17464733 | 150113 | 1027337.3 | 60.04 | 32915303 | 5587317 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 4563814 | 9019 | 268459.6 | 229.76 | 125959783 | 21381459 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 4962653 | 15640 | 291920.8 | 211.29 | 115836634 | 19663072 | 5852944.0 | 273.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 4601519 | 17441 | 270677.6 | 227.88 | 124927662 | 21206259 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 31834844 | 561686 | 31834843.7 | 32.94 | 18057478 | 3064787 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-at | 17005463 | 98786 | 17005463.2 | 61.66 | 33804254 | 5737391 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 17636163 | 161584 | 17636163.1 | 59.46 | 32595355 | 5532212 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 17457639 | 152779 | 17457639.4 | 60.06 | 32928679 | 5588785 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 17719851 | 340819 | 17719850.6 | 59.18 | 32441414 | 5506085 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 4583742 | 40177 | 4583742.0 | 228.76 | 125412162 | 21285447 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 5056911 | 43447 | 5056911.1 | 207.35 | 113677499 | 19293794 | 8942832.0 | 29.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 4612918 | 40359 | 4612918.0 | 227.31 | 124618950 | 21150820 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 92891257 | 553203 | 2081.6 | 17.50 | 9126047 | 3066349 | N/A | N/A | 10 |
| synthetic | 0 | onig-at | 46310331 | 495005 | 1037.8 | 35.09 | 18305419 | 6150615 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 56669134 | 511529 | 1269.9 | 28.68 | 14959290 | 5026316 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 54642942 | 354724 | 1224.5 | 29.74 | 15513989 | 5212695 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 55204248 | 733944 | 1237.1 | 29.44 | 15356246 | 5159694 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 11840625 | 73367 | 265.3 | 137.26 | 71595035 | 24055908 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 14497864 | 55381 | 324.9 | 112.10 | 58472753 | 19646826 | 13054320.0 | 113867.0 | 10 |
| synthetic | 0 | pretok-reused | 11952311 | 36696 | 267.8 | 135.98 | 70926036 | 23831125 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 22965616 | 139433 | 5583.7 | 45.66 | 16990749 | 2602151 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-at | 12317521 | 66423 | 2994.8 | 85.13 | 31678695 | 4851625 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 13974698 | 143801 | 3397.7 | 75.03 | 27922105 | 4276300 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 13373691 | 73597 | 3251.6 | 78.41 | 29176911 | 4468475 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 13371284 | 68401 | 3251.0 | 78.42 | 29182164 | 4469279 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 4081270 | 12339 | 992.3 | 256.92 | 95608238 | 14642502 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 4748750 | 20523 | 1154.6 | 220.81 | 82169630 | 12584365 | 2606960.0 | 20962.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 4101180 | 6399 | 997.1 | 255.68 | 95144089 | 14571417 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 19773985 | 102141 | 76941.6 | 53.03 | 19733150 | 2856632 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-at | 11853006 | 79296 | 46120.6 | 88.46 | 32920171 | 4765626 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 12378143 | 86394 | 48164.0 | 84.71 | 31523548 | 4563447 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 12154111 | 102144 | 47292.3 | 86.27 | 32104612 | 4647563 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 12263847 | 85368 | 47719.2 | 85.50 | 31817341 | 4605977 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 3992198 | 11048 | 15533.8 | 262.66 | 97741397 | 14149349 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 4296978 | 39611 | 16719.8 | 244.03 | 90808712 | 13145752 | 2562800.0 | 2355.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 4005071 | 14964 | 15583.9 | 261.81 | 97427237 | 14103870 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 19660453 | 405657 | 1156497.2 | 53.33 | 19847101 | 2862294 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-at | 11864615 | 303792 | 697918.5 | 88.38 | 32887961 | 4743011 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 12096888 | 288419 | 711581.6 | 86.68 | 32256478 | 4651940 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 12098226 | 214094 | 711660.4 | 86.67 | 32252910 | 4651426 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 12013089 | 265959 | 706652.3 | 87.29 | 32481488 | 4684391 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 3979361 | 14911 | 234080.1 | 263.50 | 98056699 | 14141467 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 4250244 | 32554 | 250014.4 | 246.71 | 91807192 | 13240180 | 3436336.0 | 249.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 3999846 | 27044 | 235285.1 | 262.15 | 97554497 | 14069040 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 21646415 | 348514 | 21646414.9 | 48.44 | 18026218 | 2599183 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-at | 12086328 | 47777 | 12086328.4 | 86.76 | 32284660 | 4655094 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 12386795 | 162602 | 12386795.2 | 84.65 | 31501530 | 4542176 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 12279069 | 109514 | 12279069.1 | 85.40 | 31777898 | 4582025 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 12415037 | 113967 | 12415037.1 | 84.46 | 31429870 | 4531843 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 3981259 | 6055 | 3981259.4 | 263.38 | 98009941 | 14131960 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 4268605 | 23245 | 4268604.5 | 245.65 | 91412310 | 13180654 | 4363504.0 | 26.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 4001489 | 10342 | 4001488.8 | 262.05 | 97514454 | 14060517 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.548x |
| ar-wikipedia | 4096 | 3.394x |
| ar-wikipedia | 65536 | 3.309x |
| ar-wikipedia | 1048576 | 3.271x |
| en-wikitext2 | 256 | 6.312x |
| en-wikitext2 | 4096 | 6.210x |
| en-wikitext2 | 65536 | 6.214x |
| en-wikitext2 | 1048576 | 6.295x |
| ja-wikipedia | 256 | 2.694x |
| ja-wikipedia | 4096 | 2.505x |
| ja-wikipedia | 65536 | 2.486x |
| ja-wikipedia | 1048576 | 2.487x |
| ru-wikipedia | 256 | 4.011x |
| ru-wikipedia | 4096 | 3.829x |
| ru-wikipedia | 65536 | 3.793x |
| ru-wikipedia | 1048576 | 3.809x |
| synthetic | 0 | 4.615x |
| zh-wikipedia | 256 | 3.277x |
| zh-wikipedia | 4096 | 3.044x |
| zh-wikipedia | 65536 | 3.040x |
| zh-wikipedia | 1048576 | 3.084x |
| Combined selected natural corpora | 256 | 3.925x |
| Combined selected natural corpora | 4096 | 3.753x |
| Combined selected natural corpora | 65536 | 3.723x |
| Combined selected natural corpora | 1048576 | 3.743x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.548x |
| ar-wikipedia | 4096 | 3.387x |
| ar-wikipedia | 65536 | 3.302x |
| ar-wikipedia | 1048576 | 3.291x |
| en-wikitext2 | 256 | 6.244x |
| en-wikitext2 | 4096 | 6.193x |
| en-wikitext2 | 65536 | 6.081x |
| en-wikitext2 | 1048576 | 6.255x |
| ja-wikipedia | 256 | 2.684x |
| ja-wikipedia | 4096 | 2.510x |
| ja-wikipedia | 65536 | 2.476x |
| ja-wikipedia | 1048576 | 2.484x |
| ru-wikipedia | 256 | 4.037x |
| ru-wikipedia | 4096 | 3.835x |
| ru-wikipedia | 65536 | 3.795x |
| ru-wikipedia | 1048576 | 3.841x |
| synthetic | 0 | 4.619x |
| zh-wikipedia | 256 | 3.260x |
| zh-wikipedia | 4096 | 3.062x |
| zh-wikipedia | 65536 | 3.003x |
| zh-wikipedia | 1048576 | 3.103x |
| Combined selected natural corpora | 256 | 3.916x |
| Combined selected natural corpora | 4096 | 3.757x |
| Combined selected natural corpora | 65536 | 3.692x |
| Combined selected natural corpora | 1048576 | 3.752x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.237x |
| ar-wikipedia | 4096 | 3.195x |
| ar-wikipedia | 65536 | 3.049x |
| ar-wikipedia | 1048576 | 2.989x |
| en-wikitext2 | 256 | 5.178x |
| en-wikitext2 | 4096 | 5.373x |
| en-wikitext2 | 65536 | 5.238x |
| en-wikitext2 | 1048576 | 5.422x |
| ja-wikipedia | 256 | 2.504x |
| ja-wikipedia | 4096 | 2.396x |
| ja-wikipedia | 65536 | 2.359x |
| ja-wikipedia | 1048576 | 2.348x |
| ru-wikipedia | 256 | 3.582x |
| ru-wikipedia | 4096 | 3.533x |
| ru-wikipedia | 65536 | 3.514x |
| ru-wikipedia | 1048576 | 3.488x |
| synthetic | 0 | 3.909x |
| zh-wikipedia | 256 | 2.943x |
| zh-wikipedia | 4096 | 2.881x |
| zh-wikipedia | 65536 | 2.846x |
| zh-wikipedia | 1048576 | 2.902x |
| Combined selected natural corpora | 256 | 3.496x |
| Combined selected natural corpora | 4096 | 3.477x |
| Combined selected natural corpora | 65536 | 3.404x |
| Combined selected natural corpora | 1048576 | 3.425x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 5.987x |
| ar-wikipedia | 4096 | 5.393x |
| ar-wikipedia | 65536 | 5.267x |
| ar-wikipedia | 1048576 | 5.804x |
| en-wikitext2 | 256 | 10.872x |
| en-wikitext2 | 4096 | 10.343x |
| en-wikitext2 | 65536 | 10.466x |
| en-wikitext2 | 1048576 | 11.456x |
| ja-wikipedia | 256 | 4.347x |
| ja-wikipedia | 4096 | 3.833x |
| ja-wikipedia | 65536 | 3.803x |
| ja-wikipedia | 1048576 | 4.144x |
| ru-wikipedia | 256 | 7.009x |
| ru-wikipedia | 4096 | 6.215x |
| ru-wikipedia | 65536 | 6.185x |
| ru-wikipedia | 1048576 | 6.901x |
| synthetic | 0 | 7.772x |
| zh-wikipedia | 256 | 5.600x |
| zh-wikipedia | 4096 | 4.937x |
| zh-wikipedia | 65536 | 4.915x |
| zh-wikipedia | 1048576 | 5.410x |
| Combined selected natural corpora | 256 | 6.685x |
| Combined selected natural corpora | 4096 | 6.068x |
| Combined selected natural corpora | 65536 | 6.047x |
| Combined selected natural corpora | 1048576 | 6.660x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 5.205x |
| ar-wikipedia | 4096 | 5.010x |
| ar-wikipedia | 65536 | 4.851x |
| ar-wikipedia | 1048576 | 5.311x |
| en-wikitext2 | 256 | 8.680x |
| en-wikitext2 | 4096 | 8.907x |
| en-wikitext2 | 65536 | 8.946x |
| en-wikitext2 | 1048576 | 9.859x |
| ja-wikipedia | 256 | 3.878x |
| ja-wikipedia | 4096 | 3.616x |
| ja-wikipedia | 65536 | 3.626x |
| ja-wikipedia | 1048576 | 3.931x |
| ru-wikipedia | 256 | 5.952x |
| ru-wikipedia | 4096 | 5.667x |
| ru-wikipedia | 65536 | 5.735x |
| ru-wikipedia | 1048576 | 6.295x |
| synthetic | 0 | 6.407x |
| zh-wikipedia | 256 | 4.836x |
| zh-wikipedia | 4096 | 4.602x |
| zh-wikipedia | 65536 | 4.626x |
| zh-wikipedia | 1048576 | 5.071x |
| Combined selected natural corpora | 256 | 5.715x |
| Combined selected natural corpora | 4096 | 5.557x |
| Combined selected natural corpora | 65536 | 5.555x |
| Combined selected natural corpora | 1048576 | 6.087x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 58911.9 | 172.5 | 10 |
| onig-compile | 58942.2 | 215.3 | 10 |
| pretok-compile | 11.9 | 0.1 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
