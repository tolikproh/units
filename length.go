package units

// Длина
type Length struct {
	Quantiter
}

func NewLength(val uint64, pref Prefix) *Length {
	return &Length{NewQuantity(val, pref, Nano, LengthType, LengthNames)}
}

func LengthNames(pref Prefix) (short string, full string) {
	switch pref {
	case Nano:
		short = "нм"
		full = "нанометр"
	case Micro:
		short = "μм"
		full = "микрометр"
	case Milli:
		short = "мм"
		full = "миллиметр"
	case Normal:
		short = "м"
		full = "метр"
	case Kilo:
		short = "км"
		full = "километр"
	case Mega:
		short = "тыс.км"
		full = "тысяч киллометров"
	}
	return
}
