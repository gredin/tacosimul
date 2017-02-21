package main

import (
	_"math/rand"
	"time"
	"fmt"
)

const time_contraction = float64(simulation_duration_in_seconds) / (3600 * real_life_duration_in_hours)

type RealLifeDuration struct {
	hour   int
	minute int
	second int
}

func (d RealLifeDuration) String() string {
	return fmt.Sprintf("%02d:%02d:%02d", d.hour, d.minute, d.second)
}

func (d RealLifeDuration) toDuration() time.Duration {
	nanoseconds := 1e9*(d.hour*3600 + d.minute*60 + d.second)
	simulation_nanoseconds := float64(nanoseconds) * time_contraction

	return time.Duration(int64(simulation_nanoseconds))
}

/*
type TimeRange struct {
	start int
	end   int
}

func (tr TimeRange) RandomTime() int {
	return tr.start + rand.Intn(tr.end-tr.start+1)
}
*/
