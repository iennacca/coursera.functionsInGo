package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	var a, vo, so float64

	scanner := bufio.NewScanner(os.Stdin)
	//asking for values
	fmt.Print("What's the acceleration (a)? \t")
	scanner.Scan()
	a, _ = strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("What's the initial velocity (v0)? \t")
	scanner.Scan()
	vo, _ = strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("What's the initial displacement (s0)? \t")
	scanner.Scan()
	so, _ = strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("What's the time ('t' in seconds...)? \t")
	scanner.Scan()
	t, _ := strconv.ParseFloat(scanner.Text(), 64)

	fn := GenDisplaceFn(a, vo, so)

	fmt.Println(fn(t))

}

func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	fn := func(t float64) float64 {
		return ((a * math.Pow(t, 2) / 2) + (v0 * t) + s0)
	}
	return fn
}