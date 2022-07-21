package builder

import "fmt"

type Item interface {
	Name() string
	Price() float64
	Packing() Packing
}

type Packing interface {
	Pack() string
}

type Wrapper struct{}

func (w *Wrapper) Pack() string {
	return "Wrapper"
}

type Bottle struct{}

func (b *Bottle) Pack() string {
	return "Bottle"
}

type VegBurger struct{}

func (v *VegBurger) Name() string {
	return "VegBurger"
}

func (v *VegBurger) Price() float64 {
	return 2.3
}

func (v *VegBurger) Packing() Packing {
	return &Wrapper{}
}

type NonVegBurger struct{}

func (v *NonVegBurger) Name() string {
	return "NonVegBurger"
}

func (v *NonVegBurger) Price() float64 {
	return 3
}

func (v *NonVegBurger) Packing() Packing {
	return &Bottle{}
}

type Meal struct {
	items []Item
}

func (m *Meal) AddItem(item Item) {
	m.items = append(m.items, item)
}

func (m *Meal) GetCost() float64 {
	cost := 0.0
	for _, v := range m.items {
		cost += v.Price()
	}
	return cost
}

func (m *Meal) DisplayItems() {
	for _, v := range m.items {
		fmt.Println(v.Name())
	}
}
