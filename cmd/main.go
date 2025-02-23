package main

import (
	"fmt"

	"github.com/tolikproh/units"
)

func main() {

	l1 := units.NewLength(2, units.Normal)
	t1 := units.NewThings(24, units.Normal)
	p1 := units.NewPackage(4, units.Normal)

	sum := l1.Mul(t1).Mul(p1)
	sum.SetPrefix(units.Normal)
	sum.SetDecimals(3)
	fmt.Println(sum.String())

	t2 := units.NewThings(25, units.Normal)
	p2 := units.NewPackage(10, units.Normal)

	sum2 := t2.Mul(p2)
	sum2.SetPrefix(units.Normal)
	sum2.SetDecimals(3)
	fmt.Println(sum2.String())
}
