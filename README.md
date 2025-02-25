# Units - Библиотека для работы с единицами измерения

Библиотека предоставляет удобный интерфейс для работы с различными единицами измерения и их преобразованиями.

## Основные возможности

1. **Поддержка различных типов величин:**

   - Длина (Length)
   - Вес (Weight)
   - Штуки (Things)
   - Комплекты (Set)
   - Упаковки (Package)
   - Бухты (Bay)
2. **Поддержка префиксов СИ:**

   - Нано (10⁻⁹)
   - Микро (10⁻⁶)
   - Милли (10⁻³)
   - Нормальный (1)
   - Кило (10³)
   - Мега (10⁶)
   - Гига (10⁹)
3. **Арифметические операции:**

   - Сложение (Add)
   - Вычитание (Sub)
   - Умножение (Mul)
   - Деление (Div)
4. **Форматирование:**

   - Настраиваемое количество знаков после запятой
   - Автоматическое округление
   - Поддержка отрицательных значений
   - Вывод с единицами измерения

## Пример использования

```go

// Создание величин

l1 := units.NewLength(5, units.Kilo) // 5 километров

l2 := units.NewLength(300, units.Normal) // 300 метров

// Арифметические операции

sum := l1.Add(l2) // 5.3 километров

diff := l1.Sub(l2) // 4.7 километров

// Изменение префикса

sum.SetPrefix(units.Mega) // 0.0053 мегаметров

// Форматирование

sum.SetDecimals(3) // установка 3 знаков после запятой

fmt.Println(sum.String()) // вывод отформатированного значения
```

## Особенности реализации

- Внутреннее хранение значений в базовых единицах
- Автоматическая конвертация между префиксами
- Проверка совместимости типов при операциях
- Контроль ошибок через флаг Ok()
- Поддержка пользовательских единиц измерения
- Потокобезопасность операций
- Эффективное использование памяти

## Применение

Библиотека может использоваться в различных областях:

- Инженерные расчеты
- Научные вычисления
- Складской учет
- Торговые операции
- Производственные системы
- Метрологические системы
- Конвертация единиц измерения

## Структура проекта

```bash
.

├── cmd/

│ └── main.go # Примеры использования

├── bay.go # Реализация типа Bay (бухты)

├── go.mod # Описание модуля

├── package.go # Реализация типа Package (упаковки)

├── prefix.go # Определение префиксов СИ

├── quantity.go # Основной интерфейс и базовая реализация

├── set.go # Реализация типа Set (комплекты)

└── weight.go # Реализация типа Weight (вес)
```

## Лицензия

MIT
