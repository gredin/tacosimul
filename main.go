package main

import (
	_"fmt"
	_ "encoding/json"
	"net/http"
	"time"
	_"database/sql/driver"
)

const map_w = 100
const map_h = 100

const real_life_duration_in_hours = 10
const simulation_duration_in_seconds = 5

const driver_count = 100
const driver_speed = 1
const update_interval_in_seconds = 30 // position update interval in real life

const passenger_count = 200

//const passenger_max_waiting_time = time.Second

/*
var map_ Map = Map{map_w, map_h}
var time_range TimeRange = TimeRange{time_start, time_end}

// TODO: 2 copies : une pour modéliser les données du serveur et une pour les données des clients
var server_drivers []Driver = GenerateDrivers(driver_count, driver_speed, map_)
var server_passengers []Passenger = GeneratePassengers(passenger_count, m, time_range)
*/
func rootHandler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(200)
}

/*
func animateDriver(driver Driver, time_channel <-chan time.Time) {
	for t := range time_channel {
		driver.moveTo(dri)
	}
}

func animatePassenger(passenger Passenger, time_channel <-chan time.Time) {
	for t := range time_channel {

	}
}
*/
func main() {
	map_ := Map{map_w, map_h}
	total_duration := RealLifeDuration{hour: real_life_duration_in_hours}

	drivers := GenerateRandomDrivers(driver_count, &map_)

	drivers_to_positions := make(map[DriverId]Position)
	drivers_availability := make(map[DriverId]bool)
	driverids_to_drivers := make(map[DriverId]*Driver)

	//fmt.Printf("%p\n", &(drivers[0]))

	for i := range drivers {
		driver := drivers[i]
		driver_id := driver.id

		drivers_to_positions[driver_id] = driver.position
		drivers_availability[driver_id] = true
		driverids_to_drivers[driver_id] = driver

		//fmt.Printf("%p\n", driver)
	}

	supervisor := MemorySupervisor{
		map_:                 &map_,
		drivers_to_positions: drivers_to_positions,
		drivers_availability: drivers_availability,
		drivers:              driverids_to_drivers,
	}

	update_interval := RealLifeDuration{second: update_interval_in_seconds}.toDuration()

	for i := range drivers {
		driver := drivers[i]
		go driver.animate(&supervisor, update_interval)
		//fmt.Printf("%p\n", driver)
	}

	passengers := GenerateRandomPassengers(passenger_count, &map_, total_duration)

	dispatcher := GreedyDispatcher{&supervisor}

	for i := range passengers {
		passenger := passengers[i]
		go passenger.animate(&dispatcher)
	}

	simulation_duration := total_duration.toDuration()
	time.Sleep(simulation_duration)

	// TODO: simulation results and stats

	/*
	var i int
	fmt.Printf("%v %T\n", i, i)
	i2 := new(int)
	fmt.Printf("%v %T\n", *i2, *i2)
	fmt.Println("")

	var f float64
	fmt.Printf("%v %T\n", f, f)
	fmt.Println("")

	var s string
	fmt.Printf("\"%v\" %T\n", s, s)
	fmt.Println("")

	var b bool
	fmt.Printf("%v %T\n", b, b)
	fmt.Println("")

	var a [5]int
	fmt.Printf("%v %T\n", a, a)
	fmt.Println("")

	var sl []int
	fmt.Printf("%v %T\n", sl, sl)
	if sl == nil {
		fmt.Println("is nil")
	} else {
		fmt.Println("is NOT nil")
	}
	sl2 := make([]int, 5)
	fmt.Printf("%v %T\n", sl2, sl2)
	if sl2 == nil {
		fmt.Println("is nil")
	} else {
		fmt.Println("is NOT nil")
	}
	fmt.Println("")

	var i_p *int
	fmt.Printf("%v %T\n", i_p, i_p)
	if i_p == nil {
		fmt.Println("is nil")
	} else {
		fmt.Println("is NOT nil")
	}
	fmt.Println("")

	var fu func(int) int
	fmt.Printf("%v %T\n", fu, fu)
	if fu == nil {
		fmt.Println("is nil")
	} else {
		fmt.Println("is NOT nil")
	}
	fmt.Println("")

	var c chan int
	fmt.Printf("%v %T\n", c, c)
	if c == nil {
		fmt.Println("is nil")
	} else {
		fmt.Println("is NOT nil")
	}
	c2 := make(chan int)
	fmt.Printf("%v %T\n", c2, c2)
	if c2 == nil {
		fmt.Println("is nil")
	} else {
		fmt.Println("is NOT nil")
	}
	fmt.Println("")

	var m map[int]int
	fmt.Printf("%v %T\n", m, m)
	if m == nil {
		fmt.Println("is nil")
	} else {
		fmt.Println("is NOT nil")
	}
	m2 := make(map[int]int)
	fmt.Printf("%v %T\n", m2, m2)
	if m2 == nil {
		fmt.Println("is nil")
	} else {
		fmt.Println("is NOT nil")
	}
	fmt.Println("")

	var o struct {
		i   int
		i_p *int
	}
	fmt.Printf("%v %T\n", o, o)
	o2 := new(struct {
		i   int
		i_p *int
	})
	fmt.Printf("%v %T\n", *o2, *o2)
	fmt.Println("")
	*/

	//time.Sleep(RealLifeTime{2, 43,3}.toDurationFromStart())

	//mux := http.NewServeMux()
	//mux.HandleFunc("/", rootHandler)

	//ticker := time.NewTicker(time_quantum)

	/*
	for _, driver := range drivers {
		go animateDriver(driver, ticker.C)
	}

	for _, passenger := range passengers {
		go animatePassenger(passenger, ticker.C)
	}


	err := http.ListenAndServe(":9999", mux)

	fmt.Println(err)
	*/
	/*

	d := Driver{Position{0, 0}}

	d.moveTo(Position{3, 3})
	fmt.Println(d)
	d.moveTo(Position{3, 3})
	fmt.Println(d)
	d.moveTo(Position{3, 3})
	fmt.Println(d)
	d.moveTo(Position{3, 3})
	fmt.Println(d)
	d.moveTo(Position{3, 3})
	fmt.Println(d)
	d.moveTo(Position{3, 3})
	fmt.Println(d)
	d.moveTo(Position{3, 3})
	fmt.Println(d)
	*/
}
