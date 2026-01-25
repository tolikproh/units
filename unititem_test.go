package units

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewUnitItem проверяет создание новой единицы измерения
func TestNewUnitItem(t *testing.T) {
	tests := []struct {
		name     string
		unitName string
		fullName string
		toBase   decimal.Decimal
		wantBase string
	}{
		{
			name:     "valid positive",
			unitName: "км",
			fullName: "километр",
			toBase:   decimal.NewFromInt(1000),
			wantBase: "1000",
		},
		{
			name:     "valid fraction",
			unitName: "см",
			fullName: "сантиметр",
			toBase:   decimal.NewFromFloat(0.01),
			wantBase: "0.01",
		},
		{
			name:     "zero toBase defaults to 1",
			unitName: "шт",
			fullName: "штука",
			toBase:   decimal.Zero,
			wantBase: "1",
		},
		{
			name:     "negative toBase defaults to 1",
			unitName: "упак",
			fullName: "упаковка",
			toBase:   decimal.NewFromInt(-10),
			wantBase: "1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := NewUnitItem(tt.unitName, tt.fullName, tt.toBase)
			require.NotNil(t, item)
			assert.Equal(t, tt.unitName, item.Name)
			assert.Equal(t, tt.fullName, item.FullName)
			assert.Equal(t, tt.wantBase, item.ToBase.String())
		})
	}
}

// TestNewUnitItemFromFloat проверяет создание единицы из float64
func TestNewUnitItemFromFloat(t *testing.T) {
	tests := []struct {
		name     string
		unitName string
		fullName string
		toBase   float64
		wantBase string
	}{
		{
			name:     "valid float",
			unitName: "км",
			fullName: "километр",
			toBase:   1000.0,
			wantBase: "1000",
		},
		{
			name:     "valid fraction",
			unitName: "мм",
			fullName: "миллиметр",
			toBase:   0.001,
			wantBase: "0.001",
		},
		{
			name:     "zero defaults to 1",
			unitName: "шт",
			fullName: "штука",
			toBase:   0.0,
			wantBase: "1",
		},
		{
			name:     "negative defaults to 1",
			unitName: "уп",
			fullName: "упаковка",
			toBase:   -5.0,
			wantBase: "1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := NewUnitItemFromFloat(tt.unitName, tt.fullName, tt.toBase)
			require.NotNil(t, item)
			assert.Equal(t, tt.unitName, item.Name)
			assert.Equal(t, tt.fullName, item.FullName)
			assert.Equal(t, tt.wantBase, item.ToBase.String())
		})
	}
}

// TestNewUnitItemFromInt проверяет создание единицы из int64
func TestNewUnitItemFromInt(t *testing.T) {
	tests := []struct {
		name     string
		unitName string
		fullName string
		toBase   int64
		wantBase string
	}{
		{
			name:     "valid positive",
			unitName: "км",
			fullName: "километр",
			toBase:   1000,
			wantBase: "1000",
		},
		{
			name:     "one",
			unitName: "м",
			fullName: "метр",
			toBase:   1,
			wantBase: "1",
		},
		{
			name:     "zero defaults to 1",
			unitName: "шт",
			fullName: "штука",
			toBase:   0,
			wantBase: "1",
		},
		{
			name:     "negative defaults to 1",
			unitName: "пак",
			fullName: "пакет",
			toBase:   -100,
			wantBase: "1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := NewUnitItemFromInt(tt.unitName, tt.fullName, tt.toBase)
			require.NotNil(t, item)
			assert.Equal(t, tt.unitName, item.Name)
			assert.Equal(t, tt.fullName, item.FullName)
			assert.Equal(t, tt.wantBase, item.ToBase.String())
		})
	}
}

// TestConvertToBase проверяет конвертацию в базовые единицы
func TestConvertToBase(t *testing.T) {
	tests := []struct {
		name  string
		item  *UnitItem
		value decimal.Decimal
		want  string
	}{
		{
			name:  "km to m",
			item:  NewUnitItemFromInt("км", "километр", 1000),
			value: decimal.NewFromInt(5),
			want:  "5000",
		},
		{
			name:  "cm to m",
			item:  NewUnitItemFromFloat("см", "сантиметр", 0.01),
			value: decimal.NewFromInt(250),
			want:  "2.5",
		},
		{
			name:  "base unit",
			item:  NewUnitItemFromInt("м", "метр", 1),
			value: decimal.NewFromInt(100),
			want:  "100",
		},
		{
			name:  "fractional value",
			item:  NewUnitItemFromInt("км", "километр", 1000),
			value: decimal.NewFromFloat(2.5),
			want:  "2500",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.item.ConvertToBase(tt.value)
			assert.Equal(t, tt.want, result.String())
		})
	}
}

// TestConvertFromBase проверяет конвертацию из базовых единиц
func TestConvertFromBase(t *testing.T) {
	tests := []struct {
		name      string
		item      *UnitItem
		baseValue decimal.Decimal
		want      string
	}{
		{
			name:      "m to km",
			item:      NewUnitItemFromInt("км", "километр", 1000),
			baseValue: decimal.NewFromInt(5000),
			want:      "5",
		},
		{
			name:      "m to cm",
			item:      NewUnitItemFromFloat("см", "сантиметр", 0.01),
			baseValue: decimal.NewFromFloat(2.5),
			want:      "250",
		},
		{
			name:      "base unit",
			item:      NewUnitItemFromInt("м", "метр", 1),
			baseValue: decimal.NewFromInt(100),
			want:      "100",
		},
		{
			name:      "fractional result",
			item:      NewUnitItemFromInt("км", "километр", 1000),
			baseValue: decimal.NewFromInt(2500),
			want:      "2.5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.item.ConvertFromBase(tt.baseValue)
			assert.Equal(t, tt.want, result.String())
		})
	}
}

// TestUnitItemString проверяет строковое представление единицы
func TestUnitItemString(t *testing.T) {
	tests := []struct {
		name string
		item *UnitItem
		want string
	}{
		{
			name: "kilometer",
			item: NewUnitItemFromInt("км", "километр", 1000),
			want: "км (1 км = 1000 базовых единиц)",
		},
		{
			name: "centimeter",
			item: NewUnitItemFromFloat("см", "сантиметр", 0.01),
			want: "см (1 см = 0.01 базовых единиц)",
		},
		{
			name: "base unit",
			item: NewUnitItemFromInt("м", "метр", 1),
			want: "м (1 м = 1 базовых единиц)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.item.String()
			assert.Equal(t, tt.want, result)
		})
	}
}
