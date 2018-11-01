package tax

import "testing"

type BonusTestSet struct {
	income      int //前月給与から社会保険を控除した所得
	kou_or_otsu int //甲か乙か
	support     int //扶養
	follow      int //扶養の申請書を出しているか？
	result      int //テスト実行時の正解の値
}

func NewBonusTestSet() []BonusTestSet {
	return []BonusTestSet{
		BonusTestSet{90000, 0, 0, 0, 3675},
		BonusTestSet{90000, 0, 1, 0, 0},
		BonusTestSet{300000, 0, 3, 0, 12252},
		BonusTestSet{300000, 1, 0, 0, 91890},
	}
}

func TestCalcBonusTax(t *testing.T) {

	sets := NewBonusTestSet()
	for i, set := range sets {
		tax := CalcBonusTax(set.income, set.kou_or_otsu, set.support, set.follow)
		if tax != set.result {
			t.Errorf("failed index:%d ok:%d prac:%d", i, set.result, tax)
		}
	}
}
