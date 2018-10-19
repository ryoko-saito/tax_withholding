package tax

import (
	"testing"
)

func TestCalcHealthInsurance(t *testing.T) {
	//1200000円の場合、源泉徴収は7800 + 3200円
	s := 1200000
	tax := CalcHealthInsurance(s)
	if tax != 11000 {
		t.Error("failed")
	}
	//2100000円の場合、11800 + 3200
	s = 2100000
	tax = CalcHealthInsurance(s)
	if tax != 15000 {
		t.Error("failed")
	}

	//収入が4500000円以上 21800 + 3200
	s = 4500000
	tax = CalcHealthInsurance(s)
	if tax != 25000 {
		t.Error("failed")
	}
}

func TestCalcCareInsurance(t *testing.T) {
	//年齢が40未満
	age := 39
	i := CalcCareInsurance(age)
	if i != 0 {
		t.Error("failed")
	}

	age = 40
	i = CalcCareInsurance(age)
	if i != 3300 {
		t.Error("failed")
	}
}
