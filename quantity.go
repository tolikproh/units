package units

// Typed unit
type UnitType int

const (
	LengthType UnitType = iota
	WeigthType
	ThingsType
	SetType
	PackageType
	BayType
)

// Quantity
type Quantiter interface {
	Value() uint64
	String() string
	ShortName(pref Prefix) string
	FullName(pref Prefix) string
	Add(qa Quantiter) Quantiter
	Sub(qa Quantiter) Quantiter
	Mul(qa Quantiter) Quantiter
	Div(qa Quantiter) Quantiter
	SetPrefix(pref Prefix)
	SetDecimals(dec uint)
}

type Quantity struct {
	value    uint64
	prefix   Prefix
	divisor  Prefix
	decimals uint
	unit     func(p Prefix) (string, string)
	types    UnitType
}

func NewQuantity(val uint64, pref, div Prefix, types UnitType, unit func(p Prefix) (string, string)) Quantiter {
	if pref == 0 {
		pref = Normal
	}
	val = val * pref.Uint() / div.Uint()
	return &Quantity{value: val, prefix: pref, divisor: div, types: types, unit: unit}
}

func (q *Quantity) Value() uint64 {
	return q.value
}

func (q *Quantity) String() string {
	str, _ := FormatWithDecimals(q, q.Value(), q.prefix, q.decimals)
	return str
}

func (q *Quantity) ShortName(pref Prefix) string {
	name, _ := q.unit(pref)
	return name
}

func (q *Quantity) FullName(pref Prefix) string {
	_, fname := q.unit(pref)
	return fname
}

func (q *Quantity) Add(qa Quantiter) Quantiter {
	if qa == nil {
		return nil
	}
	return &Quantity{
		value:   q.Value() + qa.Value(),
		prefix:  q.prefix,
		divisor: q.divisor,
		types:   q.types,
		unit:    q.unit,
	}
}

func (q *Quantity) Sub(qa Quantiter) Quantiter {
	if qa == nil {
		return nil
	}
	if q.Value() < qa.Value() {
		return nil
	}
	return &Quantity{
		value:   q.Value() - qa.Value(),
		prefix:  q.prefix,
		divisor: q.divisor,
		types:   q.types,
		unit:    q.unit,
	}
}

func (q *Quantity) Mul(qa Quantiter) Quantiter {
	if qa == nil {
		return nil
	}
	return &Quantity{
		value:   q.Value() * qa.Value(),
		prefix:  q.prefix,
		divisor: q.divisor,
		types:   q.types,
		unit:    q.unit,
	}
}

func (q *Quantity) Div(qa Quantiter) Quantiter {
	if qa == nil || qa.Value() == 0 {
		return nil
	}
	return &Quantity{
		value:   q.Value() / qa.Value(),
		prefix:  q.prefix,
		divisor: q.divisor,
		types:   q.types,
		unit:    q.unit,
	}
}

func (q *Quantity) SetPrefix(pref Prefix) {
	if pref == 0 || pref < q.divisor {
		pref = q.divisor
	}
	if q.prefix != pref {
		q.value = q.Value() * pref.Uint() / q.prefix.Uint()
		q.prefix = pref
	}
}

func (q *Quantity) SetDecimals(dec uint) {
	q.decimals = dec
}
