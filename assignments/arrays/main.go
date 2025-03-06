package main

import "fmt"

func main() {
	values := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var oddCount int
	var evenCount int

	for i := range len(values) {
		if (i % 2 == 0) {
			oddCount += 1
			fmt.Println("Odd Number =", values[i])
		} else {
			evenCount += 1
			fmt.Println("Even Number =", values[i])
		}
	}

	for i := len(values) - 1; i >= 0; i-- {
		if (i % 2 == 0) {
			oddCount += 1
			fmt.Println("Odd Number =", values[i])
		} else {
			evenCount += 1
			fmt.Println("Even Number =", values[i])
		}
	}
}