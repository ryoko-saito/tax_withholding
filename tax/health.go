package tax

type HealthSet struct {
	index  int //2年前所得の上限値
	amount int //保険料
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

//健康保険 fは家族の人数,bは2年前所得
func CalcHealthInsurance(b int, f int) int {
	ss := int(b / 10000)
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
		totalAmount := 5500 + int(float64(b-330000)*0.0055/10)*10 //1桁目は必ず0にするため、合算を÷10してから×10で戻す
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

//介護保険、oは40歳以上64歳未満の扶養人数
func CalcCareInsurance(o int) int {
	s := 3300 * o
	if s > 13200 {
		s = 13200
	}
	return s
}
