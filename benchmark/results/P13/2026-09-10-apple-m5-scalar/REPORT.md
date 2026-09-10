# P13 benchmark

Native Oniguruma and HF Split versus the public scalar pretok API. All request boundaries passed exact oracle comparison before timing. UTF-8 validation is inside each public operation; native Rust allocation counters are unavailable.

Environment and source fingerprints: [environment.json](environment.json). Raw samples: [samples.jsonl](samples.jsonl).

## Per-workload measurements

One batch processes every request for a corpus/size once. MB/s uses decimal units. Allocation columns are per batch; Go allocations were measured separately from timing. IQR is across sample batch means, not request latency.

| Corpus | Request bytes | Engine | ns/batch | IQR ns | ns/request | MB/s | scalars/s | spans/s | B/batch | allocs/batch | samples |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| ar-wikipedia | 256 | hf-split | 36647857 | 419533 | 8929.8 | 28.61 | 15951137 | 3466478 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-at | 19207355 | 229346 | 4680.2 | 54.59 | 30434956 | 6614081 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-fresh | 21070735 | 311471 | 5134.2 | 49.76 | 27743455 | 6029168 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-iter | 20374379 | 172159 | 4964.5 | 51.47 | 28691672 | 6235233 | N/A | N/A | 10 |
| ar-wikipedia | 256 | onig-reused | 20289872 | 216050 | 4943.9 | 51.68 | 28811173 | 6261203 | N/A | N/A | 10 |
| ar-wikipedia | 256 | pretok-cursor | 6383208 | 132833 | 1555.4 | 164.27 | 91580134 | 19902063 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 256 | pretok-fresh | 7493340 | 46916 | 1825.9 | 139.93 | 78012613 | 16953589 | 5425056.0 | 25855.0 | 10 |
| ar-wikipedia | 256 | pretok-reused | 6424709 | 192577 | 1565.5 | 163.21 | 90988559 | 19773503 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | hf-split | 32204286 | 160990 | 125308.5 | 32.56 | 18152087 | 3841849 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-at | 18679676 | 152585 | 72683.6 | 56.13 | 31294708 | 6623455 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-fresh | 19441349 | 170339 | 75647.3 | 53.94 | 30068645 | 6363962 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-iter | 19044111 | 117817 | 74101.6 | 55.06 | 30695841 | 6496706 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | onig-reused | 19097046 | 70076 | 74307.6 | 54.91 | 30610755 | 6478698 | N/A | N/A | 10 |
| ar-wikipedia | 4096 | pretok-cursor | 6258091 | 23291 | 24350.5 | 167.56 | 93411072 | 19770246 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 4096 | pretok-fresh | 6851400 | 74637 | 26659.1 | 153.05 | 85321983 | 18058208 | 5045488.0 | 2628.0 | 10 |
| ar-wikipedia | 4096 | pretok-reused | 6323593 | 50973 | 24605.4 | 165.82 | 92443482 | 19565458 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | hf-split | 31351094 | 475285 | 1844182.0 | 33.45 | 18646080 | 3939512 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-at | 18474118 | 258014 | 1086712.8 | 56.76 | 31642919 | 6685461 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-fresh | 18889764 | 160708 | 1111162.6 | 55.51 | 30946655 | 6538356 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-iter | 18884110 | 102934 | 1110830.0 | 55.53 | 30955920 | 6540313 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | onig-reused | 18929342 | 150771 | 1113490.7 | 55.39 | 30881951 | 6524685 | N/A | N/A | 10 |
| ar-wikipedia | 65536 | pretok-cursor | 6245788 | 44272 | 367399.3 | 167.89 | 93595069 | 19774605 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 65536 | pretok-fresh | 6882028 | 22337 | 404825.2 | 152.36 | 84942259 | 17946454 | 8081200.0 | 290.0 | 10 |
| ar-wikipedia | 65536 | pretok-reused | 6294070 | 25168 | 370239.4 | 166.60 | 92877112 | 19622916 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | hf-split | 35269739 | 962284 | 35269738.9 | 29.73 | 16574407 | 3501387 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-at | 18347693 | 375280 | 18347693.4 | 57.15 | 31860953 | 6730710 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-fresh | 18756688 | 208679 | 18756687.5 | 55.90 | 31166217 | 6583945 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-iter | 18675547 | 248895 | 18675547.1 | 56.15 | 31301627 | 6612551 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | onig-reused | 18724239 | 119679 | 18724239.2 | 56.00 | 31220227 | 6595355 | N/A | N/A | 10 |
| ar-wikipedia | 1048576 | pretok-cursor | 6166909 | 50211 | 6166908.8 | 170.03 | 94792224 | 20025106 | 0.0 | 0.0 | 10 |
| ar-wikipedia | 1048576 | pretok-fresh | 6753476 | 98782 | 6753476.1 | 155.26 | 86559128 | 18285842 | 11302128.0 | 30.0 | 10 |
| ar-wikipedia | 1048576 | pretok-reused | 6205660 | 83312 | 6205660.2 | 168.97 | 94200291 | 19900058 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | hf-split | 50282058 | 1064769 | 12272.9 | 20.85 | 20811340 | 4336278 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-at | 25746620 | 188498 | 6284.3 | 40.73 | 40643665 | 8468568 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-fresh | 28040670 | 341112 | 6844.2 | 37.39 | 37318545 | 7775741 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-iter | 27153536 | 348556 | 6627.7 | 38.62 | 38537780 | 8029783 | N/A | N/A | 10 |
| en-wikitext2 | 256 | onig-reused | 27168727 | 171626 | 6631.4 | 38.59 | 38516232 | 8025293 | N/A | N/A | 10 |
| en-wikitext2 | 256 | pretok-cursor | 4772948 | 48462 | 1165.0 | 219.69 | 219243345 | 45681834 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 256 | pretok-fresh | 6147221 | 210669 | 1500.4 | 170.58 | 170229278 | 35469198 | 9046128.0 | 29028.0 | 10 |
| en-wikitext2 | 256 | pretok-reused | 4831561 | 21152 | 1179.3 | 217.03 | 216583613 | 45127649 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | hf-split | 46359117 | 1187567 | 181090.3 | 22.62 | 22572410 | 4637642 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-at | 25631181 | 286091 | 100121.8 | 40.91 | 40826718 | 8388103 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-fresh | 26820763 | 591229 | 104768.6 | 39.10 | 39015929 | 8016066 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-iter | 26317680 | 434698 | 102803.4 | 39.84 | 39761749 | 8169299 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | onig-reused | 26579344 | 269686 | 103825.6 | 39.45 | 39370309 | 8088875 | N/A | N/A | 10 |
| en-wikitext2 | 4096 | pretok-cursor | 4719349 | 67198 | 18435.0 | 222.19 | 221733335 | 45556495 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 4096 | pretok-fresh | 5609978 | 64764 | 21914.0 | 186.91 | 186531383 | 38324035 | 9977856.0 | 2929.0 | 10 |
| en-wikitext2 | 4096 | pretok-reused | 4774621 | 81551 | 18650.9 | 219.61 | 219166504 | 45029124 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | hf-split | 46356460 | 517583 | 2897278.8 | 22.62 | 22573704 | 4633853 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-at | 25759833 | 227761 | 1609989.6 | 40.71 | 40622817 | 8338913 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-fresh | 26192890 | 153755 | 1637055.6 | 40.03 | 39951186 | 8201042 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-iter | 26149047 | 179053 | 1634315.4 | 40.10 | 40018170 | 8214793 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | onig-reused | 26097627 | 172151 | 1631101.7 | 40.18 | 40097017 | 8230978 | N/A | N/A | 10 |
| en-wikitext2 | 65536 | pretok-cursor | 4719940 | 14924 | 294996.3 | 222.16 | 221705551 | 45510956 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 65536 | pretok-fresh | 5657484 | 25860 | 353592.8 | 185.34 | 184965081 | 37968998 | 14634752.0 | 320.0 | 10 |
| en-wikitext2 | 65536 | pretok-reused | 4798890 | 14818 | 299930.6 | 218.50 | 218058140 | 44762227 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | hf-split | 53035896 | 612762 | 53035895.9 | 19.77 | 19730731 | 4049993 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-at | 25677488 | 207850 | 25677487.5 | 40.84 | 40753092 | 8365110 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-fresh | 26319298 | 133966 | 26319298.2 | 39.84 | 39759305 | 8161122 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-iter | 26360621 | 158183 | 26360620.6 | 39.78 | 39696979 | 8148329 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | onig-reused | 26364487 | 256846 | 26364486.8 | 39.77 | 39691158 | 8147134 | N/A | N/A | 10 |
| en-wikitext2 | 1048576 | pretok-cursor | 4755115 | 17928 | 4755115.0 | 220.52 | 220065552 | 45171358 | 0.0 | 0.0 | 10 |
| en-wikitext2 | 1048576 | pretok-fresh | 5631873 | 64531 | 5631872.7 | 186.19 | 185806225 | 38139179 | 17962224.0 | 32.0 | 10 |
| en-wikitext2 | 1048576 | pretok-reused | 4810966 | 15549 | 4810966.3 | 217.96 | 217510771 | 44646955 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | hf-split | 23186657 | 28392 | 5637.4 | 45.22 | 16357684 | 2592784 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-at | 12792105 | 97984 | 3110.2 | 81.97 | 29649538 | 4699618 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-fresh | 14126560 | 76472 | 3434.6 | 74.23 | 26848717 | 4255672 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-iter | 13644935 | 66641 | 3317.5 | 76.85 | 27796394 | 4405884 | N/A | N/A | 10 |
| ja-wikipedia | 256 | onig-reused | 13673320 | 52288 | 3324.4 | 76.69 | 27738690 | 4396737 | N/A | N/A | 10 |
| ja-wikipedia | 256 | pretok-cursor | 5376901 | 15961 | 1307.3 | 195.01 | 70538779 | 11180791 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 256 | pretok-fresh | 6049333 | 22432 | 1470.8 | 173.34 | 62697823 | 9937955 | 2600816.0 | 20602.0 | 10 |
| ja-wikipedia | 256 | pretok-reused | 5410039 | 21448 | 1315.4 | 193.82 | 70106701 | 11112304 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | hf-split | 20148438 | 106795 | 78398.6 | 52.04 | 18824288 | 2816397 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-at | 12268116 | 142155 | 47735.9 | 85.47 | 30915912 | 4625486 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-fresh | 12730974 | 55771 | 49536.9 | 82.36 | 29791908 | 4457318 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-iter | 12560669 | 62598 | 48874.2 | 83.48 | 30195844 | 4517753 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | onig-reused | 12576939 | 42813 | 48937.5 | 83.37 | 30156781 | 4511909 | N/A | N/A | 10 |
| ja-wikipedia | 4096 | pretok-cursor | 5270780 | 28639 | 20508.9 | 198.94 | 71958994 | 10766149 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 4096 | pretok-fresh | 5592127 | 30435 | 21759.2 | 187.51 | 67823924 | 10147480 | 2591472.0 | 2354.0 | 10 |
| ja-wikipedia | 4096 | pretok-reused | 5299296 | 20505 | 20619.8 | 197.87 | 71571764 | 10708214 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | hf-split | 19782291 | 402088 | 1163664.2 | 53.01 | 19172704 | 2857606 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-at | 12449959 | 192654 | 732350.5 | 84.22 | 30464358 | 4540577 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-fresh | 12629483 | 360365 | 742910.8 | 83.03 | 30031316 | 4476034 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-iter | 12552211 | 654011 | 738365.3 | 83.54 | 30216191 | 4503589 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | onig-reused | 12555132 | 261007 | 738537.2 | 83.52 | 30209161 | 4502541 | N/A | N/A | 10 |
| ja-wikipedia | 65536 | pretok-cursor | 5307993 | 43552 | 312234.9 | 197.55 | 71454503 | 10649976 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 65536 | pretok-fresh | 5590357 | 80514 | 328844.5 | 187.57 | 67845400 | 10112056 | 3493648.0 | 249.0 | 10 |
| ja-wikipedia | 65536 | pretok-reused | 5356675 | 64892 | 315098.5 | 195.75 | 70805116 | 10553188 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | hf-split | 22123648 | 322712 | 22123647.7 | 47.40 | 17143647 | 2554461 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-at | 12369072 | 245703 | 12369071.6 | 84.77 | 30663579 | 4568977 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-fresh | 12637567 | 593657 | 12637567.2 | 82.97 | 30012106 | 4471905 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-iter | 12590266 | 204139 | 12590265.6 | 83.28 | 30124861 | 4488706 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | onig-reused | 12554230 | 99127 | 12554229.7 | 83.52 | 30211332 | 4501590 | N/A | N/A | 10 |
| ja-wikipedia | 1048576 | pretok-cursor | 5339341 | 87432 | 5339341.1 | 196.39 | 71034982 | 10584452 | 0.0 | 0.0 | 10 |
| ja-wikipedia | 1048576 | pretok-fresh | 5649486 | 36070 | 5649486.2 | 185.61 | 67135309 | 10003388 | 4363504.0 | 26.0 | 10 |
| ja-wikipedia | 1048576 | pretok-reused | 5369436 | 51584 | 5369436.2 | 195.29 | 70636839 | 10525127 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | hf-split | 31820824 | 728557 | 7753.6 | 32.95 | 18065434 | 3334672 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-at | 16020212 | 468145 | 3903.6 | 65.45 | 35883234 | 6623633 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-fresh | 17782565 | 371169 | 4333.0 | 58.97 | 32327002 | 5967193 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-iter | 17254544 | 384374 | 4204.3 | 60.77 | 33316268 | 6149800 | N/A | N/A | 10 |
| ru-wikipedia | 256 | onig-reused | 17184065 | 365298 | 4187.2 | 61.02 | 33452911 | 6175023 | N/A | N/A | 10 |
| ru-wikipedia | 256 | pretok-cursor | 4780314 | 54256 | 1164.8 | 219.35 | 120255079 | 22197706 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 256 | pretok-fresh | 5650538 | 206753 | 1376.8 | 185.57 | 101734907 | 18779095 | 4684160.0 | 25105.0 | 10 |
| ru-wikipedia | 256 | pretok-reused | 4827576 | 20031 | 1176.3 | 217.21 | 119077765 | 21980388 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | hf-split | 27630889 | 115465 | 107513.2 | 37.95 | 20804868 | 3715914 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-at | 15482881 | 89330 | 60244.7 | 67.72 | 37128554 | 6631453 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-fresh | 16207899 | 84065 | 63065.8 | 64.70 | 35467706 | 6334812 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-iter | 15908771 | 41725 | 61901.8 | 65.91 | 36134594 | 6453924 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | onig-reused | 15987588 | 64887 | 62208.5 | 65.59 | 35956456 | 6422107 | N/A | N/A | 10 |
| ru-wikipedia | 4096 | pretok-cursor | 4643509 | 38619 | 18068.1 | 225.82 | 123797981 | 22111297 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 4096 | pretok-fresh | 5082566 | 15602 | 19776.5 | 206.31 | 113103704 | 20201215 | 4299760.0 | 2574.0 | 10 |
| ru-wikipedia | 4096 | pretok-reused | 4684405 | 35585 | 18227.3 | 223.84 | 122717197 | 21918260 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | hf-split | 26786742 | 173690 | 1575690.7 | 39.15 | 21460504 | 3825064 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-at | 15479796 | 78075 | 910576.2 | 67.74 | 37135954 | 6619015 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-fresh | 15806537 | 55826 | 929796.3 | 66.34 | 36368307 | 6482191 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-iter | 15724545 | 40591 | 924973.2 | 66.68 | 36557942 | 6515991 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | onig-reused | 15774196 | 36872 | 927893.9 | 66.47 | 36442872 | 6495482 | N/A | N/A | 10 |
| ru-wikipedia | 65536 | pretok-cursor | 4629019 | 28751 | 272295.3 | 226.52 | 124185479 | 22134493 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 65536 | pretok-fresh | 5058214 | 37295 | 297542.0 | 207.30 | 113648221 | 20256360 | 6549264.0 | 278.0 | 10 |
| ru-wikipedia | 65536 | pretok-reused | 4678488 | 38389 | 275205.2 | 224.13 | 122872391 | 21900452 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | hf-split | 30597555 | 253149 | 30597555.1 | 34.27 | 18787678 | 3348209 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-at | 15465843 | 72693 | 15465842.8 | 67.80 | 37169458 | 6624081 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-fresh | 15818238 | 107225 | 15818238.3 | 66.29 | 36341405 | 6476511 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-iter | 15770964 | 100271 | 15770964.2 | 66.49 | 36450340 | 6495925 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | onig-reused | 15821904 | 79210 | 15821904.3 | 66.27 | 36332984 | 6475011 | N/A | N/A | 10 |
| ru-wikipedia | 1048576 | pretok-cursor | 4595507 | 43650 | 4595507.1 | 228.17 | 125091092 | 22292861 | 0.0 | 0.0 | 10 |
| ru-wikipedia | 1048576 | pretok-fresh | 5108101 | 33091 | 5108100.8 | 205.28 | 112538305 | 20055791 | 8942832.0 | 29.0 | 10 |
| ru-wikipedia | 1048576 | pretok-reused | 4659044 | 49288 | 4659043.8 | 225.06 | 123385189 | 21988847 | 0.0 | 0.0 | 10 |
| synthetic | 0 | hf-split | 75342461 | 664176 | 1730.9 | 13.94 | 8200741 | 3129895 | N/A | N/A | 10 |
| synthetic | 0 | onig-at | 33931361 | 233113 | 779.5 | 30.96 | 18209231 | 6949736 | N/A | N/A | 10 |
| synthetic | 0 | onig-fresh | 43148615 | 214244 | 991.3 | 24.35 | 14319440 | 5465158 | N/A | N/A | 10 |
| synthetic | 0 | onig-iter | 41610337 | 623501 | 956.0 | 25.25 | 14848810 | 5667198 | N/A | N/A | 10 |
| synthetic | 0 | onig-reused | 41614421 | 379360 | 956.1 | 25.24 | 14847353 | 5666641 | N/A | N/A | 10 |
| synthetic | 0 | pretok-cursor | 9634265 | 32669 | 221.3 | 109.03 | 64131928 | 24476594 | 0.0 | 0.0 | 10 |
| synthetic | 0 | pretok-fresh | 12188633 | 118057 | 280.0 | 86.18 | 50691820 | 19347042 | 10138304.0 | 117364.0 | 10 |
| synthetic | 0 | pretok-reused | 9769171 | 32641 | 224.4 | 107.53 | 63246305 | 24138587 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | hf-split | 23152895 | 138543 | 5629.2 | 45.29 | 16853314 | 2834678 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-at | 11990334 | 65329 | 2915.2 | 87.45 | 32543129 | 5473659 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-fresh | 13501697 | 74414 | 3282.7 | 77.66 | 28900293 | 4860945 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-iter | 12904150 | 47231 | 3137.4 | 81.26 | 30238568 | 5086038 | N/A | N/A | 10 |
| zh-wikipedia | 256 | onig-reused | 12950072 | 41739 | 3148.6 | 80.97 | 30131339 | 5068003 | N/A | N/A | 10 |
| zh-wikipedia | 256 | pretok-cursor | 4148471 | 12807 | 1008.6 | 252.76 | 94059468 | 15820527 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 256 | pretok-fresh | 4871802 | 18315 | 1184.5 | 215.23 | 80094190 | 13471608 | 2889584.0 | 21367.0 | 10 |
| zh-wikipedia | 256 | pretok-reused | 4173961 | 8536 | 1014.8 | 251.22 | 93485055 | 15723912 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | hf-split | 20005036 | 136804 | 77840.6 | 52.42 | 19505238 | 3117715 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-at | 11563089 | 84213 | 44992.6 | 90.68 | 33745567 | 5393887 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-fresh | 12057681 | 66440 | 46917.0 | 86.96 | 32361365 | 5172637 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-iter | 11886439 | 37493 | 46250.7 | 88.22 | 32827578 | 5247156 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | onig-reused | 11879830 | 43456 | 46225.0 | 88.27 | 32845839 | 5250075 | N/A | N/A | 10 |
| zh-wikipedia | 4096 | pretok-cursor | 4048668 | 16342 | 15753.6 | 258.99 | 96378113 | 15405066 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 4096 | pretok-fresh | 4400517 | 23334 | 17122.6 | 238.28 | 88672074 | 14173333 | 2891760.0 | 2394.0 | 10 |
| zh-wikipedia | 4096 | pretok-reused | 4074403 | 21352 | 15853.7 | 257.36 | 95769369 | 15307764 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | hf-split | 19732942 | 111748 | 1160761.3 | 53.14 | 19774193 | 3149911 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-at | 11599250 | 131002 | 682308.8 | 90.40 | 33640366 | 5358709 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-fresh | 11871048 | 138122 | 698297.0 | 88.33 | 32870138 | 5236016 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-iter | 11788811 | 55323 | 693459.4 | 88.95 | 33099438 | 5272542 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | onig-reused | 11819474 | 131667 | 695263.2 | 88.72 | 33013566 | 5258863 | N/A | N/A | 10 |
| zh-wikipedia | 65536 | pretok-cursor | 4083767 | 43646 | 240221.6 | 256.77 | 95549776 | 15220507 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 65536 | pretok-fresh | 4422427 | 35333 | 260142.8 | 237.10 | 88232765 | 14054951 | 3764016.0 | 253.0 | 10 |
| zh-wikipedia | 65536 | pretok-reused | 4122972 | 29080 | 242527.8 | 254.32 | 94641200 | 15075776 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | hf-split | 23205723 | 395624 | 23205723.5 | 45.19 | 16814947 | 2678046 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-at | 11757231 | 121858 | 11757230.6 | 89.19 | 33188343 | 5285769 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-fresh | 11886936 | 45426 | 11886935.6 | 88.21 | 32826206 | 5228093 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-iter | 11938041 | 159033 | 11938040.7 | 87.83 | 32685682 | 5205712 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | onig-reused | 11898406 | 143396 | 11898405.9 | 88.13 | 32794561 | 5223053 | N/A | N/A | 10 |
| zh-wikipedia | 1048576 | pretok-cursor | 4121746 | 15003 | 4121746.1 | 254.40 | 94669345 | 15077591 | 0.0 | 0.0 | 10 |
| zh-wikipedia | 1048576 | pretok-fresh | 4484954 | 31344 | 4484953.5 | 233.80 | 87002686 | 13856554 | 5559536.0 | 27.0 | 10 |
| zh-wikipedia | 1048576 | pretok-reused | 4127293 | 22625 | 4127292.7 | 254.06 | 94542120 | 15057328 | 0.0 | 0.0 | 10 |

