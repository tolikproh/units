package main

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/tolikproh/units"
)

type Product struct {
	name string
	unit *units.Unit
	val  decimal.Decimal
}

func New(name string) *Product {
	u := units.New("м", "метр")
	u.AddFromInt("км", "километр", 1000)
	u.AddFromInt("бух-50", "бухта 50 м", 50)
	u.AddFromInt("бух-100", "бухта 100 м", 100)
	u.AddFromInt("бух-200", "бухта 200 м", 200)

	return &Product{
		name: name,
		unit: u,
		val:  decimal.NewFromInt(2500),
	}
}

func (p *Product) ListUnit() []*units.UnitItem {
	return p.unit.List()
}

func (p *Product) String() string {
	value, _ := p.unit.GetInBaseUnit(p.val)
	return strings.TrimSpace(value)
}

func (p *Product) StringFromUnit(unitName string) string {

	value, _ := p.unit.GetInUnit(unitName, p.val)

	return strings.TrimSpace(value)
}

func (p *Product) Add(unitName string, val any) error {
	// Передаем текущее значение, единицу и добавляемое значение
	v, err := p.unit.ToBase(unitName, val)
	if err != nil {
		return err
	}

	// Обновляем значение результатом операции
	p.val = p.val.Add(v)

	return nil
}

func (p *Product) Sub(unitName string, val any) error {
	// Передаем текущее значение, единицу и добавляемое значение
	v, err := p.unit.ToBase(unitName, val)
	if err != nil {
		return err
	}

	// Обновляем значение результатом операции
	p.val = p.val.Sub(v)

	return nil
}

func main() {
	prod1 := New("КПС 3х10")

	fmt.Println("=== Информация о единицах измерения ===")
	fmt.Println(prod1.unit)
	fmt.Println("Кол-во: ", prod1)
	prod1.Add("км", 1.10125)
	fmt.Println("Кол-во: ", prod1)
	prod1.Sub("км", 0.10125)
	fmt.Println("Кол-во: ", prod1)
	for _, v := range prod1.ListUnit() {
		fmt.Println("Кол-во: ", prod1.StringFromUnit(v.Name), " ", v.FullName)
	}

}
