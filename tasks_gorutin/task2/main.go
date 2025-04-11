package main

import (
	"fmt"
	"sync"
)

func Goroutine(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutine number", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Goroutine(i, &wg)
	}
	wg.Wait()
}
