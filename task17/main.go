package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	result := ConcatSlices(slice1, slice2)
	fmt.Println(result) // [1, 2, 3, 4, 5, 6]
}

func ConcatSlices(slice1 []int, slice2 []int) []int {
	newSlice := make([]int, 0, 6)
	newSlice = append(slice1, slice2...)
	return newSlice
}