## Regex iteration

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.192x |
| ar-wikipedia | 4096 | 3.043x |
| ar-wikipedia | 65536 | 3.023x |
| ar-wikipedia | 1048576 | 3.028x |
| en-wikitext2 | 256 | 5.689x |
| en-wikitext2 | 4096 | 5.577x |
| en-wikitext2 | 65536 | 5.540x |
| en-wikitext2 | 1048576 | 5.544x |
| ja-wikipedia | 256 | 2.538x |
| ja-wikipedia | 4096 | 2.383x |
| ja-wikipedia | 65536 | 2.365x |
| ja-wikipedia | 1048576 | 2.358x |
| ru-wikipedia | 256 | 3.610x |
| ru-wikipedia | 4096 | 3.426x |
| ru-wikipedia | 65536 | 3.397x |
| ru-wikipedia | 1048576 | 3.432x |
| synthetic | 0 | 4.319x |
| zh-wikipedia | 256 | 3.111x |
| zh-wikipedia | 4096 | 2.936x |
| zh-wikipedia | 65536 | 2.887x |
| zh-wikipedia | 1048576 | 2.896x |
| Combined selected natural corpora | 256 | 3.587x |
| Combined selected natural corpora | 4096 | 3.437x |
| Combined selected natural corpora | 65536 | 3.406x |
| Combined selected natural corpora | 1048576 | 3.416x |

