package routing

import (
	"projector_hw/hw_6/entity"
)

type PublicTransport interface {
	BoardPassenger(p *entity.Passenger)
	UnboardPassenger(p *entity.Passenger)
	Move()
}

type Route struct {
	transports []PublicTransport
}

func (r *Route) MoveByRoute(p *entity.Passenger) {
	for _, transportEntity := range r.transports {
		transportEntity.BoardPassenger(p)
		transportEntity.Move()
		transportEntity.UnboardPassenger(p)
	}
}

func (r *Route) AddTransport(t PublicTransport) []PublicTransport {
	r.transports = append(r.transports, t)

	return r.transports
}
