package units

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/shopspring/decimal"
)

type Unit struct {
	Base       *UnitItem            `json:"base"`
	Additional map[string]*UnitItem `json:"additional"`
	Precision  int32                `json:"precision"` // Точность вывода (количество знаков после запятой), по умолчанию 3
}

// New создает новый набор единиц с базовой единицей
func New(baseUnitName, baseUnitFullName string) *Unit {
	return &Unit{
		Base:       NewUnitItemFromInt(baseUnitName, baseUnitFullName, 1),
		Additional: make(map[string]*UnitItem),
		Precision:  3, // По умолчанию 3 знака после запятой
	}
}

// NewJSON создает новый набор единиц с базовой единицей из JSON
func NewJSON(data []byte) (*Unit, error) {
	var parsed Unit
	if err := json.Unmarshal(data, &parsed); err != nil {
		return nil, err
	}
	// Инициализируем карту если она nil
	if parsed.Additional == nil {
		parsed.Additional = make(map[string]*UnitItem)
	}
	// Устанавливаем precision по умолчанию если не задан
	if parsed.Precision == 0 {
		parsed.Precision = 3
	}

	return &parsed, nil
}

// SetPrecision устанавливает точность вывода чисел (количество знаков после запятой)
func (u *Unit) SetPrecision(precision int32) {
	if u != nil {
		u.Precision = precision
	}
}

// GetPrecision возвращает текущую точность вывода чисел
func (u *Unit) GetPrecision() int32 {
	if u == nil {
		return 3
	}
	return u.Precision
}

// AddUnit добавляет дополнительную единицу измерения
func (u *Unit) AddUnit(name, fullName string, toBase any) error {
	dec, err := ToDecimalValue(toBase)
	if err != nil {
		return fmt.Errorf("invalid toBase value: %w", err)
	}
	return u.addItem(NewUnitItem(name, fullName, dec))
}

// ToJSON сериализует Unit в JSON-строку
func (u *Unit) ToJSON() ([]byte, error) {
	if u == nil {
		return []byte("null"), nil
	}
	return json.Marshal(u)
}

// ToBase конвертирует значение из указанной единицы измерения в базовые единицы.
// unitName - название единицы измерения
// val - значение для конвертации
// Возвращает значение в базовых единицах.
func (u *Unit) ToBase(unitName string, val any) (decimal.Decimal, error) {

	unit, dec, err := u.resolveUnitValue(unitName, val)
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
func (u *Unit) StringBase(quantity any) (string, error) {
	if u == nil || u.Base == nil {
		return "", fmt.Errorf("unit set is not initialized")
	}

	dec, err := ToDecimalValue(quantity)
	if err != nil {
		return "", err
	}

	return formatDecimal(dec, u.Precision), nil
}

// StringUnit возвращает количество в указанной единице измерения
// unitName - название единицы, в которой нужно получить значение
// quantity - количество в базовых единицах
func (u *Unit) StringUnit(unitName string, quantity any) (string, error) {
	if u == nil || u.Base == nil {
		return "", fmt.Errorf("unit set is not initialized")
	}

	dec, err := ToDecimalValue(quantity)
	if err != nil {
		return "", err
	}

	var unit *UnitItem
	if unitName == u.Base.Name {
		unit = u.Base
	} else {
		found, exists := u.Additional[unitName]
		if !exists {
			return "", fmt.Errorf("unit not found: %s", unitName)
		}
		unit = found
	}

	// Конвертируем из базовых единиц в указанную единицу
	converted := unit.ConvertFromBase(dec)
	return formatDecimal(converted, u.Precision), nil
}

// List возвращает список доступных единиц измерения (сначала базовая, затем дополнительные, по алфавиту)
func (u *Unit) List() []*UnitItem {
	if u == nil || u.Base == nil {
		return nil
	}
	names := make([]string, 0, len(u.Additional))
	for name := range u.Additional {
		names = append(names, name)
	}
	sort.Strings(names)
	result := make([]*UnitItem, 0, 1+len(names))
	// Базовая единица первой
	result = append(result, u.Base)
	// Дополнительные в алфавитном порядке по краткому имени
	for _, n := range names {
		result = append(result, u.Additional[n])
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
func (u *Unit) addItem(unit *UnitItem) error {
	if unit == nil {
		return fmt.Errorf("cannot add nil unit")
	}
	if u == nil || u.Base == nil {
		return fmt.Errorf("unit set is not initialized")
	}
	if unit.Name == u.Base.Name {
		return fmt.Errorf("unit name conflicts with base unit: %s", unit.Name)
	}
	if _, exists := u.Additional[unit.Name]; exists {
		return fmt.Errorf("unit already exists: %s", unit.Name)
	}
	u.Additional[unit.Name] = unit
	return nil
}

// resolveUnitValue находит единицу измерения по имени и конвертирует значение в decimal.
// unitName - название единицы измерения
// val - значение для конвертации
// Возвращает найденную единицу, сконвертированное значение и ошибку при необходимости.
func (u *Unit) resolveUnitValue(unitName string, val any) (*UnitItem, decimal.Decimal, error) {
	if unitName == "" {
		return nil, decimal.Zero, fmt.Errorf("unit name cannot be empty")
	}
	if u == nil || u.Base == nil {
		return nil, decimal.Zero, fmt.Errorf("unit set is not initialized")
	}
	var unit *UnitItem
	if unitName == u.Base.Name {
		unit = u.Base
	} else {
		found, exists := u.Additional[unitName]
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
