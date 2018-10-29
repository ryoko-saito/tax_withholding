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

//健康保険 sは収入、fは家族の人数,bは2年前所得
func CalcHealthInsurance(s int, f int, b int) int {
	ss := int(s / 10000)
	if f == 0 {
		list := NewHealthList()

		min := 0
		for _, set := range list {
			if ss >= min && ss < set.index {
				return set.amount + 3200
			}
			min = set.index
		}
		return max + 3200
	} else if f > 0 {
		totalAmount := 5500 + int(float64(b/10000-33)*0.0055/10)*10 //1桁目は必ず0にするため、合算を÷10してから×10で戻す
		totalAmount += f * 200

		//合計が45000円以上にさせたくない
		if totalAmount > 45000 {
			totalAmount = 45000
		}
		h := f * 3200 //hは後期高齢者支援金額
		if h > 12800 {
			h = 12800
		}
		return totalAmount + h
	}

	return 0
}

//介護保険
func CalcCareInsurance(age int, f int) int {
	if age >= 40 && f == 0 {
		return 3300
	} else {
		s := 3300 * f
		if s > 13200 {
			return 13200
		}
		return s
	}
}
