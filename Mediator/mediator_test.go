package mediator

import (
	"testing"
)

func TestInit(t *testing.T) {
	stationManager := newStationManager()
	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}

	freightTrain := &FreightTrain{
		mediator: stationManager,
	}
	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
