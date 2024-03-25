package transport

import "projector_hw/hw_6/entity"

type PublicTransport interface {
	BoardPassenger(p *entity.Passenger)
	UnboardPassenger(p *entity.Passenger)
	Move()
}
