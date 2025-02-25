package units

// Prefix определяет множитель единицы измерения
type Prefix uint64

// Константы для стандартных префиксов СИ
const (
	Nano   Prefix = 1             // 10^-9
	Micro  Prefix = 1000          // 10^-6
	Milli  Prefix = 1000 * Micro  // 10^-3
	Normal Prefix = 1000 * Milli  // 10^0
	Kilo   Prefix = 1000 * Normal // 10^3
	Mega   Prefix = 1000 * Kilo   // 10^6
	Giga   Prefix = 1000 * Mega   // 10^9
)

// Uint возвращает значение префикса как uint64
func (p Prefix) Uint() uint64 {
	return uint64(p)
}

// Name возвращает строковое представление префикса
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
