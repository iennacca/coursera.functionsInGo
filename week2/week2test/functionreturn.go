package week2test

import (
	"fmt"
)

func fA() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func testFunctionReturn() {
	fB := fA()
	fmt.Print(fB())
	fmt.Print(fB())
}
