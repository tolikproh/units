package main

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/tolikproh/units"
)

func main() {
	// Создаём единицу измерения для шурупов (базовая: штуки)
	screwsUnit := units.New("шт", "штука")
	// 1 кг = 500 штук (1 шуруп 3.5x25 = 0.002 кг = 2 грамма)
	screwsUnit.AddUnit("кг", "килограмм", decimal.NewFromFloat(1.0/0.002))

	// Есть 5 кг на складе
	quantity, _ := screwsUnit.ToBase("кг", 5)
	unitStr, _ := screwsUnit.StringUnit("кг", quantity)
	baseStr, _ := screwsUnit.StringBase(quantity)
	fmt.Printf("На складе: %s %s (%s %s)\n",
		unitStr,
		screwsUnit.List()[1].FullName,
		baseStr,
		screwsUnit.List()[0].FullName)

	// Отгрузили 30 штук
	quantity, _ = screwsUnit.Sub("шт", quantity, 30)
	baseStr, _ = screwsUnit.StringBase(quantity)
	unitStr, _ = screwsUnit.StringUnit("кг", quantity)
	fmt.Printf("После отгрузки: %s %s (%s %s)\n",
		baseStr,
		screwsUnit.List()[0].FullName,
		unitStr,
		screwsUnit.List()[1].FullName)
}
