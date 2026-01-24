package units

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func (us *Unit) Add(unitName string, val any) error {
	unit, dec, err := us.resolveUnitValue(unitName, val)
	if err != nil {
		return err
	}
	result := unit.ToBase.Add(dec)
	if !result.IsPositive() {
		return fmt.Errorf("resulting ToBase must be positive: %s", result.String())
	}
	unit.ToBase = result
	return nil
}

func (us *Unit) Sub(unitName string, val any) error {
	unit, dec, err := us.resolveUnitValue(unitName, val)
	if err != nil {
		return err
	}
	result := unit.ToBase.Sub(dec)
	if !result.IsPositive() {
		return fmt.Errorf("resulting ToBase must be positive: %s", result.String())
	}
	unit.ToBase = result
	return nil
}

func (us *Unit) Mul(unitName string, val any) error {
	unit, dec, err := us.resolveUnitValue(unitName, val)
	if err != nil {
		return err
	}
	result := unit.ToBase.Mul(dec)
	if !result.IsPositive() {
		return fmt.Errorf("resulting ToBase must be positive: %s", result.String())
	}
	unit.ToBase = result
	return nil
}

func (us *Unit) Div(unitName string, val any) error {
	unit, dec, err := us.resolveUnitValue(unitName, val)
	if err != nil {
		return err
	}
	if dec.IsZero() {
		return fmt.Errorf("division by zero")
	}
	result := unit.ToBase.Div(dec)
	if !result.IsPositive() {
		return fmt.Errorf("resulting ToBase must be positive: %s", result.String())
	}
	unit.ToBase = result
	return nil
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
