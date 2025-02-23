package units

import (
	"fmt"
	"strconv"
)

// Typed unit
type UnitType int

const (
	LengthType UnitType = iota
	WeigthType
	ThingsType
	SetType
	PackageType
	BayType
)

// Quantity
type Quantiter interface {
	Value() uint64
	String() string
	ShortName(pref Prefix) string
	FullName(pref Prefix) string
	SetPrefix(pref Prefix)
	GetPrefix() Prefix
	SetDecimals(dec int)
	Types() UnitType
	Ok() bool
	Add(qa Quantiter) Quantiter
	Sub(qa Quantiter) Quantiter
	Mul(qa Quantiter) Quantiter
	Div(qa Quantiter) Quantiter
}

type Quantity struct {
	value    uint64
	prefix   Prefix
	divisor  Prefix
	decimals int
	unit     func(p Prefix) (string, string)
	types    UnitType
	negative bool
	ok       bool
}

func NewQuantity(val uint64, pref, div Prefix, types UnitType, unit func(p Prefix) (string, string)) *Quantity {
	if pref == 0 || pref < div {
		pref = div
	}
	val = val * pref.Uint()
	return &Quantity{value: val, prefix: pref, divisor: div, types: types, unit: unit}
}

func (q *Quantity) Value() uint64 {
	return q.value / q.divisor.Uint()
}

func (q *Quantity) String() string {
	sname := q.ShortName(q.prefix)
	if sname == "" {
		return ""
	}

	quotient := q.Value() / q.prefix.Uint()
	remainder := q.Value() % q.prefix.Uint()

	// Подготавливаем дробную часть
	decimalStr := fmt.Sprintf("%d", remainder)

	// Добавляем ведущие нули если нужно
	for len(decimalStr) < 3 {
		decimalStr = "0" + decimalStr
	}

	// Если нужно округление
	if q.decimals < len(decimalStr) {
		// Получаем следующую цифру после места округления
		nextDigit := 0
		if q.decimals < len(decimalStr) {
			nextDigit = int(decimalStr[q.decimals] - '0')
		}

		// Округляем если следующая цифра >= 5
		if nextDigit >= 5 {
			// Конвертируем обрезанную часть в число для округления
			numStr := decimalStr[:q.decimals]
			num, _ := strconv.ParseUint(numStr, 10, 64)
			num++ // увеличиваем на 1 для округления

			// Проверяем переполнение
			if len(fmt.Sprint(num)) > int(q.decimals) {
				quotient++
				num = 0
			}

			// Форматируем обратно в строку с ведущими нулями
			decimalStr = fmt.Sprintf("%0*d", q.decimals, num)
		} else {
			decimalStr = decimalStr[:q.decimals]
		}
	}

	negativ := ""
	if q.negative {
		negativ = "-"
	}
	if q.decimals == 0 {
		return fmt.Sprintf("%s%d %s", negativ, quotient, sname)
	}
	return fmt.Sprintf("%s%d.%s %s", negativ, quotient, decimalStr, sname)
}

func (q *Quantity) ShortName(pref Prefix) string {
	name, _ := q.unit(pref)
	return name
}

func (q *Quantity) FullName(pref Prefix) string {
	_, fname := q.unit(pref)
	return fname
}

func (q *Quantity) SetPrefix(pref Prefix) {
	if pref == 0 || pref < q.divisor {
		pref = q.divisor
	}
	if q.prefix != pref {
		q.prefix = pref
	}
}

func (q *Quantity) GetPrefix() Prefix {
	return q.prefix
}

func (q *Quantity) SetDecimals(dec int) {
	if dec < 0 {
		dec = 0
	}
	q.decimals = dec
}

func (q *Quantity) Types() UnitType {
	return q.types
}

func (q *Quantity) Ok() bool {
	return q.ok
}

func (q *Quantity) Add(qa Quantiter) Quantiter {
	if qa != nil && q.Types() == qa.Types() {
		q.value = q.Value() + qa.Value()
		q.ok = true
	} else {
		q.ok = false
	}

	return q
}

func (q *Quantity) Sub(qa Quantiter) Quantiter {
	if qa != nil && q.Types() == qa.Types() {
		if qa.Value() > q.Value() {
			q.value = qa.Value() - q.Value()
			q.negative = true
		} else {
			q.value = q.Value() - qa.Value()
		}
		q.ok = true
	} else {
		q.ok = false
	}

	return q
}

func (q *Quantity) Mul(qa Quantiter) Quantiter {
	if qa != nil {
		q.value = q.Value() * qa.Value()
		q.ok = true
	} else {
		q.ok = false
	}

	return q
}

func (q *Quantity) Div(qa Quantiter) Quantiter {
	if qa != nil && qa.Value() > 0 {

		q.value = q.Value() * qa.GetPrefix().Uint() / qa.Value()
		q.ok = true
	} else {
		q.ok = false
	}

	return q
}
