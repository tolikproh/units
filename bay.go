package units

// Бухта
type Bay struct {
	Quantiter
}

func NewBay(val uint64, pref Prefix) *Bay {
	return &Bay{NewQuantity(val, pref, Normal, LengthType, BayNames)}
}

func BayNames(pref Prefix) (short string, full string) {
	switch pref {
	case Normal:
		short = "бух"
		full = "бухта"
	case Kilo:
		short = "тыс.бух"
		full = "тысяч бухт"
	}
	return
}