## Regex span collection, reused

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 3.158x |
| ar-wikipedia | 4096 | 3.020x |
| ar-wikipedia | 65536 | 3.007x |
| ar-wikipedia | 1048576 | 3.017x |
| en-wikitext2 | 256 | 5.623x |
| en-wikitext2 | 4096 | 5.567x |
| en-wikitext2 | 65536 | 5.438x |
| en-wikitext2 | 1048576 | 5.480x |
| ja-wikipedia | 256 | 2.527x |
| ja-wikipedia | 4096 | 2.373x |
| ja-wikipedia | 65536 | 2.344x |
| ja-wikipedia | 1048576 | 2.338x |
| ru-wikipedia | 256 | 3.560x |
| ru-wikipedia | 4096 | 3.413x |
| ru-wikipedia | 65536 | 3.372x |
| ru-wikipedia | 1048576 | 3.396x |
| synthetic | 0 | 4.260x |
| zh-wikipedia | 256 | 3.103x |
| zh-wikipedia | 4096 | 2.916x |
| zh-wikipedia | 65536 | 2.867x |
| zh-wikipedia | 1048576 | 2.883x |
| Combined selected natural corpora | 256 | 3.556x |
| Combined selected natural corpora | 4096 | 3.423x |
| Combined selected natural corpora | 65536 | 3.373x |
| Combined selected natural corpora | 1048576 | 3.391x |

