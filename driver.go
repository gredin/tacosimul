package main

import (
	"fmt"
	_ "sync"
	_ "sync/atomic"
	"time"
	"sync"
//	"image"
)

type DriverId int

func (id DriverId) String() string {
	return fmt.Sprintf("Driver#%v", int(id))
}

type Driver struct {
	id                DriverId
	position          Position
	passenger_onboard bool
	passenger         *Passenger

	//destination Position
	//speed       int
	mutex sync.Mutex
}

func (d *Driver) animate(supervisor *MemorySupervisor, update_interval time.Duration) {
	c := time.Tick(update_interval)

	//fmt.Printf("%p\n", d)

	for range c {
		current_position := d.position
		new_position := d.position

		//fmt.Printf("%p, %v, %v\n", d, t, d.passenger)

		if d.passenger == nil {
			continue
		}

		var destination Position
		if d.passenger_onboard {
			destination = d.passenger.dropoffPosition
		} else {
			destination = d.passenger.pickupPosition
		}

		new_position = d.position.MoveTo(destination, 1)

		if !new_position.Equals(current_position) {
			d.position = new_position
			supervisor.RecordPosition(d.id, d.position)
		}

		if new_position.Equals(d.passenger.pickupPosition) {
			d.pickupPassenger()
		}

		if new_position.Equals(d.passenger.dropoffPosition) {
			d.dropoffPassenger()
			supervisor.RecordAvailability(d.id, true)
		}
	}
}

// retourner une erreur si taxi non disponible
func (d *Driver) book(p *Passenger) bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.passenger != nil {
		fmt.Printf("booking failed for %v\n", p.id)

		return false
	}

	d.passenger = p

	//fmt.Printf("%p\n", d)

	fmt.Printf("%v books %v\n", *p, *d)

	return true
}

func (d *Driver) pickupPassenger() {
	d.passenger_onboard = true
}

// retourner erreur si aucun passager (?)
func (d *Driver) dropoffPassenger() bool {
	/*
	d.mutex.Lock()
	defer d.mutex.Unlock()
	*/

	fmt.Printf("%v drops off %v\n", d.id, d.passenger.id)

	if d.passenger == nil {
		// TODO: no passenger => symptom of error?

		return false
	}

	d.passenger = nil

	return true
}

/*
func (d *Driver) moveTo(position Position) {
	for i := 0; i < d.speed; i++ {
		if d.position.x < position.x {
			d.position.x++
		} else if d.position.x > position.x {
			d.position.x--
		} else if d.position.y < position.y {
			d.position.y++
		} else if d.position.y > position.y {
			d.position.y--
		} else {
			break
		}
	}
}
*/

func (d Driver) String() string {
	return fmt.Sprintf("Driver{id: %v, position: %v}", d.id, d.position)
}

func GenerateRandomDrivers(n int, map_ *Map) []*Driver {
	drivers := make([]*Driver, n)

	for i := 0; i < n; i++ {
		drivers[i] = &Driver{
			id:                DriverId(i),
			position:          map_.RandomPosition(),
			passenger_onboard: false,
			passenger:         nil,
		}
	}

	return drivers
}
