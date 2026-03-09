# Units - Библиотека для работы с единицами измерения

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Coverage](https://img.shields.io/badge/coverage-99.4%25-brightgreen.svg)](https://github.com/tolikproh/units)

Библиотека для точной работы с различными единицами измерения в Go. Использует `decimal.Decimal` для высокоточных вычислений без потери точности при работе с дробными числами.

## Возможности

- ✅ Точные вычисления с использованием `decimal.Decimal`
- ✅ Поддержка произвольных единиц измерения
- ✅ Автоматическая конвертация между единицами
- ✅ Математические операции (сложение, вычитание, умножение, деление)
- ✅ Настраиваемая точность отображения (по умолчанию 3 знака)
- ✅ Умное форматирование (целые числа без десятичных знаков, удаление конечных нулей)
- ✅ Поддержка JSON сериализации/десериализации
- ✅ Покрытие тестами 99.4%

## Установка

```bash
go get github.com/tolikproh/units
```

## Быстрый старт

```go
package main

import (
    "fmt"
    "github.com/tolikproh/units"
)

func main() {
    // Создаём единицу измерения для кабеля
    cable := units.New("м", "метр")
    cable.AddUnit("км", "километр", 1000)
    cable.AddUnit("бухта", "бухта", 200)
    
    // Конвертируем 2.5 км в метры
    meters, _ := cable.ToBase("км", 2.5)
    fmt.Println(meters) // 2500
    
    // Выводим в разных единицах
    result, _ := cable.StringBase(2500)
    fmt.Println(result) // "2500"
    
    result, _ = cable.StringUnit("км", 2500)
    fmt.Println(result) // "2.5"
    
    result, _ = cable.StringUnit("бухта", 2500)
    fmt.Println(result) // "12.5"
}
```

## Основные типы

### Unit

Основная структура для работы с единицами измерения.

```go
type Unit struct {
    Base       *UnitItem            // Базовая единица измерения
    Additional map[string]*UnitItem // Дополнительные единицы
    Precision  *int32               // Точность отображения (количество знаков после запятой), nil = 3
}
```

### UnitItem

Представляет одну единицу измерения с её названием и коэффициентом конвертации.

```go
type UnitItem struct {
    Name     string          // Краткое название ("м", "км")
    FullName string          // Полное название ("метр", "километр")
    ToBase   decimal.Decimal // Коэффициент конвертации в базовую единицу
}
```

## API

### Создание единицы измерения

#### `New(name, fullName string) *Unit`

Создаёт новую единицу измерения с базовой единицей.

```go
distance := units.New("м", "метр")
```

#### `NewJSON(data []byte) (*Unit, error)`

Создаёт единицу измерения из JSON.

```go
jsonData := []byte(`{
    "base": {
        "Name": "м",
        "FullName": "метр",
        "ToBase": "1"
    },
    "additional": {
        "км": {
            "Name": "км",
            "FullName": "километр",
            "ToBase": "1000"
        }
    },
    "precision": 3
}`)
unit, err := units.NewJSON(jsonData)
```

### Добавление дополнительных единиц

#### `AddUnit(name, fullName string, toBase any) error`

Добавляет дополнительную единицу измерения.

```go
// toBase может быть: int, int64, float64, string, decimal.Decimal
distance.AddUnit("км", "километр", 1000)
distance.AddUnit("см", "сантиметр", 0.01)
distance.AddUnit("миля", "миля", 1609.34)
```

### Конвертация значений

#### `ToBase(unitName string, value any) (decimal.Decimal, error)`

Конвертирует значение из указанной единицы в базовую.

```go
meters, err := distance.ToBase("км", 2.5) // 2500 метров
```

#### `StringBase(value any) (string, error)`

Возвращает строковое представление значения в базовой единице.

```go
result, _ := distance.StringBase(1500) // "1500"
```

#### `StringUnit(unitName string, value any) (string, error)`

Возвращает строковое представление значения в указанной единице.

```go
result, _ := distance.StringUnit("км", 1500) // "1.5"
```

### Математические операции

#### `Add(unitName string, currentQuantity any, value any) (decimal.Decimal, error)`

Складывает текущее количество со значением в указанной единице.

```go
// Есть 1000м, добавляем 2 км
result, _ := distance.Add("км", 1000, 2) // 3000 метров
```

#### `Sub(unitName string, currentQuantity any, value any) (decimal.Decimal, error)`

Вычитает значение из текущего количества.

```go
// Есть 5000м, вычитаем 1.5 км
result, _ := distance.Sub("км", 5000, 1.5) // 3500 метров
```

#### `Mul(unitName string, currentQuantity any, multiplier any) (decimal.Decimal, error)`

Умножает текущее количество на коэффициент.

```go
// Удваиваем количество
result, _ := distance.Mul("м", 1000, 2) // 2000 метров
```

#### `Div(unitName string, currentQuantity any, divisor any) (decimal.Decimal, error)`

Делит текущее количество на делитель.

```go
// Делим на 4 части
result, _ := distance.Div("м", 1000, 4) // 250 метров
```

### Настройка точности

#### `SetPrecision(precision int32)`

Устанавливает точность отображения (количество знаков после запятой).

```go
distance.SetPrecision(2) // 2 знака после запятой
```

#### `GetPrecision() int32`

Возвращает текущую точность отображения.

```go
precision := distance.GetPrecision() // по умолчанию 3
```

### Дополнительные методы

#### `List() []*UnitItem`

Возвращает список всех единиц измерения.

```go
units := distance.List()
for _, u := range units {
    fmt.Printf("%s (%s): %s\n", u.Name, u.FullName, u.ToBase)
}
```

#### `ToJSON() ([]byte, error)`

Сериализует единицу измерения в JSON. Сохраняет все данные включая точность.

```go
jsonData, err := distance.ToJSON()
// {"base":{"Name":"м","FullName":"метр","ToBase":"1"},"additional":{"км":{"Name":"км","FullName":"километр","ToBase":"1000"}},"precision":3}
```

### JSON сериализация и точность

При сохранении Unit в JSON сохраняется текущая точность. При восстановлении из JSON:
- Если `precision` присутствует в JSON, используется сохранённое значение (включая 0)
- Если `precision` отсутствует в JSON, используется дефолт (3 знака)

```go
// Создаём и устанавливаем кастомную точность
distance := units.New("м", "метр")
distance.SetPrecision(5)
distance.AddUnit("км", "километр", 1000)

// Сохраняем в JSON (precision=5 будет сохранён)
jsonData, _ := distance.ToJSON()

// Восстанавливаем из JSON (precision автоматически восстановится)
restored, _ := units.NewJSON(jsonData)
fmt.Println(restored.GetPrecision()) // 5
```

## Примеры использования

В папке `examples/` доступно 11 примеров:

1. **01_cable_ppg** - Кабель ППГ (метры, километры, бухты)
2. **02_cable_utp** - Кабель UTP (метры, коробки, паллеты)
3. **03_fiber_optic** - Оптоволокно (метры, катушки)
4. **04_cable_duct** - Кабель-канал (метры, штуки)
5. **05_tray** - Лоток (метры, штуки)
6. **06_smoke_detector** - Дымовые извещатели (штуки, коробки)
7. **07_cable_ties** - Кабельные стяжки (штуки, упаковки, коробки)
8. **08_screws** - Шурупы (штуки, килограммы)
9. **09_warehouse_complex** - Комплексный склад (несколько типов товаров)
10. **10_precision** - Настройка точности отображения
11. **11_math_operations** - Математические операции

### Запуск примеров

```bash
# Запустить все примеры
make examples

# Запустить конкретный пример
make example-01
make example-11
```

### Пример: Управление складом кабеля

```go
package main

import (
    "fmt"
    "github.com/shopspring/decimal"
    "github.com/tolikproh/units"
)

func main() {
    // Создаём единицу для кабеля
    cable := units.New("м", "метр")
    cable.AddUnit("км", "километр", 1000)
    cable.AddUnit("бухта", "бухта", 200)
    
    // Начальный запас: 5000 метров
    stock := decimal.NewFromFloat(5000)
    
    // Поступление: 3 бухты (600 метров)
    stock, _ = cable.Add("бухта", stock, 3)
    fmt.Printf("После поступления: %s м\n", stock) // 5600 м
    
    // Отгрузка: 2 км
    stock, _ = cable.Sub("км", stock, 2)
    fmt.Printf("После отгрузки: %s м\n", stock) // 3600 м
    
    // Списание 5% на отходы
    stock, _ = cable.Mul("м", stock, 0.95)
    
    // Вывод в разных единицах
    meters, _ := cable.StringBase(stock)
    km, _ := cable.StringUnit("км", stock)
    coils, _ := cable.StringUnit("бухта", stock)
    
    fmt.Printf("Остаток: %s м = %s км = %s бухт\n", meters, km, coils)
    // Остаток: 3420 м = 3.42 км = 17.1 бухт
}
```

### Пример: Настройка точности

```go
package main

import (
    "fmt"
    "github.com/tolikproh/units"
)

func main() {
    weight := units.New("кг", "килограмм")
    weight.AddUnit("г", "грамм", 0.001)
    
    value := 123.456789
    
    // По умолчанию: 3 знака
    result, _ := weight.StringBase(value)
    fmt.Println(result) // "123.457"
    
    // Точность 0 знаков
    weight.SetPrecision(0)
    result, _ = weight.StringBase(value)
    fmt.Println(result) // "123"
    
    // Точность 5 знаков
    weight.SetPrecision(5)
    result, _ = weight.StringBase(value)
    fmt.Println(result) // "123.45679"
    
    // Целые числа всегда без десятичных знаков
    result, _ = weight.StringBase(100)
    fmt.Println(result) // "100" (не "100.000")
    
    // Удаление конечных нулей
    weight.SetPrecision(3)
    result, _ = weight.StringBase(2.5)
    fmt.Println(result) // "2.5" (не "2.500")
}
```

## Форматирование чисел

Библиотека использует умное форматирование:

- **Целые числа**: выводятся без десятичной части (`100`, а не `100.000`)
- **Дробные числа**: округляются до заданной точности (`123.457`)
- **Конечные нули**: автоматически удаляются (`2.5`, а не `2.500`)

## Тестирование

```bash
# Запустить все тесты
make test

# Запустить тесты с покрытием
make coverage

# Запустить тесты напрямую
go test -v ./...
```

Текущее покрытие: **99.4%**

## Makefile команды

```bash
make help          # Показать все доступные команды
make test          # Запустить тесты
make coverage      # Запустить тесты с покрытием
make examples      # Запустить все примеры
make example-01    # Запустить пример 01
make example-11    # Запустить пример 11
make clean         # Очистить бинарные файлы
```

## Зависимости

- [github.com/shopspring/decimal](https://github.com/shopspring/decimal) - Точные десятичные вычисления
- [github.com/stretchr/testify](https://github.com/stretchr/testify) - Тестирование (dev)

## Структура проекта

```
.
├── go.mod                          # Зависимости проекта
├── units.go                        # Основная логика библиотеки
├── unit_math.go                    # Математические операции
├── unititem.go                     # Структура UnitItem
├── units_test.go                   # Тесты
├── Makefile                        # Команды для сборки и тестирования
├── README.md                       # Документация (этот файл)
├── LICENSE                         # Лицензия MIT (английский)
├── LICENSE.ru                      # Лицензия MIT (русский)
└── examples/                       # Примеры использования
    ├── 01_cable_ppg/
    ├── 02_cable_utp/
    ├── 03_fiber_optic/
    ├── 04_cable_duct/
    ├── 05_tray/
    ├── 06_smoke_detector/
    ├── 07_cable_ties/
    ├── 08_screws/
    ├── 09_warehouse_complex/
    ├── 10_precision/
    └── 11_math_operations/
```

## Лицензия

MIT License - см. [LICENSE](LICENSE) или [LICENSE.ru](LICENSE.ru) для русской версии.

## Автор

[tolikproh](https://github.com/tolikproh)

## Вклад в проект

Приветствуются Pull Request'ы и Issue'ы! Пожалуйста, убедитесь, что:

1. Все тесты проходят: `make test`
2. Покрытие тестами не снижается
3. Код соответствует стилю Go
4. Добавлены комментарии на русском языке

## Дорожная карта

- [ ] Поддержка валют с курсами обмена
- [ ] Кеширование конвертаций
- [ ] Batch операции для массовых конвертаций
- [ ] CLI утилита для конвертации
- [ ] Web API для работы с единицами

## Changelog

### v1.0.0 (2026-01-25)

- ✅ Первый релиз
- ✅ Основная функциональность конвертации
- ✅ Математические операции (Add, Sub, Mul, Div)
- ✅ Настраиваемая точность
- ✅ Умное форматирование чисел
- ✅ 11 примеров использования
- ✅ Покрытие тестами 99.4%
- ✅ JSON сериализация/десериализация
