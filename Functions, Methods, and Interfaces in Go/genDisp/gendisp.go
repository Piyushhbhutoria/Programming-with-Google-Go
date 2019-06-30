package main

import (
	"fmt"
)

func main() {
	var a, v0, s0, t float64
	fmt.Println("Enter the acceleration: ")
	fmt.Scanln(&a)
	fmt.Println("Enter the initial velocity: ")
	fmt.Scanln(&v0)
	fmt.Println("Enter the initial displacement: ")
	fmt.Scanln(&s0)
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println("Enter the value of time t: ")
	fmt.Scanln(&t)
	fmt.Println("Displacement:", fn(t))
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return s0 + v0*t + 0.5*a*t*t
	}
}
