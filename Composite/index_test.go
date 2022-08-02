package composite

import (
	"testing"
)

func TestInit(t *testing.T) {
	mealIns := &Meal{}
	mealIns.AddItem(&VegBurger{})
	// mealIns.AddItem(&VegBurger{})
	cost := mealIns.GetCost()
	if cost != 2.3 {
		t.Fatal("cost is not correct")
	}
}
