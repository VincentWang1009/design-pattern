package facade

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	ins := &FacadeStruct{
		Shape1: &Circle{},
	}
	fmt.Printf("ins: %v\n", ins.GetShape().CreateShape())
	res, expect := ins.GetShape().CreateShape(), "Circle create shape"
	if res != expect {
		t.Fatalf("expect %s, return %s", "Circle create shape", res)
	}

}
