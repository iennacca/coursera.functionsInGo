package main

import "fmt"

func BubbleSort(a []int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9-i; j++ {
			if a[j] > a[j+1] {
				Swap(&a[j], &a[j+1])
			}
		}
	}
}

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	var arr [10]int
	for i := range arr {
		fmt.Printf("Enter a number a[%d]:\n", i)
		_, _ = fmt.Scanln(&arr[i])
	}

	BubbleSort(arr[:])

	for i := 0; i < 10; i++ {
		fmt.Println(arr[i])
	}

}