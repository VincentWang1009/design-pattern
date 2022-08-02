package decorator

type IPizza interface {
	GetPrice() int
}

type VeggeMania struct{}

func (p *VeggeMania) GetPrice() int {
	return 15
}

type TomatoTopping struct {
	pizza IPizza
}

func (t *TomatoTopping) GetPrice() int {
	return t.pizza.GetPrice() + 7
}

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) GetPrice() int {
	return c.pizza.GetPrice() + 10
}
