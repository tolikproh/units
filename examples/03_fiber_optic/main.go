package main

import (
	"fmt"

	"github.com/tolikproh/units"
)

func main() {
	// Создаём единицу измерения для оптоволокна (базовая: метры)
	fiberUnit := units.New("м", "метр")
	fiberUnit.AddUnit("км", "километр", 1000)            // 1 км = 1000 метров
	fiberUnit.AddUnit("барабан", "барабан 3000 м", 3000) // 1 барабан = 3000 метров

	// Есть 2 барабана на складе
	quantity, _ := fiberUnit.ToBase("барабан", 2)
	baseStr, _ := fiberUnit.StringBase(quantity)
	fmt.Printf("На складе: %s %s\n", baseStr, fiberUnit.List()[0].FullName)

	// Отгрузили 2.5 км
	quantity, _ = fiberUnit.Sub("км", quantity, 2.5)
	baseStr, _ = fiberUnit.StringBase(quantity)
	unitStr, _ := fiberUnit.StringUnit("барабан", quantity)
	fmt.Printf("После отгрузки: %s %s (%s барабанов)\n",
		baseStr,
		fiberUnit.List()[0].FullName,
		unitStr)
}
