package tax

import (
	"testing"
)

func TestCalc(t *testing.T) {
	//1200000円の場合、源泉徴収は7800 + 3200円
	age := 39
	s := 1200000
	tax := Calc(s, age)
	if tax != 11000 {
		t.Error("failed")
	}

	//40歳以上の加算が入る +3300
	age = 40
	tax = Calc(s, age)
	if tax != 14300 {
		t.Error("failed")
	}

	//2100000円の場合、11800 + 3200
	age = 39
	s = 2100000
	tax = Calc(s, age)
	if tax != 15000 {
		t.Error("failed")
	}

	age = 40
	tax = Calc(s, age)
	if tax != 18300 {
		t.Error("failed")
	}
}
