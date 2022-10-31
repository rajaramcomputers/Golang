package main

import "fmt"

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,

		//panic("Please implement the NewCar function")
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
	//panic("Please implement the NewTrack function")
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.

//Add a new field to keep track of the distance driven
//You need to check if there is enough battery to drive the car
func Drive(car Car) Car {
	return Car{
		speed:        car.speed,
		batteryDrain: car.batteryDrain,
		battery:      car.battery - car.batteryDrain,
		distance:     car.distance + car.speed,
	}
	//panic("Please implement the Drive function")
}

func CanFinish(car Car, track Track) bool {
	max_distance := car.battery / car.batteryDrain
	car_can_cover := car.speed * track.distance

	if car_can_cover > max_distance {
		return true
	} else {
		return false
	}
	//panic("Please implement the CanFinish function")
}

func main() {
	//Task1
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)
	fmt.Println(car)
	//Task2
	distance := 800
	track := NewTrack(distance)
	fmt.Println(track)
	//Task3
	speed = 18
	batteryDrain = 10
	car1 := NewCar(speed, batteryDrain)
	car1 = Drive(car1)
	fmt.Println(car1.speed, car1.batteryDrain, car1.battery, car1.distance)
	// => Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}
	//Task4
	speed = 5
	batteryDrain = 2
	car = NewCar(speed, batteryDrain)

	distance = 100
	track = NewTrack(distance)

	fmt.Println(CanFinish(car, track))
	// => true
}

/* Assume the car is just starting the race
You need to calculate the maximum distance a car can drive with the current level of battery
The number of times a car can be driven can be calculated by battery / batteryDrain.
The maximum distance the car can cover is the product of the car's speed and the number of times it can be driven.
Knowing the maximum distance the car can drive, compare it with the distance of the race track */
// CanFinish checks if a car is able to finish a certain track.
