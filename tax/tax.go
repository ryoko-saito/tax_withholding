package tax

type TaxSet struct {
	index int   //(月収-社会保険)/1000
	kou   []int //甲
	otsu  int   //乙
}

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
	}

	return list
}
