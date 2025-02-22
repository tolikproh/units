package units

// Вес
type Weigth struct {
	Quantity
}

func NewWeigth(val uint64, pref Prefix) Quantiter {
	return &Weigth{NewQuantity(val, pref, WeigthType)}
}

func (w *Weigth) Value() uint64 {
	return uint64(w.value * uint64(w.prefix))
}

func (w *Weigth) String() string {
	str, _ := FormatWithDecimals(w, w.Value(), w.prefix, w.decimals)
	return str
}

func (w *Weigth) SuffixNames(pref Prefix) (short string, full string) {
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
func (w *Weigth) ShortName(pref Prefix) string {
	name, _ := w.SuffixNames(pref)
	return name
}
func (w *Weigth) FullName(pref Prefix) string {
	_, fname := w.SuffixNames(pref)
	return fname
}

func (w *Weigth) Add(q Quantiter) Quantiter {
	if q == nil || q.(*Weigth).types != w.types {
		return nil
	}
	return &Weigth{
		Quantity{
			value:  w.value + q.Value()/uint64(w.prefix),
			prefix: w.prefix,
			types:  w.types,
		},
	}
}

func (w *Weigth) Sub(q Quantiter) Quantiter {
	if q == nil || q.(*Weigth).types != w.types {
		return nil
	}
	if w.Value() < q.Value() {
		return nil
	}
	return &Weigth{
		Quantity{
			value:  w.value - q.Value()/uint64(w.prefix),
			prefix: w.prefix,
			types:  w.types,
		},
	}
}

func (w *Weigth) SetPrefix(pref Prefix) {
	if pref == 0 {
		pref = Normal
	}
	if w.prefix != pref {
		w.value = w.Value() / pref.Uint()
		w.prefix = pref
	}
}

func (w *Weigth) SetDecimals(dec uint) {
	w.decimals = dec
}
