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

//Задача: Напишите программу, которая запускает 5 горутин, каждая из которых печатает свой номер (от 1 до 5),
//и использует sync.WaitGroup для их синхронизации(нужно подождать их выполнения).
