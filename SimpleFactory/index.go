package simplefactory

type Shape interface {
	DrawShape() string
}

type Circle struct {
}

func (*Circle) DrawShape() string {
	return "Circle drawshape"
}

type Rectangle struct {
}

func (*Rectangle) DrawShape() string {
	return "Rectangle drawshape"

}

type Square struct {
}

func (*Square) DrawShape() string {
	return "Square drawshape"
}

func GetShape(shapeType string) Shape {
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
