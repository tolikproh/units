package main

import (
	"fmt"

	"github.com/tolikproh/units"
)

func main() {
	// Создаём единицу измерения для дымовых извещателей (базовая: штуки)
	detectorUnit := units.New("шт", "штука")
	detectorUnit.AddUnit("упаковка", "упаковка", 10) // 1 упаковка = 10 штук
	detectorUnit.AddUnit("коробка", "коробка", 120)  // 1 коробка = 120 штук (12 упаковок)

	// Есть 2 коробки на складе
	quantity, _ := detectorUnit.ToBase("коробка", 2)
	baseStr, _ := detectorUnit.StringBase(quantity)
	fmt.Printf("На складе: %s %s\n", baseStr, detectorUnit.List()[0].FullName)

	// Отгрузили 5 упаковок
	quantity, _ = detectorUnit.Sub("упаковка", quantity, 5)
	baseStr, _ = detectorUnit.StringBase(quantity)
	unitStr, _ := detectorUnit.StringUnit("коробка", quantity)
	fmt.Printf("После отгрузки: %s %s (%s коробок)\n",
		baseStr,
		detectorUnit.List()[0].FullName,
		unitStr)
}
