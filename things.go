package units

import "errors"

// Things представляет количество штук
type Things struct {
	Quantiter
}

// NewThings создает новую величину в штуках
func NewThings(val uint64, pref Prefix) *Things {
	return &Things{NewQuantity(val, pref, Normal, ThingsType, ThingsNames)}
}

// NewThingsJSON создает величину в штуках из JSON
func NewThingsJSON(data []byte) (*Things, error) {
	things := NewThings(0, 0)
	if err := things.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	if things.Types() != ThingsType {
		return nil, errors.New("new things json: unmarshal types is not things")
	}
	return things, nil
}

// ThingsNames возвращает названия единиц измерения штук для разных префиксов
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

func (q *Things) Add(b Quantiter) *Things {
	return &Things{add(q, b)}
}

func (q *Things) Sub(b Quantiter) *Things {
	return &Things{sub(q, b)}
}

func (q *Things) Mul(b Quantiter) *Things {
	return &Things{mul(q, b)}
}

func (q *Things) Div(b Quantiter) *Things {
	return &Things{div(q, b)}
}
