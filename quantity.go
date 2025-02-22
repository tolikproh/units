package units

// Quantity
type Quantiter interface {
	Value() uint64
	String() string
	SuffixNames(pref Prefix) (short string, full string)
	ShortName(pref Prefix) string
	FullName(pref Prefix) string
	Add(q Quantiter) Quantiter
	Sub(q Quantiter) Quantiter
	SetPrefix(pref Prefix)
	SetDecimals(dec uint)
}

type Quantity struct {
	value    uint64
	prefix   Prefix
	types    UnitType
	decimals uint
}

func NewQuantity(val uint64, pref Prefix, types UnitType) Quantity {
	if pref == 0 {
		pref = Normal
	}
	return Quantity{value: val, prefix: pref, types: types}
}
