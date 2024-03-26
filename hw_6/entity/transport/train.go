package transport

import (
	"fmt"
	"projector_hw/hw_6/entity"
)

type Train struct {
	Name       string
	Passengers []entity.Passenger
}

func (t *Train) BoardPassenger(p *entity.Passenger) {
	t.Passengers = append(t.Passengers, *p)
	fmt.Printf("Passenger %s boarded the train: %s.\n", p.Name, t.Name)
}

func (t *Train) UnboardPassenger(p *entity.Passenger) {
	for i, passenger := range t.Passengers {
		if passenger.ID == p.ID {
			t.Passengers = append(t.Passengers[:i], t.Passengers[i+1:]...)
			fmt.Printf("Passenger %s left the train: %s.\n", passenger.Name, t.Name)
			break
		}
	}
}

func (t *Train) Move() {
	fmt.Printf("Train %s is moving.\n", t.Name)
}
