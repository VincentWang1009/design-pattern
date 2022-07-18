package simplefactory

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	ins1 := GetShape("Circle")
	fmt.Printf("ins1: %v\n", ins1)
	ins2 := GetShape("Rectangle")
	ins3 := GetShape("square")
	if ins1.DrawShape() != "Circle drawshape" {
		t.Fatal("instance is not circletype")
	}

	if reflect.TypeOf(ins2).String() != "*simplefactory.Rectangle" {
		t.Fatal("instance is not Rectangle type")
	}

	if ins3 != nil {
		t.Fatal("instance is not nil")
	}

}
