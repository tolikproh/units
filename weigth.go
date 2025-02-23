package units

// Вес
type Weigth struct {
	Quantiter
}

func NewWeigth(val uint64, pref Prefix) *Weigth {
	return &Weigth{NewQuantity(val, pref, Nano, WeigthType, WeigthNames)}
}

func WeigthNames(pref Prefix) (short string, full string) {
	switch pref {
	case Nano:
		short = "нгр"
		full = "нанограмм"
	case Micro:
		short = "μгр"
		full = "микрограмм"
	case Milli:
		short = "мгр"
		full = "миллиграмм"
	case Normal:
		short = "гр"
		full = "грамм"
	case Kilo:
		short = "кг"
		full = "килограмм"
	case Mega:
		short = "т"
		full = "тонна"
	case Giga:
		short = "Мт"
		full = "мегатонн"
	}
	return
}
