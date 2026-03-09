package units

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNew проверяет создание нового набора единиц
func TestNew(t *testing.T) {
	u := New("м", "метр")
	require.NotNil(t, u)
	require.NotNil(t, u.Base)
	assert.Equal(t, "м", u.Base.Name)
	assert.Equal(t, "метр", u.Base.FullName)
	assert.True(t, u.Base.ToBase.Equal(decimal.NewFromInt(1)))
	assert.NotNil(t, u.Additional)
	assert.Empty(t, u.Additional)
}

// TestNewJSON проверяет создание набора единиц из JSON
func TestNewJSON(t *testing.T) {
	tests := []struct {
		name     string
		json     string
		wantErr  bool
		wantPrec int32
	}{
		{
			name: "valid json with precision",
			json: `{
				"base": {"Name": "м", "FullName": "метр", "ToBase": "1"},
				"additional": {
					"км": {"Name": "км", "FullName": "километр", "ToBase": "1000"}
				},
				"precision": 5
			}`,
			wantErr:  false,
			wantPrec: 5,
		},
		{
			name: "valid json without precision (default)",
			json: `{
				"base": {"Name": "м", "FullName": "метр", "ToBase": "1"},
				"additional": {}
			}`,
			wantErr:  false,
			wantPrec: 3,
		},
		{
			name:     "invalid json",
			json:     `{invalid}`,
			wantErr:  true,
			wantPrec: 0,
		},
		{
			name: "nil additional units",
			json: `{
				"base": {"Name": "шт", "FullName": "штука", "ToBase": "1"},
				"precision": 2
			}`,
			wantErr:  false,
			wantPrec: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := NewJSON([]byte(tt.json))
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, u)
			} else {
				assert.NoError(t, err)
				require.NotNil(t, u)
				assert.NotNil(t, u.Additional)
				assert.Equal(t, tt.wantPrec, u.Precision)
			}
		})
	}
}

// TestToJSON проверяет сохранение Unit в JSON с сохранением Precision
func TestToJSONWithPrecision(t *testing.T) {
	tests := []struct {
		name      string
		setupUnit func() *Unit
		wantPrec  int32
	}{
		{
			name: "default precision",
			setupUnit: func() *Unit {
				return New("м", "метр")
			},
			wantPrec: 3,
		},
		{
			name: "custom precision",
			setupUnit: func() *Unit {
				u := New("м", "метр")
				u.SetPrecision(5)
				u.AddUnit("км", "километр", 1000)
				return u
			},
			wantPrec: 5,
		},
		{
			name: "zero precision",
			setupUnit: func() *Unit {
				u := New("м", "метр")
				u.SetPrecision(0)
				return u
			},
			wantPrec: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.setupUnit()

			// Сериализуем в JSON
			data, err := u.ToJSON()
			assert.NoError(t, err)
			assert.NotNil(t, data)

			// Десериализуем обратно
			u2, err := NewJSON(data)
			assert.NoError(t, err)
			require.NotNil(t, u2)

			// Проверяем что Precision сохранился
			if tt.wantPrec == 0 {
				// Если был 0, должен стать 3 (default)
				assert.Equal(t, int32(3), u2.Precision)
			} else {
				assert.Equal(t, tt.wantPrec, u2.Precision)
			}
		})
	}
}

