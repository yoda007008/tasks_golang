package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mu      sync.Mutex
)

func IncrementCount() {
	defer wg.Done()
	mu.Lock()
	counter += 1
	mu.Unlock()
}

func main() {
	count := 10 // количество горутин
	wg.Add(count)
	for i := 0; i < count; i++ {
		go IncrementCount()
	}
	wg.Wait()
	fmt.Println(count)
}

//Потокобезопасный инкремент - Mutex.
//Задача: Напишите программу, где 10 горутин инкрементируют один счётчик, защищая его sync.Mutex.
