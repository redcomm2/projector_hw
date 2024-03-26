package main

import (
	"projector_hw/hw_6/entity"
	"projector_hw/hw_6/entity/transport"
	"projector_hw/hw_6/routing"
)

func main() {
	bus := transport.Bus{Name: "Маршрутка №535"}
	train := transport.Train{Name: "№245 (Киев - Варшава)"}
	plane := transport.Plane{Name: "Рейс №4436 (Варшава - Банкок)"}

	route := routing.Route{}

	passenger := entity.Passenger{ID: 1, Name: "Alex"}

	route.AddTransport(&bus)
	route.AddTransport(&train)
	route.AddTransport(&plane)

	route.MoveByRoute(&passenger)
}
