package units

import "errors"

// Package представляет величину в упаковках
type Package struct {
	Quantiter
}

// NewPackage создает новую величину в упаковках
func NewPackage(val uint64, pref Prefix) *Package {
	return &Package{NewQuantity(val, pref, Normal, PackageType, PackageNames)}
}

// NewPackageJSON создает величину в упаковках из JSON
func NewPackageJSON(data []byte) (*Package, error) {
	pack := NewPackage(0, 0)
	if err := pack.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	if pack.Types() != PackageType {
		return nil, errors.New("new package json: unmarshal types is not package")
	}
	return pack, nil
}

// PackageNames возвращает названия единиц измерения упаковок для разных префиксов
func PackageNames(pref Prefix) (short string, full string) {
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

func (q *Package) Add(b Quantiter) *Package {
	return &Package{add(q, b)}
}

func (q *Package) Sub(b Quantiter) *Package {
	return &Package{sub(q, b)}
}

func (q *Package) Mul(b Quantiter) *Package {
	return &Package{mul(q, b)}
}

func (q *Package) Div(b Quantiter) *Package {
	return &Package{div(q, b)}
}
