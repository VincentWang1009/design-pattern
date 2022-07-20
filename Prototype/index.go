package prototype

type Cloneable interface {
	clone() struct{}
}

type Customers struct{}

func (c *Customers) Cloneable() Customers {

}
