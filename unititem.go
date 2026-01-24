package units

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// UnitItem представляет единицу измерения с коэффициентом конвертации в базовую единицу
type UnitItem struct {
	// Name - короткое название единицы (например, "м", "шт", "упак")
	Name string
	// FullName - полное название единицы (например, "метр", "штука", "упаковка")
	FullName string
	// ToBase - коэффициент конвертации в базовую единицу
	// Например, если базовая единица - метры, а эта единица - километры, то ToBase = 1000
	ToBase decimal.Decimal
}

// NewUnitItem создает новую единицу измерения
func NewUnitItem(name, fullName string, toBase decimal.Decimal) *UnitItem {
	if toBase.IsZero() || toBase.IsNegative() {
		toBase = decimal.NewFromInt(1)
	}
	return &UnitItem{
		Name:     name,
		FullName: fullName,
		ToBase:   toBase,
	}
}

// NewUnitItemFromFloat создает новую единицу измерения из float64
func NewUnitItemFromFloat(name, fullName string, toBase float64) *UnitItem {
	return NewUnitItem(name, fullName, decimal.NewFromFloat(toBase))
}

// NewUnitItemFromInt создает новую единицу измерения из int64
func NewUnitItemFromInt(name, fullName string, toBase int64) *UnitItem {
	return NewUnitItem(name, fullName, decimal.NewFromInt(toBase))
}

// ConvertToBase конвертирует значение из текущей единицы в базовую
func (u *UnitItem) ConvertToBase(value decimal.Decimal) decimal.Decimal {
	return value.Mul(u.ToBase)
}

// ConvertFromBase конвертирует значение из базовой единицы в текущую
func (u *UnitItem) ConvertFromBase(baseValue decimal.Decimal) decimal.Decimal {
	return baseValue.Div(u.ToBase)
}

// String возвращает строковое представление единицы
func (u *UnitItem) String() string {
	return fmt.Sprintf("%s (1 %s = %s базовых единиц)", u.Name, u.Name, u.ToBase.String())
}
