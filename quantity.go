// Package units предоставляет типы и функции для работы с различными единицами измерения
package units

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

// Quantiter определяет интерфейс для работы с величинами
type Quantiter interface {
	// Copy создает новый метод и делает его копию
	Copy() *Quantity
	// Value возвращает значение в базовых единицах
	Value() uint64
	// String возвращает отформатированную строку с единицей измерения
	String() string
	// MarshalJSON сериализует величину в JSON
	MarshalJSON() ([]byte, error)
	// UnmarshalJSON десериализует величину из JSON
	UnmarshalJSON(data []byte) error
	// ShortName возвращает короткое имя единицы измерения
	ShortName(pref Prefix) string
	// FullName возвращает полное имя единицы измерения
	FullName(pref Prefix) string
	// SetPrefix устанавливает новый префикс
	SetPrefix(pref Prefix)
	// Prefix возвращает текущий префикс
	Prefix() Prefix
	// SetDecimals устанавливает количество знаков после запятой
	SetDecimals(dec int32)
	// Types возвращает тип величины
	Types() UnitType
	// IsNegative проверяет, является ли значение отрицательным
	IsNegative() bool
	// Ok возвращает статус последней операции
	Ok() bool
}

// Quantity базовая структура для всех величин
type Quantity struct {
	value    uint64                          // значение в базовых единицах
	prefix   Prefix                          // текущий префикс отображения
	divisor  Prefix                          // базовый делитель
	decimals int32                           // количество знаков после запятой
	unit     func(p Prefix) (string, string) // функция для получения имен единиц
	types    UnitType                        // тип величины
	negative bool                            // флаг отрицательного значения
	ok       bool                            // статус последней операции
}

// quantityJSON вспомогательная структура для сериализации
type quantityJSON struct {
	Value    uint64   `json:"value"`
	Prefix   Prefix   `json:"prefix"`
	Divisor  Prefix   `json:"divisor"`
	Decimals int32    `json:"decimals"`
	Type     UnitType `json:"type"`
	Negative bool     `json:"negative"`
}

// NewQuantity создает новую величину
func NewQuantity(val uint64, pref, div Prefix, types UnitType, unit func(p Prefix) (string, string)) *Quantity {
	if pref == 0 || pref < div {
		pref = div
	}
	val = val * pref.Uint()
	return &Quantity{value: val, prefix: pref, divisor: div, types: types, unit: unit}
}

// Value возвращает значение в базовых единицах с учетом делителя
func (q *Quantity) Copy() *Quantity {
	return &Quantity{
		value:    q.value,
		prefix:   q.prefix,
		divisor:  q.divisor,
		decimals: q.decimals,
		unit:     q.unit,
		types:    q.types,
		negative: q.negative,
		ok:       q.ok,
	}
}

// Value возвращает значение в базовых единицах с учетом делителя
func (q *Quantity) Value() uint64 {
	return q.value / q.divisor.Uint()
}

// String возвращает отформатированную строку с учетом префикса и точности
func (q *Quantity) String() string {
	sname := q.ShortName(q.prefix)
	if sname == "" {
		return ""
	}

	// Конвертируем значения в decimal для точных вычислений
	baseValue := decimal.NewFromInt(int64(q.value))
	prefixDiv := decimal.NewFromInt(int64(q.prefix.Uint()))
	divisorDiv := decimal.NewFromInt(int64(q.divisor.Uint()))

	// Приводим к базовым единицам и затем к нужному префиксу
	value := baseValue.Div(divisorDiv)
	quotient := value.Div(prefixDiv)

	// Получаем целую и дробную части
	intPart := quotient.IntPart()
	fracPart := quotient.Sub(decimal.NewFromInt(intPart))

	// Форматируем дробную часть
	decimalStr := ""
	if q.decimals > 0 {
		// Масштабируем до нужной точности
		fracPart = fracPart.Mul(decimal.NewFromInt(1)).Round(q.decimals)
		decimalStr = fracPart.StringFixed(q.decimals)

		// Убираем начальный "0" и точку
		if len(decimalStr) > 0 && decimalStr[0] == '0' {
			decimalStr = decimalStr[1:]
		}
		// Убираем лишние нули в конце
		for len(decimalStr) > 0 && decimalStr[len(decimalStr)-1] == '0' {
			decimalStr = decimalStr[:len(decimalStr)-1]
		}
		// Если осталась только точка, убираем её
		if decimalStr == "." {
			decimalStr = ""
		}
	} else {
		// Если decimals == 0, округляем до целого числа
		quotient = quotient.Round(0) // Округляем до целого
		intPart = quotient.IntPart() // Обновляем целую часть
	}

	// Формируем результат
	sign := ""
	if q.negative {
		sign = "-"
	}

	if decimalStr == "" {
		return fmt.Sprintf("%s%d %s", sign, intPart, sname)
	}
	return fmt.Sprintf("%s%d%s %s", sign, intPart, decimalStr, sname)
}

