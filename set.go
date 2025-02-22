package units

// Комплект
type Set struct {
	Quantity
}

func NewSet(val uint64, pref Prefix) Quantiter {
	return &Set{NewQuantity(val, pref, SetType)}
}

func (s *Set) Value() uint64 {
	return uint64(s.value * uint64(s.prefix))
}

func (s *Set) String() string {
	str, _ := FormatWithDecimals(s, s.Value(), s.prefix, s.decimals)
	return str
}

func (s *Set) SuffixNames(pref Prefix) (short string, full string) {
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
func (s *Set) ShortName(pref Prefix) string {
	name, _ := s.SuffixNames(pref)
	return name
}
func (s *Set) FullName(pref Prefix) string {
	_, fname := s.SuffixNames(pref)
	return fname
}

func (s *Set) Add(q Quantiter) Quantiter {
	if q == nil || q.(*Set).types != s.types {
		return nil
	}
	return &Set{
		Quantity{
			value:  s.value + q.Value()/uint64(s.prefix),
			prefix: s.prefix,
			types:  s.types,
		},
	}
}

func (s *Set) Sub(q Quantiter) Quantiter {
	if q == nil || q.(*Set).types != s.types {
		return nil
	}
	if s.Value() < q.Value() {
		return nil
	}
	return &Set{
		Quantity{
			value:  s.value - q.Value()/uint64(s.prefix),
			prefix: s.prefix,
			types:  s.types,
		},
	}
}

func (s *Set) SetPrefix(pref Prefix) {
	if pref == 0 {
		pref = Normal
	}
	if s.prefix != pref {
		s.value = s.Value() / pref.Uint()
		s.prefix = pref
	}
}

func (s *Set) SetDecimals(dec uint) {
	s.decimals = dec
}
