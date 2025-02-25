package units

import "errors"

// Set представляет величину в комплектах
type Set struct {
	Quantiter
}

// NewSet создает новую величину в комплектах
func NewSet(val uint64, pref Prefix) *Set {
	return &Set{NewQuantity(val, pref, Normal, SetType, SetNames)}
}

// NewSetJSON создает величину в комплектах из JSON
func NewSetJSON(data []byte) (*Set, error) {
	set := NewSet(0, 0)
	if err := set.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	if set.Types() != SetType {
		return nil, errors.New("new set json: unmarshal types is not set")
	}
	return set, nil
}

// SetNames возвращает названия единиц измерения комплектов для разных префиксов
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
