package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4, 5}
	result := ReverseSlice(input)
	fmt.Println(result) // [5, 4, 3, 2, 1]
}

func ReverseSlice(input []int) []int {
	var newArray []int // reverse массива
	for i := len(input) - 1; i >= 0; i-- {
		newArray = append(newArray, input[i])
	}
	return newArray
}
