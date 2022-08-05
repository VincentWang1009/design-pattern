package chainofresponsibility

import (
	"testing"
)

func TestInit(t *testing.T) {
	cashier := &Cashier{}

	medical := &Medical{}
	medical.setNext(cashier)

	doctor := &Doctor{}
	doctor.setNext(medical)

	reception := &Reception{}
	reception.setNext(doctor)
	// t.Fatal("Aaa")
}
