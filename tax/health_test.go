package tax

import (
	"testing"
)

type HealthTestSet struct {
	before int //2年前所得
	family int //家族の人数
	result int //テスト実行時の正解の値
}

func NewHealthTestSet() []HealthTestSet {
	return []HealthTestSet{
		HealthTestSet{1200000, 4, 23880}, //1200000円の場合、源泉徴収は7800 + 3200円
		HealthTestSet{2100000, 3, 25430},
		HealthTestSet{2500000, 0, 16500},
	}
}

func TestCalcHealthInsurance(t *testing.T) {
	sets := NewHealthTestSet()
	for i, set := range sets {
		ins := CalcHealthInsurance(set.before, set.family)
		if ins != set.result {
			t.Errorf("failed index:%d ok:%d prac:%d", i, set.result, ins)
		}
	}
}

func NewCareSet() []int {
	//インデックスに人数、値は正解の金額
	return []int{0, 3300, 6600, 9900, 13200, 13200, 13200}
}

func TestCalcCareInsurance(t *testing.T) {
	sets := NewCareSet()
	for i, res := range sets {
		ins := CalcCareInsurance(i)
		if ins != res {
			t.Errorf("failed index:%d ok:%d prac:%d", i, res, ins)
		}
	}
}
