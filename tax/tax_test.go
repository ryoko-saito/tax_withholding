package tax

import (
	"testing"
)

func TestCalcTax(t *testing.T) {
	income := 200000
	kou_or_otsu := 0
	support := 0

	tax := CalcTax(income, kou_or_otsu, support)
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
}
