package units

import "errors"

// Вес
// Weigth представляет величину веса
type Weigth struct {
	Quantiter
}

// NewWeigth создает новую величину веса
func NewWeigth(val uint64, pref Prefix) *Weigth {
	return &Weigth{NewQuantity(val, pref, Nano, WeigthType, WeigthNames)}
}

// NewWeigthJSON создает величину веса из JSON
func NewWeigthJSON(data []byte) (*Weigth, error) {
	weigth := NewWeigth(0, 0)
	if err := weigth.UnmarshalJSON(data); err != nil {
		return nil, err
	}

	if weigth.Types() != WeigthType {
		return nil, errors.New("new weigth json: unmarshal types is not weigth")
	}

	return weigth, nil
}

// WeigthNames возвращает названия единиц измерения веса для разных префиксов
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

func (q *Weigth) Add(b Quantiter) *Weigth {
	return &Weigth{add(q, b)}
}

func (q *Weigth) Sub(b Quantiter) *Weigth {
	return &Weigth{sub(q, b)}
}

func (q *Weigth) Mul(b Quantiter) *Weigth {
	return &Weigth{mul(q, b)}
}

func (q *Weigth) Div(b Quantiter) *Weigth {
	return &Weigth{div(q, b)}
}