// ShortName возвращает короткое имя единицы измерения для указанного префикса
func (q *Quantity) ShortName(pref Prefix) string {
	name, _ := q.unit(pref)
	return name
}

// FullName возвращает полное имя единицы измерения для указанного префикса
func (q *Quantity) FullName(pref Prefix) string {
	_, fname := q.unit(pref)
	return fname
}

// SetPrefix устанавливает новый префикс отображения
func (q *Quantity) SetPrefix(pref Prefix) {
	if pref == 0 || pref < q.divisor {
		pref = q.divisor
	}
	if q.prefix != pref {
		q.prefix = pref
	}
}

// Prefix возвращает текущий префикс отображения
func (q *Quantity) Prefix() Prefix {
	return q.prefix
}

// SetDecimals устанавливает количество знаков после запятой
func (q *Quantity) SetDecimals(dec int32) {
	if dec < 0 {
		dec = 0
	}
	q.decimals = dec
}

// Types возвращает тип величины
func (q *Quantity) Types() UnitType {
	return q.types
}

// IsNegative возвращает true если значение отрицательное
func (q *Quantity) IsNegative() bool {
	return q.negative
}

// Ok возвращает статус последней операции
func (q *Quantity) Ok() bool {
	return q.ok
}

// MarshalJSON реализует интерфейс json.Marshaler
func (q *Quantity) MarshalJSON() ([]byte, error) {
	qj := quantityJSON{
		Value:    q.value,
		Prefix:   q.prefix,
		Divisor:  q.divisor,
		Decimals: q.decimals,
		Type:     q.types,
		Negative: q.negative,
	}
	return json.Marshal(qj)
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler
func (q *Quantity) UnmarshalJSON(data []byte) error {
	var qj quantityJSON
	if err := json.Unmarshal(data, &qj); err != nil {
		return err
	}

	// Проверяем корректность данных
	if qj.Prefix == 0 || qj.Prefix < qj.Divisor {
		qj.Prefix = qj.Divisor
	}

	// Обновляем поля структуры
	q.value = qj.Value
	q.prefix = qj.Prefix
	q.divisor = qj.Divisor
	q.decimals = qj.Decimals
	q.types = qj.Type
	q.negative = qj.Negative
	q.ok = true

	return nil
}

// Арифметические операции
// Add выполняет сложение с другой величиной того же типа
func add(a, b Quantiter) Quantiter {
	res := a.Copy()

	if b != nil && a.Types() == b.Types() {
		res.value = a.Value() + b.Value()
		res.ok = true
	} else {
		res.value = 0
		res.ok = false
	}
	return res
}

// Sub выполняет вычитание другой величины того же типа
func sub(a, b Quantiter) Quantiter {
	res := a.Copy()

	if b != nil && a.Types() == b.Types() {
		if b.Value() > a.Value() {
			res.value = b.Value() - a.Value()
			res.negative = true
		} else {
			res.value = a.Value() - b.Value()
		}
		res.ok = true
	} else {
		res.value = 0
		res.ok = false
	}
	return res
}

// Mul выполняет умножение на другую величину
func mul(a, b Quantiter) Quantiter {
	res := a.Copy()

	if b != nil {
		res.value = a.Value() * b.Value()
		res.ok = true
	} else {
		res.value = 0
		res.ok = false
	}
	return res
}

// Div выполняет деление на другую величину
func div(a, b Quantiter) Quantiter {
	res := a.Copy()

	if b != nil && b.Value() > 0 {
		// Убедитесь, что делитель не равен нулю
		if b.Value() == 0 {
			res.ok = false
			return res
		}
		if a.Value() >= b.Value() {
			res.value = a.Value() / b.Value() * 1000000000

		} else {
			res.value = (a.Value() * 1000) / b.Value() * 1000000
		}
		res.ok = true
	} else {
		res.value = 0
		res.ok = false
	}
	return res
}
