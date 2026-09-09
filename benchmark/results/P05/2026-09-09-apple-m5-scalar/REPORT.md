# P05 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 37503025 | 132113 | 9138.2 | 27.96 | 15587409 | 3387433 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-at | 20525572 | 122837 | 5001.4 | 51.09 | 28480327 | 6189304 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 22772600 | 40546 | 5548.9 | 46.05 | 25670103 | 5578590 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 21861129 | 66972 | 5326.8 | 47.97 | 26740385 | 5811182 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 21876703 | 47409 | 5330.6 | 47.93 | 26721348 | 5807045 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 6371117 | 44288 | 1552.4 | 164.58 | 91753921 | 19939830 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 7161976 | 68162 | 1745.1 | 146.41 | 81622025 | 17737981 | 5425056.0 | 25855.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 6390964 | 50903 | 1557.3 | 164.07 | 91468992 | 19877910 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 33328086 | 120440 | 129681.3 | 31.46 | 17540011 | 3712304 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-at | 20003324 | 110486 | 77833.9 | 52.42 | 29223893 | 6185172 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 21060397 | 72145 | 81947.1 | 49.79 | 27757075 | 5874723 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 20568788 | 42022 | 80034.2 | 50.98 | 28420489 | 6015133 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 20871241 | 114196 | 81211.1 | 50.24 | 28008636 | 5927966 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 6262188 | 40050 | 24366.5 | 167.45 | 93349968 | 19757313 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 6649270 | 95524 | 25872.6 | 157.70 | 87915665 | 18607155 | 5045488.0 | 2628.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 6309512 | 83446 | 24550.6 | 166.19 | 92649790 | 19609122 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 32548155 | 168975 | 1914597.3 | 32.22 | 17960311 | 3794624 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-at | 19885957 | 126771 | 1169762.2 | 52.73 | 29396373 | 6210815 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 20551556 | 52766 | 1208915.0 | 51.02 | 28444319 | 6009667 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 20360078 | 47735 | 1197651.7 | 51.50 | 28711825 | 6066185 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 20489558 | 93181 | 1205268.1 | 51.18 | 28530386 | 6027851 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 6229231 | 75058 | 366425.4 | 168.33 | 93843847 | 19827167 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 6574315 | 83009 | 386724.4 | 159.50 | 88918013 | 18786445 | 8081200.0 | 290.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 6279576 | 38406 | 369386.8 | 166.98 | 93091483 | 19668208 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 36570085 | 255289 | 36570084.8 | 28.67 | 15985060 | 3376886 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-at | 20104972 | 165900 | 20104972.5 | 52.16 | 29076140 | 6142411 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 20566878 | 123283 | 20566878.3 | 50.98 | 28423127 | 6004460 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 20502776 | 87407 | 20502775.8 | 51.14 | 28511993 | 6023233 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 20527029 | 136769 | 20527029.2 | 51.08 | 28478305 | 6016117 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 6207677 | 77596 | 6207677.5 | 168.92 | 94169680 | 19893592 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 6654464 | 69173 | 6654463.8 | 157.57 | 87847048 | 18557919 | 11302128.0 | 30.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 6236696 | 96231 | 6236695.9 | 168.13 | 93731523 | 19801030 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 52550142 | 688007 | 12826.5 | 19.95 | 19913115 | 4149123 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-at | 28361157 | 312719 | 6922.4 | 36.97 | 36896837 | 7687874 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 31484246 | 146766 | 7684.7 | 33.30 | 33236845 | 6925273 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 30149759 | 222371 | 7359.0 | 34.78 | 34707973 | 7231799 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 30374094 | 179472 | 7413.7 | 34.52 | 34451628 | 7178387 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 4263873 | 46970 | 1040.7 | 245.92 | 245419381 | 51135907 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 5264741 | 51672 | 1285.0 | 199.17 | 198763262 | 41414577 | 9046128.0 | 29028.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 4329112 | 33864 | 1056.7 | 242.22 | 241720951 | 50365298 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 48941178 | 290031 | 191176.5 | 21.43 | 21381525 | 4392967 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-at | 27966615 | 84500 | 109244.6 | 37.49 | 37417364 | 7687631 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 30177589 | 347817 | 117881.2 | 34.75 | 34675964 | 7124393 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 29252450 | 74567 | 114267.4 | 35.85 | 35772627 | 7349709 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 29889663 | 152085 | 116756.5 | 35.08 | 35009997 | 7193022 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 4177303 | 27862 | 16317.6 | 251.02 | 250505382 | 51467891 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 4727353 | 41668 | 18466.2 | 221.81 | 221357906 | 45479361 | 9977856.0 | 2929.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 4236038 | 25038 | 16547.0 | 247.54 | 247032027 | 50754269 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 47720877 | 214962 | 2982554.8 | 21.97 | 21928285 | 4501363 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-at | 27879542 | 87129 | 1742471.4 | 37.61 | 37534225 | 7704897 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 29007804 | 147686 | 1812987.8 | 36.15 | 36074326 | 7405214 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 28813677 | 165153 | 1800854.8 | 36.39 | 36317371 | 7455105 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 29005142 | 148370 | 1812821.4 | 36.15 | 36077637 | 7405894 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 4166746 | 21394 | 260421.6 | 251.65 | 251140121 | 51553183 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 4688683 | 33844 | 293042.7 | 223.64 | 223183585 | 45814361 | 14634752.0 | 320.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 4202016 | 19100 | 262626.0 | 249.54 | 249032103 | 51120456 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 53418819 | 2965346 | 53418818.8 | 19.63 | 19589295 | 4020961 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-at | 28189003 | 146451 | 28189003.4 | 37.20 | 37122171 | 7619815 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 29153304 | 78284 | 29153304.4 | 35.97 | 35894284 | 7367775 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 29146098 | 42463 | 29146098.4 | 35.98 | 35903159 | 7369597 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 29173119 | 95257 | 29173119.2 | 35.94 | 35869905 | 7362771 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 4237419 | 18928 | 4237418.9 | 247.46 | 246951513 | 50690056 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 4820680 | 21146 | 4820679.9 | 217.52 | 217072493 | 44556993 | 17962224.0 | 32.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 4283579 | 25721 | 4283578.7 | 244.79 | 244290364 | 50143820 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 24016049 | 129551 | 5839.1 | 43.66 | 15792773 | 2503243 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-at | 13673961 | 91899 | 3324.6 | 76.68 | 27737391 | 4396531 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 15136746 | 33323 | 3680.2 | 69.27 | 25056905 | 3971660 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 14604184 | 51885 | 3550.7 | 71.80 | 25970640 | 4116492 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 14621938 | 25298 | 3555.1 | 71.71 | 25939106 | 4111493 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 5413814 | 7932 | 1316.3 | 193.69 | 70057821 | 11104556 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 5980839 | 18774 | 1454.1 | 175.32 | 63415854 | 10051767 | 2600816.0 | 20602.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 5437838 | 16351 | 1322.1 | 192.83 | 69748310 | 11055497 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 20985870 | 208925 | 81657.1 | 49.97 | 18073113 | 2704010 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-at | 13113765 | 76195 | 51026.3 | 79.96 | 28922281 | 4327209 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 13668367 | 969106 | 53184.3 | 76.72 | 27748744 | 4151630 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 13401214 | 27461 | 52144.8 | 78.24 | 28301914 | 4234393 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 13467724 | 56975 | 52403.6 | 77.86 | 28162146 | 4213481 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 5286855 | 6413 | 20571.4 | 198.34 | 71740193 | 10733413 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 5531776 | 20689 | 21524.4 | 189.55 | 68563874 | 10258188 | 2591472.0 | 2354.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 5306851 | 11707 | 20649.2 | 197.59 | 71469881 | 10692971 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 20525833 | 72908 | 1207402.0 | 51.09 | 18478178 | 2754090 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-at | 12986156 | 63074 | 763891.5 | 80.75 | 29206487 | 4353097 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 13370578 | 44885 | 786504.6 | 78.42 | 28366762 | 4227940 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 13235370 | 39447 | 778551.1 | 79.23 | 28656548 | 4271131 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 13313477 | 46750 | 783145.7 | 78.76 | 28488426 | 4246073 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 5287836 | 6807 | 311049.2 | 198.30 | 71726882 | 10690573 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 5456369 | 14049 | 320962.9 | 192.17 | 69511429 | 10360370 | 3493648.0 | 249.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 5292475 | 4923 | 311322.0 | 198.13 | 71664016 | 10681203 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 23243081 | 1011111 | 23243081.5 | 45.11 | 16317974 | 2431433 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-at | 12988524 | 44711 | 12988524.1 | 80.73 | 29201162 | 4351072 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 13323444 | 105641 | 13323443.5 | 78.70 | 28467115 | 4241696 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 13258094 | 52172 | 13258094.3 | 79.09 | 28607430 | 4262604 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 13287342 | 54885 | 13287342.1 | 78.92 | 28544460 | 4253221 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 5284954 | 8015 | 5284953.5 | 198.41 | 71766005 | 10693377 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 5473211 | 14444 | 5473211.1 | 191.58 | 69297529 | 10325566 | 4363504.0 | 26.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 5297271 | 5296 | 5297271.5 | 197.95 | 71599124 | 10668511 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 33188823 | 208349 | 8086.9 | 31.59 | 17320801 | 3197221 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-at | 17506440 | 46894 | 4265.7 | 59.90 | 32836888 | 6061312 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 19633939 | 130042 | 4784.1 | 53.41 | 29278740 | 5404519 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 18735952 | 47221 | 4565.3 | 55.97 | 30682028 | 5663550 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 18865734 | 88310 | 4596.9 | 55.58 | 30470959 | 5624589 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 4711669 | 21978 | 1148.1 | 222.55 | 122007090 | 22521108 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 5422121 | 19377 | 1321.2 | 193.39 | 106020689 | 19570201 | 4684160.0 | 25105.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 4747952 | 9439 | 1156.9 | 220.85 | 121074738 | 22349006 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 29109274 | 160043 | 113265.7 | 36.02 | 19748242 | 3527192 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-at | 17033619 | 109733 | 66278.7 | 61.56 | 33748376 | 6027727 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 18037132 | 78073 | 70183.4 | 58.13 | 31870755 | 5692368 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 17527175 | 108646 | 68199.1 | 59.83 | 32798042 | 5857989 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 17641438 | 75894 | 68643.7 | 59.44 | 32585609 | 5820047 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 4612972 | 9834 | 17949.3 | 227.31 | 124617505 | 22257670 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 4940141 | 16508 | 19222.3 | 212.26 | 116364499 | 20783618 | 4299760.0 | 2574.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 4639450 | 4504 | 18052.3 | 226.01 | 123906291 | 22130642 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 28369061 | 78047 | 1668768.3 | 36.96 | 20263519 | 3611716 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-at | 16963400 | 60775 | 997847.1 | 61.81 | 33888077 | 6040122 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 17475304 | 38269 | 1027959.1 | 60.00 | 32895394 | 5863188 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 17366001 | 78843 | 1021529.5 | 60.38 | 33102440 | 5900092 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 17386871 | 59083 | 1022757.1 | 60.31 | 33062706 | 5893010 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 4599829 | 6498 | 270578.2 | 227.96 | 124973555 | 22274958 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 4885881 | 26816 | 287404.8 | 214.61 | 117656778 | 20970835 | 6549264.0 | 278.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 4631828 | 14186 | 272460.5 | 226.38 | 124110186 | 22121073 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 32305941 | 307508 | 32305941.4 | 32.46 | 17794157 | 3171150 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-at | 17030053 | 100228 | 17030052.8 | 61.57 | 33755444 | 6015660 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 17477351 | 78417 | 17477351.3 | 60.00 | 32891540 | 5861701 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 17446656 | 68404 | 17446655.9 | 60.10 | 32949409 | 5872014 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 17480021 | 125075 | 17480020.8 | 59.99 | 32886517 | 5860805 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 4594668 | 5624 | 4594668.0 | 228.22 | 125113936 | 22296932 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 4951850 | 14739 | 4951849.8 | 211.75 | 116089344 | 20688632 | 8942832.0 | 29.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 4626384 | 13025 | 4626384.2 | 226.65 | 124256216 | 22144075 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 9221809 | 53200 | 24015.1 | 29.30 | 18844893 | 4229431 | N/A | N/A | 10 |
| synthetic | 0 | onig-at | 5587006 | 43129 | 14549.5 | 48.37 | 31105031 | 6981020 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 5887680 | 37647 | 15332.5 | 45.90 | 29516550 | 6624511 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 5814584 | 18633 | 15142.1 | 46.47 | 29887608 | 6707789 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 5884496 | 25947 | 15324.2 | 45.92 | 29532522 | 6628095 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 1106369 | 17415 | 2881.2 | 244.24 | 157075972 | 35253154 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 1208584 | 7930 | 3147.4 | 223.59 | 143791470 | 32271663 | 2317088.0 | 1065.0 | 10 |
| synthetic | 0 | pretok-reused | 1104894 | 17360 | 2877.3 | 244.57 | 157285689 | 35300222 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 23859643 | 78009 | 5801.0 | 43.95 | 16354101 | 2750712 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-at | 12914693 | 89036 | 3140.0 | 81.19 | 30213880 | 5081886 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 14552429 | 61391 | 3538.2 | 72.05 | 26813599 | 4509969 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 13915906 | 82365 | 3383.4 | 75.35 | 28040071 | 4716258 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 13979711 | 67135 | 3398.9 | 75.01 | 27912093 | 4694732 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 4194199 | 23529 | 1019.7 | 250.01 | 93033975 | 15648042 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 4787822 | 30530 | 1164.1 | 219.01 | 81499061 | 13707903 | 2889584.0 | 21367.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 4218434 | 24430 | 1025.6 | 248.57 | 92499495 | 15558144 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 20845119 | 88157 | 81109.4 | 50.30 | 18719154 | 2992067 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-at | 12518619 | 176293 | 48710.6 | 83.76 | 31169813 | 4982179 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 13055297 | 18356 | 50798.8 | 80.32 | 29888481 | 4777371 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 12858405 | 76720 | 50032.7 | 81.55 | 30346142 | 4850524 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 12862431 | 53273 | 50048.4 | 81.52 | 30336646 | 4849006 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 4085298 | 10465 | 15896.1 | 256.67 | 95513952 | 15266938 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 4318226 | 32330 | 16802.4 | 242.83 | 90361877 | 14443431 | 2891760.0 | 2394.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 4094531 | 22606 | 15932.0 | 256.09 | 95298575 | 15232513 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 20329939 | 52580 | 1195878.8 | 51.58 | 19193515 | 3057412 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-at | 12309847 | 11145 | 724108.6 | 85.18 | 31698446 | 5049372 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 12709031 | 59067 | 747590.1 | 82.51 | 30702812 | 4890774 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 12591676 | 28497 | 740686.8 | 83.28 | 30988966 | 4936357 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 12642526 | 23631 | 743678.0 | 82.94 | 30864322 | 4916502 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 4077023 | 21976 | 239824.9 | 257.19 | 95707827 | 15245683 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 4272728 | 9215 | 251336.9 | 245.41 | 91324095 | 14547381 | 3764016.0 | 253.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 4100382 | 22765 | 241198.9 | 255.73 | 95162605 | 15158833 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 23184827 | 976590 | 23184826.7 | 45.23 | 16830102 | 2680460 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-at | 12444764 | 76183 | 12444763.7 | 84.26 | 31354794 | 4993747 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 12711293 | 21024 | 12711292.7 | 82.49 | 30697350 | 4889039 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 12676103 | 68069 | 12676103.1 | 82.72 | 30782567 | 4902611 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 12695434 | 37042 | 12695434.4 | 82.59 | 30735695 | 4895146 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 4065764 | 35101 | 4065764.3 | 257.90 | 95972854 | 15285195 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 4305896 | 36236 | 4305896.4 | 243.52 | 90620620 | 14432767 | 5559536.0 | 27.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 4088401 | 37332 | 4088401.1 | 256.48 | 95441468 | 15200563 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.431x |
| ar-wikipedia | 4096 | 3.285x |
| ar-wikipedia | 65536 | 3.268x |
| ar-wikipedia | 1048576 | 3.303x |
| en-wikitext2 | 256 | 7.071x |
| en-wikitext2 | 4096 | 7.003x |
| en-wikitext2 | 65536 | 6.915x |
| en-wikitext2 | 1048576 | 6.878x |
| ja-wikipedia | 256 | 2.698x |
| ja-wikipedia | 4096 | 2.535x |
| ja-wikipedia | 65536 | 2.503x |
| ja-wikipedia | 1048576 | 2.509x |
| ru-wikipedia | 256 | 3.977x |
| ru-wikipedia | 4096 | 3.800x |
| ru-wikipedia | 65536 | 3.775x |
| ru-wikipedia | 1048576 | 3.797x |
| synthetic | 0 | 5.256x |
| zh-wikipedia | 256 | 3.318x |
| zh-wikipedia | 4096 | 3.147x |
| zh-wikipedia | 65536 | 3.088x |
| zh-wikipedia | 1048576 | 3.118x |
| Combined selected natural corpora | 256 | 3.978x |
| Combined selected natural corpora | 4096 | 3.833x |
| Combined selected natural corpora | 65536 | 3.792x |
| Combined selected natural corpora | 1048576 | 3.814x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.423x |
| ar-wikipedia | 4096 | 3.308x |
| ar-wikipedia | 65536 | 3.263x |
| ar-wikipedia | 1048576 | 3.291x |
| en-wikitext2 | 256 | 7.016x |
| en-wikitext2 | 4096 | 7.056x |
| en-wikitext2 | 65536 | 6.903x |
| en-wikitext2 | 1048576 | 6.810x |
| ja-wikipedia | 256 | 2.689x |
| ja-wikipedia | 4096 | 2.538x |
| ja-wikipedia | 65536 | 2.516x |
| ja-wikipedia | 1048576 | 2.508x |
| ru-wikipedia | 256 | 3.973x |
| ru-wikipedia | 4096 | 3.802x |
| ru-wikipedia | 65536 | 3.754x |
| ru-wikipedia | 1048576 | 3.778x |
| synthetic | 0 | 5.326x |
| zh-wikipedia | 256 | 3.314x |
| zh-wikipedia | 4096 | 3.141x |
| zh-wikipedia | 65536 | 3.083x |
| zh-wikipedia | 1048576 | 3.105x |
| Combined selected natural corpora | 256 | 3.969x |
| Combined selected natural corpora | 4096 | 3.853x |
| Combined selected natural corpora | 65536 | 3.788x |
| Combined selected natural corpora | 1048576 | 3.798x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.180x |
| ar-wikipedia | 4096 | 3.167x |
| ar-wikipedia | 65536 | 3.126x |
| ar-wikipedia | 1048576 | 3.091x |
| en-wikitext2 | 256 | 5.980x |
| en-wikitext2 | 4096 | 6.384x |
| en-wikitext2 | 65536 | 6.187x |
| en-wikitext2 | 1048576 | 6.048x |
| ja-wikipedia | 256 | 2.531x |
| ja-wikipedia | 4096 | 2.471x |
| ja-wikipedia | 65536 | 2.450x |
| ja-wikipedia | 1048576 | 2.434x |
| ru-wikipedia | 256 | 3.621x |
| ru-wikipedia | 4096 | 3.651x |
| ru-wikipedia | 65536 | 3.577x |
| ru-wikipedia | 1048576 | 3.529x |
| synthetic | 0 | 4.872x |
| zh-wikipedia | 256 | 3.039x |
| zh-wikipedia | 4096 | 3.023x |
| zh-wikipedia | 65536 | 2.974x |
| zh-wikipedia | 1048576 | 2.952x |
| Combined selected natural corpora | 256 | 3.619x |
| Combined selected natural corpora | 4096 | 3.669x |
| Combined selected natural corpora | 65536 | 3.598x |
| Combined selected natural corpora | 1048576 | 3.558x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 5.868x |
| ar-wikipedia | 4096 | 5.282x |
| ar-wikipedia | 65536 | 5.183x |
| ar-wikipedia | 1048576 | 5.864x |
| en-wikitext2 | 256 | 12.139x |
| en-wikitext2 | 4096 | 11.554x |
| en-wikitext2 | 65536 | 11.357x |
| en-wikitext2 | 1048576 | 12.471x |
| ja-wikipedia | 256 | 4.416x |
| ja-wikipedia | 4096 | 3.954x |
| ja-wikipedia | 65536 | 3.878x |
| ja-wikipedia | 1048576 | 4.388x |
| ru-wikipedia | 256 | 6.990x |
| ru-wikipedia | 4096 | 6.274x |
| ru-wikipedia | 65536 | 6.125x |
| ru-wikipedia | 1048576 | 6.983x |
| synthetic | 0 | 8.346x |
| zh-wikipedia | 256 | 5.656x |
| zh-wikipedia | 4096 | 5.091x |
| zh-wikipedia | 65536 | 4.958x |
| zh-wikipedia | 1048576 | 5.671x |
| Combined selected natural corpora | 256 | 6.811x |
| Combined selected natural corpora | 4096 | 6.231x |
| Combined selected natural corpora | 65536 | 6.100x |
| Combined selected natural corpora | 1048576 | 6.878x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 5.236x |
| ar-wikipedia | 4096 | 5.012x |
| ar-wikipedia | 65536 | 4.951x |
| ar-wikipedia | 1048576 | 5.496x |
| en-wikitext2 | 256 | 9.982x |
| en-wikitext2 | 4096 | 10.353x |
| en-wikitext2 | 65536 | 10.178x |
| en-wikitext2 | 1048576 | 11.081x |
| ja-wikipedia | 256 | 4.015x |
| ja-wikipedia | 4096 | 3.794x |
| ja-wikipedia | 65536 | 3.762x |
| ja-wikipedia | 1048576 | 4.247x |
| ru-wikipedia | 256 | 6.121x |
| ru-wikipedia | 4096 | 5.892x |
| ru-wikipedia | 65536 | 5.806x |
| ru-wikipedia | 1048576 | 6.524x |
| synthetic | 0 | 7.630x |
| zh-wikipedia | 256 | 4.983x |
| zh-wikipedia | 4096 | 4.827x |
| zh-wikipedia | 65536 | 4.758x |
| zh-wikipedia | 1048576 | 5.384x |
| Combined selected natural corpora | 256 | 5.979x |
| Combined selected natural corpora | 4096 | 5.855x |
| Combined selected natural corpora | 65536 | 5.777x |
| Combined selected natural corpora | 1048576 | 6.438x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 58480.8 | 53.0 | 10 |
| onig-compile | 58618.0 | 181.5 | 10 |
| pretok-compile | 9.4 | 0.0 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
