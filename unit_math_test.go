package units

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

// TestAdd проверяет сложение значений
func TestAdd(t *testing.T) {
	u := New("м", "метр")
	u.AddUnit("км", "километр", 1000)
	u.AddUnit("см", "сантиметр", 0.01)

	tests := []struct {
		name            string
		unitName        string
		currentQuantity any
		val             any
		want            string
		wantErr         bool
	}{
		{
			name:            "add meters",
			unitName:        "м",
			currentQuantity: 100,
			val:             50,
			want:            "150",
			wantErr:         false,
		},
		{
			name:            "add kilometers",
			unitName:        "км",
			currentQuantity: 1000,
			val:             2,
			want:            "3000",
			wantErr:         false,
		},
		{
			name:            "add centimeters",
			unitName:        "см",
			currentQuantity: 10,
			val:             250,
			want:            "12.5",
			wantErr:         false,
		},
		{
			name:            "add float",
			unitName:        "км",
			currentQuantity: 5000,
			val:             1.5,
			want:            "6500",
			wantErr:         false,
		},
		{
			name:            "invalid current quantity",
			unitName:        "м",
			currentQuantity: []int{1, 2, 3},
			val:             10,
			wantErr:         true,
		},
		{
			name:            "invalid value",
			unitName:        "м",
			currentQuantity: 100,
			val:             []int{1, 2, 3},
			wantErr:         true,
		},
		{
			name:            "unit not found",
			unitName:        "мм",
			currentQuantity: 100,
			val:             10,
			wantErr:         true,
		},
		{
			name:            "negative result",
			unitName:        "м",
			currentQuantity: -100,
			val:             50,
			wantErr:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := u.Add(tt.unitName, tt.currentQuantity, tt.val)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result.String())
			}
		})
	}
}

// TestSub проверяет вычитание значений
func TestSub(t *testing.T) {
	u := New("м", "метр")
	u.AddUnit("км", "километр", 1000)
	u.AddUnit("см", "сантиметр", 0.01)

	tests := []struct {
		name            string
		unitName        string
		currentQuantity any
		val             any
		want            string
		wantErr         bool
	}{
		{
			name:            "subtract meters",
			unitName:        "м",
			currentQuantity: 100,
			val:             30,
			want:            "70",
			wantErr:         false,
		},
		{
			name:            "subtract kilometers",
			unitName:        "км",
			currentQuantity: 5000,
			val:             2,
			want:            "3000",
			wantErr:         false,
		},
		{
			name:            "subtract centimeters",
			unitName:        "см",
			currentQuantity: 10,
			val:             250,
			want:            "7.5",
			wantErr:         false,
		},
		{
			name:            "subtract float",
			unitName:        "км",
			currentQuantity: 5000,
			val:             1.5,
			want:            "3500",
			wantErr:         false,
		},
		{
			name:            "result is zero",
			unitName:        "м",
			currentQuantity: 100,
			val:             100,
			want:            "0",
			wantErr:         false,
		},
		{
			name:            "negative result",
			unitName:        "м",
			currentQuantity: 50,
			val:             100,
			wantErr:         true,
		},
		{
			name:            "invalid current quantity",
			unitName:        "м",
			currentQuantity: []int{1, 2, 3},
			val:             10,
			wantErr:         true,
		},
		{
			name:            "invalid value",
			unitName:        "м",
			currentQuantity: 100,
			val:             []int{1, 2, 3},
			wantErr:         true,
		},
		{
			name:            "unit not found",
			unitName:        "мм",
			currentQuantity: 100,
			val:             10,
			wantErr:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := u.Sub(tt.unitName, tt.currentQuantity, tt.val)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result.String())
			}
		})
	}
}

// TestMul проверяет умножение значений
func TestMul(t *testing.T) {
	u := New("м", "метр")

	tests := []struct {
		name            string
		unitName        string
		currentQuantity any
		val             any
		want            string
		wantErr         bool
	}{
		{
			name:            "multiply by int",
			unitName:        "м",
			currentQuantity: 100,
			val:             3,
			want:            "300",
			wantErr:         false,
		},
		{
			name:            "multiply by float",
			unitName:        "м",
			currentQuantity: 50,
			val:             2.5,
			want:            "125",
			wantErr:         false,
		},
		{
			name:            "multiply by zero",
			unitName:        "м",
			currentQuantity: 100,
			val:             0,
			want:            "0",
			wantErr:         false,
		},
		{
			name:            "multiply by fraction",
			unitName:        "м",
			currentQuantity: 100,
			val:             0.5,
			want:            "50",
			wantErr:         false,
		},
		{
			name:            "multiply by decimal",
			unitName:        "м",
			currentQuantity: decimal.NewFromInt(75),
			val:             decimal.NewFromFloat(1.5),
			want:            "112.5",
			wantErr:         false,
		},
		{
			name:            "negative result",
			unitName:        "м",
			currentQuantity: 100,
			val:             -2,
			wantErr:         true,
		},
		{
			name:            "invalid current quantity",
			unitName:        "м",
			currentQuantity: []int{1, 2, 3},
			val:             2,
			wantErr:         true,
		},
		{
			name:            "invalid multiplier",
			unitName:        "м",
			currentQuantity: 100,
			val:             []int{1, 2, 3},
			wantErr:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := u.Mul(tt.unitName, tt.currentQuantity, tt.val)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result.String())
			}
		})
	}
}

// TestDiv проверяет деление значений
func TestDiv(t *testing.T) {
	u := New("м", "метр")

	tests := []struct {
		name            string
		unitName        string
		currentQuantity any
		val             any
		want            string
		wantErr         bool
	}{
		{
			name:            "divide by int",
			unitName:        "м",
			currentQuantity: 100,
			val:             4,
			want:            "25",
			wantErr:         false,
		},
		{
			name:            "divide by float",
			unitName:        "м",
			currentQuantity: 100,
			val:             2.5,
			want:            "40",
			wantErr:         false,
		},
		{
			name:            "divide by fraction",
			unitName:        "м",
			currentQuantity: 50,
			val:             0.5,
			want:            "100",
			wantErr:         false,
		},
		{
			name:            "divide by decimal",
			unitName:        "м",
			currentQuantity: decimal.NewFromInt(150),
			val:             decimal.NewFromInt(3),
			want:            "50",
			wantErr:         false,
		},
		{
			name:            "result with fraction",
			unitName:        "м",
			currentQuantity: 100,
			val:             3,
			want:            "33.3333333333333333",
			wantErr:         false,
		},
		{
			name:            "divide by zero",
			unitName:        "м",
			currentQuantity: 100,
			val:             0,
			wantErr:         true,
		},
		{
			name:            "negative result",
			unitName:        "м",
			currentQuantity: 100,
			val:             -2,
			wantErr:         true,
		},
		{
			name:            "invalid current quantity",
			unitName:        "м",
			currentQuantity: []int{1, 2, 3},
			val:             2,
			wantErr:         true,
		},
		{
			name:            "invalid divisor",
			unitName:        "м",
			currentQuantity: 100,
			val:             []int{1, 2, 3},
			wantErr:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := u.Div(tt.unitName, tt.currentQuantity, tt.val)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result.String())
			}
		})
	}
}
