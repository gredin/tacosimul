package main

import (
	"fmt"
	"time"
	"math/rand"
)

type PassengerId int

func (id PassengerId) String() string {
	return fmt.Sprintf("Passenger#%v", int(id))
}

type Passenger struct {
	id PassengerId
	pickupPosition        Position
	dropoffPosition       Position
	durationBeforeRequest RealLifeDuration
}

func (p *Passenger) animate(dispatcher *GreedyDispatcher) {
	time.Sleep(p.durationBeforeRequest.toDuration())

	fmt.Printf("passenger %v requests a ride from %v to %v\n", p, p.pickupPosition, p.dropoffPosition)

	dispatcher.MatchWithDriver(p)
}

func (p Passenger) String() string {
	return fmt.Sprintf("Passenger{id: %v, pick up: %v, drop off: %v, request after: %v}", p.id, p.pickupPosition, p.dropoffPosition, p.durationBeforeRequest)
}

func GenerateRandomPassengers(n int, map_ *Map, max_duration RealLifeDuration) []*Passenger {
	passengers := make([]*Passenger, n)

	for i := 0; i < n; i++ {
		duration_before_request := RealLifeDuration{
			rand.Intn(int(max_duration.hour + 1)),
			rand.Intn(max_duration.minute + 1),
			rand.Intn(max_duration.second + 1),
		}

		pickup_position := map_.RandomPosition()

		var dropoff_position Position
		for {
			dropoff_position = map_.RandomPosition()
			if !dropoff_position.Equals(pickup_position) {
				break
			}
		}

		passengers[i] = &Passenger{
			PassengerId(i),
			pickup_position,
			dropoff_position, // TODO: ensure it's <> pickupPosition
			duration_before_request,
		}
	}

	return passengers
}
