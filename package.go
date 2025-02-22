package units

// Упаковка
type Package struct {
	Quantity
}

func NewPackage(val uint64, pref Prefix) Quantiter {
	return &Package{NewQuantity(val, pref, PackageType)}
}

func (p *Package) Value() uint64 {
	return uint64(p.value * uint64(p.prefix))
}

func (p *Package) String() string {
	str, _ := FormatWithDecimals(p, p.Value(), p.prefix, p.decimals)
	return str
}

func (p *Package) SuffixNames(pref Prefix) (short string, full string) {
	switch pref {
	case Normal:
		short = "упак"
		full = "упаковка"
	case Kilo:
		short = "тыс.упак"
		full = "тысяч упаковок"
	}
	return
}
func (p *Package) ShortName(pref Prefix) string {
	name, _ := p.SuffixNames(pref)
	return name
}
func (p *Package) FullName(pref Prefix) string {
	_, fname := p.SuffixNames(pref)
	return fname
}

func (p *Package) Add(q Quantiter) Quantiter {
	if q == nil || q.(*Package).types != p.types {
		return nil
	}
	return &Package{
		Quantity{
			value:  p.value + q.Value()/uint64(p.prefix),
			prefix: p.prefix,
			types:  p.types,
		},
	}
}

func (p *Package) Sub(q Quantiter) Quantiter {
	if q == nil || q.(*Package).types != p.types {
		return nil
	}
	if p.Value() < q.Value() {
		return nil
	}
	return &Package{
		Quantity{
			value:  p.value - q.Value()/uint64(p.prefix),
			prefix: p.prefix,
			types:  p.types,
		},
	}
}

func (p *Package) SetPrefix(pref Prefix) {
	if pref == 0 {
		pref = Normal
	}
	if p.prefix != pref {
		p.value = p.Value() / pref.Uint()
		p.prefix = pref
	}
}

func (p *Package) SetDecimals(dec uint) {
	p.decimals = dec
}
