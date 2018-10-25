package tax

import (
	"testing"
)

func TestCalcTax(t *testing.T) {
	//88000円以下
	income := 87000
	kou_or_otsu := 0
	tax := CalcTax(income, kou_or_otsu, 0)
	if tax != 0 {
		t.Error("failed")
	}

	//乙の場合は3.063％
	kou_or_otsu = 1
	tax = CalcTax(income, kou_or_otsu, 0)
	if tax != int(float64(income)*0.03063) {
		t.Error("failed")
	}

	income = 200000
	kou_or_otsu = 0
	support := 0

	tax = CalcTax(income, kou_or_otsu, support)
	if tax != 4770 {
		t.Error("failed")
	}

	support = 5
	tax = CalcTax(income, kou_or_otsu, support)
	if tax != 0 {
		t.Error("failed")
	}

	income = 478000
	kou_or_otsu = 0
	support = 5
	tax = CalcTax(income, kou_or_otsu, support)
	if tax != 7500 {
		t.Error("failed")
	}

	//860000円の時
	income = 860000
	kou_or_otsu = 0
	support = 3
	tax = CalcTax(income, kou_or_otsu, support)
	if tax != 75930 {
		t.Error("failed")
	}

	//乙
	kou_or_otsu = 1
	tax = CalcTax(income, kou_or_otsu, support)
	if tax != 320900 {
		t.Error("failed")
	}

	//甲
	income = 870000
	kou_or_otsu = 0
	support = 2
	tax = CalcTax(income, kou_or_otsu, support)
	if tax != 84828 {
		t.Error("failed")
	}

	//乙
	kou_or_otsu = 1
	tax = CalcTax(income, kou_or_otsu, support)
	if tax != 324984 {
		t.Error("failed")
	}
}
