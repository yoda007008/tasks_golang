package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var result int32

func Increment() {
	atomic.AddInt32(&result, 1)
}

func main() {
	num := 10 // количество горутин

	for i := 0; i < num; i++ {
		go Increment()
	}

	time.Sleep(2 * time.Millisecond)
	fmt.Println(result)
}

//Задача: Напишите программу, где 10 горутин инкрементируют один счётчик без использования мютексов, через атомики.
