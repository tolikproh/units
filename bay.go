package units

// Бухта
type Bay struct {
	Quantity
}

func NewBay(val uint64, pref Prefix) Quantiter {
	return &Bay{NewQuantity(val, pref, BayType)}
}

func (b *Bay) Value() uint64 {
	return b.value * uint64(b.prefix)
}

func (b *Bay) String() string {
	str, _ := FormatWithDecimals(b, b.Value(), b.prefix, b.decimals)
	return str
}

func (b *Bay) SuffixNames(pref Prefix) (short string, full string) {
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

func (b *Bay) ShortName(pref Prefix) string {
	name, _ := b.SuffixNames(pref)
	return name
}

func (b *Bay) FullName(pref Prefix) string {
	_, fname := b.SuffixNames(pref)
	return fname
}

func (b *Bay) Add(q Quantiter) Quantiter {
	if q == nil || q.(*Bay).types != b.types {
		return nil
	}
	return &Bay{
		Quantity{
			value:  b.value + q.Value()/uint64(b.prefix),
			prefix: b.prefix,
			types:  b.types,
		},
	}
}

func (b *Bay) Sub(q Quantiter) Quantiter {
	if q == nil || q.(*Bay).types != b.types {
		return nil
	}
	if b.Value() < q.Value() {
		return nil
	}
	return &Bay{
		Quantity{
			value:  b.value - q.Value()/uint64(b.prefix),
			prefix: b.prefix,
			types:  b.types,
		},
	}
}

func (b *Bay) SetPrefix(pref Prefix) {
	if pref == 0 {
		pref = Normal
	}
	if b.prefix != pref {
		b.value = b.Value() / pref.Uint()
		b.prefix = pref
	}
}

func (b *Bay) SetDecimals(dec uint) {
	b.decimals = dec
}
