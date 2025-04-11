package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4, 5, 6}
	isEven := func(x int) bool { return x%2 == 0 }
	even, odd := SplitSlice(input, isEven)
	fmt.Println(even) // [2, 4, 6]
	fmt.Println(odd)  // [1, 3, 5]
}

func SplitSlice(input []int, isEven func(x int) bool) ([]int, []int) {
	var even []int
	var odd []int
	for _, value := range input {
		if isEven(value) {
			even = append(even, value)
		} else {
			odd = append(odd, value)
		}
	}
	return even, odd
}
