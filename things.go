package units

// Количество
type Things struct {
	Quantity
}

func NewThings(val uint64, pref Prefix) Quantiter {
	return &Things{NewQuantity(val, pref, ThingsType)}
}

func (t *Things) Value() uint64 {
	return uint64(t.value * uint64(t.prefix))
}

func (t *Things) String() string {
	str, _ := FormatWithDecimals(t, t.Value(), t.prefix, uint(t.prefix))
	return str
}

func (t *Things) SuffixNames(pref Prefix) (short string, full string) {
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
func (t *Things) ShortName(pref Prefix) string {
	name, _ := t.SuffixNames(pref)
	return name
}
func (t *Things) FullName(pref Prefix) string {
	_, fname := t.SuffixNames(pref)
	return fname
}

func (t *Things) Add(q Quantiter) Quantiter {
	if q == nil || q.(*Things).types != t.types {
		return nil
	}
	return &Things{
		Quantity{
			value:  t.value + q.Value()/uint64(t.prefix),
			prefix: t.prefix,
			types:  t.types,
		},
	}
}

func (t *Things) Sub(q Quantiter) Quantiter {
	if q == nil || q.(*Things).types != t.types {
		return nil
	}
	if t.Value() < q.Value() {
		return nil
	}
	return &Things{
		Quantity{
			value:  t.value - q.Value()/uint64(t.prefix),
			prefix: t.prefix,
			types:  t.types,
		},
	}
}

func (t *Things) SetPrefix(pref Prefix) {
	if pref == 0 {
		pref = Normal
	}
	if t.prefix != pref {
		t.value = t.Value() / pref.Uint()
		t.prefix = pref
	}
}

func (t *Things) SetDecimals(dec uint) {
	t.decimals = dec
}
