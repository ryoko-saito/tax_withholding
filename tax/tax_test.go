package tax

import (
	"testing"
)

type TestSet struct {
	income      int //所得
	kou_or_otsu int //甲か乙か？
	support     int //扶養人数
	follow      int //扶養の申請書提出の有無
	result      int //テスト実行時の正解の値
}

func NewTestSet() []TestSet {
	return []TestSet{
		TestSet{87000, 0, 0, 0, 0},    //88000円以下、甲で不要なし
		TestSet{87000, 1, 0, 0, 2664}, //88000円以下、乙で不要なし
		TestSet{200000, 0, 0, 0, 4770},
		TestSet{200000, 0, 5, 0, 0},
		TestSet{478000, 0, 5, 0, 7500},
		TestSet{820000, 0, 9, 0, 38870},
		TestSet{860000, 0, 3, 0, 75930},
		TestSet{860000, 1, 3, 0, 320900},
		TestSet{870000, 0, 2, 0, 84828},
		TestSet{870000, 1, 2, 0, 324984},
		TestSet{970000, 1, 8, 0, 364214},
		TestSet{1700000, 0, 8, 0, 320258},
		TestSet{3550000, 0, 4, 0, 1095390},
		TestSet{3550000, 1, 4, 0, 1512993},
		TestSet{843000, 1, 2, 1, 308380},
		TestSet{843000, 0, 2, 1, 79040},
	}
}

func TestCalcTax(t *testing.T) {

	sets := NewTestSet()
	for i, set := range sets {
		tax := CalcTax(set.income, set.kou_or_otsu, set.support, set.follow)
		if tax != set.result {
			t.Errorf("failed index:%d ok:%d prac:%d", i, set.result, tax)
		}
	}
}
