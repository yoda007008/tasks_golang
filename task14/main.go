package main

import "fmt"

func main() {
	input := []int{4, 2, 3, 2, 4, 5, 1}
	output := RemoveDuplicates(input)
	fmt.Println(output)
}

func RemoveDuplicates(input []int) []int {
	hashTable := make(map[int]bool)
	newSlice := make([]int, 0)
	for _, i := range input {
		if !hashTable[i] { // если элемент еще не встречался, то добавляем, иначе он уже встречавшийся
			newSlice = append(newSlice, i)
			hashTable[i] = true
		}
	}
	return newSlice
}
