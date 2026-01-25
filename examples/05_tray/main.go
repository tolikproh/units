package main

import (
	"fmt"

	"github.com/tolikproh/units"
)

func main() {
	// Создаём единицу измерения для лотка (базовая: метры)
	trayUnit := units.New("м", "метр")
	trayUnit.AddUnit("шт", "штука", 3)        // 1 штука = 3 метра
	trayUnit.AddUnit("коробка", "коробка", 6) // 1 коробка = 6 метров (2 штуки)

	// Есть 10 коробок на складе
	quantity, _ := trayUnit.ToBase("коробка", 10)
	baseStr, _ := trayUnit.StringBase(quantity)
	fmt.Printf("На складе: %s %s\n", baseStr, trayUnit.List()[0].FullName)

	// Отгрузили 25 метров
	quantity, _ = trayUnit.Sub("м", quantity, 25)
	baseStr, _ = trayUnit.StringBase(quantity)
	unitStr, _ := trayUnit.StringUnit("шт", quantity)
	fmt.Printf("После отгрузки: %s %s (%s штук)\n",
		baseStr,
		trayUnit.List()[0].FullName,
		unitStr)
}
