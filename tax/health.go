package tax

type HealthSet struct {
	index  int
	amount int
}

//450万以上の税金額
const max = 21800

//配列の中に構造体が入っている
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

//健康保険
func CalcHealthInsurance(s int) int {
	ss := int(s / 10000)
	list := NewHealthList()

	min := 0
	for _, set := range list {
		if ss >= min && ss < set.index {
			return set.amount + 3200
		}
		min = set.index
	}
	return max + 3200
}

//介護保険
func CalcCareInsurance(age int) int {
	if age >= 40 {
		return 3300
	} else {
		return 0
	}
}