// TestAddUnit проверяет добавление дополнительной единицы
func TestAddUnit(t *testing.T) {
	tests := []struct {
		name     string
		unitName string
		fullName string
		toBase   any
		wantErr  bool
	}{
		{
			name:     "valid int",
			unitName: "км",
			fullName: "километр",
			toBase:   1000,
			wantErr:  false,
		},
		{
			name:     "valid float",
			unitName: "см",
			fullName: "сантиметр",
			toBase:   0.01,
			wantErr:  false,
		},
		{
			name:     "valid string",
			unitName: "дм",
			fullName: "дециметр",
			toBase:   "0.1",
			wantErr:  false,
		},
		{
			name:     "invalid type",
			unitName: "мм",
			fullName: "миллиметр",
			toBase:   []int{1, 2, 3},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := New("м", "метр")
			err := u.AddUnit(tt.unitName, tt.fullName, tt.toBase)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Contains(t, u.Additional, tt.unitName)
			}
		})
	}

	t.Run("conflict with base unit", func(t *testing.T) {
		u := New("м", "метр")
		err := u.addItem(NewUnitItemFromInt("м", "метр2", 2))
		assert.Error(t, err)
	})

	t.Run("duplicate unit", func(t *testing.T) {
		u := New("м", "метр")
		err := u.AddUnit("км", "километр", 1000)
		assert.NoError(t, err)
		err = u.AddUnit("км", "километр2", 2000)
		assert.Error(t, err)
	})

	t.Run("nil unit", func(t *testing.T) {
		u := New("м", "метр")
		err := u.addItem(nil)
		assert.Error(t, err)
	})

	t.Run("uninitialized unit", func(t *testing.T) {
		u := &Unit{}
		err := u.AddUnit("км", "километр", 1000)
		assert.Error(t, err)
	})
}

// TestToJSON проверяет сериализацию в JSON
func TestToJSON(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() *Unit
		wantErr bool
	}{
		{
			name: "valid unit",
			setup: func() *Unit {
				u := New("м", "метр")
				u.AddUnit("км", "километр", 1000)
				return u
			},
			wantErr: false,
		},
		{
			name: "nil unit",
			setup: func() *Unit {
				return nil
			},
			wantErr: false,
		},
		{
			name: "nil data",
			setup: func() *Unit {
				return &Unit{}
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.setup()
			data, err := u.ToJSON()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, data)
			}
		})
	}
}

// TestToBase проверяет конвертацию в базовые единицы
func TestToBase(t *testing.T) {
	u := New("м", "метр")
	u.AddUnit("км", "километр", 1000)
	u.AddUnit("см", "сантиметр", 0.01)

	tests := []struct {
		name     string
		unitName string
		val      any
		want     string
		wantErr  bool
	}{
		{
			name:     "from base",
			unitName: "м",
			val:      10,
			want:     "10",
			wantErr:  false,
		},
		{
			name:     "from km",
			unitName: "км",
			val:      2.5,
			want:     "2500",
			wantErr:  false,
		},
		{
			name:     "from cm",
			unitName: "см",
			val:      150,
			want:     "1.5",
			wantErr:  false,
		},
		{
			name:     "unit not found",
			unitName: "мм",
			val:      10,
			wantErr:  true,
		},
		{
			name:     "negative result",
			unitName: "м",
			val:      -10,
			wantErr:  true,
		},
		{
			name:     "empty unit name",
			unitName: "",
			val:      10,
			wantErr:  true,
		},
		{
			name:     "invalid value type",
			unitName: "м",
			val:      []int{1, 2, 3},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := u.ToBase(tt.unitName, tt.val)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result.String())
			}
		})
	}

	t.Run("uninitialized unit", func(t *testing.T) {
		u := &Unit{}
		_, err := u.ToBase("м", 10)
		assert.Error(t, err)
	})
}

// TestStringBase проверяет вывод в базовых единицах
func TestStringBase(t *testing.T) {
	u := New("м", "метр")

	tests := []struct {
		name     string
		quantity any
		want     string
		wantErr  bool
	}{
		{
			name:     "valid int",
			quantity: 100,
			want:     "100",
			wantErr:  false,
		},
		{
			name:     "valid float",
			quantity: 123.45,
			want:     "123.45",
			wantErr:  false,
		},
		{
			name:     "valid string",
			quantity: "999.99",
			want:     "999.99",
			wantErr:  false,
		},
		{
			name:     "invalid type",
			quantity: []int{1, 2, 3},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := u.StringBase(tt.quantity)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result)
			}
		})
	}

	t.Run("uninitialized unit", func(t *testing.T) {
		u := &Unit{}
		_, err := u.StringBase(10)
		assert.Error(t, err)
	})

	t.Run("nil unit", func(t *testing.T) {
		var u *Unit
		_, err := u.StringBase(10)
		assert.Error(t, err)
	})
}

