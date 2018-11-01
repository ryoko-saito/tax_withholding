package tax

/** ajustのためのパラメータ **/
// var supportOverDiduct int //扶養で7人を超えた時に増える控除額
// var FOLLOW int            //従たる給与についての扶養控除などの申告書が提出されていない場合0、提出されている場合1
// var kouOrOtsu int         //甲 or 乙のパラメータ
// var SUPPORT int

type BonusSets struct {
	set []BonusSet
}

type BonusSet struct {
	max  int
	rate float64
}

func NewBonusList() []BonusSets {

	return []BonusSets{
		//扶養が0の時のrate
		BonusSets{
			set: []BonusSet{
				BonusSet{max: 68, rate: 0},
				BonusSet{max: 79, rate: 2.042},
				BonusSet{max: 252, rate: 4.084},
				BonusSet{max: 300, rate: 6.126},
				BonusSet{max: 334, rate: 8.168},
				BonusSet{max: 363, rate: 10.210},
				BonusSet{max: 395, rate: 12.252},
				BonusSet{max: 426, rate: 14.294},
				BonusSet{max: 550, rate: 16.336},
				BonusSet{max: 647, rate: 18.378},
				BonusSet{max: 699, rate: 20.420},
				BonusSet{max: 730, rate: 22.462},
				BonusSet{max: 764, rate: 24.504},
				BonusSet{max: 804, rate: 26.546},
				BonusSet{max: 857, rate: 28.588},
				BonusSet{max: 926, rate: 30.630},
				BonusSet{max: 1321, rate: 32.672},
				BonusSet{max: 1532, rate: 35.735},
				BonusSet{max: 2661, rate: 38.798},
				BonusSet{max: 3548, rate: 41.861},
			},
		},
		//扶養が1の時のrate
		BonusSets{
			set: []BonusSet{
				BonusSet{max: 94, rate: 0},
				BonusSet{max: 243, rate: 2.042},
				BonusSet{max: 282, rate: 4.084},
				BonusSet{max: 338, rate: 6.126},
				BonusSet{max: 365, rate: 8.168},
				BonusSet{max: 394, rate: 10.210},
				BonusSet{max: 422, rate: 12.252},
				BonusSet{max: 455, rate: 14.294},
				BonusSet{max: 550, rate: 16.336},
				BonusSet{max: 663, rate: 18.378},
				BonusSet{max: 720, rate: 20.420},
				BonusSet{max: 752, rate: 22.462},
				BonusSet{max: 787, rate: 24.504},
				BonusSet{max: 826, rate: 26.546},
				BonusSet{max: 885, rate: 28.588},
				BonusSet{max: 956, rate: 30.630},
				BonusSet{max: 1346, rate: 32.672},
				BonusSet{max: 1560, rate: 35.735},
				BonusSet{max: 2685, rate: 38.798},
				BonusSet{max: 3580, rate: 41.861},
			},
		},
		//扶養が2の時のrate
		BonusSets{
			set: []BonusSet{
				BonusSet{max: 133, rate: 0},
				BonusSet{max: 269, rate: 2.042},
				BonusSet{max: 312, rate: 4.084},
				BonusSet{max: 369, rate: 6.126},
				BonusSet{max: 393, rate: 8.168},
				BonusSet{max: 420, rate: 10.210},
				BonusSet{max: 450, rate: 12.252},
				BonusSet{max: 484, rate: 14.294},
				BonusSet{max: 550, rate: 16.336},
				BonusSet{max: 678, rate: 18.378},
				BonusSet{max: 741, rate: 20.420},
				BonusSet{max: 774, rate: 22.462},
				BonusSet{max: 810, rate: 24.504},
				BonusSet{max: 852, rate: 26.546},
				BonusSet{max: 914, rate: 28.588},
				BonusSet{max: 987, rate: 30.630},
				BonusSet{max: 1370, rate: 32.672},
				BonusSet{max: 1589, rate: 35.735},
				BonusSet{max: 2708, rate: 38.798},
				BonusSet{max: 3611, rate: 41.861},
			},
		},
		//扶養が3の時のrate
		BonusSets{
			set: []BonusSet{
				BonusSet{max: 171, rate: 0},
				BonusSet{max: 295, rate: 2.042},
				BonusSet{max: 345, rate: 4.084},
				BonusSet{max: 398, rate: 6.126},
				BonusSet{max: 417, rate: 8.168},
				BonusSet{max: 445, rate: 10.210},
				BonusSet{max: 477, rate: 12.252},
				BonusSet{max: 513, rate: 14.294},
				BonusSet{max: 557, rate: 16.336},
				BonusSet{max: 693, rate: 18.378},
				BonusSet{max: 762, rate: 20.420},
				BonusSet{max: 796, rate: 22.462},
				BonusSet{max: 833, rate: 24.504},
				BonusSet{max: 879, rate: 26.546},
				BonusSet{max: 942, rate: 28.588},
				BonusSet{max: 1017, rate: 30.630},
				BonusSet{max: 1394, rate: 32.672},
				BonusSet{max: 1617, rate: 35.735},
				BonusSet{max: 2732, rate: 38.798},
				BonusSet{max: 3643, rate: 41.861},
			},
		},
		//扶養が4の時のrate
		BonusSets{
			set: []BonusSet{
				BonusSet{max: 210, rate: 0},
				BonusSet{max: 300, rate: 2.042},
				BonusSet{max: 378, rate: 4.084},
				BonusSet{max: 424, rate: 6.126},
				BonusSet{max: 444, rate: 8.168},
				BonusSet{max: 470, rate: 10.210},
				BonusSet{max: 504, rate: 12.252},
				BonusSet{max: 543, rate: 14.294},
				BonusSet{max: 591, rate: 16.336},
				BonusSet{max: 591, rate: 18.378},
				BonusSet{max: 708, rate: 20.420},
				BonusSet{max: 828, rate: 22.462},
				BonusSet{max: 859, rate: 24.504},
				BonusSet{max: 906, rate: 26.546},
				BonusSet{max: 970, rate: 28.588},
				BonusSet{max: 1048, rate: 30.630},
				BonusSet{max: 1419, rate: 32.672},
				BonusSet{max: 1645, rate: 35.735},
				BonusSet{max: 2756, rate: 38.798},
				BonusSet{max: 3675, rate: 41.861},
			},
		},

		//扶養が5の時のrate
		BonusSets{
			set: []BonusSet{
				BonusSet{max: 243, rate: 0},
				BonusSet{max: 300, rate: 2.042},
				BonusSet{max: 406, rate: 4.084},
				BonusSet{max: 450, rate: 6.126},
				BonusSet{max: 472, rate: 8.168},
				BonusSet{max: 496, rate: 10.210},
				BonusSet{max: 531, rate: 12.252},
				BonusSet{max: 574, rate: 14.294},
				BonusSet{max: 618, rate: 16.336},
				BonusSet{max: 723, rate: 18.378},
				BonusSet{max: 804, rate: 20.420},
				BonusSet{max: 841, rate: 22.462},
				BonusSet{max: 885, rate: 24.504},
				BonusSet{max: 934, rate: 26.546},
				BonusSet{max: 998, rate: 28.588},
				BonusSet{max: 1078, rate: 30.630},
				BonusSet{max: 1443, rate: 32.672},
				BonusSet{max: 1674, rate: 35.735},
				BonusSet{max: 2780, rate: 38.798},
				BonusSet{max: 3706, rate: 41.861},
			},
		},

		//扶養が6の時のrate
		BonusSets{
			set: []BonusSet{
				BonusSet{max: 275, rate: 0},
				BonusSet{max: 333, rate: 2.042},
				BonusSet{max: 431, rate: 4.084},
				BonusSet{max: 476, rate: 6.126},
				BonusSet{max: 499, rate: 8.168},
				BonusSet{max: 525, rate: 10.210},
				BonusSet{max: 559, rate: 12.252},
				BonusSet{max: 602, rate: 14.294},
				BonusSet{max: 645, rate: 16.336},
				BonusSet{max: 739, rate: 18.378},
				BonusSet{max: 825, rate: 20.420},
				BonusSet{max: 865, rate: 22.462},
				BonusSet{max: 911, rate: 24.504},
				BonusSet{max: 961, rate: 26.546},
				BonusSet{max: 1026, rate: 28.588},
				BonusSet{max: 1108, rate: 30.630},
				BonusSet{max: 1468, rate: 32.672},
				BonusSet{max: 1702, rate: 35.735},
				BonusSet{max: 2803, rate: 38.798},
				BonusSet{max: 3738, rate: 41.861},
			},
		},

		//扶養が7人以上の時のrate
		BonusSets{
			set: []BonusSet{
				BonusSet{max: 308, rate: 0},
				BonusSet{max: 372, rate: 2.042},
				BonusSet{max: 456, rate: 4.084},
				BonusSet{max: 502, rate: 6.126},
				BonusSet{max: 527, rate: 8.168},
				BonusSet{max: 553, rate: 10.210},
				BonusSet{max: 588, rate: 12.252},
				BonusSet{max: 627, rate: 14.294},
				BonusSet{max: 671, rate: 16.336},
				BonusSet{max: 754, rate: 18.378},
				BonusSet{max: 848, rate: 20.420},
				BonusSet{max: 890, rate: 22.462},
				BonusSet{max: 937, rate: 24.504},
				BonusSet{max: 988, rate: 26.546},
				BonusSet{max: 1054, rate: 28.588},
				BonusSet{max: 1139, rate: 30.630},
				BonusSet{max: 1492, rate: 32.672},
				BonusSet{max: 1730, rate: 35.735},
				BonusSet{max: 2827, rate: 38.798},
				BonusSet{max: 3770, rate: 41.861},
			},
		},
	}
}

