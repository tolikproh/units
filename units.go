package units

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/shopspring/decimal"
)

type unitData struct {
	BaseUnit        *UnitItem            `json:"base_unit"`
	AdditionalUnits map[string]*UnitItem `json:"additional_units"`
}

// Unit представляет набор единиц измерения для сохранения в БД
// Это то, что должно храниться вместе с товаром в базе данных
type Unit struct {
	data      *unitData
	precision int32 // Точность вывода (количество знаков после запятой), по умолчанию 3
}

// New создает новый набор единиц с базовой единицей
func New(baseUnitName, baseUnitFullName string) *Unit {
	return &Unit{
		data: &unitData{
			BaseUnit:        NewUnitItemFromInt(baseUnitName, baseUnitFullName, 1),
			AdditionalUnits: make(map[string]*UnitItem),
		},
		precision: 3, // По умолчанию 3 знака после запятой
	}
}

// NewJSON создает новый набор единиц с базовой единицей из JSON
func NewJSON(data []byte) (*Unit, error) {
	var parsed unitData
	if err := json.Unmarshal(data, &parsed); err != nil {
		return nil, err
	}
	// Инициализируем карту если она nil
	if parsed.AdditionalUnits == nil {
		parsed.AdditionalUnits = make(map[string]*UnitItem)
	}

	return &Unit{
		data:      &parsed,
		precision: 3, // По умолчанию 3 знака после запятой
	}, nil
}

// SetPrecision устанавливает точность вывода чисел (количество знаков после запятой)
func (us *Unit) SetPrecision(precision int32) {
	if us != nil {
		us.precision = precision
	}
}

// GetPrecision возвращает текущую точность вывода чисел
func (us *Unit) GetPrecision() int32 {
	if us == nil {
		return 3
	}
	return us.precision
}

// AddUnit добавляет дополнительную единицу измерения
func (us *Unit) AddUnit(name, fullName string, toBase any) error {
	dec, err := ToDecimalValue(toBase)
	if err != nil {
		return fmt.Errorf("invalid toBase value: %w", err)
	}
	return us.addItem(NewUnitItem(name, fullName, dec))
}

// ToJSON сериализует Unit в JSON-строку
func (us *Unit) ToJSON() ([]byte, error) {
	if us == nil || us.data == nil {
		return []byte("null"), nil
	}
	return json.Marshal(us.data)
}

// ToBase конвертирует значение из указанной единицы измерения в базовые единицы.
// unitName - название единицы измерения
// val - значение для конвертации
// Возвращает значение в базовых единицах.
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

// formatDecimal форматирует decimal с учетом точности.
// Если число целое, не выводит дробную часть.
// Убирает незначащие нули справа.
func formatDecimal(dec decimal.Decimal, precision int32) string {
	// Проверяем, является ли число целым
	if dec.Equal(dec.Truncate(0)) {
		return dec.Truncate(0).String()
	}
	// Округляем до нужной точности и используем String() для удаления незначащих нулей
	return dec.Round(precision).String()
}

// StringBase возвращает количество в базовых единицах в виде строки
func (us *Unit) StringBase(quantity any) (string, error) {
	if us == nil || us.data == nil || us.data.BaseUnit == nil {
		return "", fmt.Errorf("unit set is not initialized")
	}

	dec, err := ToDecimalValue(quantity)
	if err != nil {
		return "", err
	}

	return formatDecimal(dec, us.precision), nil
}

// StringUnit возвращает количество в указанной единице измерения
// unitName - название единицы, в которой нужно получить значение
// quantity - количество в базовых единицах
func (us *Unit) StringUnit(unitName string, quantity any) (string, error) {
	if us == nil || us.data == nil || us.data.BaseUnit == nil {
		return "", fmt.Errorf("unit set is not initialized")
	}

	dec, err := ToDecimalValue(quantity)
	if err != nil {
		return "", err
	}

	var unit *UnitItem
	if unitName == us.data.BaseUnit.Name {
		unit = us.data.BaseUnit
	} else {
		found, exists := us.data.AdditionalUnits[unitName]
		if !exists {
			return "", fmt.Errorf("unit not found: %s", unitName)
		}
		unit = found
	}

	// Конвертируем из базовых единиц в указанную единицу
	converted := unit.ConvertFromBase(dec)
	return formatDecimal(converted, us.precision), nil
}

// List возвращает список доступных единиц измерения (сначала базовая, затем дополнительные, по алфавиту)
func (us *Unit) List() []*UnitItem {
	if us == nil || us.data == nil || us.data.BaseUnit == nil {
		return nil
	}
	names := make([]string, 0, len(us.data.AdditionalUnits))
	for name := range us.data.AdditionalUnits {
		names = append(names, name)
	}
	sort.Strings(names)
	result := make([]*UnitItem, 0, 1+len(names))
	// Базовая единица первой
	result = append(result, us.data.BaseUnit)
	// Дополнительные в алфавитном порядке по краткому имени
	for _, n := range names {
		result = append(result, us.data.AdditionalUnits[n])
	}
	return result
}

// ToDecimalValue конвертирует любое значение в decimal.Decimal.
// Поддерживает типы: int, int64, uint, uint64, float32, float64, string, decimal.Decimal.
// Возвращает ошибку, если тип не поддерживается или строка не может быть распарсена.
func ToDecimalValue(val any) (decimal.Decimal, error) {
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
	case string:
		d, err := decimal.NewFromString(v)
		if err != nil {
			return decimal.Zero, fmt.Errorf("invalid decimal string: %s", v)
		}
		return d, nil
	default:
		return decimal.Zero, fmt.Errorf("unsupported value type: %T", val)
	}
}

// addItem добавляет дополнительную единицу измерения в набор
// unit - единица измерения для добавления
// Возвращает ошибку, если единица уже существует или конфликтует с базовой единицей.
func (us *Unit) addItem(unit *UnitItem) error {
	if unit == nil {
		return fmt.Errorf("cannot add nil unit")
	}
	if us == nil || us.data == nil || us.data.BaseUnit == nil {
		return fmt.Errorf("unit set is not initialized")
	}
	if unit.Name == us.data.BaseUnit.Name {
		return fmt.Errorf("unit name conflicts with base unit: %s", unit.Name)
	}
	if _, exists := us.data.AdditionalUnits[unit.Name]; exists {
		return fmt.Errorf("unit already exists: %s", unit.Name)
	}
	us.data.AdditionalUnits[unit.Name] = unit
	return nil
}

// resolveUnitValue находит единицу измерения по имени и конвертирует значение в decimal.
// unitName - название единицы измерения
// val - значение для конвертации
// Возвращает найденную единицу, сконвертированное значение и ошибку при необходимости.
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
	dec, err := ToDecimalValue(val)
	if err != nil {
		return nil, decimal.Zero, err
	}
	return unit, dec, nil
}
