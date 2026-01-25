package main

import (
	"fmt"

	"github.com/tolikproh/units"
)

func main() {
	// Создаём единицу измерения для кабеля ППГ (базовая: метры)
	cableUnit := units.New("м", "метр")
	cableUnit.AddUnit("бухта", "бухта 200 м", 200) // 1 бухта = 200 метров

	// Есть 5 бухт на складе
	quantity, _ := cableUnit.ToBase("бухта", 5)
	baseStr, _ := cableUnit.StringBase(quantity)
	fmt.Printf("На складе: %s %s\n", baseStr, cableUnit.List()[0].FullName)

	// Отгрузили 150 метров
	quantity, _ = cableUnit.Sub("м", quantity, 150)
	baseStr, _ = cableUnit.StringBase(quantity)
	unitStr, _ := cableUnit.StringUnit("бухта", quantity)
	fmt.Printf("После отгрузки: %s %s (%s бухт)\n",
		baseStr,
		cableUnit.List()[0].FullName,
		unitStr)
}
