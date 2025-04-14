package main

import (
	"crypto/md5"
	"fmt"
	"hash/crc32"
	"strconv"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	data := []int{0, 1, 2, 3, 4, 52}
	//ch := make(chan int, len(data))
	//
	//for _, v := range data {
	//
	//}
}

func SingleHash(data int) string {
	strData := strconv.Itoa(data)

	crcDate := CountDataSingerCrc32(strData)
	mdDate := CountDataSignerMd5(strData)

	res := crcDate + "~" + mdDate
	return res
}

func ExecutePipeline(data int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
}

func CountDataSignerMd5(data string) string {
	crcH := crc32.ChecksumIEEE([]byte(data))
	dateHash := strconv.FormatInt(uint64(crcH), 10)
	time.Sleep(time.Second)
	return dateHash
}

func CountDataSingerCrc32(data string) string {
	dataHash := fmt.Sprintf("%x", md5.Sum([]byte(data)))
	time.Sleep(time.Millisecond * 10)
	return dataHash
}
