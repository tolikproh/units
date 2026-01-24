package units

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func (us *Unit) ToBase(unitName string, val any) (decimal.Decimal, error) {

	unit, dec, err := us.resolveUnitValue(unitName, val)
	if err != nil {
		return decimal.Zero, err
	}

	// Конвертируем значение из указанной единицы в базовые единицы
	valueInBase := unit.ConvertToBase(dec)

	if valueInBase.IsNegative() {
		return decimal.Zero, fmt.Errorf("resulting quantity cannot be negative: %s", valueInBase.String())
	}

	return valueInBase, nil
}

// Add складывает текущее количество (в базовых единицах) со значением в указанной единице.
// currentQuantity - текущее количество в базовых единицах
// unitName - единица измерения для добавляемого значения
// val - значение в указанной единице
// Возвращает результат в базовых единицах.
func (us *Unit) Add(currentQuantity any, unitName string, val any) (decimal.Decimal, error) {
	current, err := toDecimal(currentQuantity)
	if err != nil {
		return decimal.Zero, fmt.Errorf("invalid current quantity: %w", err)
	}

	unit, dec, err := us.resolveUnitValue(unitName, val)
	if err != nil {
		return decimal.Zero, err
	}

	// Конвертируем значение из указанной единицы в базовые единицы
	valueInBase := unit.ConvertToBase(dec)
	result := current.Add(valueInBase)

	if result.IsNegative() {
		return decimal.Zero, fmt.Errorf("resulting quantity cannot be negative: %s", result.String())
	}

	return result, nil
}

// Sub вычитает значение (в указанной единице) из текущего количества (в базовых единицах).
// currentQuantity - текущее количество в базовых единицах
// unitName - единица измерения для вычитаемого значения
// val - значение в указанной единице
// Возвращает результат в базовых единицах.
func (us *Unit) Sub(currentQuantity any, unitName string, val any) (decimal.Decimal, error) {
	current, err := toDecimal(currentQuantity)
	if err != nil {
		return decimal.Zero, fmt.Errorf("invalid current quantity: %w", err)
	}

	unit, dec, err := us.resolveUnitValue(unitName, val)
	if err != nil {
		return decimal.Zero, err
	}

	// Конвертируем значение из указанной единицы в базовые единицы
	valueInBase := unit.ConvertToBase(dec)
	result := current.Sub(valueInBase)

	if result.IsNegative() {
		return decimal.Zero, fmt.Errorf("resulting quantity cannot be negative: %s", result.String())
	}

	return result, nil
}

// Mul умножает текущее количество на числовое значение.
// currentQuantity - текущее количество в базовых единицах
// val - множитель (безразмерное число)
// Возвращает результат в базовых единицах.
func (us *Unit) Mul(currentQuantity any, val any) (decimal.Decimal, error) {
	current, err := toDecimal(currentQuantity)
	if err != nil {
		return decimal.Zero, fmt.Errorf("invalid current quantity: %w", err)
	}

	multiplier, err := toDecimal(val)
	if err != nil {
		return decimal.Zero, fmt.Errorf("invalid multiplier: %w", err)
	}

	result := current.Mul(multiplier)

	if result.IsNegative() {
		return decimal.Zero, fmt.Errorf("resulting quantity cannot be negative: %s", result.String())
	}

	return result, nil
}

// Div делит текущее количество на числовое значение.
// currentQuantity - текущее количество в базовых единицах
// val - делитель (безразмерное число)
// Возвращает результат в базовых единицах.
func (us *Unit) Div(currentQuantity any, val any) (decimal.Decimal, error) {
	current, err := toDecimal(currentQuantity)
	if err != nil {
		return decimal.Zero, fmt.Errorf("invalid current quantity: %w", err)
	}

	divisor, err := toDecimal(val)
	if err != nil {
		return decimal.Zero, fmt.Errorf("invalid divisor: %w", err)
	}

	if divisor.IsZero() {
		return decimal.Zero, fmt.Errorf("division by zero")
	}

	result := current.Div(divisor)

	if result.IsNegative() {
		return decimal.Zero, fmt.Errorf("resulting quantity cannot be negative: %s", result.String())
	}

	return result, nil
}

func (us *Unit) resolveUnitValue(unitName string, val any) (*UnitItem, decimal.Decimal, error) {
	if unitName == "" {
		return nil, decimal.Zero, fmt.Errorf("unit name cannot be empty")
	}
	if us == nil || us.data == nil || us.data.BaseUnit == nil {
		return nil, decimal.Zero, fmt.Errorf("unit set is not initialized")
	}
	var unit *UnitItem
	if unitName == us.data.BaseUnit.Name {
		unit = us.data.BaseUnit
	} else {
		found, exists := us.data.AdditionalUnits[unitName]
		if !exists {
			return nil, decimal.Zero, fmt.Errorf("unit not found: %s", unitName)
		}
		unit = found
	}
	dec, err := toDecimal(val)
	if err != nil {
		return nil, decimal.Zero, err
	}
	return unit, dec, nil
}

func toDecimal(val any) (decimal.Decimal, error) {
	switch v := val.(type) {
	case decimal.Decimal:
		return v, nil
	case *decimal.Decimal:
		if v == nil {
			return decimal.Zero, fmt.Errorf("decimal value cannot be nil")
		}
		return *v, nil
	case int:
		return decimal.NewFromInt(int64(v)), nil
	case int64:
		return decimal.NewFromInt(v), nil
	case int32:
		return decimal.NewFromInt(int64(v)), nil
	case int16:
		return decimal.NewFromInt(int64(v)), nil
	case int8:
		return decimal.NewFromInt(int64(v)), nil
	case uint:
		return decimal.NewFromInt(int64(v)), nil
	case uint64:
		if v > uint64(^uint(0)) {
			return decimal.Zero, fmt.Errorf("uint64 value overflows int64: %d", v)
		}
		return decimal.NewFromInt(int64(v)), nil
	case uint32:
		return decimal.NewFromInt(int64(v)), nil
	case uint16:
		return decimal.NewFromInt(int64(v)), nil
	case uint8:
		return decimal.NewFromInt(int64(v)), nil
	case float64:
		return decimal.NewFromFloat(v), nil
	case float32:
		return decimal.NewFromFloat(float64(v)), nil
	default:
		return decimal.Zero, fmt.Errorf("unsupported value type: %T", val)
	}
}
