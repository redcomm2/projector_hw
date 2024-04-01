package transport

import (
	"fmt"
	"projector_hw/hw_6/entity"
)

type Bus struct {
	Name       string
	Passengers []entity.Passenger
}

func (b *Bus) BoardPassenger(p *entity.Passenger) {
	b.Passengers = append(b.Passengers, *p)
	fmt.Printf("Passenger %s boarded the bus: %s.\n", p.Name, b.Name)
}

func (b *Bus) UnboardPassenger(p *entity.Passenger) {
	for i, passenger := range b.Passengers {
		if passenger.ID == p.ID {
			b.Passengers = append(b.Passengers[:i], b.Passengers[i+1:]...)
			fmt.Printf("Passenger %s left the bus: %s.\n", passenger.Name, b.Name)
			break
		}
	}
}

func (b *Bus) Move() {
	fmt.Printf("Bus %s is moving.\n", b.Name)
}
