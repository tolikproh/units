package units

import (
	"testing"
)

// TestAllFunctional проверяет весь функционал библиотеки
func TestAllFunctional(t *testing.T) {
	// Создание величин
	length1 := NewLength(100, Normal) // 100 метров
	length2 := NewLength(2, Kilo)     // 2 километра

	// Проверка значений
	if length1.Value() != 100 {
		t.Errorf("Expected value 100, got %d", length1.Value())
	}
	if length2.Value() != 2000 {
		t.Errorf("Expected value 2000, got %d", length2.Value())
	}

	// Сложение
	totalLength := length1.Add(length2)
	if totalLength.Value() != 2100 {
		t.Errorf("Expected total length value 2100, got %d", totalLength.Value())
	}

	// Вычитание
	difference := length2.Sub(length1)
	if difference.Value() != 1900 {
		t.Errorf("Expected difference value 1900, got %d", difference.Value())
	}

	// Умножение
	things := NewThings(5, Normal) // 5 штук
	multiplied := length1.Mul(things)
	if multiplied.Value() != 500 {
		t.Errorf("Expected multiplied value 500, got %d", multiplied.Value())
	}

	// Деление
	divided := length2.Div(NewLength(1, Normal)) // Делим на 1 метр
	if divided.Value() != 2000 {
		t.Errorf("Expected divided value 2000, got %d", divided.Value())
	}

	// Проверка строкового представления
	if length1.String() != "100 м" {
		t.Errorf("Expected string representation '100 м', got '%s'", length1.String())
	}
	if length2.String() != "2000 м" {
		t.Errorf("Expected string representation '2000 м', got '%s'", length2.String())
	}

	// Сериализация в JSON
	jsonData, err := length1.MarshalJSON()
	if err != nil {
		t.Errorf("Error marshaling JSON: %v", err)
	}

	expectedJSON := `{"value":100,"prefix":1,"divisor":1,"decimals":0,"type":0,"negative":false}`
	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}

	// Десериализация из JSON
	var newLength Length
	err = newLength.UnmarshalJSON(jsonData)
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}
	if newLength.Value() != 100 {
		t.Errorf("Expected value 100 after unmarshal, got %d", newLength.Value())
	}

	// Проверка типов
	if newLength.Types() != LengthType {
		t.Errorf("Expected type LengthType after unmarshal, got %d", newLength.Types())
	}
}
