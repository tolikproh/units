package main

import (
	"fmt"

	"github.com/tolikproh/units"
)

func main() {
	// Создаём единицу измерения для кабельных стяжек (базовая: штуки)
	tiesUnit := units.New("шт", "штука")
	tiesUnit.AddUnit("упаковка", "упаковка", 100) // 1 упаковка = 100 штук
	tiesUnit.AddUnit("коробка", "коробка", 12000) // 1 коробка = 12000 штук (120 упаковок)

	// Есть 1 коробка на складе
	quantity, _ := tiesUnit.ToBase("коробка", 1)
	baseStr, _ := tiesUnit.StringBase(quantity)
	fmt.Printf("На складе: %s %s\n", baseStr, tiesUnit.List()[0].FullName)

	// Отгрузили 50 упаковок
	quantity, _ = tiesUnit.Sub("упаковка", quantity, 50)
	baseStr, _ = tiesUnit.StringBase(quantity)
	unitStr, _ := tiesUnit.StringUnit("упаковка", quantity)
	fmt.Printf("После отгрузки: %s %s (%s упаковок)\n",
		baseStr,
		tiesUnit.List()[0].FullName,
		unitStr)
}
