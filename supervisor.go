package main

import (
	_"fmt"
	"sync"
)

//import "database/sql/driver"

//import()

type Supervisor interface {
	RecordPosition(driver *Driver)
	//ListDrivers() *[]Driver
}

var mutex sync.Mutex

type MemorySupervisor struct {
	map_                 *Map
	drivers_to_positions map[DriverId]Position
	drivers_availability map[DriverId]bool
	drivers              map[DriverId]*Driver
}

func (supervisor MemorySupervisor) RecordPosition(driver_id DriverId, position Position) {
	mutex.Lock() // prevents concurrent map writes
	defer mutex.Unlock()

	supervisor.drivers_to_positions[driver_id] = position

	//fmt.Printf("%v is now at %v\n", driver_id, position)
}

func (supervisor MemorySupervisor) RecordAvailability(driver_id DriverId, is_available bool) {
	mutex.Lock() // prevents concurrent map writes
	defer mutex.Unlock()

	supervisor.drivers_availability[driver_id] = is_available
}

/*
func (supervisor MemorySupervisor) ListDrivers() []*Driver {
	drivers := make([]*Driver, len(supervisor.ids_to_drivers))

	for _, driver := range supervisor.ids_to_drivers {
		drivers = append(drivers, &driver)
	}

	return drivers
}
*/
