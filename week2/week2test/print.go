package week2test

import "fmt"

// TestPrint shows how defer works
func TestPrint() {
	i := 1
	fmt.Print(i)
	i++
	defer fmt.Print(i + 1)

	fmt.Print(i)
}
