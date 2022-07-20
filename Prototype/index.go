package prototype

type Cloneable interface {
	clone() struct{}
}

type Customers struct {
	customerList []string
}

func NewCustomers() *Customers {
	return &Customers{
		customerList: []string{},
	}
}

func NewCustomers2(customerList []string) *Customers {
	return &Customers{
		customerList: customerList,
	}
}

func (c *Customers) LoadDataFromDB() {
	c.customerList[0] = "Bharat"
	c.customerList[1] = "Sahdev"
	c.customerList[2] = "Richi"
}

func (c *Customers) GetCustomerList() []string {
	return c.customerList
}

func (c *Customers) Cloneable() Customers {
	var cl []string = c.GetCustomerList()

	return *NewCustomers2(cl)
}
