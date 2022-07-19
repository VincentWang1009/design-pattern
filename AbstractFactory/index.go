package abstractfactory

type Shape interface {
	DrawShape() string
}

type Circle struct {
}

func (c *Circle) DrawShape() string {
	return "Circle draw shape"
}

type Rectangle struct {
}

func (r *Rectangle) DrawShape() string {
	return "Rectangle draw shape"
}

type Square struct{}

func (s *Square) DrawShape() string {
	return "Square draw shape"
}

type Color interface {
	FillColor() string
}

type Blue struct{}

func (b *Blue) FillColor() string {
	return "Blue Fill Color"
}

type Green struct{}

func (g *Green) FillColor() string {
	return "Green Fill Color"
}

type Red struct{}

func (r *Red) FillColor() string {
	return "Red Fill Color"
}

type AbstractFactory interface {
	GetColor(color string) Color
	GetShape(shape string) Shape
}

type ShapeFactory struct{}

func (s *ShapeFactory) GetShape(shapeType string) Shape {
	if shapeType == "Circle" {
		return &Circle{}
	} else if shapeType == "Rectangle" {
		return &Rectangle{}
	} else if shapeType == "Square" {
		return &Square{}
	} else {
		return nil
	}
}

func (s *ShapeFactory) GetColor(colorType string) Color {
	return nil
}

type ColorFactory struct{}

func (c *ColorFactory) GetColor(colorType string) Color {
	if colorType == "Red" {
		return &Red{}
	} else if colorType == "Blue" {
		return &Blue{}
	} else if colorType == "Green" {
		return &Green{}
	} else {
		return nil
	}
}

func (s *ColorFactory) GetShape(shapeType string) Shape {
	return nil
}

type FactoryGenerator struct{}

func (f *FactoryGenerator) GetFactory(choice string) AbstractFactory {
	if choice == "Shape" {
		return &ShapeFactory{}
	} else if choice == "Color" {
		return &ColorFactory{}
	} else {
		return nil
	}
}
