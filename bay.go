package units

import "errors"

// Bay представляет величину в бухтах
type Bay struct {
	Quantiter
}

// NewBay создает новую величину в бухтах
func NewBay(val uint64, pref Prefix) *Bay {
	return &Bay{NewQuantity(val, pref, Normal, BayType, BayNames)}
}

// NewBayJSON создает величину в бухтах из JSON
func NewBayJSON(data []byte) (*Bay, error) {
	bay := NewBay(0, 0)
	if err := bay.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	if bay.Types() != BayType {
		return nil, errors.New("new bay json: unmarshal types is not bay")
	}
	return bay, nil
}

// BayNames возвращает названия единиц измерения бухт для разных префиксов
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
