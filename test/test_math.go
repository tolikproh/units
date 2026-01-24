package main

import (
	"fmt"

	"github.com/tolikproh/units"
)

func testMath() {
	// Создаем единицы измерения: базовая - метр
	u := units.New("м", "метр")
	u.AddFromInt("км", "километр", 1000)
	u.AddFromFloat("см", "сантиметр", 0.01)

	fmt.Println("=== Тестирование математических операций ===")
	fmt.Println("Базовая единица: метр")
	fmt.Println()

	// Начинаем с 1000 метров
	quantity := 1000.0
	fmt.Printf("Начальное количество: %v метров\n", quantity)
	fmt.Println()

	// Тест 1: Добавление километров
	fmt.Println("Тест 1: Add")
	fmt.Printf("Добавляем 2 км к %v м\n", quantity)
	if result, err := u.Add(quantity, "км", 2); err == nil {
		quantity = result.InexactFloat64()
		fmt.Printf("Результат: %v м (ожидается 3000, т.к. 1000 + 2*1000)\n", quantity)
	} else {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println()

	// Тест 2: Вычитание
	fmt.Println("Тест 2: Sub")
	fmt.Printf("Вычитаем 1 км из %v м\n", quantity)
	if result, err := u.Sub(quantity, "км", 1); err == nil {
		quantity = result.InexactFloat64()
		fmt.Printf("Результат: %v м (ожидается 2000, т.к. 3000 - 1*1000)\n", quantity)
	} else {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println()

	// Тест 3: Умножение
	fmt.Println("Тест 3: Mul")
	fmt.Printf("Умножаем %v м на 3\n", quantity)
	if result, err := u.Mul(quantity, 3); err == nil {
		quantity = result.InexactFloat64()
		fmt.Printf("Результат: %v м (ожидается 6000, т.к. 2000 * 3)\n", quantity)
	} else {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println()

	// Тест 4: Деление
	fmt.Println("Тест 4: Div")
	fmt.Printf("Делим %v м на 2\n", quantity)
	if result, err := u.Div(quantity, 2); err == nil {
		quantity = result.InexactFloat64()
		fmt.Printf("Результат: %v м (ожидается 3000, т.к. 6000 / 2)\n", quantity)
	} else {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println()

	// Тест 5: Добавление сантиметров
	fmt.Println("Тест 5: Add с маленькими единицами")
	fmt.Printf("Добавляем 50 см к %v м\n", quantity)
	if result, err := u.Add(quantity, "см", 50); err == nil {
		quantity = result.InexactFloat64()
		fmt.Printf("Результат: %v м (ожидается 3000.5, т.к. 3000 + 50*0.01)\n", quantity)
	} else {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println()

	// Тест 6: Вывод в разных единицах
	fmt.Println("Тест 6: Вывод результата в различных единицах")
	if formatted, err := u.FormatInBaseUnit(quantity); err == nil {
		fmt.Println("В метрах:", formatted)
	}
	if formatted, err := u.FormatInUnit("км", quantity); err == nil {
		fmt.Println("В километрах:", formatted)
	}
	if formatted, err := u.FormatInUnit("см", quantity); err == nil {
		fmt.Println("В сантиметрах:", formatted)
	}
}

func main() {
	testMath()
}
