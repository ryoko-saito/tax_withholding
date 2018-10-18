package tax

type TaxSet struct {
	index int
	tax   int
}

//450万以上の税金額
const maxTax = 21800

func NewList() []TaxSet {
	list := []TaxSet{
		TaxSet{100, 6900},
		TaxSet{150, 7800},
		TaxSet{200, 10300},
		TaxSet{250, 11800},
		TaxSet{300, 13300},
		TaxSet{350, 14800},
		TaxSet{400, 16800},
		TaxSet{450, 18800},
	}
	return list
}

//1000000
func Calc(s int, age int) int {
	ss := int(s / 10000)
	list := NewList()

	var total int

	//40歳以上
	if age >= 40 {
		total = 3300
	}

	min := 0
	for _, set := range list {
		if ss >= min && ss < set.index {
			total += set.tax + 3200
			return total
		}
		min = set.index
	}
	return total + maxTax + 3200
}
