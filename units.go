package units

import (
	"encoding/json"
	"fmt"
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
