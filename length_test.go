package units

import (
	"testing"
)

// TestNewLength проверяет создание новой величины длины
func TestNewLength(t *testing.T) {
	length := NewLength(100, Normal)

	if length.Value() != 100000000000 {
		t.Errorf("Expected value 100000000000, got %d", length.Value())
	}
	if length.Types() != LengthType {
		t.Errorf("Expected type LengthType, got %d", length.Types())
	}
}

// TestNewLengthJSON проверяет создание величины длины из JSON
func TestNewLengthJSON(t *testing.T) {
	jsonData := `{"value":100,"prefix":1,"divisor":1,"decimals":0,"type":0,"negative":false}`
	length, err := NewLengthJSON([]byte(jsonData))
	if err != nil {
		t.Errorf("Error creating Length from JSON: %v", err)
	}

	if length.Value() != 100 {
		t.Errorf("Expected value 100, got %d", length.Value())
	}
}

// TestLengthNames проверяет правильность названий единиц измерения длины
func TestLengthNames(t *testing.T) {
	short, full := LengthNames(Normal)
	if short != "м" || full != "метр" {
		t.Errorf("Expected short name 'м' and full name 'метр', got '%s' and '%s'", short, full)
	}

	short, full = LengthNames(Kilo)
	if short != "км" || full != "километр" {
		t.Errorf("Expected short name 'км' and full name 'километр', got '%s' and '%s'", short, full)
	}
}

// TestLengthString проверяет метод String для длины
func TestLengthString(t *testing.T) {
	length := NewLength(100, Normal)

	expected := "100 м"
	if length.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, length.String())
	}
}

// TestLengthMarshalJSON проверяет сериализацию в JSON
func TestLengthMarshalJSON(t *testing.T) {
	length := NewLength(100, Normal)

	data, err := length.MarshalJSON()
	if err != nil {
		t.Errorf("Error marshaling JSON: %v", err)
	}

	expected := `{"value":100000000000,"prefix":1000000000,"divisor":1,"decimals":0,"type":0,"negative":false}`
	if string(data) != expected {
		t.Errorf("Expected %s, got %s", expected, string(data))
	}
}

// TestLengthUnmarshalJSON проверяет десериализацию из JSON
func TestLengthUnmarshalJSON(t *testing.T) {
	jsonData := `{"value":100,"prefix":1,"divisor":1,"decimals":0,"type":0,"negative":false}`

	length, err := NewLengthJSON([]byte(jsonData))
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	if length.Value() != 100 {
		t.Errorf("Expected value 100, got %d", length.Value())
	}
}
