package units

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// Add складывает текущее количество (в базовых единицах) со значением в указанной единице.
// currentQuantity - текущее количество в базовых единицах
// unitName - единица измерения для добавляемого значения
// val - значение в указанной единице
// Возвращает результат в базовых единицах.
func (us *Unit) Add(unitName string, currentQuantity, val any) (decimal.Decimal, error) {
	current, err := ToDecimalValue(currentQuantity)
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
func (us *Unit) Sub(unitName string, currentQuantity, val any) (decimal.Decimal, error) {
	current, err := ToDecimalValue(currentQuantity)
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
func (us *Unit) Mul(unitName string, currentQuantity, val any) (decimal.Decimal, error) {
	current, err := ToDecimalValue(currentQuantity)
	if err != nil {
		return decimal.Zero, fmt.Errorf("invalid current quantity: %w", err)
	}

	multiplier, err := ToDecimalValue(val)
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
func (us *Unit) Div(unitName string, currentQuantity, val any) (decimal.Decimal, error) {
	current, err := ToDecimalValue(currentQuantity)
	if err != nil {
		return decimal.Zero, fmt.Errorf("invalid current quantity: %w", err)
	}

	divisor, err := ToDecimalValue(val)
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
