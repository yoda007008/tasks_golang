package main

import (
	"crypto/md5"
	"fmt"
	"hash/crc32"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	var wg sync.WaitGroup
	data := []int{0, 1, 2, 3, 4, 52}
	ch := make(chan string, len(data))

	for _, v := range data {
		wg.Add(1)
		go ExecutePipeline(v, &wg, ch)
	}
	wg.Wait()
	close(ch)

	res := ""
	for result := range ch {
		res += result
	}

	res = res[:len(res)-1]

	fmt.Println(res)
	fmt.Println("Duration:", time.Since(start))
}

func ExecutePipeline(data int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	strData := strconv.Itoa(data)

	crcDateHash := CountDataSingerCrc32(strData)
	mdDateHash := CountDataSignerMd5(strData)
	crcMdHash := CountDataSingerCrc32(mdDateHash)

	combo := crcDateHash + "~" + crcMdHash // SingleHash()

	var innerWg sync.WaitGroup
	res := make([]string, 6)
	for i := 0; i < 6; i++ {
		innerWg.Add(1)
		go func(i int) {
			defer innerWg.Done()
			crc := CountDataSingerCrc32(strconv.Itoa(i) + combo)
			res[i] = crc
		}(i)
	}
	innerWg.Wait() // MultiHash()

	comboRes := strings.Join(res, "")
	ch <- comboRes + "_" // CombineResults()
}

func CountDataSignerMd5(data string) string { // функция хэширует md5
	crcH := crc32.ChecksumIEEE([]byte(data))
	dateHash := strconv.FormatInt(int64(crcH), 10)
	time.Sleep(time.Second)
	return dateHash
}

func CountDataSingerCrc32(data string) string { // функция хэширует crc32
	dataHash := fmt.Sprintf("%x", md5.Sum([]byte(data)))
	time.Sleep(time.Millisecond * 10)
	return dataHash
}
