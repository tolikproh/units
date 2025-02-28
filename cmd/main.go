package main

import (
	"fmt"

	"github.com/tolikproh/units"
)

func main() {
	// Создание величин
	length1 := units.NewLength(100, units.Normal) // 100 метров
	length2 := units.NewLength(20, units.Kilo)    // 2 километра

	// Сложение
	totalLength := length1.Add(length2)
	fmt.Printf("Сумма: %s\n", totalLength.String()) // Ожидается: 2100 м

	// Вычитание
	difference := length2.Sub(length1)
	difference.SetDecimals(2)
	fmt.Printf("Разность: %s\n", difference.String()) // Ожидается: 1.9 км
	difference.SetDecimals(0)
	fmt.Printf("Разность: %s\n", difference.String()) // Ожидается: 2 км
	difference.SetPrefix(units.Normal)
	fmt.Printf("Разность: %s\n", difference.String()) // Ожидается: 1900 м

	// Умножение
	things := units.NewThings(5, units.Normal) // 5 штук
	multiplied := length1.Mul(things)
	fmt.Printf("Умножение: %s\n", multiplied.String()) // Ожидается: 500 м

	// Деление
	length3 := units.NewLength(100, units.Kilo)            // 100 километров
	divided := length3.Div(units.NewLength(2, units.Kilo)) // Делим на 2  километра
	divided.SetDecimals(10)
	divided.SetPrefix(units.Normal)
	fmt.Printf("Деление: %s, результат: %v\n", divided.String(), divided.Ok()) // Ожидается: 50 м

	// Сериализация в JSON
	jsonData, err := length1.MarshalJSON()
	if err != nil {
		fmt.Printf("Ошибка сериализации: %v\n", err)
		return
	}
	fmt.Printf("Сериализованный JSON: %s\n", jsonData)

	jsd, err := units.NewJSON(jsonData)
	if err != nil {
		fmt.Println("Error: %w", err)
	} else {
		fmt.Println("JSON")
		print(jsd)
	}

	// Десериализация из JSON

	newLength, err := units.NewLengthJSON(jsonData)
	if err != nil {
		fmt.Printf("Ошибка десериализации: %v\n", err)
		return
	}
	fmt.Printf("Десериализованная длина: %s\n", newLength.String()) // Ожидается: 100 м

	// Проверка значений
	fmt.Printf("Значение длины: %d\n", newLength.Value()) // Ожидается: 100
	fmt.Printf("Тип величины: %d\n", newLength.Types())   // Ожидается: 0 (LengthType)

	print(length2)
}

func add() {
	fmt.Println("Add")
	l1 := units.NewLength(1, units.Nano)
	l2 := units.NewLength(2, units.Micro)
	l3 := units.NewLength(3, units.Milli)
	l4 := units.NewLength(4, units.Normal)
	l5 := units.NewLength(5, units.Kilo)
	l6 := units.NewLength(6, units.Mega)
	l7 := units.NewLength(7, units.Giga)

	l1.Add(l2).Add(l3).Add(l4).Add(l5).Add(l6).Add(l7)

	print(l1)
}

func sub() {
	fmt.Println("Sub")
	l1 := units.NewLength(400, units.Normal)
	l2 := units.NewLength(5, units.Kilo)

	s1 := l2.Sub(l1)

	print(s1)

	l3 := units.NewLength(400, units.Normal)
	l4 := units.NewLength(410, units.Normal)

	l3.Sub(l4)

	print(l3)

}

func mul() {
	fmt.Println("Mul")
	l := units.NewLength(4, units.Normal)
	t := units.NewThings(24, units.Normal)
	p := units.NewPackage(2, units.Normal)

	l.Mul(t).Mul(p)

	print(l)
}

func div() {
	fmt.Println("Div")
	l := units.NewLength(300, units.Milli)
	t := units.NewLength(4, units.Normal)

	l.Div(t)
	fmt.Println(l.Value())

	print(l)

	t1 := units.NewLength(300, units.Milli)
	b, _ := t1.MarshalJSON()
	fmt.Printf("%s\n", b)

	l1, err := units.NewLengthJSON(b)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	print(l1)

}

func print(q units.Quantiter) {
	q.SetDecimals(10)
	q.SetPrefix(units.Nano)
	fmt.Println(q)
	q.SetPrefix(units.Micro)
	fmt.Println(q)
	q.SetPrefix(units.Milli)
	fmt.Println(q)
	q.SetPrefix(units.Normal)
	fmt.Println(q)
	q.SetPrefix(units.Kilo)
	fmt.Println(q)
	q.SetPrefix(units.Mega)
	fmt.Println(q)
	q.SetPrefix(units.Giga)
	fmt.Println(q)
	fmt.Println("Value: ", q.Value())
}
