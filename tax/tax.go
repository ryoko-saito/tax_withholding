package tax

type TaxSet struct {
	index int   //(月収-社会保険)/1000
	kou   []int //甲
	otsu  int   //乙
}

func NewTaxList() []TaxSet {
	list := []TaxSet{
		TaxSet{89, []int{130, 0, 0, 0, 0, 0, 0, 0}, 3200},
		TaxSet{90, []int{180, 0, 0, 0, 0, 0, 0, 0}, 3200},
		TaxSet{91, []int{230, 0, 0, 0, 0, 0, 0, 0}, 3200},
		TaxSet{92, []int{290, 0, 0, 0, 0, 0, 0, 0}, 3200},
		TaxSet{93, []int{340, 0, 0, 0, 0, 0, 0, 0}, 3300},
	}
	return list
}
