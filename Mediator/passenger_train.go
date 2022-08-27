package mediator

import "fmt"

type PassengerTrain struct {
	mediator Mediator
}

func (g *PassengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	g.arrive()
}

func (c *PassengerTrain) arrive() {
	if !c.mediator.canArrive(c) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (c *PassengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	c.mediator.notifyAboutDeparture()
}
