package tax

type TaxSet struct {
	index int   //(月収-社会保険)/1000
	kou   []int //甲
	otsu  int   //乙
}

const kouMax = 500000
const otsuMax = 400000

func NewTaxList() []TaxSet {
	list := []TaxSet{
		TaxSet{89, []int{130, 0, 0, 0, 0, 0, 0, 0}, 3200},
		TaxSet{90, []int{180, 0, 0, 0, 0, 0, 0, 0}, 3200},
		TaxSet{91, []int{230, 0, 0, 0, 0, 0, 0, 0}, 3200},
		TaxSet{92, []int{290, 0, 0, 0, 0, 0, 0, 0}, 3200},
		TaxSet{93, []int{340, 0, 0, 0, 0, 0, 0, 0}, 3300},
		TaxSet{94, []int{390, 0, 0, 0, 0, 0, 0, 0}, 3300},
		TaxSet{95, []int{440, 0, 0, 0, 0, 0, 0, 0}, 3300},
		TaxSet{96, []int{490, 0, 0, 0, 0, 0, 0, 0}, 3400},
		TaxSet{97, []int{540, 0, 0, 0, 0, 0, 0, 0}, 3400},
		TaxSet{98, []int{590, 0, 0, 0, 0, 0, 0, 0}, 3500},
		TaxSet{99, []int{640, 0, 0, 0, 0, 0, 0, 0}, 3500},
		TaxSet{101, []int{720, 0, 0, 0, 0, 0, 0, 0}, 3600},
		TaxSet{103, []int{830, 0, 0, 0, 0, 0, 0, 0}, 3600},
		TaxSet{105, []int{930, 0, 0, 0, 0, 0, 0, 0}, 3700},
		TaxSet{107, []int{1030, 0, 0, 0, 0, 0, 0, 0}, 3800},
		TaxSet{109, []int{1130, 0, 0, 0, 0, 0, 0, 0}, 3800},
		TaxSet{111, []int{1240, 0, 0, 0, 0, 0, 0, 0}, 3900},
		TaxSet{113, []int{1340, 0, 0, 0, 0, 0, 0, 0}, 4000},
		TaxSet{115, []int{1440, 0, 0, 0, 0, 0, 0, 0}, 4100},
		TaxSet{117, []int{1540, 0, 0, 0, 0, 0, 0, 0}, 4100},
		TaxSet{119, []int{1640, 0, 0, 0, 0, 0, 0, 0}, 4200},
		TaxSet{121, []int{1750, 120, 0, 0, 0, 0, 0, 0}, 4300},
		TaxSet{123, []int{1850, 220, 0, 0, 0, 0, 0, 0}, 4500},
		TaxSet{125, []int{1950, 330, 0, 0, 0, 0, 0, 0}, 4800},
		TaxSet{127, []int{2050, 430, 0, 0, 0, 0, 0, 0}, 5100},
		TaxSet{129, []int{2150, 530, 0, 0, 0, 0, 0, 0}, 5400},
		TaxSet{131, []int{2250, 630, 0, 0, 0, 0, 0, 0}, 5700},
		TaxSet{133, []int{2360, 740, 0, 0, 0, 0, 0, 0}, 6000},
		TaxSet{135, []int{2460, 840, 0, 0, 0, 0, 0, 0}, 6300},
		TaxSet{137, []int{2550, 930, 0, 0, 0, 0, 0, 0}, 6600},
		TaxSet{139, []int{2610, 990, 0, 0, 0, 0, 0, 0}, 6800},
		TaxSet{141, []int{2680, 1050, 0, 0, 0, 0, 0, 0}, 7100},
		TaxSet{143, []int{2740, 1110, 0, 0, 0, 0, 0, 0}, 7500},
		TaxSet{145, []int{2800, 1170, 0, 0, 0, 0, 0, 0}, 7800},
		TaxSet{147, []int{2860, 1240, 0, 0, 0, 0, 0, 0}, 8100},
		TaxSet{149, []int{2920, 1300, 0, 0, 0, 0, 0, 0}, 8400},
		TaxSet{151, []int{2920, 1360, 0, 0, 0, 0, 0, 0}, 8700},
		TaxSet{153, []int{3050, 1430, 0, 0, 0, 0, 0, 0}, 9000},
		TaxSet{153, []int{3120, 1500, 0, 0, 0, 0, 0, 0}, 9300},
		TaxSet{157, []int{3200, 1570, 0, 0, 0, 0, 0, 0}, 9600},
		TaxSet{159, []int{3270, 1640, 0, 0, 0, 0, 0, 0}, 9900},
		TaxSet{161, []int{3340, 1720, 100, 0, 0, 0, 0, 0}, 10200},
		TaxSet{163, []int{3410, 1790, 170, 0, 0, 0, 0, 0}, 10500},
		TaxSet{165, []int{3480, 1860, 250, 0, 0, 0, 0, 0}, 10800},
		TaxSet{167, []int{3550, 1930, 320, 0, 0, 0, 0, 0}, 11100},
		TaxSet{169, []int{3620, 2000, 390, 0, 0, 0, 0, 0}, 11400},
		TaxSet{171, []int{3700, 2070, 460, 0, 0, 0, 0, 0}, 11700},
		TaxSet{173, []int{3770, 2140, 530, 0, 0, 0, 0, 0}, 12000},
		TaxSet{175, []int{3840, 2220, 600, 0, 0, 0, 0, 0}, 12400},
		TaxSet{177, []int{3910, 2290, 670, 0, 0, 0, 0, 0}, 12700},
		TaxSet{179, []int{3980, 2360, 750, 0, 0, 0, 0, 0}, 13200},
		TaxSet{181, []int{4050, 2430, 820, 0, 0, 0, 0, 0}, 13900},
		TaxSet{183, []int{4120, 2500, 890, 0, 0, 0, 0, 0}, 14600},
		TaxSet{185, []int{4200, 2570, 960, 0, 0, 0, 0, 0}, 15300},
		TaxSet{187, []int{4270, 2640, 1030, 0, 0, 0, 0, 0}, 16000},
		TaxSet{189, []int{4340, 2720, 1100, 0, 0, 0, 0, 0}, 16700},
		TaxSet{191, []int{4410, 2790, 1170, 0, 0, 0, 0, 0}, 17500},
		TaxSet{193, []int{4480, 2860, 1250, 0, 0, 0, 0, 0}, 18100},
		TaxSet{195, []int{4550, 2930, 1320, 0, 0, 0, 0, 0}, 18800},
		TaxSet{197, []int{4630, 3000, 1390, 0, 0, 0, 0, 0}, 19500},
		TaxSet{199, []int{4700, 3070, 1460, 0, 0, 0, 0, 0}, 20200},
		TaxSet{201, []int{4770, 3140, 1530, 0, 0, 0, 0, 0}, 20900},
		TaxSet{203, []int{4840, 3220, 1600, 0, 0, 0, 0, 0}, 21500},
		TaxSet{205, []int{4910, 3290, 1670, 0, 0, 0, 0, 0}, 22200},
		TaxSet{207, []int{4980, 3360, 1750, 130, 0, 0, 0, 0}, 22700},
		TaxSet{209, []int{5050, 3430, 1820, 200, 0, 0, 0, 0}, 23300},
		TaxSet{211, []int{5130, 3500, 1890, 280, 0, 0, 0, 0}, 23900},
		TaxSet{213, []int{5200, 3570, 1960, 350, 0, 0, 0, 0}, 24400},
		TaxSet{215, []int{5270, 3640, 2030, 420, 0, 0, 0, 0}, 25000},
		TaxSet{217, []int{5340, 3720, 2100, 490, 0, 0, 0, 0}, 25500},
		TaxSet{219, []int{5410, 3790, 2170, 560, 0, 0, 0, 0}, 26100},
		TaxSet{221, []int{5480, 3860, 2250, 630, 0, 0, 0, 0}, 26800},
		TaxSet{224, []int{5560, 3950, 2340, 710, 0, 0, 0, 0}, 27400},
		TaxSet{227, []int{5680, 4060, 2440, 830, 0, 0, 0, 0}, 28400},
		TaxSet{230, []int{5780, 4170, 2550, 9830, 0, 0, 0, 0}, 29300},
		TaxSet{233, []int{5890, 4280, 2650, 1040, 0, 0, 0, 0}, 30300},
		TaxSet{236, []int{5990, 4380, 2770, 1140, 0, 0, 0, 0}, 31300},
		TaxSet{239, []int{6110, 4490, 2870, 1260, 0, 0, 0, 0}, 32400},
		TaxSet{242, []int{6210, 4590, 2970, 1360, 0, 0, 0, 0}, 33400},
		TaxSet{245, []int{6320, 4710, 3080, 1470, 0, 0, 0, 0}, 34400},
		TaxSet{248, []int{6420, 4810, 3200, 1570, 0, 0, 0, 0}, 35400},
		TaxSet{251, []int{6530, 4920, 3300, 1680, 0, 0, 0, 0}, 36400},
		TaxSet{254, []int{6640, 5020, 3410, 1790, 0, 0, 0, 0}, 37500},
		TaxSet{257, []int{6750, 5140, 3510, 1900, 290, 0, 0, 0}, 38500},
		TaxSet{260, []int{6850, 5240, 3620, 2000, 390, 0, 0, 0}, 39400},
		TaxSet{263, []int{6960, 5350, 3730, 2110, 500, 0, 0, 0}, 40400},
		TaxSet{266, []int{7070, 5450, 3840, 2220, 600, 0, 0, 0}, 41500},
		TaxSet{269, []int{7180, 5560, 3940, 2330, 710, 0, 0, 0}, 42500},
		TaxSet{272, []int{7280, 5670, 4050, 2430, 820, 0, 0, 0}, 43500},
		TaxSet{275, []int{7390, 5780, 4160, 2540, 930, 0, 0, 0}, 44500},
		TaxSet{278, []int{7490, 5880, 4270, 2640, 1030, 0, 0, 0}, 45500},
		TaxSet{281, []int{7610, 5990, 4370, 2760, 1140, 0, 0, 0}, 46600},
		TaxSet{284, []int{7710, 6100, 4480, 2860, 1250, 0, 0, 0}, 47600},
		TaxSet{287, []int{7820, 6210, 4580, 2970, 1360, 0, 0, 0}, 48600},
		TaxSet{290, []int{7920, 6310, 4700, 3070, 1460, 0, 0, 0}, 49500},
		TaxSet{293, []int{8040, 6420, 4800, 3190, 1570, 0, 0, 0}, 50500},
		TaxSet{296, []int{8140, 6520, 4910, 3290, 1670, 0, 0, 0}, 51600},
		TaxSet{299, []int{8250, 6640, 5010, 3400, 1790, 160, 0, 0}, 52300},
		TaxSet{302, []int{8420, 6740, 5130, 3510, 1890, 280, 0, 0}, 52900},
		TaxSet{305, []int{8670, 6860, 5250, 3630, 2010, 400, 0, 0}, 53500},
		TaxSet{308, []int{8910, 6980, 5370, 3760, 2130, 520, 0, 0}, 54200},
		TaxSet{311, []int{9160, 7110, 5490, 3880, 2260, 640, 0, 0}, 54800},
		TaxSet{314, []int{9400, 7230, 5620, 4000, 2380, 770, 0, 0}, 55400},
		TaxSet{317, []int{9650, 7350, 5740, 4120, 2500, 890, 0, 0}, 56100},
		TaxSet{320, []int{9890, 7470, 5860, 4250, 2620, 1010, 0, 0}, 56800},
		TaxSet{323, []int{10140, 7600, 5980, 4370, 2750, 1130, 0, 0}, 57700},
		TaxSet{326, []int{10380, 7720, 6110, 4490, 2870, 1260, 0, 0}, 58500},
		TaxSet{329, []int{10630, 7840, 6230, 4610, 2990, 1380, 0, 0}, 59300},
		TaxSet{332, []int{10870, 7960, 6350, 4740, 3110, 1500, 0, 0}, 60200},
		TaxSet{335, []int{11120, 8090, 6470, 4860, 3240, 1620, 0, 0}, 61100},
		TaxSet{338, []int{11360, 8210, 6600, 4980, 3360, 1750, 130, 0}, 62000},
		TaxSet{341, []int{11610, 8370, 6720, 5110, 3480, 1870, 260, 0}, 63000},
		TaxSet{344, []int{11850, 8260, 6840, 5230, 3600, 1990, 380, 0}, 64000},
		TaxSet{347, []int{12100, 8860, 6960, 5350, 3730, 2110, 500, 0}, 65000},
		TaxSet{350, []int{12340, 9110, 7090, 5470, 3850, 2240, 620, 0}, 66200},
		TaxSet{353, []int{12590, 9350, 7210, 5600, 3970, 2360, 750, 0}, 67200},
		TaxSet{356, []int{12830, 9600, 7330, 5720, 4090, 2480, 870, 0}, 68200},
		TaxSet{359, []int{13080, 9840, 7450, 5840, 4220, 2600, 990, 0}, 69200},
		TaxSet{362, []int{13320, 10090, 7580, 5960, 4340, 2730, 1110, 0}, 70200},
		TaxSet{365, []int{13570, 10330, 7700, 6090, 4460, 2850, 1240, 0}, 71300},
		TaxSet{368, []int{13810, 10580, 7820, 6210, 4580, 2970, 1360, 0}, 72300},
		TaxSet{371, []int{14060, 10820, 7940, 6330, 4710, 3090, 1480, 0}, 73200},
		TaxSet{374, []int{14360, 11070, 8070, 6450, 4830, 3220, 1600, 0}, 74200},
		TaxSet{377, []int{14550, 11310, 8190, 6580, 4950, 3340, 1730, 100}, 75100},
		TaxSet{380, []int{14790, 11560, 8320, 6700, 5070, 3460, 1850, 220}, 76100},
		TaxSet{383, []int{15040, 11800, 8750, 6820, 5200, 3580, 1970, 350}, 77000},
		TaxSet{386, []int{15280, 12050, 8810, 6940, 5320, 3710, 2090, 470}, 77900},
		TaxSet{389, []int{15530, 12290, 9060, 7070, 5440, 3830, 2220, 590}, 78800},
		TaxSet{392, []int{15770, 12540, 9300, 7190, 5560, 3950, 2340, 710}, 80600},
		TaxSet{395, []int{16020, 12780, 9550, 7310, 5690, 4070, 2460, 840}, 82300},
		TaxSet{398, []int{16260, 13030, 9790, 7430, 5810, 4200, 2580, 960}, 83900},
		TaxSet{401, []int{16510, 13270, 10400, 7560, 5930, 4320, 2710, 1080}, 85700},
		TaxSet{404, []int{16750, 13520, 10280, 7680, 6050, 4440, 2830, 1200}, 87400},
		TaxSet{407, []int{17000, 13760, 10530, 7800, 6180, 4560, 2950, 1330}, 89000},
		TaxSet{410, []int{17240, 14010, 10770, 7920, 6300, 4690, 3070, 1450}, 90800},
		TaxSet{413, []int{17249, 14250, 11020, 8050, 6420, 4810, 3200, 1570}, 92500},
		TaxSet{416, []int{17730, 14500, 11260, 8170, 6540, 4930, 3320, 1690}, 94100},
		TaxSet{419, []int{17980, 14740, 11510, 8290, 6670, 5050, 3440, 1820}, 95900},
		TaxSet{422, []int{18220, 14990, 11750, 8530, 6790, 5180, 3560, 1940}, 97600},
		TaxSet{425, []int{18470, 15230, 12000, 8770, 6910, 5300, 3690, 2060}, 99200},
		TaxSet{428, []int{18710, 15480, 12240, 9020, 7030, 5420, 3810, 2180}, 101000},
		TaxSet{431, []int{18960, 15720, 12490, 9260, 7160, 5540, 3930, 2310}, 102600},
		TaxSet{434, []int{19210, 15970, 12730, 9510, 7280, 5670, 4050, 2430}, 104300},
		TaxSet{437, []int{19450, 16210, 12980, 9750, 7400, 5790, 4180, 2550}, 106100},
		TaxSet{440, []int{19700, 16460, 13220, 10000, 7520, 5910, 4300, 2680}, 107700},
		TaxSet{443, []int{20090, 16700, 13470, 10240, 7650, 6030, 4420, 2800}, 109500},
		TaxSet{446, []int{20580, 16950, 13710, 10490, 7770, 6160, 4540, 2920}, 111200},
		TaxSet{449, []int{21070, 171900, 13960, 10730, 7890, 6280, 4670, 3040}, 112800},
		TaxSet{452, []int{21560, 17440, 14200, 10980, 8010, 6400, 4790, 3170}, 114600},
		TaxSet{455, []int{22050, 17680, 14450, 11220, 8140, 6520, 4910, 3290}, 116300},
		TaxSet{458, []int{22540, 17930, 14690, 11470, 8260, 6650, 5030, 3410}, 117900},
		TaxSet{461, []int{23030, 18170, 14940, 11710, 8470, 6770, 5160, 3530}, 119700},
		TaxSet{464, []int{23520, 18420, 15180, 11960, 8720, 6890, 5280, 3660}, 121400},
		TaxSet{467, []int{24010, 18660, 15430, 12200, 8960, 7010, 5400, 3780}, 123000},
		TaxSet{470, []int{24500, 18910, 15670, 12450, 9210, 7140, 5520, 3900}, 124800},
		TaxSet{473, []int{24990, 19150, 15920, 12690, 9450, 726, 5650, 4020}, 126500},
		TaxSet{476, []int{25480, 19400, 16160, 12940, 9700, 7380, 57700, 4150}, 128100},
		TaxSet{479, []int{25970, 19640, 16410, 13180, 9940, 7500, 5890, 4270}, 129900},
		TaxSet{482, []int{26460, 20000, 16650, 13430, 10190, 7630, 6010, 4390}, 131600},
		TaxSet{485, []int{26950, 20490, 16900, 13670, 10430, 7750, 6140, 4510}, 133200},
		TaxSet{488, []int{27440, 20980, 17140, 13920, 10680, 7870, 6260, 4640}, 135000},
		TaxSet{491, []int{27930, 21470, 17390, 14160, 10920, 7990, 6380, 4760}, 136600},
		TaxSet{497, []int{28910, 21450, 17880, 14650, 11410, 8240, 6630, 5000}, 140100},
		TaxSet{500, []int{29400, 22940, 18120, 14900, 11660, 8420, 6750, 5130}, 141700},
		TaxSet{503, []int{29890, 23430, 18370, 15140, 11900, 8670, 6870, 5250}, 143500},
		TaxSet{506, []int{30380, 23920, 18610, 15390, 12150, 8910, 6990, 5370}, 145200},
		TaxSet{509, []int{30880, 24410, 18860, 15630, 12390, 9160, 7120, 5490}, 146800},
		TaxSet{512, []int{31370, 24900, 19100, 15880, 12640, 9400, 7240, 5620}, 148600},
		TaxSet{515, []int{31860, 25390, 19350, 16120, 12890, 9650, 7360, 5740}, 150300},
		TaxSet{518, []int{32350, 25880, 19590, 16370, 13130, 9890, 7480, 5860}, 151900},
		TaxSet{521, []int{32840, 26370, 19900, 16610, 13380, 10140, 7610, 5980}, 153700},
		TaxSet{524, []int{33330, 26860, 20390, 16860, 13620, 10380, 7730, 6110}, 155400},
		TaxSet{527, []int{33820, 27350, 20880, 17100, 13870, 10630, 7850, 6230}, 157000},
		TaxSet{530, []int{34310, 27840, 21370, 17350, 14110, 10870, 7970, 6350}, 158800},
		TaxSet{533, []int{34800, 28330, 21860, 17590, 14360, 11120, 8100, 6470}, 160300},
		TaxSet{536, []int{35290, 28820, 22350, 17840, 14600, 11360, 8220, 6600}, 161900},
		TaxSet{539, []int{35780, 29310, 22840, 18080, 14850, 11610, 8380, 6720}, 163500},
		TaxSet{542, []int{36270, 29800, 23330, 18330, 15090, 11850, 8630, 6840}, 165000},
		TaxSet{545, []int{36760, 30290, 23820, 18570, 15340, 12100, 8870, 6960}, 166000},
		TaxSet{548, []int{37250, 30780, 24310, 18820, 15580, 12340, 9120, 7090}, 168200},
		TaxSet{551, []int{37740, 31270, 24800, 19060, 15830, 12590, 9360, 7210}, 169800},
		TaxSet{554, []int{38280, 31810, 25340, 19330, 16100, 12869, 9630, 7350}, 171300},
		TaxSet{557, []int{38830, 32370, 25890, 19600, 16380, 13140, 9900, 7480}, 173000},
		TaxSet{560, []int{39380, 32920, 26440, 19980, 16650, 13420, 10180, 7630}, 174500},
		TaxSet{563, []int{39930, 33470, 27000, 20530, 16930, 13690, 10460, 7760}, 175900},
		TaxSet{566, []int{40480, 34020, 27550, 21080, 17200, 13970, 10730, 7900}, 177300},
		TaxSet{569, []int{41030, 34570, 28100, 21630, 17480, 14240, 11010, 8040}, 178900},
		TaxSet{572, []int{415930, 35120, 28650, 22190, 17760, 14520, 11280, 8180}, 180300},
		TaxSet{575, []int{42140, 35670, 29200, 22740, 18030, 14790, 11560, 8330}, 181800},
		TaxSet{578, []int{42690, 36230, 29750, 23290, 18310, 15070, 11830, 8610}, 183300},
		TaxSet{581, []int{43240, 36780, 30300, 23840, 18580, 15350, 12110, 8880}, 184700},
		TaxSet{584, []int{43790, 37330, 30850, 24390, 18860, 15620, 12380, 9160}, 186200},
		TaxSet{587, []int{44340, 37880, 31410, 24940, 19130, 15900, 12660, 9430}, 187700},
		TaxSet{590, []int{44890, 38430, 31960, 25490, 19410, 16170, 12940, 9710}, 189200},
		TaxSet{593, []int{45440, 38980, 32510, 26050, 19680, 16450, 13210, 9990}, 190600},
		TaxSet{596, []int{46000, 39530, 33060, 26600, 20130, 16720, 13490, 10260}, 192100},
		TaxSet{599, []int{46550, 40080, 33610, 27150, 20690, 17000, 13760, 10540}, 193600},
		TaxSet{602, []int{47100, 40640, 34160, 27700, 21240, 17280, 14040, 10810}, 195000},
		TaxSet{605, []int{47650, 41190, 34710, 28250, 21790, 17550, 14310, 11090}, 196500},
		TaxSet{608, []int{48200, 41740, 35270, 28800, 22340, 17830, 14590, 11360}, 198000},
		TaxSet{611, []int{48750, 42290, 35820, 29350, 22890, 18100, 14870, 11640}, 199400},
		TaxSet{614, []int{49300, 42840, 36370, 29910, 23440, 18380, 15140, 11920}, 200900},
		TaxSet{617, []int{49860, 43390, 36920, 30460, 23990, 18650, 15420, 12190}, 202400},
		TaxSet{620, []int{50410, 43940, 37470, 31010, 24540, 18930, 15690, 12470}, 203900},
		TaxSet{623, []int{50960, 44500, 38020, 31560, 25100, 19210, 15970, 12740}, 205300},
		TaxSet{626, []int{51510, 45050, 38570, 32110, 25650, 19480, 16240, 13020}, 206800},
		TaxSet{629, []int{52060, 45600, 39120, 32660, 26200, 19760, 16520, 13290}, 208300},
		TaxSet{632, []int{52610, 46150, 39680, 33210, 26750, 20280, 16800, 13570}, 209700},
		TaxSet{635, []int{53160, 46700, 40230, 33760, 27300, 20830, 17070, 13840}, 211200},
		TaxSet{638, []int{53710, 47250, 40780, 34320, 27850, 21380, 17350, 14120}, 212700},
		TaxSet{641, []int{54270, 47800, 41330, 34870, 28400, 21930, 17620, 14400}, 214100},
		TaxSet{644, []int{54820, 48350, 41880, 35420, 28960, 22480, 17900, 14670}, 215600},
		TaxSet{647, []int{55370, 48910, 42430, 35970, 29510, 23030, 18170, 14950}, 217000},
		TaxSet{650, []int{55920, 49460, 42980, 36520, 30060, 23590, 18450, 15220}, 218000},
		TaxSet{653, []int{56470, 50010, 43540, 37070, 30610, 24140, 18730, 15500}, 219000},
		TaxSet{656, []int{57020, 50560, 44090, 37620, 31160, 24690, 19000, 15770}, 220000},
		TaxSet{659, []int{57570, 51110, 44640, 38180, 31710, 25240, 19280, 16050}, 221000},
		TaxSet{662, []int{58130, 51660, 45190, 38730, 32260, 25790, 19550, 16330}, 222100},
		TaxSet{665, []int{58680, 52210, 45740, 39280, 32810, 26340, 19880, 16600}, 223100},
		TaxSet{668, []int{59230, 52770, 46290, 39830, 33370, 26890, 20430, 16880}, 224100},
		TaxSet{671, []int{59780, 53320, 46840, 40380, 33920, 27440, 20980, 17150}, 225000},
		TaxSet{674, []int{60330, 53870, 47390, 40930, 34470, 28000, 21530, 17430}, 226000},
		TaxSet{677, []int{60880, 54420, 47950, 41480, 35020, 28550, 22080, 17700}, 227100},
		TaxSet{680, []int{61430, 54970, 48590, 42030, 35570, 29100, 22640, 17980}, 228100},
		TaxSet{683, []int{61980, 55520, 49050, 42590, 36120, 29650, 23190, 18260}, 229100},
		TaxSet{686, []int{62540, 56070, 49600, 43140, 36670, 30200, 23740, 18530}, 230100},
		TaxSet{689, []int{63090, 56620, 50150, 43690, 37230, 30750, 24290, 18810}, 231200},
		TaxSet{692, []int{63640, 57180, 50700, 44240, 37780, 31300, 24840, 19080}, 232700},
		TaxSet{695, []int{64190, 57730, 51250, 44790, 38330, 31860, 25390, 19360}, 234200},
		TaxSet{698, []int{64740, 58280, 51810, 45340, 38880, 32410, 25940, 19630}, 235700},
		TaxSet{701, []int{65290, 58830, 52360, 45890, 39430, 32960, 26490, 20030}, 237300},
		TaxSet{704, []int{65840, 59380, 52910, 46450, 39980, 33510, 27050, 20580}, 238900},
		TaxSet{707, []int{66400, 59930, 53460, 47000, 40530, 34060, 27600, 21130}, 240400},
		TaxSet{710, []int{66950, 60480, 54010, 47550, 41090, 34610, 28150, 21690}, 242000},
		TaxSet{710, []int{66950, 60480, 54010, 47550, 41090, 34610, 28150, 21690}, 242000},
		TaxSet{713, []int{67500, 61040, 54560, 48100, 41640, 35160, 28700, 22240}, 243500},
		TaxSet{716, []int{68050, 61590, 55110, 48650, 42190, 35710, 29250, 22790}, 245000},
		TaxSet{719, []int{68600, 62140, 55660, 49200, 42740, 36270, 29800, 23340}, 246600},
		TaxSet{722, []int{69150, 62690, 56220, 49750, 43290, 36820, 30350, 23890}, 248100},
		TaxSet{725, []int{69700, 63240, 56770, 50300, 43840, 37370, 30910, 24440}, 249700},
		TaxSet{728, []int{70260, 63790, 57320, 50860, 44390, 37920, 31460, 24990}, 251300},
		TaxSet{731, []int{70810, 64340, 57870, 51410, 44940, 38470, 32010, 25550}, 252800},
		TaxSet{734, []int{71360, 64890, 58420, 51960, 45500, 39020, 32560, 26100}, 254300},
		TaxSet{737, []int{71910, 65450, 58970, 52510, 46050, 39570, 33110, 26650}, 255900},
		TaxSet{740, []int{72460, 66000, 59520, 53060, 46600, 40130, 33660, 27200}, 257400},
		TaxSet{743, []int{73010, 66550, 60080, 53610, 47150, 40680, 34210, 27750}, 259000},
		TaxSet{746, []int{73560, 67100, 60630, 54160, 47700, 41230, 34770, 28300}, 260600},
		TaxSet{749, []int{74110, 67650, 61180, 54720, 48250, 41780, 35320, 28850}, 262100},
		TaxSet{752, []int{74670, 68200, 61730, 55270, 48800, 42330, 35870, 29400}, 263600},
		TaxSet{755, []int{75220, 68750, 62280, 55820, 49360, 42880, 36420, 29960}, 265200},
		TaxSet{758, []int{75770, 69310, 62830, 56370, 49910, 43430, 36970, 30510}, 266700},
		TaxSet{761, []int{76320, 69860, 63380, 56920, 50460, 43980, 37520, 31060}, 268200},
		TaxSet{764, []int{76870, 70410, 63940, 57470, 51010, 44540, 38070, 31610}, 269900},
		TaxSet{767, []int{77420, 70960, 64490, 58020, 51560, 45090, 38620, 32160}, 271400},
		TaxSet{770, []int{77970, 71510, 65040, 58570, 52110, 45640, 39180, 32710}, 272900},
		TaxSet{773, []int{78530, 72060, 65590, 59130, 52660, 46190, 397380, 33260}, 274400},
		TaxSet{776, []int{79080, 7610, 66140, 59680, 53210, 46740, 40280, 33820}, 276000},
		TaxSet{779, []int{79630, 73160, 66690, 60230, 53770, 47290, 40830, 34370}, 277500},
		TaxSet{782, []int{80180, 73720, 67240, 60780, 54320, 47840, 41380, 34920}, 279000},
		TaxSet{785, []int{80730, 74270, 67790, 61330, 54870, 48440, 41930, 35470}, 280700},
		TaxSet{788, []int{81280, 74820, 68350, 61880, 55420, 48950, 42480, 36020}, 282200},
		TaxSet{791, []int{81830, 75370, 68900, 62430, 55970, 49500, 43040, 36570}, 283700},
		TaxSet{794, []int{82460, 75920, 69450, 62990, 56520, 50050, 43590, 37120}, 285300},
		TaxSet{797, []int{83100, 76470, 70000, 63540, 57070, 50600, 44140, 37670}, 286800},
		TaxSet{800, []int{83730, 77020, 70550, 64090, 57630, 51150, 44690, 38230}, 288300},
		TaxSet{803, []int{84370, 77580, 71100, 64640, 58180, 51700, 45240, 38780}, 290000},
		TaxSet{806, []int{85000, 78130, 71650, 645190, 58730, 52250, 45790, 39330}, 291500},
		TaxSet{809, []int{85630, 78680, 72210, 64740, 59280, 52810, 46340, 39880}, 293000},
		TaxSet{812, []int{86260, 79230, 72760, 66290, 59830, 53360, 46890, 40430}, 294600},
		TaxSet{815, []int{86900, 79780, 73310, 66840, 60380, 53910, 47450, 40980}, 296100},
		TaxSet{818, []int{87530, 80330, 73860, 67400, 60930, 54460, 48000, 41530}, 297600},
		TaxSet{821, []int{88160, 80880, 74410, 67950, 61480, 55010, 48550, 42090}, 299200},
		TaxSet{824, []int{88800, 81430, 74960, 68500, 62040, 55560, 49100, 42640}, 300800},
		TaxSet{827, []int{89440, 82000, 75510, 69050, 62590, 56110, 49650, 43190}, 302300},
		TaxSet{830, []int{90070, 82630, 76060, 69600, 63140, 56670, 50200, 43740}, 303800},
		TaxSet{833, []int{90710, 83260, 76620, 70150, 63690, 57220, 50750, 44290}, 305400},
		TaxSet{836, []int{91360, 83930, 77200, 70720, 64260, 57800, 51330, 44860}, 306900},
		TaxSet{839, []int{92060, 84630, 77810, 71340, 64870, 58410, 51940, 45480}, 308400},
		TaxSet{842, []int{92770, 85340, 78420, 71950, 65490, 59020, 52550, 46090}, 310000},
		TaxSet{845, []int{93470, 86040, 79040, 72560, 66100, 59640, 53160, 46700}, 311600},
		TaxSet{848, []int{94180, 86740, 79650, 73180, 66710, 60250, 53780, 47310}, 313100},
		TaxSet{851, []int{94880, 87450, 80260, 73790, 67320, 60860, 54390, 47930}, 314700},
		TaxSet{854, []int{95590, 88150, 80870, 74400, 67940, 61470, 55000, 48540}, 316200},
		TaxSet{857, []int{96290, 88860, 81490, 75010, 68550, 62090, 55610, 49150}, 317700},
		TaxSet{860, []int{97000, 89560, 82130, 75630, 69160, 62700, 56230, 48760}, 319300},
	}

	return list
}

// kou_or_otsu kouは0、otsuは1にする
func CalcTax(income int, kou_or_otsu int, support int) int {
	i := int(income / 1000)
	list := NewTaxList()

	//supportは7まで。7より大きな数字を指定した場合は7にする
	if support > 7 {
		support = 7
	}

	min := 0
	for _, set := range list {
		if i >= min && i < set.index {
			if kou_or_otsu == 0 {
				return set.kou[support]
			} else {
				return set.otsu
			}
		}
		min = set.index
	}

	if kou_or_otsu == 0 {
		return kouMax
	} else {
		return otsuMax
	}
}
