package facade

type Shape interface {
	CreateShape() string
}

type Rectangle struct{}

func (r *Rectangle) CreateShape() string {
	return "Rectangle create shape"
}

type Circle struct{}

func (c *Circle) CreateShape() string {
	return "Circle create shape"
}

type Facade interface {
	GetShape() Shape
}

type FacadeStruct struct {
	Shape1 Shape
}

func (f *FacadeStruct) GetShape() Shape {
	return f.Shape1
}
