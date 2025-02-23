package main

import (
	"fmt"

	"github.com/tolikproh/units"
)

func main() {

	add()
	sub()
	mul()
	div()

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
}

func print(q units.Quantiter) {
	q.SetDecimals(16)
	q.SetPrefix(units.Nano)
	fmt.Println(q.String())
	q.SetPrefix(units.Micro)
	fmt.Println(q.String())
	q.SetPrefix(units.Milli)
	fmt.Println(q.String())
	q.SetPrefix(units.Normal)
	fmt.Println(q.String())
	q.SetPrefix(units.Kilo)
	fmt.Println(q.String())
	q.SetPrefix(units.Mega)
	fmt.Println(q.String())
	q.SetPrefix(units.Giga)
	fmt.Println(q.String())
	fmt.Println("Value: ", q.Value())
}
