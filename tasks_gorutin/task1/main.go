package main

import (
	"fmt"
	"sync"
)

func NewGoroutine(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello from goroutine")
}

func main() {
	var wg sync.WaitGroup
	for {
		wg.Add(1)
		go NewGoroutine(&wg)
		break
	}
	wg.Wait()
}

//Задача: Напишите функцию, которая запускает горутину, выполняющую fmt.Println("Hello from goroutine!"),
//и использует sync.WaitGroup для ожидания её завершения.
