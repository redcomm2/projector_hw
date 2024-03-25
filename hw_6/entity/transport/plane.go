package transport

import (
	"fmt"
	"projector_hw/hw_6/entity"
)

type Plane struct {
	Name       string
	Passengers []entity.Passenger
}

func (p *Plane) BoardPassenger(pas *entity.Passenger) {
	p.Passengers = append(p.Passengers, *pas)
	fmt.Printf("Passenger %s boarded the plain: %s.\n", pas.Name, p.Name)
}

func (p *Plane) UnboardPassenger(pas *entity.Passenger) {
	for i, passenger := range p.Passengers {
		if passenger.ID == pas.ID {
			p.Passengers = append(p.Passengers[:i], p.Passengers[i+1:]...)
			fmt.Printf("Passenger %s left the plain: %s.\n", passenger.Name, p.Name)
			break
		}
	}
}

func (p *Plane) Move() {
	fmt.Printf("Plane %s is taking off.\n", p.Name)
}
