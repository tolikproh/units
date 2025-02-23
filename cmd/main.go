package main

import (
	"fmt"

	"github.com/tolikproh/units"
)

func main() {

	l1 := units.NewLength(2, units.Normal)
	t1 := units.NewThings(24, units.Normal)
	p1 := units.NewPackage(4, units.Normal)

	sum := l1.Mul(t1)
	sadd := l1.Add(l1)
	//sum.SetPrefix(units.Normal)
	sum.SetDecimals(3)
	fmt.Println(l1.Value())
	fmt.Println(t1.Value())
	fmt.Println(p1.Value())
	fmt.Println(sum.Value())
	fmt.Println(sum.String())
	fmt.Println(sadd.Value())
	fmt.Println(sadd.String())

	t2 := units.NewThings(25, units.Normal)
	p2 := units.NewPackage(10, units.Normal)

	sum2 := t2.Mul(p2)
	//sum2.SetPrefix(units.Normal)
	sum2.SetDecimals(3)
	fmt.Println(sum2.String())
	fmt.Println(sum2.Value())
}
