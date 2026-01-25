package main

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/tolikproh/units"
)

// Product представляет товар на складе
type Product struct {
	ID       string
	Name     string
	Units    *units.Unit
	Quantity decimal.Decimal
}

func main() {
	// Создаём склад с разными товарами
	warehouse := make(map[string]*Product)

	// 1. Кабель ППГ
	cableUnit := units.New("м", "метр")
	cableUnit.AddUnit("бухта", "бухта", 200)
	cableQuantity, _ := cableUnit.ToBase("бухта", 10)
	warehouse["cable_ppg"] = &Product{
		ID:       "cable_ppg",
		Name:     "Кабель ППГ",
		Units:    cableUnit,
		Quantity: cableQuantity,
	}

	// 2. Дымовые извещатели
	detectorUnit := units.New("шт", "штука")
	detectorUnit.AddUnit("упаковка", "упаковка", 10)
	detectorUnit.AddUnit("коробка", "коробка", 120)
	detectorQuantity, _ := detectorUnit.ToBase("коробка", 3)
	warehouse["detector"] = &Product{
		ID:       "detector",
		Name:     "Дымовой извещатель",
		Units:    detectorUnit,
		Quantity: detectorQuantity,
	}

	// 3. Шурупы
	screwsUnit := units.New("шт", "штука")
	screwsUnit.AddUnit("кг", "килограмм", decimal.NewFromFloat(1.0/0.002))
	screwsQuantity, _ := screwsUnit.ToBase("кг", 50) // 50 кг = 25000 штук
	warehouse["screws"] = &Product{
		ID:       "screws",
		Name:     "Шурупы 3.5x25",
		Units:    screwsUnit,
		Quantity: screwsQuantity,
	}

	// Вывод текущего состояния склада
	fmt.Println("=== Склад (начальное состояние) ===")
	for _, p := range []string{"cable_ppg", "detector", "screws"} {
		product := warehouse[p]
		baseStr, _ := product.Units.StringBase(product.Quantity)
		fmt.Printf("%s: %s %s\n",
			product.Name,
			baseStr,
			product.Units.List()[0].FullName)
	}

	// Обработка заказа
	fmt.Println("\n=== Обработка заказа ===")

	// Отгрузка 3 бухт кабеля (600 метров)
	cable := warehouse["cable_ppg"]
	orderQty, _ := cable.Units.ToBase("бухта", 3)
	cable.Quantity, _ = cable.Units.Sub("м", cable.Quantity, orderQty)
	baseStr, _ := cable.Units.StringBase(orderQty)
	unitStr, _ := cable.Units.StringUnit("бухта", orderQty)
	fmt.Printf("Отгружено: %s %s кабеля (%s бухт)\n",
		baseStr,
		cable.Units.List()[0].FullName,
		unitStr)

	// Отгрузка 15 упаковок извещателей (150 штук)
	detector := warehouse["detector"]
	orderQty, _ = detector.Units.ToBase("упаковка", 15)
	detector.Quantity, _ = detector.Units.Sub("шт", detector.Quantity, orderQty)
	baseStr, _ = detector.Units.StringBase(orderQty)
	unitStr, _ = detector.Units.StringUnit("упаковка", orderQty)
	fmt.Printf("Отгружено: %s %s извещателей (%s упаковок)\n",
		baseStr,
		detector.Units.List()[0].FullName,
		unitStr)

	// Отгрузка 200 штук шурупов
	screws := warehouse["screws"]
	screws.Quantity, _ = screws.Units.Sub("шт", screws.Quantity, 200)
	baseStr, _ = screws.Units.StringBase(decimal.NewFromFloat(200))
	unitStr, _ = screws.Units.StringUnit("кг", decimal.NewFromFloat(200))
	fmt.Printf("Отгружено: %s %s шурупов (%s кг)\n",
		baseStr,
		screws.Units.List()[0].FullName,
		unitStr)

	// Вывод итогового состояния склада
	fmt.Println("\n=== Склад (после отгрузки) ===")
	for _, p := range []string{"cable_ppg", "detector", "screws"} {
		product := warehouse[p]
		baseStr, _ := product.Units.StringBase(product.Quantity)
		fmt.Printf("%s: %s %s\n",
			product.Name,
			baseStr,
			product.Units.List()[0].FullName)
	}

	// Демонстрация JSON-сериализации
	fmt.Println("\n=== JSON-представление ===")
	for _, p := range []string{"cable_ppg"} {
		product := warehouse[p]
		jsonData, _ := product.Units.ToJSON()
		fmt.Printf("%s:\n%s\n", product.Name, string(jsonData))
	}
}
