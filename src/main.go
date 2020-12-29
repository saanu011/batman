package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Vehicle struct {
	Name       string
	Speed      float64
	CraterTime float64
}
type Orbit struct {
	Name         string
	Length       float64
	Crater       int
	TrafficSpeed float64
}

// Vehicle's initial data
var BatBoat = Vehicle{Name: "BatBoat", Speed: 10, CraterTime: 2}
var BatMobile = Vehicle{Name: "BatMobile", Speed: 12, CraterTime: 1}
var BatCycle = Vehicle{Name: "BatCycle", Speed: 20, CraterTime: 3}

// Orbit's initial data
var Orbit1 = Orbit{Name: "Orbit1", Length: 18, Crater: 20}
var Orbit2 = Orbit{Name: "Orbit2", Length: 20, Crater: 10}
var Orbit3 = Orbit{Name: "Orbit3", Length: 22, Crater: 14}

const sunny = "Sunny"
const rainy = "Rainy"
const windy = "Windy"

var weathers = []string{sunny, rainy, windy}

var Sunny = []Vehicle{BatBoat, BatMobile, BatCycle}
var Rainy = []Vehicle{BatMobile, BatCycle}
var Windy = []Vehicle{BatBoat, BatCycle}

func main() {

	weather := weathers[rand.Intn(len(weathers))]

	var Orbits []Orbit
	var TrafficSpeed1 float64 = 12
	var TrafficSpeed2 float64 = 10
	var TrafficSpeed3 float64 = 14

	// available vehicles on days
	var availableVehicles []Vehicle
	if weather == sunny {
		availableVehicles = Sunny
		Orbits = finalCraterForOrbit(0.9, TrafficSpeed1, TrafficSpeed2, TrafficSpeed3)
	} else if weather == rainy {
		availableVehicles = Rainy
		Orbits = finalCraterForOrbit(1.2, TrafficSpeed1, TrafficSpeed2, TrafficSpeed3)
	} else if weather == windy {
		availableVehicles = Windy
		Orbits = finalCraterForOrbit(1, TrafficSpeed1, TrafficSpeed2, TrafficSpeed3)
	}

	var smallestTime float64
	var vehicle string
	var orbit string
	for k := 0; k < len(Orbits); k++ {
		var timeTaken float64
		for i := 0; i < len(availableVehicles); i += 1 {
			timeTaken = timeTakenInOrbit(Orbits[k], availableVehicles[i])

			// initial value for result storing variables
			if k == 0 && i == 0 {
				smallestTime = timeTaken
				vehicle = availableVehicles[i].Name
				orbit = Orbits[k].Name
			}
			// if time taken on current path with current vehicle is less than earlier ones
			if timeTaken < smallestTime {
				smallestTime = timeTaken
				vehicle = availableVehicles[i].Name
				orbit = Orbits[k].Name
			}
		}
	}
	fmt.Println(fmt.Sprintf("Output: %s on %s", orbit, vehicle))
}

func finalSpeed(maxSpeed, defaultSpeed float64) float64 {
	if maxSpeed < defaultSpeed {
		return maxSpeed
	}
	return defaultSpeed
}

func timeTakenInOrbit(orbit Orbit, vehicle Vehicle) float64 {
	var timeTaken float64
	timeTaken = (orbit.Length / finalSpeed(orbit.TrafficSpeed, vehicle.Speed)) + (float64(orbit.Crater) * (vehicle.CraterTime / 60))
	return timeTaken
}

func finalCraterForOrbit(ratio, trafficSpeed1, trafficSpeed2, trafficSpeed3 float64) (Orbits []Orbit) {

	// Orbit1
	FinalOrbit1 := Orbit1
	FinalOrbit1.Crater = int(math.Ceil(ratio * float64(Orbit1.Crater)))
	FinalOrbit1.TrafficSpeed = trafficSpeed1
	Orbits = append(Orbits, FinalOrbit1)
	// Orbit2
	FinalOrbit2 := Orbit2
	FinalOrbit2.Crater = int(math.Ceil(ratio * float64(Orbit2.Crater)))
	FinalOrbit2.TrafficSpeed = trafficSpeed2
	Orbits = append(Orbits, FinalOrbit2)
	// Orbit3
	FinalOrbit3 := Orbit3
	FinalOrbit3.Crater = int(math.Ceil(ratio * float64(Orbit3.Crater)))
	FinalOrbit3.TrafficSpeed = trafficSpeed3
	Orbits = append(Orbits, FinalOrbit3)

	return Orbits
}
