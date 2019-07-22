package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	a := getInputFloat("Enter acceleration: ")
	v0 := getInputFloat("Enter initial velocity: ")
	s0 := getInputFloat("Enter initial displacement: ")

	fn := GenDisplaceFn(a, v0, s0)
	t := getInputFloat("Enter time: ")

	fmt.Printf("Displacement at %0.02f seconds: %0.02f\n", t, fn(t))
}

func getInputFloat(prompt string) float64 {
	var s string
	fmt.Printf(prompt)
	fmt.Scan(&s)
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return (0.5 * a * t * t) + (v0 * t) + s0
	}
}
