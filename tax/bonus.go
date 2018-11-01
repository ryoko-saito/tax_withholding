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
			},
		},
		//扶養が1の時のrate
		BonusSets{
			set: []BonusSet{
				BonusSet{max: 94, rate: 0},
				BonusSet{max: 243, rate: 2.042},
				BonusSet{max: 282, rate: 4.084},
			},
		},
	}

}

// type BonusSet struct {
//   rate int //賞与の金額に乗ずべき率
//   kou int   ///甲
//   otsu  int   //乙
// }
//
// //固定長配列のBonusSetという配列の中の要素が構造体になっている
// func NewBonuslist() []BonusSet{
//   return[]BonusSet{
// BonusSet{2.042, []int{ 79  243 133 269 171 295 210 300 243 300 275 333 308 372}, 239},
// BonusSet{4.084, []int{252 282 269 312 295 345 300 378 300 406 333 431 372 456},239},
// BonusSet{6.126, []int{  300  338 312 369 345 398 378 424 406 450 431 476 456 502},239},
// BonusSet{8.168, []int{ 334  365 369 393 398 417 424 444 450 472 476 499 502 527},239},
// BonusSet{10.210, []int{ 363  394 393 420 417 445 444 470 472 496 499 525 527 553},239},
// BonusSet{12.252, []int{ 395  422 420 450 445 477 470 504 496 531 525 559 553 588},239},
// BonusSet{14.294, []int{ 426  455 450 484 477 513 504 543 531 574 559 602 588 627},239},
// BonusSet{16.336, []int{ 550  550 484 550 513 557 543 591 574 618 602 645 627 671},239},
// BonusSet{18.378, []int{ 647  663 550 678 557 693 591 708 618 723 645 739 671 754},239},
// BonusSet{20.420, []int{ 699  720 678 741 693 762 708 783 723 804 739 825 754 848},296},
// BonusSet{22.462, []int{ 730  752 741 774 762 796 783 818 804 841 825 865 848 890},296},
// BonusSet{22.504, []int{ 764  787 774 810 796 833 818 859 841 885 865 911 890 937},296},
// BonusSet{26.546, []int{ 804  826 810 852 833 879 859 906 885 934 911 961 937 988},296},
// BonusSet{28.588, []int{ 857  885 852 914 879 942 906 970 934 998 961 1,026 988 1,054},296},
// BonusSet{30.630, []int{ 926  956 914 987 942 1,017 970 1,048 998 1,078 1,026 1,108 1,054 1,139},528},
// BonusSet{32.672, []int{ 1,321  1,346 987 1,370 1,017 1,394 1,048 1,419 1,078 1,443 1,108 1,468 1,139 1,492},528},
// BonusSet{35.735, []int{ 1,532  1,560 1,370 1,589 1,394 1,617 1,419 1,645 1,443 1,674 1,468 1,702 1,492 1,730},528},
// BonusSet{38.798, []int{ 2,661  2,685 1,589 2,708 1,617 2,732 1,645 2,756 1,674 2,780 1,702 2,803 1,730 2,827},1135},
// BonusSet{41.861, []int{ 3,548  3,580 2,708 3,611 2,732 3,643 2,756 3,675 2,780 3,706 2,803 3,738 2,827 3,770},1135},
// }
// }

// kou_or_otsu kouは0、otsuは1にする
func CalcBonusTax(income int, kou_or_otsu int, support int, follow int) int {
	i := int(income / 1000)

	if support > 7 {
		support = 7
	}

	/** ここからは重い処理 **/
	l := NewBonusList()
	list := l[support]

	min := 0
	for _, set := range list.set {
		if i >= min && i < set.max {
			return CalcIntMulFloat(income, set.rate/100)
		}
		min = set.max
	}

	//ここを通ることはあり得ないので、0を返すことにする
	return 0
}
