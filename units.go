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
	data *unitData
}

// New создает новый набор единиц с базовой единицей
func New(baseUnitName, baseUnitFullName string) *Unit {
	return &Unit{
		data: &unitData{
			BaseUnit:        NewUnitItemFromInt(baseUnitName, baseUnitFullName, 1),
			AdditionalUnits: make(map[string]*UnitItem),
		},
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

	return &Unit{data: &parsed}, nil
}

// AddFromFloat добавляет дополнительную единицу из float64 (совместимость с примерами)
func (us *Unit) AddFromFloat(name, fullName string, toBase float64) error {
	return us.addItem(NewUnitItemFromFloat(name, fullName, toBase))
}

// AddFromInt добавляет дополнительную единицу из int64 (совместимость с примерами)
func (us *Unit) AddFromInt(name, fullName string, toBase int64) error {
	return us.addItem(NewUnitItemFromInt(name, fullName, toBase))
}

// ToJSON сериализует Unit в JSON-строку
func (us *Unit) ToJSON() ([]byte, error) {
	if us == nil || us.data == nil {
		return []byte("null"), nil
	}
	return json.Marshal(us.data)
}

// Add добавляет дополнительную единицу измерения
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

// GetInUnit возвращает количество в указанной единице измерения
// unitName - название единицы, в которой нужно получить значение
// quantity - количество в базовых единицах
func (us *Unit) GetInUnit(unitName string, quantity any) (string, error) {
	if us == nil || us.data == nil || us.data.BaseUnit == nil {
		return "", fmt.Errorf("unit set is not initialized")
	}

	dec, err := toDecimalValue(quantity)
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
	return converted.String(), nil
}

// GetInBaseUnit возвращает количество в базовых единицах в виде строки
func (us *Unit) GetInBaseUnit(quantity any) (string, error) {
	if us == nil || us.data == nil || us.data.BaseUnit == nil {
		return "", fmt.Errorf("unit set is not initialized")
	}

	dec, err := toDecimalValue(quantity)
	if err != nil {
		return "", err
	}

	return dec.String(), nil
}

// FormatInUnit возвращает отформатированную строку с количеством и единицей измерения
// unitName - название единицы, в которой нужно вывести значение
// quantity - количество в базовых единицах
func (us *Unit) FormatInUnit(unitName string, quantity any) (string, error) {
	if us == nil || us.data == nil || us.data.BaseUnit == nil {
		return "", fmt.Errorf("unit set is not initialized")
	}

	dec, err := toDecimalValue(quantity)
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
	return fmt.Sprintf("%s %s", converted.String(), unit.Name), nil
}

// FormatInBaseUnit возвращает отформатированную строку с количеством в базовых единицах
func (us *Unit) FormatInBaseUnit(quantity any) (string, error) {
	if us == nil || us.data == nil || us.data.BaseUnit == nil {
		return "", fmt.Errorf("unit set is not initialized")
	}

	dec, err := toDecimalValue(quantity)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", dec.String(), us.data.BaseUnit.Name), nil
}

// String возвращает строковое представление всего набора единиц
func (us *Unit) String() string {
	if us == nil || us.data == nil || us.data.BaseUnit == nil {
		return "Unit(uninitialized)"
	}

	result := fmt.Sprintf("Базовая единица: %s (%s)", us.data.BaseUnit.Name, us.data.BaseUnit.FullName)
	if len(us.data.AdditionalUnits) > 0 {
		result += "\nДополнительные единицы:"
		for _, unit := range us.data.AdditionalUnits {
			result += fmt.Sprintf("\n  - %s (%s): 1 %s = %s %s",
				unit.Name, unit.FullName, unit.Name, unit.ToBase.String(), us.data.BaseUnit.Name)
		}
	}
	return result
}

// UnitItems возвращает список доступных единиц измерения (сначала базовая, затем дополнительные, по алфавиту)
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

// toDecimalValue конвертирует любое значение в decimal.Decimal
func toDecimalValue(val any) (decimal.Decimal, error) {
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
