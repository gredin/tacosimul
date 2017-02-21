package main

import "fmt"

type Dispatcher interface {
	MatchWithDriver(passenger *Passenger)
}

type GreedyDispatcher struct {
	supervisor *MemorySupervisor
}

func (dispatcher *GreedyDispatcher) MatchWithDriver(passenger *Passenger) {
	//closest_driver := dispatcher.supervisor.closestDriver(&passenger.pickupPosition)

	//_ := (*closest_driver).pickupPassenger(passenger)

	map_ := dispatcher.supervisor.map_
	shortest_distance := map_.w + map_.h + 1
	var closest_driver_id DriverId

	for driver_id, driver_position := range dispatcher.supervisor.drivers_to_positions {
		if dist := driver_position.ManhattanDistance((*passenger).pickupPosition); dist < shortest_distance {
			shortest_distance = dist
			closest_driver_id = driver_id
		}
	}

	fmt.Printf("closest driver from %v is %v\n", passenger, closest_driver_id)

	//if !dispatcher.supervisor.drivers_availability[closest_driver_id] {
	//	return
	//}

	// TODO: handle concurrency

	dispatcher.supervisor.drivers[closest_driver_id].book(passenger)
}

/*
func findClosestDriver(passenger *Passenger, drivers *[]Driver, map_ *Map) *Driver {
	shortest_distance := m.w + m.h + 1
	var closest_driver Driver

	for _, driver := range *drivers {
		if dist := driver.position.ManhattanDistance((*passenger).pickupPosition); dist < shortest_distance {
			shortest_distance = dist
			closest_driver = driver
		}
	}

	return &closest_driver
}
*/
