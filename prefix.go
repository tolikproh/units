package units

// Prefix
type Prefix uint64

const (
	Nano   Prefix = 1
	Micro  Prefix = 1000
	Milli  Prefix = 1000 * Micro
	Normal Prefix = 1000 * Milli
	Kilo   Prefix = 1000 * Normal
	Mega   Prefix = 1000 * Kilo
	Giga   Prefix = 1000 * Mega
)

func (p Prefix) Uint() uint64 {
	return uint64(p)
}

func (p Prefix) Name() (name string) {
	switch p {
	case Nano:
		name = "Nano"
	case Micro:
		name = "Micro"
	case Milli:
		name = "Milli"
	case Normal:
		name = "Normal"
	case Kilo:
		name = "Kilo"
	case Mega:
		name = "Mega"
	case Giga:
		name = "Giga"
	default:
		name = "empty"
	}
	return
}
