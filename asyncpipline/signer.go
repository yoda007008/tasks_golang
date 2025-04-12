package main

import (
	"sync"
)

func main() {
	//start := time.Now()

}

func ExecutePipeline(data int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
}

func CountDataSignerMd5() {

}

func CountDataSingerCrc32() {

}
