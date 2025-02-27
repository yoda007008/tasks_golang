package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{5, 3, 8, 1, 2}
	result := SortSlice(input)
	fmt.Println(result) // [1, 2, 3, 5, 8]
}

func SortSlice(input []int) []int {
	sort.Ints(input)
	return input
}
