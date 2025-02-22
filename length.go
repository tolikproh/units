package units

// Длина
type Length struct {
	Quantity
}

func NewLength(val uint64, pref Prefix) Quantiter {
	return &Length{NewQuantity(val, pref, LengthType)}
}

func (l *Length) Value() uint64 {
	return l.value * uint64(l.prefix)
}

func (l *Length) String() string {
	str, _ := FormatWithDecimals(l, l.Value(), l.prefix, l.decimals)
	return str
}

func (l *Length) SuffixNames(pref Prefix) (short string, full string) {
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

func (l *Length) ShortName(pref Prefix) string {
	name, _ := l.SuffixNames(pref)
	return name
}

func (l *Length) FullName(pref Prefix) string {
	_, fname := l.SuffixNames(pref)
	return fname
}

func (l *Length) Add(q Quantiter) Quantiter {
	if q == nil || q.(*Length).types != l.types {
		return nil
	}
	return &Length{
		Quantity{
			value:  l.value + q.Value()/uint64(l.prefix),
			prefix: l.prefix,
			types:  l.types,
		},
	}
}

func (l *Length) Sub(q Quantiter) Quantiter {
	if q == nil || q.(*Length).types != l.types {
		return nil
	}
	if l.Value() < q.Value() {
		return nil
	}
	return &Length{
		Quantity{
			value:  l.value - q.Value()/uint64(l.prefix),
			prefix: l.prefix,
			types:  l.types,
		},
	}
}

func (l *Length) SetPrefix(pref Prefix) {
	if pref == 0 {
		pref = Normal
	}
	if l.prefix != pref {
		l.value = l.Value() / pref.Uint()
		l.prefix = pref
	}
}

func (l *Length) SetDecimals(dec uint) {
	l.decimals = dec
}
