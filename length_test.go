package units

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLengthValue(t *testing.T) {
	testCases := []struct {
		name     string
		input    uint64
		valueExp uint64
		strExp   string
		prefix   Prefix
	}{
		{
			name:     "100 метров",
			input:    100,
			valueExp: 100000000000,
			strExp:   "100 м",
			prefix:   Normal,
		},
		{
			name:     "200 метров",
			input:    200,
			valueExp: 200000000000,
			strExp:   "200 м",
			prefix:   Normal,
		},
	}

	for _, tc := range testCases {
		l := NewLength(tc.input, tc.prefix)

		t.Run(tc.name, func(t *testing.T) {
			actualValue := l.Value()
			require.Equal(t, tc.valueExp, actualValue, "Два значения должны быть равны. Ожидалось: %d, получено: %d", tc.valueExp, actualValue)

		})
	}

}

// TestLengthString проверяет метод String для длины
func TestLengthString(t *testing.T) {
	testCases := []struct {
		name       string
		input      uint64
		prefix     Prefix
		expected   string
		expPrefix  Prefix
		expDecimal int32
	}{
		{"100 метров", 100, Normal, "100 м", Normal, 0},
		{"200 метров", 200, Normal, "200 м", Normal, 0},
		{"1 километр", 1, Kilo, "1 км", Kilo, 0},
		{"1500 метров", 1500, Normal, "1.5 км", Kilo, 2},
		{"0 метров", 0, Normal, "0 м", Normal, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLength(tc.input, tc.prefix)
			l.SetPrefix(tc.expPrefix)
			l.SetDecimals(tc.expDecimal)
			actual := l.String() + " " + l.Suffix()
			require.Equal(t, tc.expected, actual, "Ожидалось: %s, получено: %s", tc.expected, actual)
		})
	}
}

// TestAdd проверяет сложение величин
func TestLengthAdd(t *testing.T) {
	testCases := []struct {
		name     string
		length1  uint64
		prefix1  Prefix
		length2  uint64
		prefix2  Prefix
		expected uint64
	}{
		{"Сложение 100 метров и 200 метров", 100, Normal, 200, Normal, 300000000000},
		{"Сложение 1 километр и 500 метров", 1, Kilo, 500, Normal, 1500000000000},
		{"Сложение 0 метров и 0 метров", 0, Normal, 0, Normal, 0},
		{"Сложение 1 метр и 100 миллиметров", 1, Normal, 100, Milli, 1100000000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l1 := NewLength(tc.length1, tc.prefix1)
			l2 := NewLength(tc.length2, tc.prefix2)
			result := l1.Add(l2)
			require.EqualValues(t, tc.expected, result.Value(), "Ожидалось: %d, получено: %d", tc.expected, result.Value())
		})
	}
}

// TestSub проверяет вычитание величин
func TestLengthSub(t *testing.T) {
	testCases := []struct {
		name     string
		length1  uint64
		prefix1  Prefix
		length2  uint64
		prefix2  Prefix
		expected uint64
	}{
		{"Вычитание 200 метров из 300 метров", 300, Normal, 200, Normal, 100000000000},
		{"Вычитание 1 километр из 1.5 километра", 1500, Kilo, 1000, Kilo, 500000000000},
		{"Вычитание 0 метров из 0 метров", 0, Normal, 0, Normal, 0},
		{"Вычитание 1 метр из 1.5 метра", 150, Normal, 1, Normal, 50000000000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l1 := NewLength(tc.length1, tc.prefix1)
			l2 := NewLength(tc.length2, tc.prefix2)
			result := l1.Sub(l2)
			require.EqualValues(t, tc.expected, result.Value(), "Ожидалось: %d, получено: %d", tc.expected, result.Value())
		})
	}
}

// TestMul проверяет умножение величин
func TestLengthMul(t *testing.T) {
	testCases := []struct {
		name     string
		length   uint64
		prefix   Prefix
		factor   uint64
		expected uint64
	}{
		{"Умножение 100 метров на 2", 100, Normal, 2, 200000000000},
		{"Умножение 1 километр на 3", 1, Kilo, 3, 3000000000000},
		{"Умножение 0 метров на 5", 0, Normal, 5, 0},
		{"Умножение 1 метр на 1.5", 1, Normal, 15, 150000000000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLength(tc.length, tc.prefix)
			result := l.Mul(NewLength(tc.factor, Normal)) // Умножаем на длину
			require.EqualValues(t, tc.expected, result.Value(), "Ожидалось: %d, получено: %d", tc.expected, result.Value())
		})
	}
}

