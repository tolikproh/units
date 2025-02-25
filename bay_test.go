package units

import (
	"testing"
)

// TestNewBay проверяет создание новой величины в бухтах
func TestNewBay(t *testing.T) {
	bay := NewBay(100, Normal)

	if bay.Value() != 100 {
		t.Errorf("Expected value 100, got %d", bay.Value())
	}
	if bay.Types() != BayType {
		t.Errorf("Expected type BayType, got %d", bay.Types())
	}
}

// TestNewBayJSON проверяет создание величины в бухтах из JSON
func TestNewBayJSON(t *testing.T) {
	jsonData := `{"value":100,"prefix":1,"divisor":1,"decimals":0,"type":5,"negative":false}`
	bay, err := NewBayJSON([]byte(jsonData))
	if err != nil {
		t.Errorf("Error creating Bay from JSON: %v", err)
	}

	if bay.Value() != 100 {
		t.Errorf("Expected value 100, got %d", bay.Value())
	}
}

// TestBayNames проверяет правильность названий единиц измерения бухт
func TestBayNames(t *testing.T) {
	short, full := BayNames(Normal)
	if short != "бух" || full != "бухта" {
		t.Errorf("Expected short name 'бух' and full name 'бухта', got '%s' and '%s'", short, full)
	}

	short, full = BayNames(Kilo)
	if short != "тыс.бух" || full != "тысяч бухт" {
		t.Errorf("Expected short name 'тыс.бух' and full name 'тысяч бухт', got '%s' and '%s'", short, full)
	}
}

// TestMarshalJSON проверяет сериализацию в JSON
func TestMarshalJSONBay(t *testing.T) {
	bay := NewBay(100, Normal)

	data, err := bay.MarshalJSON()
	if err != nil {
		t.Errorf("Error marshaling JSON: %v", err)
	}

	expected := `{"value":100000000000,"prefix":1000000000,"divisor":1000000000,"decimals":0,"type":5,"negative":false}`
	if string(data) != expected {
		t.Errorf("Expected %s, got %s", expected, string(data))
	}
}

// TestUnmarshalJSON проверяет десериализацию из JSON
func TestUnmarshalJSONBay(t *testing.T) {
	jsonData := `{"value":100000000000,"prefix":1000000000,"divisor":1000000000,"decimals":0,"type":5,"negative":false}`

	bay, err := NewBayJSON([]byte(jsonData))
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	if bay.Value() != 100 {
		t.Errorf("Expected value 100, got %d", bay.Value())
	}
}
