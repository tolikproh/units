package main

import (
	"fmt"

	"github.com/tolikproh/units"
)

func main() {
	// Создаём единицу измерения для кабеля UTP (базовая: метры)
	cableUnit := units.New("м", "метр")
	cableUnit.AddUnit("коробка", "коробка 305 м", 305) // 1 коробка = 305 метров

	// Есть 3 коробки на складе
	quantity, _ := cableUnit.ToBase("коробка", 3)
	baseStr, _ := cableUnit.StringBase(quantity)
	fmt.Printf("На складе: %s %s\n", baseStr, cableUnit.List()[0].FullName)

	// Отгрузили 500 метров
	quantity, _ = cableUnit.Sub("м", quantity, 500)
	baseStr, _ = cableUnit.StringBase(quantity)
	unitStr, _ := cableUnit.StringUnit("коробка", quantity)
	fmt.Printf("После отгрузки: %s %s (%s коробок)\n",
		baseStr,
		cableUnit.List()[0].FullName,
		unitStr)
}
