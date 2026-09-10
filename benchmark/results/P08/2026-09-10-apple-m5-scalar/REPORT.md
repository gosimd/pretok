# P08 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 14122922 | 196412 | 3441.3 | 74.25 | 41391929 | 405015 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 10340938 | 107148 | 2519.7 | 101.40 | 56530170 | 553141 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 10277778 | 101798 | 2504.3 | 102.02 | 56877566 | 88346 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 10296159 | 138931 | 2508.8 | 101.84 | 56776024 | 555547 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 6163216 | 78817 | 1501.8 | 170.13 | 94849025 | 147326 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 6222139 | 65464 | 1516.1 | 168.52 | 93950804 | 919298 | 141824.0 | 5040.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 6158658 | 46561 | 1500.6 | 170.26 | 94919213 | 928774 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 12379752 | 89139 | 48170.2 | 84.70 | 47220253 | 151780 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 9784937 | 39263 | 38073.7 | 107.16 | 59742338 | 192030 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 9758466 | 26591 | 37970.7 | 107.45 | 59904393 | 92125 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 9736915 | 29402 | 37886.8 | 107.69 | 60036985 | 192977 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 6055956 | 24101 | 23564.0 | 173.15 | 96528945 | 148449 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 6114318 | 22498 | 23791.1 | 171.50 | 95607555 | 307311 | 79312.0 | 657.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 6078921 | 13590 | 23653.4 | 172.49 | 96164271 | 309101 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 12313888 | 46709 | 724346.3 | 85.15 | 47472822 | 133102 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 9755795 | 57507 | 573870.3 | 107.48 | 59920797 | 168003 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 9759786 | 111908 | 574105.0 | 107.44 | 59896295 | 92010 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 9766831 | 67145 | 574519.5 | 107.36 | 59853086 | 167813 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 6054858 | 23842 | 356168.1 | 173.18 | 96546444 | 148311 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 6110225 | 19640 | 359425.0 | 171.61 | 95671602 | 268239 | 73136.0 | 113.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 6074592 | 22809 | 357328.9 | 172.62 | 96232799 | 269812 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 13621568 | 297240 | 13621567.6 | 76.98 | 42915398 | 119149 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 9587244 | 108861 | 9587244.5 | 109.37 | 60974246 | 169287 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 9666398 | 94072 | 9666398.2 | 108.48 | 60474955 | 92899 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 9627082 | 135154 | 9627082.1 | 108.92 | 60721929 | 168587 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 6005102 | 44082 | 6005102.2 | 174.61 | 97346387 | 149540 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 6029219 | 61496 | 6029219.4 | 173.92 | 96956996 | 269189 | 79088.0 | 13.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 5999271 | 72760 | 5999271.1 | 174.78 | 97441004 | 270533 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 40822896 | 1012640 | 9964.1 | 25.69 | 25633581 | 4697315 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 21333313 | 354292 | 5207.1 | 49.15 | 49051780 | 8988665 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 20213391 | 472970 | 4933.7 | 51.88 | 51769493 | 8070195 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 20518120 | 326514 | 5008.1 | 51.10 | 51000628 | 9345788 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 3513808 | 34080 | 857.7 | 298.42 | 297807110 | 46424278 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 4697037 | 115531 | 1146.5 | 223.24 | 222786611 | 40825310 | 8323184.0 | 28673.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 3575653 | 49307 | 872.7 | 293.25 | 292656222 | 53628811 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 37628025 | 549847 | 146984.5 | 27.87 | 27810043 | 5015676 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 20146466 | 166740 | 78697.1 | 52.05 | 51941468 | 9367896 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 19761829 | 111437 | 77194.6 | 53.06 | 52952436 | 8140238 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 19889558 | 84307 | 77693.6 | 52.72 | 52612380 | 9488898 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 3475922 | 31692 | 13577.8 | 301.67 | 301053087 | 46280097 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 4240375 | 46027 | 16564.0 | 247.28 | 246779333 | 44507852 | 7663616.0 | 2816.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 3542517 | 20729 | 13838.0 | 296.00 | 295393616 | 53275675 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 37003921 | 1020845 | 2312745.1 | 28.34 | 28279084 | 5095190 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 20336827 | 1294497 | 1271051.7 | 51.56 | 51455275 | 9270965 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 19804713 | 478274 | 1237794.6 | 52.95 | 52837776 | 8115644 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 20205963 | 1361553 | 1262872.7 | 51.89 | 51788523 | 9331008 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 3437636 | 46760 | 214852.2 | 305.03 | 304405994 | 46755387 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 4409057 | 213243 | 275566.0 | 237.82 | 237338071 | 42762435 | 14634752.0 | 320.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 3563645 | 266188 | 222727.8 | 294.24 | 293642292 | 52907060 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 39254311 | 695865 | 39254310.9 | 26.71 | 26657887 | 4802734 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 18524303 | 5444870 | 18524302.8 | 56.61 | 56489953 | 10177333 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 18544856 | 900173 | 18544856.4 | 56.54 | 56427345 | 8666392 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 19097360 | 4045383 | 19097360.1 | 54.91 | 54794851 | 9871940 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 3294099 | 418663 | 3294098.5 | 318.32 | 317670219 | 48789372 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 4347509 | 1452746 | 4347508.9 | 241.19 | 240698070 | 43364604 | 17962224.0 | 32.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 3345341 | 151296 | 3345340.6 | 313.44 | 312804326 | 56355398 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 10589803 | 185584 | 2574.7 | 99.02 | 35815585 | 943549 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 6990193 | 62450 | 1699.5 | 150.01 | 54258876 | 1429431 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 6867035 | 79548 | 1669.6 | 152.70 | 55231988 | 496138 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 6873310 | 17758 | 1671.1 | 152.56 | 55181563 | 1453739 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 3722332 | 17397 | 905.0 | 281.70 | 101893116 | 915286 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 3840827 | 41515 | 933.8 | 273.01 | 98749575 | 2601523 | 339088.0 | 6768.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 3749214 | 28198 | 911.6 | 279.68 | 101162543 | 2665092 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 9027342 | 90057 | 35125.8 | 116.16 | 42014582 | 682925 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 6263692 | 22439 | 24372.3 | 167.41 | 60552145 | 984244 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 6255447 | 22912 | 24340.3 | 167.63 | 60631954 | 538251 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 6210441 | 55898 | 24165.1 | 168.84 | 61071344 | 992683 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 3665007 | 26422 | 14260.7 | 286.10 | 103486829 | 918688 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 3708593 | 12607 | 14430.3 | 282.74 | 102270599 | 1662356 | 268112.0 | 1121.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 3663915 | 14268 | 14256.5 | 286.19 | 103517690 | 1682626 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 9062952 | 162167 | 533114.8 | 115.70 | 41849499 | 654312 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 6253100 | 56740 | 367829.4 | 167.69 | 60654715 | 948330 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 6228304 | 114115 | 366370.8 | 168.36 | 60896197 | 539954 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 6219665 | 125498 | 365862.6 | 168.59 | 60980778 | 953428 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 3668684 | 24316 | 215805.0 | 285.82 | 103383116 | 916677 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 3706344 | 19937 | 218020.2 | 282.91 | 102332662 | 1599960 | 258832.0 | 155.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 3671451 | 22040 | 215967.7 | 285.60 | 103305205 | 1615165 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 10741504 | 252175 | 10741504.4 | 97.62 | 35309765 | 550575 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 6210214 | 62132 | 6210214.0 | 168.85 | 61073580 | 952302 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 6247600 | 99147 | 6247599.8 | 167.84 | 60708114 | 538287 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 6228181 | 39826 | 6228180.8 | 168.36 | 60897397 | 949555 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 3690391 | 19928 | 3690390.6 | 284.14 | 102775028 | 911286 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 3715938 | 19479 | 3715937.5 | 282.18 | 102068455 | 1591523 | 365808.0 | 17.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 3692438 | 40589 | 3692438.1 | 283.98 | 102718039 | 1601652 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 24919562 | 867456 | 6072.0 | 42.08 | 23068503 | 3756968 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 13171299 | 434461 | 3209.4 | 79.61 | 43644671 | 7108031 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 12637863 | 1092656 | 3079.4 | 82.97 | 45486883 | 5959947 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 12409567 | 457158 | 3023.8 | 84.50 | 46323697 | 7544341 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 4513192 | 95612 | 1099.7 | 232.34 | 127372611 | 16689076 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 5230480 | 214082 | 1274.5 | 200.47 | 109905204 | 17899312 | 4168576.0 | 24614.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 4546837 | 774394 | 1107.9 | 230.62 | 126430076 | 20590576 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 22147627 | 194611 | 86177.5 | 47.34 | 25955693 | 4074342 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 12011331 | 291539 | 46736.7 | 87.30 | 47859557 | 7512656 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 11518669 | 216972 | 44819.7 | 91.03 | 49906547 | 6289008 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 11715044 | 209754 | 45583.8 | 89.51 | 49069984 | 7702660 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 4499782 | 30082 | 17508.9 | 233.03 | 127752179 | 16098779 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 4999534 | 85495 | 19453.4 | 209.73 | 114982123 | 18049083 | 4190704.0 | 2565.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 4548478 | 94074 | 17698.4 | 230.53 | 126384468 | 19838943 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 21825529 | 294324 | 1283854.6 | 48.04 | 26338743 | 4124528 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 11644050 | 183519 | 684944.1 | 90.05 | 49369161 | 7730987 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 11529804 | 111407 | 678223.8 | 90.94 | 49858348 | 6267062 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 11604880 | 156688 | 682640.0 | 90.36 | 49535800 | 7757082 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 4547556 | 32598 | 267503.3 | 230.58 | 126410092 | 15889413 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 5033917 | 40387 | 296112.8 | 208.30 | 114196760 | 17882695 | 5852944.0 | 273.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 4631428 | 23999 | 272436.9 | 226.40 | 124120909 | 19436772 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 26767068 | 220520 | 26767068.0 | 39.17 | 21476278 | 3362602 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 11986969 | 101782 | 11986968.8 | 87.48 | 47956828 | 7508737 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 11845600 | 152698 | 11845600.3 | 88.52 | 48529157 | 6099142 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 11906770 | 101888 | 11906769.7 | 88.07 | 48279845 | 7559313 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 4603535 | 77602 | 4603534.8 | 227.78 | 124872957 | 15694027 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 5090359 | 57519 | 5090358.8 | 205.99 | 112930546 | 17681858 | 7058672.0 | 28.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 4686515 | 45632 | 4686515.0 | 223.74 | 122661936 | 19205529 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 72189655 | 1057301 | 990.6 | 31.69 | 17264468 | 2742623 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 37445918 | 355664 | 513.9 | 61.10 | 33283094 | 5287332 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 35211544 | 212516 | 483.2 | 64.97 | 35395096 | 2446357 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 35923149 | 439207 | 493.0 | 63.69 | 34693952 | 5511460 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 10430053 | 53877 | 143.1 | 219.35 | 119492783 | 8258827 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 12845586 | 105016 | 176.3 | 178.10 | 97022900 | 15412999 | 6888768.0 | 142146.0 | 10 |
| synthetic | 0 | pretok-reused | 10656608 | 48624 | 146.2 | 214.69 | 116952409 | 18578988 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 11183045 | 65325 | 2719.0 | 93.76 | 34892374 | 1268349 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 7234097 | 34499 | 1758.8 | 144.95 | 53939423 | 1960715 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 7048988 | 38309 | 1713.8 | 148.76 | 55355890 | 880694 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 7107018 | 50078 | 1727.9 | 147.54 | 54903902 | 1995774 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 3818479 | 5372 | 928.4 | 274.61 | 102188057 | 1625778 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 3975916 | 14901 | 966.7 | 263.73 | 98141665 | 3567480 | 530992.0 | 7666.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 3824728 | 10977 | 929.9 | 274.16 | 102021118 | 3708499 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 9610452 | 26716 | 37394.8 | 109.11 | 40601942 | 1080698 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 6504131 | 66342 | 25307.9 | 161.22 | 59993106 | 1596831 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 6466589 | 48120 | 25161.8 | 162.15 | 60341395 | 947486 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 6463160 | 42543 | 25148.5 | 162.24 | 60373411 | 1606954 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 3748275 | 15829 | 14584.7 | 279.75 | 104102019 | 1634619 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 3827178 | 14386 | 14891.7 | 273.98 | 101955794 | 2713749 | 487952.0 | 1144.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 3759684 | 17645 | 14629.1 | 278.90 | 103786114 | 2762466 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 9519823 | 64522 | 559989.6 | 110.15 | 40988472 | 1066091 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 6454814 | 59879 | 379694.9 | 162.45 | 60451472 | 1572315 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 6399572 | 104258 | 376445.4 | 163.85 | 60973294 | 957095 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 6400528 | 49088 | 376501.6 | 163.83 | 60964192 | 1585651 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 3738832 | 46565 | 219931.3 | 280.45 | 104364943 | 1638212 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 3803126 | 18437 | 223713.3 | 275.71 | 102600608 | 2668594 | 482576.0 | 166.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 3744709 | 35574 | 220277.0 | 280.01 | 104201147 | 2710224 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 11310449 | 166144 | 11310449.1 | 92.71 | 34499337 | 895897 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 6487091 | 89565 | 6487091.1 | 161.64 | 60150689 | 1562025 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 6483504 | 33659 | 6483504.0 | 161.73 | 60183968 | 944705 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 6483200 | 53199 | 6483199.8 | 161.74 | 60186792 | 1562963 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 3780007 | 8876 | 3780006.6 | 277.40 | 103228127 | 1620368 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 3850244 | 29798 | 3850243.8 | 272.34 | 101345012 | 2631781 | 685296.0 | 19.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 3780054 | 25803 | 3780053.7 | 277.40 | 103226840 | 2680650 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.668x |
| ar-wikipedia | 4096 | 1.611x |
| ar-wikipedia | 65536 | 1.612x |
| ar-wikipedia | 1048576 | 1.610x |
| en-wikitext2 | 256 | 5.753x |
| en-wikitext2 | 4096 | 5.685x |
| en-wikitext2 | 65536 | 5.761x |
| en-wikitext2 | 1048576 | 5.630x |
| ja-wikipedia | 256 | 1.845x |
| ja-wikipedia | 4096 | 1.707x |
| ja-wikipedia | 65536 | 1.698x |
| ja-wikipedia | 1048576 | 1.693x |
| ru-wikipedia | 256 | 2.800x |
| ru-wikipedia | 4096 | 2.560x |
| ru-wikipedia | 65536 | 2.535x |
| ru-wikipedia | 1048576 | 2.573x |
| synthetic | 0 | 3.376x |
| zh-wikipedia | 256 | 1.846x |
| zh-wikipedia | 4096 | 1.725x |
| zh-wikipedia | 65536 | 1.712x |
| zh-wikipedia | 1048576 | 1.715x |
| Combined selected natural corpora | 256 | 2.625x |
| Combined selected natural corpora | 4096 | 2.507x |
| Combined selected natural corpora | 65536 | 2.505x |
| Combined selected natural corpora | 1048576 | 2.470x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.672x |
| ar-wikipedia | 4096 | 1.602x |
| ar-wikipedia | 65536 | 1.608x |
| ar-wikipedia | 1048576 | 1.605x |
| en-wikitext2 | 256 | 5.738x |
| en-wikitext2 | 4096 | 5.615x |
| en-wikitext2 | 65536 | 5.670x |
| en-wikitext2 | 1048576 | 5.709x |
| ja-wikipedia | 256 | 1.833x |
| ja-wikipedia | 4096 | 1.695x |
| ja-wikipedia | 65536 | 1.694x |
| ja-wikipedia | 1048576 | 1.687x |
| ru-wikipedia | 256 | 2.729x |
| ru-wikipedia | 4096 | 2.576x |
| ru-wikipedia | 65536 | 2.506x |
| ru-wikipedia | 1048576 | 2.541x |
| synthetic | 0 | 3.371x |
| zh-wikipedia | 256 | 1.858x |
| zh-wikipedia | 4096 | 1.719x |
| zh-wikipedia | 65536 | 1.709x |
| zh-wikipedia | 1048576 | 1.715x |
| Combined selected natural corpora | 256 | 2.617x |
| Combined selected natural corpora | 4096 | 2.501x |
| Combined selected natural corpora | 65536 | 2.499x |
| Combined selected natural corpora | 1048576 | 2.481x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 1.662x |
| ar-wikipedia | 4096 | 1.600x |
| ar-wikipedia | 65536 | 1.597x |
| ar-wikipedia | 1048576 | 1.590x |
| en-wikitext2 | 256 | 4.542x |
| en-wikitext2 | 4096 | 4.751x |
| en-wikitext2 | 65536 | 4.613x |
| en-wikitext2 | 1048576 | 4.261x |
| ja-wikipedia | 256 | 1.820x |
| ja-wikipedia | 4096 | 1.689x |
| ja-wikipedia | 65536 | 1.687x |
| ja-wikipedia | 1048576 | 1.671x |
| ru-wikipedia | 256 | 2.518x |
| ru-wikipedia | 4096 | 2.402x |
| ru-wikipedia | 65536 | 2.313x |
| ru-wikipedia | 1048576 | 2.355x |
| synthetic | 0 | 2.915x |
| zh-wikipedia | 256 | 1.819x |
| zh-wikipedia | 4096 | 1.699x |
| zh-wikipedia | 65536 | 1.697x |
| zh-wikipedia | 1048576 | 1.685x |
| Combined selected natural corpora | 256 | 2.465x |
| Combined selected natural corpora | 4096 | 2.390x |
| Combined selected natural corpora | 65536 | 2.361x |
| Combined selected natural corpora | 1048576 | 2.292x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.293x |
| ar-wikipedia | 4096 | 2.037x |
| ar-wikipedia | 65536 | 2.027x |
| ar-wikipedia | 1048576 | 2.271x |
| en-wikitext2 | 256 | 11.417x |
| en-wikitext2 | 4096 | 10.622x |
| en-wikitext2 | 65536 | 10.384x |
| en-wikitext2 | 1048576 | 11.734x |
| ja-wikipedia | 256 | 2.825x |
| ja-wikipedia | 4096 | 2.464x |
| ja-wikipedia | 65536 | 2.468x |
| ja-wikipedia | 1048576 | 2.909x |
| ru-wikipedia | 256 | 5.481x |
| ru-wikipedia | 4096 | 4.869x |
| ru-wikipedia | 65536 | 4.712x |
| ru-wikipedia | 1048576 | 5.712x |
| synthetic | 0 | 6.774x |
| zh-wikipedia | 256 | 2.924x |
| zh-wikipedia | 4096 | 2.556x |
| zh-wikipedia | 65536 | 2.542x |
| zh-wikipedia | 1048576 | 2.992x |
| Combined selected natural corpora | 256 | 4.651x |
| Combined selected natural corpora | 4096 | 4.205x |
| Combined selected natural corpora | 65536 | 4.138x |
| Combined selected natural corpora | 1048576 | 4.729x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.270x |
| ar-wikipedia | 4096 | 2.025x |
| ar-wikipedia | 65536 | 2.015x |
| ar-wikipedia | 1048576 | 2.259x |
| en-wikitext2 | 256 | 8.691x |
| en-wikitext2 | 4096 | 8.874x |
| en-wikitext2 | 65536 | 8.393x |
| en-wikitext2 | 1048576 | 9.029x |
| ja-wikipedia | 256 | 2.757x |
| ja-wikipedia | 4096 | 2.434x |
| ja-wikipedia | 65536 | 2.445x |
| ja-wikipedia | 1048576 | 2.891x |
| ru-wikipedia | 256 | 4.764x |
| ru-wikipedia | 4096 | 4.430x |
| ru-wikipedia | 65536 | 4.336x |
| ru-wikipedia | 1048576 | 5.258x |
| synthetic | 0 | 5.620x |
| zh-wikipedia | 256 | 2.813x |
| zh-wikipedia | 4096 | 2.511x |
| zh-wikipedia | 65536 | 2.503x |
| zh-wikipedia | 1048576 | 2.938x |
| Combined selected natural corpora | 256 | 4.241x |
| Combined selected natural corpora | 4096 | 3.967x |
| Combined selected natural corpora | 65536 | 3.891x |
| Combined selected natural corpora | 1048576 | 4.415x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 2626.4 | 10.5 | 10 |
| onig-compile | 2577.3 | 12.8 | 10 |
| pretok-compile | 16.5 | 0.1 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
