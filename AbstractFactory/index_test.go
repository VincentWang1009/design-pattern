package abstractfactory

import (
	"testing"
)

func TestInit(t *testing.T) {
	factory := FactoryGenerator{}
	var shapeFactory AbstractFactory = factory.GetFactory("Shape")

	var shape = shapeFactory.GetShape("Circle")

	if shape.DrawShape() != "Circle draw shape" {
		t.Fatal("类型不是circle")
	}
}
