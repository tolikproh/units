package units

// Упаковка
type Package struct {
	Quantiter
}

func NewPackage(val uint64, pref Prefix) *Package {
	return &Package{NewQuantity(val, pref, Normal, PackageType, PackageNames)}
}

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
