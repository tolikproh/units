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
	SetPrefix(pref Prefix)
	SetDecimals(dec int)
	Types() UnitType
	Add(qa Quantiter) Quantiter
	Sub(qa Quantiter) Quantiter
	Mul(qa Quantiter) Quantiter
	Div(qa Quantiter) Quantiter
}

type Quantity struct {
	value    uint64
	prefix   Prefix
	divisor  Prefix
	decimals int
	unit     func(p Prefix) (string, string)
	types    UnitType
}

func NewQuantity(val uint64, pref, div Prefix, types UnitType, unit func(p Prefix) (string, string)) *Quantity {
	if pref == 0 || pref < div {
		pref = div
	}
	val = val * pref.Uint()
	return &Quantity{value: val, prefix: pref, divisor: div, types: types, unit: unit}
}

func (q *Quantity) Value() uint64 {
	return q.value / q.divisor.Uint()
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

func (q *Quantity) SetPrefix(pref Prefix) {
	if pref == 0 || pref < q.divisor {
		pref = q.divisor
	}
	if q.prefix != pref {
		q.prefix = pref
	}
}

func (q *Quantity) SetDecimals(dec int) {
	if dec < 0 {
		dec = 0
	}
	q.decimals = dec
}

func (q *Quantity) Types() UnitType {
	return q.types
}

func (q *Quantity) Add(qa Quantiter) Quantiter {
	if qa != nil && q.Types() == qa.Types() {
		q.value = q.Value() + qa.Value()
	}

	return q
}

func (q *Quantity) Sub(qa Quantiter) Quantiter {
	if qa != nil && q.Types() == qa.Types() {
		if qa.Value() > q.Value() {
			q.value = qa.Value() - q.Value()
		} else {
			q.value = q.Value() - qa.Value()
		}
	}

	return q
}

func (q *Quantity) Mul(qa Quantiter) Quantiter {
	if qa != nil {
		q.value = q.Value() * qa.Value()
	}

	return q
}

func (q *Quantity) Div(qa Quantiter) Quantiter {
	if qa != nil && qa.Value() > 0 {
		q.value = q.Value() / qa.Value()
	}

	return q
}
