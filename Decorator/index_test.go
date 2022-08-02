package decorator

import (
	"testing"
)

func TestInit(t *testing.T) {
	pizza := &VeggeMania{}

	if pizza.GetPrice() != 15 {
		t.Fatal("pizza price not equal 15")
	}

	pizza2 := &TomatoTopping{pizza: pizza}
	if pizza2.GetPrice() != 22 {
		t.Fatalf("pizza price is %d", pizza2.GetPrice())
	}

	pizza3 := &CheeseTopping{pizza: &TomatoTopping{pizza: pizza}}
	if pizza3.GetPrice() != 32 {
		t.Fatalf("pizza price is %d", pizza3.GetPrice())
	}

}
