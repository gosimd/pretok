# P19 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 38729869 | 574475 | 9437.1 | 27.07 | 15093648 | 3153613 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-at | 21270253 | 124503 | 5182.8 | 49.30 | 27483219 | 5742245 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 23306247 | 140704 | 5678.9 | 44.99 | 25082331 | 5240612 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 22524569 | 238725 | 5488.4 | 46.55 | 25952772 | 5422479 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 22606816 | 234606 | 5508.5 | 46.38 | 25858352 | 5402751 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 7082385 | 22654 | 1725.7 | 148.05 | 82539286 | 17245462 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 8056577 | 38988 | 1963.1 | 130.15 | 72558731 | 15160161 | 5149600.0 | 25592.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 7118017 | 12561 | 1734.4 | 147.31 | 82126103 | 17159133 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 33931675 | 131602 | 132029.9 | 30.90 | 17228003 | 3500859 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-at | 20543482 | 138943 | 79935.7 | 51.04 | 28455498 | 5782369 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 21574127 | 118757 | 83946.0 | 48.60 | 27096114 | 5506132 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 21115705 | 268000 | 82162.3 | 49.66 | 27684371 | 5625671 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 21283027 | 148025 | 82813.3 | 49.27 | 27466723 | 5581443 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 6967710 | 24685 | 27111.7 | 150.49 | 83897718 | 17048642 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 7466931 | 29639 | 29054.2 | 140.43 | 78288527 | 15908813 | 4719856.0 | 2604.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 6990167 | 14850 | 27199.1 | 150.01 | 83628185 | 16993871 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 33158966 | 404965 | 1950527.4 | 31.62 | 17629470 | 3575926 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-at | 20411789 | 124718 | 1200693.5 | 51.37 | 28639087 | 5809094 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 20966050 | 336481 | 1233297.1 | 50.01 | 27881980 | 5655524 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 20910865 | 172662 | 1230050.9 | 50.15 | 27955562 | 5670449 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 21092184 | 279542 | 1240716.7 | 49.71 | 27715243 | 5621703 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 6962722 | 28145 | 409571.9 | 150.60 | 83957827 | 17029834 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 7484398 | 39203 | 440258.7 | 140.10 | 78105817 | 15842825 | 8081200.0 | 290.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 6985684 | 15710 | 410922.6 | 150.10 | 83681852 | 16973856 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 37016732 | 320013 | 37016732.1 | 28.33 | 15792183 | 3202849 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-at | 20542801 | 103909 | 20542800.8 | 51.04 | 28456441 | 5771316 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 20982612 | 196919 | 20982612.0 | 49.97 | 27859973 | 5650345 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 20997395 | 122123 | 20997395.0 | 49.94 | 27840358 | 5646367 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 21046940 | 147788 | 21046940.1 | 49.82 | 27774821 | 5633075 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 6966825 | 28609 | 6966824.9 | 150.51 | 83908381 | 17017652 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 7531340 | 8633 | 7531339.9 | 139.23 | 77618991 | 15742086 | 11302128.0 | 30.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 6991364 | 33525 | 6991364.0 | 149.98 | 83613870 | 16957921 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 53890938 | 745831 | 13153.8 | 19.46 | 19417680 | 4048807 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-at | 28556051 | 151425 | 6970.0 | 36.72 | 36645018 | 7640902 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 31310311 | 110513 | 7642.3 | 33.49 | 33421482 | 6968759 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 30294249 | 178365 | 7394.3 | 34.61 | 34542431 | 7202489 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 30255173 | 306538 | 7384.7 | 34.66 | 34587044 | 7211792 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 4922312 | 19662 | 1201.4 | 213.03 | 212590525 | 44327539 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 6224761 | 61386 | 1519.3 | 168.45 | 168108782 | 35052590 | 9048176.0 | 29029.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 4989350 | 37530 | 1217.8 | 210.16 | 209734149 | 43731952 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 49867595 | 680667 | 194795.3 | 21.03 | 20984309 | 4314525 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-at | 28055439 | 176459 | 109591.6 | 37.38 | 37298900 | 7668923 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 29849195 | 327733 | 116598.4 | 35.13 | 35057461 | 7208067 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 29094792 | 243437 | 113651.5 | 36.04 | 35966472 | 7394966 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 29205204 | 203584 | 114082.8 | 35.90 | 35830498 | 7367009 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 4884092 | 25766 | 19078.5 | 214.69 | 214254170 | 44052204 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 5728907 | 23346 | 22378.5 | 183.03 | 182659109 | 37556031 | 9977856.0 | 2929.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 4931278 | 17849 | 19262.8 | 212.64 | 212204000 | 43630674 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 49708866 | 709258 | 3106804.1 | 21.09 | 21051315 | 4324520 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-at | 28001087 | 152523 | 1750067.9 | 37.45 | 37371299 | 7677095 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 29018289 | 151584 | 1813643.1 | 36.14 | 36061292 | 7407983 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 28930291 | 156506 | 1808143.2 | 36.24 | 36170981 | 7430516 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 29051036 | 132986 | 1815689.7 | 36.09 | 36020643 | 7399633 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 4884803 | 16715 | 305300.2 | 214.66 | 214222951 | 44007298 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 5740146 | 34874 | 358759.1 | 182.67 | 182301473 | 37449747 | 14634752.0 | 320.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 4943746 | 23413 | 308984.1 | 212.10 | 211668856 | 43482617 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 53565721 | 1045488 | 53565720.8 | 19.58 | 19535572 | 4012884 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-at | 28165961 | 251494 | 28165960.7 | 37.23 | 37152541 | 7631659 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 29202409 | 231838 | 29202408.6 | 35.91 | 35833928 | 7360797 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 29039412 | 180756 | 29039412.1 | 36.11 | 36035062 | 7402113 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 29251751 | 321944 | 29251751.1 | 35.85 | 35773482 | 7348381 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 4885576 | 12839 | 4885575.8 | 214.63 | 214189081 | 43997475 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 5719390 | 54324 | 5719389.7 | 183.34 | 182963053 | 37583206 | 17962224.0 | 32.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 4970502 | 14398 | 4970501.9 | 210.96 | 210529446 | 43245734 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 25676383 | 152564 | 6242.7 | 40.84 | 14771551 | 2346086 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-at | 14944297 | 86066 | 3633.4 | 70.17 | 25379581 | 4030902 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 16357876 | 145745 | 3977.1 | 64.10 | 23186384 | 3682568 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 15839536 | 114889 | 3851.1 | 66.20 | 23945145 | 3803078 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 15908470 | 166942 | 3867.9 | 65.91 | 23841388 | 3786599 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 6405006 | 25758 | 1557.3 | 163.71 | 59216185 | 9404988 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 7078969 | 18716 | 1721.1 | 148.13 | 53578422 | 8509572 | 2607344.0 | 20615.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 6396187 | 18967 | 1555.1 | 163.94 | 59297826 | 9417954 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 22452069 | 198648 | 87362.1 | 46.70 | 16892875 | 2532996 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-at | 14325977 | 104777 | 55743.1 | 73.19 | 26474982 | 3969782 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 14923078 | 148202 | 58066.5 | 70.27 | 25415668 | 3810943 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 14620135 | 68844 | 56887.7 | 71.72 | 25942305 | 3889909 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 14709533 | 90402 | 57235.5 | 71.29 | 25784639 | 3866268 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 6243359 | 10477 | 24293.2 | 167.95 | 60749349 | 9109039 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 6614884 | 15939 | 25738.8 | 158.52 | 57337359 | 8597429 | 2599664.0 | 2355.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 6276318 | 33180 | 24421.5 | 167.07 | 60430335 | 9061204 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 22007675 | 210881 | 1294569.1 | 47.65 | 17233988 | 2574238 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-at | 14223622 | 92013 | 836683.6 | 73.72 | 26665501 | 3983022 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 14523327 | 74399 | 854313.4 | 72.20 | 26115228 | 3900828 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 14422426 | 83293 | 848378.0 | 72.70 | 26297934 | 3928119 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 14429992 | 56203 | 848823.1 | 72.67 | 26284144 | 3926059 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 6255192 | 32441 | 367952.5 | 167.63 | 60634430 | 9056956 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 6488486 | 23248 | 381675.6 | 161.61 | 58454318 | 8731313 | 3575568.0 | 250.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 6246993 | 20630 | 367470.2 | 167.85 | 60714013 | 9068843 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 24672790 | 1219039 | 24672789.7 | 42.50 | 15372400 | 2295525 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-at | 14216882 | 68269 | 14216881.9 | 73.76 | 26678142 | 3983785 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 14471062 | 127373 | 14471062.5 | 72.46 | 26209548 | 3913811 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 14465526 | 85405 | 14465526.2 | 72.49 | 26219579 | 3915309 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 14430214 | 128033 | 14430214.3 | 72.67 | 26283740 | 3924890 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 6210486 | 9897 | 6210486.1 | 168.84 | 61070904 | 9119576 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 6501947 | 11947 | 6501947.0 | 161.27 | 58333296 | 8710775 | 4363504.0 | 26.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 6232895 | 13406 | 6232895.1 | 168.23 | 60851337 | 9086789 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 35362993 | 206016 | 8616.7 | 29.65 | 16255892 | 2998049 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-at | 19154980 | 161532 | 4667.4 | 54.74 | 30010838 | 5534853 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 21125015 | 165306 | 5147.4 | 49.64 | 27212147 | 5018695 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 20313398 | 143910 | 4949.7 | 51.62 | 28299402 | 5219216 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 20334622 | 209032 | 4954.8 | 51.57 | 28269865 | 5213768 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 7406057 | 20998 | 1804.6 | 141.58 | 77619846 | 14315310 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 8302302 | 20961 | 2023.0 | 126.30 | 69240676 | 12769952 | 4676480.0 | 25097.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 7419984 | 13908 | 1808.0 | 141.32 | 77474159 | 14288441 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 31056969 | 275447 | 120844.2 | 33.76 | 18509758 | 3302930 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-at | 18599350 | 176133 | 72371.0 | 56.38 | 30907371 | 5515193 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 19443056 | 116584 | 75653.9 | 53.93 | 29566185 | 5275868 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 19119873 | 237189 | 74396.4 | 54.84 | 30065942 | 5365046 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 19174873 | 191328 | 74610.4 | 54.68 | 29979704 | 5349657 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 7278127 | 15360 | 28319.6 | 144.07 | 78984197 | 14094148 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 7718711 | 14093 | 30033.9 | 135.85 | 74475778 | 13289654 | 4299760.0 | 2574.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 7299710 | 24360 | 28403.5 | 143.65 | 78750668 | 14052477 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 30436303 | 323215 | 1790370.7 | 34.45 | 18887215 | 3363286 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-at | 18577610 | 199684 | 1092800.6 | 56.44 | 30943538 | 5510181 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 18946319 | 170368 | 1114489.3 | 55.34 | 30341356 | 5402949 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 18932042 | 97780 | 1113649.6 | 55.39 | 30364236 | 5407024 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 18979954 | 280153 | 1116467.9 | 55.25 | 30287586 | 5393374 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 7259658 | 31005 | 427038.7 | 144.44 | 79185136 | 14100664 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 7685900 | 45247 | 452111.8 | 136.43 | 74793714 | 13318675 | 6549264.0 | 278.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 7282843 | 18121 | 428402.5 | 143.98 | 78933052 | 14055775 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 34235099 | 303760 | 34235098.6 | 30.63 | 16791452 | 2989680 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-at | 18567828 | 117323 | 18567827.9 | 56.47 | 30959841 | 5512330 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 19116785 | 128639 | 19116785.5 | 54.85 | 30070798 | 5354038 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 19043692 | 153054 | 19043692.1 | 55.06 | 30186216 | 5374588 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 19045373 | 201121 | 19045372.7 | 55.06 | 30183552 | 5374114 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 7229406 | 28712 | 7229406.0 | 145.04 | 79516492 | 14157733 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 7740765 | 39252 | 7740764.7 | 135.46 | 74263593 | 13222466 | 8942832.0 | 29.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 7269126 | 11606 | 7269125.6 | 144.25 | 79082001 | 14080373 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 141994573 | 1084034 | 3053.9 | 15.04 | 8051498 | 3329782 | N/A | N/A | 10 |
| synthetic | 0 | onig-at | 73886006 | 430313 | 1589.1 | 28.90 | 15473417 | 6399196 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 85587587 | 745469 | 1840.8 | 24.94 | 13357883 | 5524294 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 83568642 | 521355 | 1797.3 | 25.55 | 13680598 | 5657756 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 83520979 | 298684 | 1796.3 | 25.56 | 13688405 | 5660985 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 20991175 | 71036 | 451.5 | 101.71 | 54464268 | 22524275 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 24565942 | 42513 | 528.3 | 86.91 | 46538780 | 19246605 | 25663408.0 | 128205.0 | 10 |
| synthetic | 0 | pretok-reused | 21187341 | 72146 | 455.7 | 100.77 | 53960003 | 22315731 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 25424162 | 235394 | 6181.4 | 41.24 | 15347723 | 2602210 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-at | 13970781 | 79081 | 3396.7 | 75.05 | 27929935 | 4735526 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 15555179 | 44777 | 3782.0 | 67.41 | 25085085 | 4253181 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 14968831 | 111691 | 3639.4 | 70.05 | 26067699 | 4419784 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 14961418 | 94818 | 3637.6 | 70.09 | 26080616 | 4421974 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 5280327 | 5415 | 1283.8 | 198.58 | 73897510 | 12529338 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 6010518 | 14297 | 1461.3 | 174.46 | 64920027 | 11007204 | 2913520.0 | 21416.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 5308960 | 13509 | 1290.8 | 197.51 | 73498948 | 12461762 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 22144098 | 100702 | 86163.8 | 47.35 | 17621084 | 2840712 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-at | 13519503 | 107467 | 52605.1 | 77.56 | 28862230 | 4652908 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 14049653 | 46278 | 54667.9 | 74.63 | 27773142 | 4477335 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 13907609 | 123463 | 54115.2 | 75.40 | 28056800 | 4523064 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 13920860 | 89424 | 54166.8 | 75.32 | 28030093 | 4518758 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 5195056 | 10224 | 20214.2 | 201.84 | 75110457 | 12108629 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 5522759 | 10576 | 21489.3 | 189.86 | 70653636 | 11390140 | 2904048.0 | 2396.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 5213847 | 12235 | 20287.3 | 201.11 | 74839743 | 12064987 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 21561722 | 162324 | 1268336.6 | 48.63 | 18097024 | 2907560 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-at | 13437938 | 62354 | 790466.9 | 78.03 | 29037416 | 4665299 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 13787684 | 71383 | 811040.2 | 76.05 | 28300837 | 4546957 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 13766234 | 141132 | 809778.5 | 76.17 | 28344934 | 4554041 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 13686572 | 129626 | 805092.5 | 76.61 | 28509915 | 4580548 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 5178105 | 12501 | 304594.4 | 202.50 | 75356335 | 12107132 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 5463250 | 12743 | 321367.7 | 191.93 | 71423234 | 11475220 | 3764016.0 | 253.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 5198468 | 13835 | 305792.2 | 201.71 | 75061153 | 12059707 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 24646208 | 432833 | 24646208.3 | 42.55 | 15832172 | 2543231 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-at | 13529357 | 127956 | 13529357.4 | 77.50 | 28841207 | 4632962 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 13753847 | 101091 | 13753847.4 | 76.24 | 28370462 | 4557343 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 13787506 | 218157 | 13787506.2 | 76.05 | 28301202 | 4546217 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 13771226 | 140872 | 13771225.8 | 76.14 | 28334660 | 4551592 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 5208908 | 37436 | 5208907.9 | 201.30 | 74910712 | 12033424 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 5523130 | 32025 | 5523130.3 | 189.85 | 70648886 | 11348818 | 5559536.0 | 27.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 5204285 | 12642 | 5204285.0 | 201.48 | 74977254 | 12044114 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.180x |
| ar-wikipedia | 4096 | 3.031x |
| ar-wikipedia | 65536 | 3.003x |
| ar-wikipedia | 1048576 | 3.014x |
| en-wikitext2 | 256 | 6.154x |
| en-wikitext2 | 4096 | 5.957x |
| en-wikitext2 | 65536 | 5.923x |
| en-wikitext2 | 1048576 | 5.944x |
| ja-wikipedia | 256 | 2.473x |
| ja-wikipedia | 4096 | 2.342x |
| ja-wikipedia | 65536 | 2.306x |
| ja-wikipedia | 1048576 | 2.329x |
| ru-wikipedia | 256 | 2.743x |
| ru-wikipedia | 4096 | 2.627x |
| ru-wikipedia | 65536 | 2.608x |
| ru-wikipedia | 1048576 | 2.634x |
| synthetic | 0 | 3.981x |
| zh-wikipedia | 256 | 2.835x |
| zh-wikipedia | 4096 | 2.677x |
| zh-wikipedia | 65536 | 2.659x |
| zh-wikipedia | 1048576 | 2.647x |
| Combined selected natural corpora | 256 | 3.343x |
| Combined selected natural corpora | 4096 | 3.201x |
| Combined selected natural corpora | 65536 | 3.175x |
| Combined selected natural corpora | 1048576 | 3.191x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.176x |
| ar-wikipedia | 4096 | 3.045x |
| ar-wikipedia | 65536 | 3.019x |
| ar-wikipedia | 1048576 | 3.010x |
| en-wikitext2 | 256 | 6.064x |
| en-wikitext2 | 4096 | 5.922x |
| en-wikitext2 | 65536 | 5.876x |
| en-wikitext2 | 1048576 | 5.885x |
| ja-wikipedia | 256 | 2.487x |
| ja-wikipedia | 4096 | 2.344x |
| ja-wikipedia | 65536 | 2.310x |
| ja-wikipedia | 1048576 | 2.315x |
| ru-wikipedia | 256 | 2.741x |
| ru-wikipedia | 4096 | 2.627x |
| ru-wikipedia | 65536 | 2.606x |
| ru-wikipedia | 1048576 | 2.620x |
| synthetic | 0 | 3.942x |
| zh-wikipedia | 256 | 2.818x |
| zh-wikipedia | 4096 | 2.670x |
| zh-wikipedia | 65536 | 2.633x |
| zh-wikipedia | 1048576 | 2.646x |
| Combined selected natural corpora | 256 | 3.332x |
| Combined selected natural corpora | 4096 | 3.201x |
| Combined selected natural corpora | 65536 | 3.172x |
| Combined selected natural corpora | 1048576 | 3.181x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.893x |
| ar-wikipedia | 4096 | 2.889x |
| ar-wikipedia | 65536 | 2.801x |
| ar-wikipedia | 1048576 | 2.786x |
| en-wikitext2 | 256 | 5.030x |
| en-wikitext2 | 4096 | 5.210x |
| en-wikitext2 | 65536 | 5.055x |
| en-wikitext2 | 1048576 | 5.106x |
| ja-wikipedia | 256 | 2.311x |
| ja-wikipedia | 4096 | 2.256x |
| ja-wikipedia | 65536 | 2.238x |
| ja-wikipedia | 1048576 | 2.226x |
| ru-wikipedia | 256 | 2.544x |
| ru-wikipedia | 4096 | 2.519x |
| ru-wikipedia | 65536 | 2.465x |
| ru-wikipedia | 1048576 | 2.470x |
| synthetic | 0 | 3.484x |
| zh-wikipedia | 256 | 2.588x |
| zh-wikipedia | 4096 | 2.544x |
| zh-wikipedia | 65536 | 2.524x |
| zh-wikipedia | 1048576 | 2.490x |
| Combined selected natural corpora | 256 | 3.018x |
| Combined selected natural corpora | 4096 | 3.021x |
| Combined selected natural corpora | 65536 | 2.959x |
| Combined selected natural corpora | 1048576 | 2.954x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 5.441x |
| ar-wikipedia | 4096 | 4.854x |
| ar-wikipedia | 65536 | 4.747x |
| ar-wikipedia | 1048576 | 5.295x |
| en-wikitext2 | 256 | 10.801x |
| en-wikitext2 | 4096 | 10.113x |
| en-wikitext2 | 65536 | 10.055x |
| en-wikitext2 | 1048576 | 10.777x |
| ja-wikipedia | 256 | 4.014x |
| ja-wikipedia | 4096 | 3.577x |
| ja-wikipedia | 65536 | 3.523x |
| ja-wikipedia | 1048576 | 3.958x |
| ru-wikipedia | 256 | 4.766x |
| ru-wikipedia | 4096 | 4.255x |
| ru-wikipedia | 65536 | 4.179x |
| ru-wikipedia | 1048576 | 4.710x |
| synthetic | 0 | 6.702x |
| zh-wikipedia | 256 | 4.789x |
| zh-wikipedia | 4096 | 4.247x |
| zh-wikipedia | 65536 | 4.148x |
| zh-wikipedia | 1048576 | 4.736x |
| Combined selected natural corpora | 256 | 5.734x |
| Combined selected natural corpora | 4096 | 5.192x |
| Combined selected natural corpora | 65536 | 5.117x |
| Combined selected natural corpora | 1048576 | 5.678x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 4.807x |
| ar-wikipedia | 4096 | 4.544x |
| ar-wikipedia | 65536 | 4.430x |
| ar-wikipedia | 1048576 | 4.915x |
| en-wikitext2 | 256 | 8.658x |
| en-wikitext2 | 4096 | 8.705x |
| en-wikitext2 | 65536 | 8.660x |
| en-wikitext2 | 1048576 | 9.366x |
| ja-wikipedia | 256 | 3.627x |
| ja-wikipedia | 4096 | 3.394x |
| ja-wikipedia | 65536 | 3.392x |
| ja-wikipedia | 1048576 | 3.795x |
| ru-wikipedia | 256 | 4.259x |
| ru-wikipedia | 4096 | 4.024x |
| ru-wikipedia | 65536 | 3.960x |
| ru-wikipedia | 1048576 | 4.423x |
| synthetic | 0 | 5.780x |
| zh-wikipedia | 256 | 4.230x |
| zh-wikipedia | 4096 | 4.010x |
| zh-wikipedia | 65536 | 3.947x |
| zh-wikipedia | 1048576 | 4.462x |
| Combined selected natural corpora | 256 | 5.020x |
| Combined selected natural corpora | 4096 | 4.824x |
| Combined selected natural corpora | 65536 | 4.774x |
| Combined selected natural corpora | 1048576 | 5.274x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 314053.7 | 1124.5 | 10 |
| onig-compile | 314289.2 | 1129.8 | 10 |
| pretok-compile | 12.6 | 0.0 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