// TestDiv проверяет деление величин
func TestLengthDiv(t *testing.T) {
	testCases := []struct {
		name     string
		length1  uint64
		prefix1  Prefix
		length2  uint64
		prefix2  Prefix
		expected uint64
	}{
		{"Деление 200 метров на 100 метров", 200, Normal, 100, Normal, 200000000000},
		{"Деление 1 километр на 500 метров", 1, Kilo, 500, Normal, 200000000000},
		{"Деление 0 метров на 1 метр", 0, Normal, 1, Normal, 0},
		{"Деление 1 метр на 0.5 метра", 1, Normal, 5, Normal, 200000000000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l1 := NewLength(tc.length1, tc.prefix1)
			l2 := NewLength(tc.length2, tc.prefix2)
			result := l1.Div(l2)
			require.EqualValues(t, tc.expected, result.Value(), "Ожидалось: %d, получено: %d", tc.expected, result.Value())
		})
	}
}

// // TestNewLength проверяет создание новой величины длины
// func TestNewLength(t *testing.T) {
// 	length := NewLength(100, Normal)

// 	if length.Value() != 100000000000 {
// 		t.Errorf("Expected value 100000000000, got %d", length.Value())
// 	}
// 	if length.Types() != LengthType {
// 		t.Errorf("Expected type LengthType, got %d", length.Types())
// 	}
// }

// // TestNewLengthJSON проверяет создание величины длины из JSON
// func TestNewLengthJSON(t *testing.T) {
// 	jsonData := `{"value":100,"prefix":1,"divisor":1,"decimals":0,"type":0,"negative":false}`
// 	length, err := NewLengthJSON([]byte(jsonData))
// 	if err != nil {
// 		t.Errorf("Error creating Length from JSON: %v", err)
// 	}

// 	if length.Value() != 100 {
// 		t.Errorf("Expected value 100, got %d", length.Value())
// 	}
// }

// // TestLengthNames проверяет правильность названий единиц измерения длины
// func TestLengthNames(t *testing.T) {
// 	short, full := LengthNames(Normal)
// 	if short != "м" || full != "метр" {
// 		t.Errorf("Expected short name 'м' and full name 'метр', got '%s' and '%s'", short, full)
// 	}

// 	short, full = LengthNames(Kilo)
// 	if short != "км" || full != "километр" {
// 		t.Errorf("Expected short name 'км' and full name 'километр', got '%s' and '%s'", short, full)
// 	}
// }

// // TestLengthString проверяет метод String для длины
// func TestLengthString(t *testing.T) {
// 	length := NewLength(100, Normal)

// 	expected := "100 м"
// 	if length.String() != expected {
// 		t.Errorf("Expected '%s', got '%s'", expected, length.String())
// 	}
// }

// // TestLengthMarshalJSON проверяет сериализацию в JSON
// func TestLengthMarshalJSON(t *testing.T) {
// 	length := NewLength(100, Normal)

// 	data, err := length.MarshalJSON()
// 	if err != nil {
// 		t.Errorf("Error marshaling JSON: %v", err)
// 	}

// 	expected := `{"value":100000000000,"prefix":1000000000,"divisor":1,"decimals":0,"type":0,"negative":false}`
// 	if string(data) != expected {
// 		t.Errorf("Expected %s, got %s", expected, string(data))
// 	}
// }

// // TestLengthUnmarshalJSON проверяет десериализацию из JSON
// func TestLengthUnmarshalJSON(t *testing.T) {
// 	jsonData := `{"value":100,"prefix":1,"divisor":1,"decimals":0,"type":0,"negative":false}`

// 	length, err := NewLengthJSON([]byte(jsonData))
// 	if err != nil {
// 		t.Errorf("Error unmarshaling JSON: %v", err)
// 	}

// 	if length.Value() != 100 {
// 		t.Errorf("Expected value 100, got %d", length.Value())
// 	}
// }
