package main

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/tolikproh/units"
)

func main() {
	// Создаём единицу измерения для кабеля (базовая: метры)
	cableUnit := units.New("м", "метр")
	cableUnit.AddUnit("км", "километр", 1000)
	cableUnit.AddUnit("бухта", "бухта", 200)

	fmt.Println("=== Математические операции с единицами измерения ===")
	fmt.Println()

	// Начальное количество
	quantity := decimal.NewFromFloat(1500) // 1500 метров
	baseStr, _ := cableUnit.StringBase(quantity)
	unitStr, _ := cableUnit.StringUnit("бухта", quantity)
	fmt.Printf("Начальное количество: %s %s (%s бухт)\n\n", baseStr, cableUnit.List()[0].FullName, unitStr)

	// 1. Сложение (Add) - добавляем 500 метров
	fmt.Println("--- 1. Сложение (Add) ---")
	quantity, _ = cableUnit.Add("м", quantity, 500)
	baseStr, _ = cableUnit.StringBase(quantity)
	unitStr, _ = cableUnit.StringUnit("бухта", quantity)
	fmt.Printf("+ 500 метров = %s %s (%s бухт)\n\n", baseStr, cableUnit.List()[0].FullName, unitStr)

	// Добавляем 2 бухты (400 метров)
	fmt.Println("--- 1.1. Сложение в других единицах ---")
	quantity, _ = cableUnit.Add("бухта", quantity, 2)
	baseStr, _ = cableUnit.StringBase(quantity)
	unitStr, _ = cableUnit.StringUnit("км", quantity)
	fmt.Printf("+ 2 бухты = %s %s (%s км)\n\n", baseStr, cableUnit.List()[0].FullName, unitStr)

	// 2. Вычитание (Sub) - отгружаем 1.5 км
	fmt.Println("--- 2. Вычитание (Sub) ---")
	quantity, _ = cableUnit.Sub("км", quantity, 1.5)
	baseStr, _ = cableUnit.StringBase(quantity)
	unitStr, _ = cableUnit.StringUnit("бухта", quantity)
	fmt.Printf("- 1.5 км = %s %s (%s бухт)\n\n", baseStr, cableUnit.List()[0].FullName, unitStr)

	// 3. Умножение (Mul) - удвоили количество
	fmt.Println("--- 3. Умножение (Mul) ---")
	fmt.Printf("Текущее количество: %s м\n", baseStr)
	quantity, _ = cableUnit.Mul("м", quantity, 2)
	baseStr, _ = cableUnit.StringBase(quantity)
	unitStr, _ = cableUnit.StringUnit("км", quantity)
	fmt.Printf("× 2 = %s %s (%s км)\n\n", baseStr, cableUnit.List()[0].FullName, unitStr)

	// Умножение на дробное число
	fmt.Println("--- 3.1. Умножение на дробное число ---")
	quantity, _ = cableUnit.Mul("м", quantity, 0.5)
	baseStr, _ = cableUnit.StringBase(quantity)
	unitStr, _ = cableUnit.StringUnit("бухта", quantity)
	fmt.Printf("× 0.5 = %s %s (%s бухт)\n\n", baseStr, cableUnit.List()[0].FullName, unitStr)

	// 4. Деление (Div) - разделили на 3 части
	fmt.Println("--- 4. Деление (Div) ---")
	fmt.Printf("Текущее количество: %s м\n", baseStr)
	quantity, _ = cableUnit.Div("м", quantity, 3)
	baseStr, _ = cableUnit.StringBase(quantity)
	unitStr, _ = cableUnit.StringUnit("бухта", quantity)
	fmt.Printf("÷ 3 = %s %s (%s бухт)\n\n", baseStr, cableUnit.List()[0].FullName, unitStr)

	// Практический пример: расчет остатка после нескольких операций
	fmt.Println("=== Практический пример: управление запасами ===")
	fmt.Println()

	// Начальный запас
	stock := decimal.NewFromFloat(5000) // 5000 метров
	baseStr, _ = cableUnit.StringBase(stock)
	fmt.Printf("Начальный запас: %s м\n", baseStr)

	// Поступление 3 бухт
	stock, _ = cableUnit.Add("бухта", stock, 3)
	baseStr, _ = cableUnit.StringBase(stock)
	fmt.Printf("+ Поступление 3 бухт: %s м\n", baseStr)

	// Отгрузка 2 км
	stock, _ = cableUnit.Sub("км", stock, 2)
	baseStr, _ = cableUnit.StringBase(stock)
	fmt.Printf("- Отгрузка 2 км: %s м\n", baseStr)

	// Списание 5% на отходы
	stock, _ = cableUnit.Mul("м", stock, 0.95)
	baseStr, _ = cableUnit.StringBase(stock)
	unitStr, _ = cableUnit.StringUnit("бухта", stock)
	fmt.Printf("- Списание 5%%: %s м (%s бухт)\n\n", baseStr, unitStr)

	// Расчет необходимого количества для заказа
	fmt.Println("=== Расчет заказа ===")
	requiredStock := decimal.NewFromFloat(10000) // Нужно 10000 метров
	needToOrder, _ := cableUnit.Sub("м", requiredStock, stock)
	baseStr, _ = cableUnit.StringBase(needToOrder)
	unitStr, _ = cableUnit.StringUnit("бухта", needToOrder)
	stockStr, _ := cableUnit.StringBase(stock)
	fmt.Printf("Требуется на складе: 10000 м\n")
	fmt.Printf("Текущий остаток: %s м\n", stockStr)
	fmt.Printf("Необходимо заказать: %s м (%s бухт)\n", baseStr, unitStr)
}