## Regex span collection, fresh

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 2.812x |
| ar-wikipedia | 4096 | 2.838x |
| ar-wikipedia | 65536 | 2.745x |
| ar-wikipedia | 1048576 | 2.777x |
| en-wikitext2 | 256 | 4.562x |
| en-wikitext2 | 4096 | 4.781x |
| en-wikitext2 | 65536 | 4.630x |
| en-wikitext2 | 1048576 | 4.673x |
| ja-wikipedia | 256 | 2.335x |
| ja-wikipedia | 4096 | 2.277x |
| ja-wikipedia | 65536 | 2.259x |
| ja-wikipedia | 1048576 | 2.237x |
| ru-wikipedia | 256 | 3.147x |
| ru-wikipedia | 4096 | 3.189x |
| ru-wikipedia | 65536 | 3.125x |
| ru-wikipedia | 1048576 | 3.097x |
| synthetic | 0 | 3.540x |
| zh-wikipedia | 256 | 2.771x |
| zh-wikipedia | 4096 | 2.740x |
| zh-wikipedia | 65536 | 2.684x |
| zh-wikipedia | 1048576 | 2.650x |
| Combined selected natural corpora | 256 | 3.129x |
| Combined selected natural corpora | 4096 | 3.169x |
| Combined selected natural corpora | 65536 | 3.093x |
| Combined selected natural corpora | 1048576 | 3.092x |