// TestStringUnit проверяет вывод в указанной единице
func TestStringUnit(t *testing.T) {
	u := New("м", "метр")
	u.AddUnit("км", "километр", 1000)
	u.AddUnit("см", "сантиметр", 0.01)

	tests := []struct {
		name     string
		unitName string
		quantity any
		want     string
		wantErr  bool
	}{
		{
			name:     "to base",
			unitName: "м",
			quantity: 1000,
			want:     "1000",
			wantErr:  false,
		},
		{
			name:     "to km",
			unitName: "км",
			quantity: 5000,
			want:     "5",
			wantErr:  false,
		},
		{
			name:     "to cm",
			unitName: "см",
			quantity: 2.5,
			want:     "250",
			wantErr:  false,
		},
		{
			name:     "unit not found",
			unitName: "мм",
			quantity: 100,
			wantErr:  true,
		},
		{
			name:     "invalid value",
			unitName: "м",
			quantity: []int{1, 2, 3},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := u.StringUnit(tt.unitName, tt.quantity)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result)
			}
		})
	}

	t.Run("uninitialized unit", func(t *testing.T) {
		u := &Unit{}
		_, err := u.StringUnit("м", 10)
		assert.Error(t, err)
	})

	t.Run("nil unit", func(t *testing.T) {
		var u *Unit
		_, err := u.StringUnit("м", 10)
		assert.Error(t, err)
	})
}

// TestList проверяет получение списка единиц
func TestList(t *testing.T) {
	t.Run("valid list", func(t *testing.T) {
		u := New("м", "метр")
		u.AddUnit("км", "километр", 1000)
		u.AddUnit("см", "сантиметр", 0.01)
		u.AddUnit("дм", "дециметр", 0.1)

		list := u.List()
		require.Len(t, list, 4)
		assert.Equal(t, "м", list[0].Name)
		assert.Equal(t, "дм", list[1].Name)
		assert.Equal(t, "км", list[2].Name)
		assert.Equal(t, "см", list[3].Name)
	})

	t.Run("nil unit", func(t *testing.T) {
		var u *Unit
		list := u.List()
		assert.Nil(t, list)
	})

	t.Run("nil data", func(t *testing.T) {
		u := &Unit{}
		list := u.List()
		assert.Nil(t, list)
	})

	t.Run("nil base unit", func(t *testing.T) {
		u := &Unit{}
		list := u.List()
		assert.Nil(t, list)
	})
}

// TestToDecimalValue проверяет конвертацию значений
func TestToDecimalValue(t *testing.T) {
	tests := []struct {
		name    string
		val     any
		want    string
		wantErr bool
	}{
		{
			name: "decimal.Decimal",
			val:  decimal.NewFromInt(100),
			want: "100",
		},
		{
			name: "*decimal.Decimal",
			val:  func() *decimal.Decimal { d := decimal.NewFromInt(200); return &d }(),
			want: "200",
		},
		{
			name:    "nil *decimal.Decimal",
			val:     (*decimal.Decimal)(nil),
			wantErr: true,
		},
		{
			name: "int",
			val:  42,
			want: "42",
		},
		{
			name: "int64",
			val:  int64(1000),
			want: "1000",
		},
		{
			name: "int32",
			val:  int32(500),
			want: "500",
		},
		{
			name: "int16",
			val:  int16(250),
			want: "250",
		},
		{
			name: "int8",
			val:  int8(127),
			want: "127",
		},
		{
			name: "uint",
			val:  uint(100),
			want: "100",
		},
		{
			name: "uint64",
			val:  uint64(5000),
			want: "5000",
		},
		{
			name: "uint64 max safe",
			val:  uint64(9223372036854775807), // max int64
			want: "9223372036854775807",
		},
		{
			name: "uint32",
			val:  uint32(2500),
			want: "2500",
		},
		{
			name: "uint16",
			val:  uint16(1000),
			want: "1000",
		},
		{
			name: "uint8",
			val:  uint8(255),
			want: "255",
		},
		{
			name: "float64",
			val:  123.45,
			want: "123.45",
		},
		{
			name: "float32",
			val:  float32(67.89),
			want: "67.88999938964844",
		},
		{
			name: "valid string",
			val:  "999.99",
			want: "999.99",
		},
		{
			name:    "invalid string",
			val:     "not-a-number",
			wantErr: true,
		},
		{
			name:    "unsupported type",
			val:     []int{1, 2, 3},
			wantErr: true,
		},
		{
			name:    "struct type",
			val:     struct{ x int }{x: 10},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToDecimalValue(tt.val)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result.String())
			}
		})
	}
}

