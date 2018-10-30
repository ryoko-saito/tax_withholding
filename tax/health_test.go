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

func TestCalcCareInsurance(t *testing.T) {
	//年齢が40未満
	age := 39
	f := 0
	i := CalcCareInsurance(age, f)
	if i != 0 {
		t.Error("failed")
	}

	age = 40
	i = CalcCareInsurance(age, f)
	if i != 3300 {
		t.Error("failed")
	}
}