## HF stage versus reused destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 5.704x |
| ar-wikipedia | 4096 | 5.093x |
| ar-wikipedia | 65536 | 4.981x |
| ar-wikipedia | 1048576 | 5.683x |
| en-wikitext2 | 256 | 10.407x |
| en-wikitext2 | 4096 | 9.709x |
| en-wikitext2 | 65536 | 9.660x |
| en-wikitext2 | 1048576 | 11.024x |
| ja-wikipedia | 256 | 4.286x |
| ja-wikipedia | 4096 | 3.802x |
| ja-wikipedia | 65536 | 3.693x |
| ja-wikipedia | 1048576 | 4.120x |
| ru-wikipedia | 256 | 6.591x |
| ru-wikipedia | 4096 | 5.898x |
| ru-wikipedia | 65536 | 5.726x |
| ru-wikipedia | 1048576 | 6.567x |
| synthetic | 0 | 7.712x |
| zh-wikipedia | 256 | 5.547x |
| zh-wikipedia | 4096 | 4.910x |
| zh-wikipedia | 65536 | 4.786x |
| zh-wikipedia | 1048576 | 5.623x |
| Combined selected natural corpora | 256 | 6.432x |
| Combined selected natural corpora | 4096 | 5.818x |
| Combined selected natural corpora | 65536 | 5.703x |
| Combined selected natural corpora | 1048576 | 6.524x |

