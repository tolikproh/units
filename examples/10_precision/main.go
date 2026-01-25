package main

import (
	"fmt"

	"github.com/tolikproh/units"
)

func main() {
	// Создаём единицу измерения
	u := units.New("м", "метр")
	u.AddUnit("км", "километр", 1000)

	// Тестовые значения
	quantity1, _ := u.ToBase("км", 2.5)        // 2500 м = 2.5 км
	quantity2, _ := u.ToBase("м", 3456.784625) // дробное значение
	quantity3, _ := u.ToBase("км", 5)          // целое значение

	// По умолчанию точность = 3
	fmt.Println("=== Точность по умолчанию (3 знака) ===")
	base1, _ := u.StringBase(quantity1)
	unit1, _ := u.StringUnit("км", quantity1)
	fmt.Printf("2500 м = %s м = %s км\n", base1, unit1)
	base2, _ := u.StringBase(quantity2)
	fmt.Printf("3456.784625 м = %s м\n", base2)
	base3, _ := u.StringBase(quantity3)
	unit3, _ := u.StringUnit("км", quantity3)
	fmt.Printf("5000 м = %s м = %s км\n", base3, unit3)

	// Устанавливаем точность 0
	fmt.Println("\n=== Точность 0 знаков ===")
	u.SetPrecision(0)
	base1, _ = u.StringBase(quantity1)
	unit1, _ = u.StringUnit("км", quantity1)
	fmt.Printf("2500 м = %s м = %s км\n", base1, unit1)
	base2, _ = u.StringBase(quantity2)
	fmt.Printf("3456.784625 м = %s м\n", base2)

	// Устанавливаем точность 1
	fmt.Println("\n=== Точность 1 знак ===")
	u.SetPrecision(1)
	base2, _ = u.StringBase(quantity2)
	fmt.Printf("3456.784625 м = %s м\n", base2)

	// Устанавливаем точность 2
	fmt.Println("\n=== Точность 2 знака ===")
	u.SetPrecision(2)
	base1, _ = u.StringBase(quantity1)
	unit1, _ = u.StringUnit("км", quantity1)
	fmt.Printf("2500 м = %s м = %s км\n", base1, unit1)
	base2, _ = u.StringBase(quantity2)
	fmt.Printf("3456.784625 м = %s м\n", base2)

	// Устанавливаем точность 5
	fmt.Println("\n=== Точность 5 знаков ===")
	u.SetPrecision(5)
	base2, _ = u.StringBase(quantity2)
	fmt.Printf("3456.784625 м = %s м\n", base2)

	// Демонстрация удаления незначащих нулей
	fmt.Println("\n=== Удаление незначащих нулей ===")
	u.SetPrecision(3)
	testValues := []float64{100.000, 100.500, 100.510, 100.551, 123.456}
	for _, val := range testValues {
		qty, _ := u.ToBase("м", val)
		str, _ := u.StringBase(qty)
		fmt.Printf("%.3f м → %s м\n", val, str)
	}

	// Проверяем текущую точность
	fmt.Printf("\nТекущая точность: %d знаков после запятой\n", u.GetPrecision())
}
