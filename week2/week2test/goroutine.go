package week2test

import "fmt"

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

// TestGoRoutine shows how goroutines work
func TestGoRoutine() {
	f("direct")
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	fmt.Scanln()
	fmt.Println("done")
}