func NewBonusOtsuList() []BonusSet {
	return []BonusSet{
		BonusSet{max: 239, rate: 10.210},
		BonusSet{max: 296, rate: 20.420},
		BonusSet{max: 528, rate: 30.630},
		BonusSet{max: 1135, rate: 38.798},
	}
}

// kou_or_otsu kouは0、otsuは1にする
func CalcBonusTax(income int, kou_or_otsu int, support int, follow int) int {
	i := int(income / 1000)

	if support > 7 {
		support = 7
	}

	/**最初に上限値の計算を行う**/
	//甲
	if kou_or_otsu == 0 {
		cnds := []int{3548, 3580, 3611, 3643, 3675, 3706, 3738, 3770}
		if i >= cnds[support] {
			return CalcIntMulFloat(income, 45.945/100)
		}
		//乙
	} else {
		if i >= 1135 {
			return CalcIntMulFloat(income, 45.945/100)
		}
	}

	/** ここからは重い処理 **/
	var list []BonusSet

	//甲
	if kou_or_otsu == 0 {
		l := NewBonusList()
		list = l[support].set

		//乙
	} else {
		list = NewBonusOtsuList()
	}

	min := 0
	for _, set := range list {
		if i >= min && i < set.max {
			return CalcIntMulFloat(income, set.rate/100)
		}
		min = set.max
	}

	//ここを通ることはあり得ないので、0を返すことにする
	return 0
}
