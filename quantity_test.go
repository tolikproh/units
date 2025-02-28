package units

import (
	"testing"
)

// TestNewQuantity проверяет создание новой величины
func TestNewQuantity(t *testing.T) {
	prefix := Prefix(1)  // Пример префикса
	divisor := Prefix(1) // Пример делителя
	quantity := NewQuantity(100, prefix, divisor, LengthType, func(p Prefix) (string, string) {
		return "m", "метр"
	})

	if quantity.Value() != 100 {
		t.Errorf("Expected value 100, got %d", quantity.Value())
	}
	if quantity.Types() != LengthType {
		t.Errorf("Expected type LengthType, got %d", quantity.Types())
	}
}

// TestAdd проверяет сложение величин
func TestAdd(t *testing.T) {
	q1 := NewQuantity(100, Prefix(1), Prefix(1), LengthType, func(p Prefix) (string, string) {
		return "m", "метр"
	})
	q2 := NewQuantity(50, Prefix(1), Prefix(1), LengthType, func(p Prefix) (string, string) {
		return "m", "метр"
	})

	result := add(q1, q2)
	if result.Value() != 150 {
		t.Errorf("Expected value 150, got %d", result.Value())
	}
}

// TestSub проверяет вычитание величин
func TestSub(t *testing.T) {
	q1 := NewQuantity(100, Prefix(1), Prefix(1), LengthType, func(p Prefix) (string, string) {
		return "m", "метр"
	})
	q2 := NewQuantity(30, Prefix(1), Prefix(1), LengthType, func(p Prefix) (string, string) {
		return "m", "метр"
	})

	result := sub(q1, q2)
	if result.Value() != 70 {
		t.Errorf("Expected value 70, got %d", result.Value())
	}
}

// TestMul проверяет умножение величин
func TestMul(t *testing.T) {
	q1 := NewQuantity(10, Prefix(1), Prefix(1), LengthType, func(p Prefix) (string, string) {
		return "m", "метр"
	})
	q2 := NewQuantity(5, Prefix(1), Prefix(1), LengthType, func(p Prefix) (string, string) {
		return "m", "метр"
	})

	result := mul(q1, q2)
	if result.Value() != 50 {
		t.Errorf("Expected value 50, got %d", result.Value())
	}
}

// TestDiv проверяет деление величин
func TestDiv(t *testing.T) {
	q1 := NewQuantity(100, Prefix(1), Prefix(1), LengthType, func(p Prefix) (string, string) {
		return "m", "метр"
	})
	q2 := NewQuantity(2, Prefix(1), Prefix(1), LengthType, func(p Prefix) (string, string) {
		return "m", "метр"
	})

	result := div(q1, q2)
	if result.Value() != 50 {
		t.Errorf("Expected value 50, got %d", result.Value())
	}
}

// TestMarshalJSON проверяет сериализацию в JSON
func TestMarshalJSON(t *testing.T) {
	q := NewQuantity(100, Prefix(1), Prefix(1), LengthType, func(p Prefix) (string, string) {
		return "m", "метр"
	})

	data, err := q.MarshalJSON()
	if err != nil {
		t.Errorf("Error marshaling JSON: %v", err)
	}

	expected := `{"value":100,"prefix":1,"divisor":1,"decimals":0,"type":0,"negative":false}`
	if string(data) != expected {
		t.Errorf("Expected %s, got %s", expected, string(data))
	}
}

// TestUnmarshalJSON проверяет десериализацию из JSON
func TestUnmarshalJSON(t *testing.T) {
	jsonData := `{"value":100,"prefix":1,"divisor":1,"decimals":0,"type":0,"negative":false}`
	var q Quantity

	err := q.UnmarshalJSON([]byte(jsonData))
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	if q.Value() != 100 {
		t.Errorf("Expected value 100, got %d", q.Value())
	}
}
