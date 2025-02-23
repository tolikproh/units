package units

// Комплект
type Set struct {
	Quantiter
}

func NewSet(val uint64, pref Prefix) *Set {
	return &Set{NewQuantity(val, pref, Normal, SetType, SetNames)}
}

func SetNames(pref Prefix) (short string, full string) {
	switch pref {
	case Normal:
		short = "компл"
		full = "комплект"
	case Kilo:
		short = "тыс.компл"
		full = "тысяч комплектов"
	}
	return
}
