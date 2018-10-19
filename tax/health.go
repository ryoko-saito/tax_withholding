package tax

type HealthSet struct {
	index  int
	amount int
}

//450万以上の税金額
const max = 21800

func NewHealthList() []HealthSet {
	list := []HealthSet{
		HealthSet{100, 6900},
		HealthSet{150, 7800},
		HealthSet{200, 10300},
		HealthSet{250, 11800},
		HealthSet{300, 13300},
		HealthSet{350, 14800},
		HealthSet{400, 16800},
		HealthSet{450, 18800},
	}
	return list
}

//1000000
func Calc(s int, age int) int {
	ss := int(s / 10000)
	list := NewHealthList()

	var total int

	//40歳以上
	if age >= 40 {
		total = 3300
	}

	min := 0
	for _, set := range list {
		if ss >= min && ss < set.index {
			total += set.amount + 3200
			return total
		}
		min = set.index
	}
	return total + max + 3200
}
