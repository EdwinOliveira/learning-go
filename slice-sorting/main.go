package main

import (
	"fmt"
	"slices"
)

func main() {

    strs := []string{"c", "a", "b"}
    slices.Sort(strs)
    fmt.Println("Strings:", strs)

		textSort := slices.IsSorted(strs)
    fmt.Println("Sorted: ", textSort)

    ints := []float32{7, 2, 4, 1, 1.5, 2.5}
    slices.Sort(ints)
    fmt.Println("Ints:   ", ints)

    s := slices.IsSorted(ints)
    fmt.Println("Sorted: ", s)
}