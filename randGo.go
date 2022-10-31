package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	SeedWithTime()
	RollADie()
	GenerateWandEnergy()

}

func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(50))

}

func RollADie() int {
	startNumber := 1
	endNumber := 20
	n := startNumber + rand.Intn(endNumber-startNumber+1) // a ≤ n ≤ b
	fmt.Println(n)
	return n
}

func GenerateWandEnergy() float64 {
	min := 0.0
	max := 12.0

	fmt.Println(min + rand.Float64()*(max-min))
	return 1.0
}
