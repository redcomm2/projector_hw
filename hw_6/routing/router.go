package routing

import (
	"projector_hw/hw_6/entity"
	"projector_hw/hw_6/entity/transport"
)

type Route struct {
	transports []transport.PublicTransport
}

func (r *Route) MoveByRoute(p *entity.Passenger) {
	for _, transportEntity := range r.transports {
		transportEntity.BoardPassenger(p)
		transportEntity.Move()
		transportEntity.UnboardPassenger(p)
	}
}

func (r *Route) AddTransport(t transport.PublicTransport) []transport.PublicTransport {
	r.transports = append(r.transports, t)

	return r.transports
}
