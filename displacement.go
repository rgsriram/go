package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (0.5 * a * math.Pow(t, 2)) + (v * t) + s
	}
	return fn
}

func main() {
	var acceleration, velocity, displacement, timeToWait float64

	fmt.Println("Enter acceleration: ")
	fmt.Scan(&acceleration)

	fmt.Println("Enter initial velocity: ")
	fmt.Scan(&velocity)

	fmt.Println("Enter inital displacement: ")
	fmt.Scan(&displacement)

	fmt.Println("Enter the time to wait: ")
	fmt.Scan(&timeToWait)

	fn := GenDisplaceFn(acceleration, velocity, displacement)
	fmt.Println(fn(timeToWait))
}
