package main

import (
	"fmt"

	"github.com/tolikproh/units"
)

func main() {
	l1 := units.NewLength(5, units.Normal)
	l2 := units.NewLength(3, units.Kilo)
	l3 := units.NewLength(600, units.Kilo)

	sum := l1.Add(l2).Add(l3)
	//diff := l1.Sub(l2)
	sum.SetPrefix(units.Normal)
	sum.SetDecimals(5)
	fmt.Println(sum.String()) // "8 км"
	//fmt.Println(diff.String()) // "2 км"

}
