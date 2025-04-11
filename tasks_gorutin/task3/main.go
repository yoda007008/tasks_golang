package main

import (
	"fmt"
)

func Sender(ch chan string) {
	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("Number %d", i)
	}
	close(ch)
}

func main() {
	ch := make(chan string)

	go Sender(ch)

	for i := range ch {
		fmt.Println(i)
	}

}
