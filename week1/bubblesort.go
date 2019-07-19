package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	strInput, err := getInputString("Enter a sequence of numbers separated by space: ")
	if err != nil {
		log.Fatal(err)
		return
	}
	listStrings := strings.Fields(strInput)
	listNumbers := mapToInt(listStrings, strconv.Atoi)

	fmt.Println(listNumbers)
	bubbleSort(listNumbers)
	fmt.Println(listNumbers)
}

func getNumbers() []int {
	return []int{1, 2, 3, 4, 5, 6, 9, 7, 8, 0}
}

func mapToInt(vs []string, f func(string) (int, error)) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i], _ = f(v)
	}
	return vsm
}

func getInputString(prompt string) (string, error) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", errors.New("Buffer empty")
}

func bubbleSort(l []int) {
	n := len(l)
	for i := 1; i < n; i++ {
		for j := n - 1; j >= i; j-- {
			if l[j-1] > l[j] {
				swap(l, j-1)
			}
		}
	}
}

func swap(numbers []int, pos int) {
	x := numbers[pos]
	numbers[pos] = numbers[pos+1]
	numbers[pos+1] = x
}