## HF stage versus fresh destination

Speedup = reference time / pretok time. Below 1 means pretok is slower. HF rows include its richer native representation and are not pure regex-engine speedups.

| Corpus | Request bytes | Speedup |
|---|---:|---:|
| ar-wikipedia | 256 | 4.891x |
| ar-wikipedia | 4096 | 4.700x |
| ar-wikipedia | 65536 | 4.556x |
| ar-wikipedia | 1048576 | 5.222x |
| en-wikitext2 | 256 | 8.180x |
| en-wikitext2 | 4096 | 8.264x |
| en-wikitext2 | 65536 | 8.194x |
| en-wikitext2 | 1048576 | 9.417x |
| ja-wikipedia | 256 | 3.833x |
| ja-wikipedia | 4096 | 3.603x |
| ja-wikipedia | 65536 | 3.539x |
| ja-wikipedia | 1048576 | 3.916x |
| ru-wikipedia | 256 | 5.631x |
| ru-wikipedia | 4096 | 5.436x |
| ru-wikipedia | 65536 | 5.296x |
| ru-wikipedia | 1048576 | 5.990x |
| synthetic | 0 | 6.181x |
| zh-wikipedia | 256 | 4.752x |
| zh-wikipedia | 4096 | 4.546x |
| zh-wikipedia | 65536 | 4.462x |
| zh-wikipedia | 1048576 | 5.174x |
| Combined selected natural corpora | 256 | 5.464x |
| Combined selected natural corpora | 4096 | 5.315x |
| Combined selected natural corpora | 65536 | 5.216x |
| Combined selected natural corpora | 1048576 | 5.944x |

## Compilation

One-time setup, including construction result lifetime/cleanup. Not included in scanning measurements.

| Engine | ns/compile | IQR ns | Samples |
|---|---:|---:|---:|
| hf-compile | 69944.1 | 664.9 | 10 |
| onig-compile | 69922.0 | 219.3 | 10 |
| pretok-compile | 13.4 | 0.1 | 10 |

Natural corpora and synthetic diagnostics are separate workloads. Results apply to these excerpts and this machine; this is not a complete BPE tokenizer comparison.
