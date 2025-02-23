package units

// Количество
type Things struct {
	Quantiter
}

func NewThings(val uint64, pref Prefix) *Things {
	return &Things{NewQuantity(val, pref, Normal, ThingsType, ThingsNames)}
}

func ThingsNames(pref Prefix) (short string, full string) {
	switch pref {
	case Normal:
		short = "шт"
		full = "штук"
	case Kilo:
		short = "тыс.шт"
		full = "тысяч штук"
	}
	return
}
