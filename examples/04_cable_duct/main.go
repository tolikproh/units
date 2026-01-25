package main

import (
	"fmt"

	"github.com/tolikproh/units"
)

func main() {
	// Создаём единицу измерения для кабель-канала (базовая: метры)
	ductUnit := units.New("м", "метр")
	ductUnit.AddUnit("шт", "штука", 2)         // 1 штука = 2 метра
	ductUnit.AddUnit("коробка", "коробка", 96) // 1 коробка = 96 метров (48 штук)

	// Есть 5 коробок на складе
	quantity, _ := ductUnit.ToBase("коробка", 5)
	baseStr, _ := ductUnit.StringBase(quantity)
	fmt.Printf("На складе: %s %s\n", baseStr, ductUnit.List()[0].FullName)

	// Отгрузили 100 метров (50 отрезков)
	quantity, _ = ductUnit.Sub("м", quantity, 100)
	baseStr, _ = ductUnit.StringBase(quantity)
	unitStr, _ := ductUnit.StringUnit("шт", quantity)
	fmt.Printf("После отгрузки: %s %s (%s штук)\n",
		baseStr,
		ductUnit.List()[0].FullName,
		unitStr)
}
