package units

import "errors"

// Length представляет величину длины
type Length struct {
	Quantiter
}

// NewLength создает новую величину длины
func NewLength(val uint64, pref Prefix) *Length {
	return &Length{NewQuantity(val, pref, Nano, LengthType, LengthNames)}
}

// NewLengthJSON создает величину длины из JSON
func NewLengthJSON(data []byte) (*Length, error) {
	length := NewLength(0, 0)
	if err := length.UnmarshalJSON(data); err != nil {
		return nil, err
	}

	if length.Types() != LengthType {
		return nil, errors.New("new length json: unmarshal types is not length")
	}

	return length, nil
}

// LengthNames возвращает названия единиц измерения длины для разных префиксов
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

func (q *Length) Add(b Quantiter) *Length {
	return &Length{add(q, b)}
}

func (q *Length) Sub(b Quantiter) *Length {
	return &Length{sub(q, b)}
}

func (q *Length) Mul(b Quantiter) *Length {
	return &Length{mul(q, b)}
}

func (q *Length) Div(b Quantiter) *Length {
	return &Length{div(q, b)}
}