// TestResolveUnitValue проверяет внутреннюю функцию разрешения единиц
func TestResolveUnitValue(t *testing.T) {
	u := New("м", "метр")
	u.AddUnit("км", "километр", 1000)

	tests := []struct {
		name     string
		unitName string
		val      any
		wantErr  bool
	}{
		{
			name:     "valid base unit",
			unitName: "м",
			val:      100,
			wantErr:  false,
		},
		{
			name:     "valid additional unit",
			unitName: "км",
			val:      5,
			wantErr:  false,
		},
		{
			name:     "empty unit name",
			unitName: "",
			val:      10,
			wantErr:  true,
		},
		{
			name:     "unit not found",
			unitName: "см",
			val:      10,
			wantErr:  true,
		},
		{
			name:     "invalid value",
			unitName: "м",
			val:      []int{1, 2, 3},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unit, dec, err := u.resolveUnitValue(tt.unitName, tt.val)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, unit)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, unit)
				assert.NotNil(t, dec)
			}
		})
	}

	t.Run("uninitialized unit", func(t *testing.T) {
		u := &Unit{}
		_, _, err := u.resolveUnitValue("м", 10)
		assert.Error(t, err)
	})

	t.Run("nil unit", func(t *testing.T) {
		var u *Unit
		_, _, err := u.resolveUnitValue("м", 10)
		assert.Error(t, err)
	})
}

// TestSetPrecision проверяет установку точности вывода
func TestSetPrecision(t *testing.T) {
	u := New("м", "метр")
	u.AddUnit("км", "километр", 1000)

	// Проверка значения по умолчанию
	assert.Equal(t, int32(3), u.GetPrecision())

	// Проверка вывода с точностью по умолчанию (3 знака)
	result, _ := u.StringBase(123.456789)
	assert.Equal(t, "123.457", result)

	// Целое число должно выводиться без дробной части
	result, _ = u.StringBase(100)
	assert.Equal(t, "100", result)

	// Число с незначащими нулями
	result, _ = u.StringBase(123.450)
	assert.Equal(t, "123.45", result)

	// Установка точности 0
	u.SetPrecision(0)
	assert.Equal(t, int32(0), u.GetPrecision())
	result, _ = u.StringBase(123.456789)
	assert.Equal(t, "123", result)

	// Установка точности 5
	u.SetPrecision(5)
	assert.Equal(t, int32(5), u.GetPrecision())
	result, _ = u.StringBase(123.456789)
	assert.Equal(t, "123.45679", result)

	// Целое число при любой точности
	result, _ = u.StringBase(1000)
	assert.Equal(t, "1000", result)

	// Проверка с StringUnit
	u.SetPrecision(2)
	result, _ = u.StringUnit("км", 5432.1)
	assert.Equal(t, "5.43", result)

	// Целое значение в StringUnit
	result, _ = u.StringUnit("км", 5000)
	assert.Equal(t, "5", result)

	// Значение с незначащими нулями в StringUnit
	result, _ = u.StringUnit("км", 2500)
	assert.Equal(t, "2.5", result)

	// Проверка для nil unit
	var nilUnit *Unit
	assert.Equal(t, int32(3), nilUnit.GetPrecision())
}
