package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	goquickdata "github.com/CaptainFallaway/GoQuickData"
)

func worker(cs <-chan []string, res chan<- int, wg *sync.WaitGroup) {
	var (
		sum  = 0
		temp = 0
	)
	for chunk := range cs {
		sum = 0
		for _, row := range chunk {
			p := strings.Split(row, ",")[3]
			temp, _ = strconv.Atoi(p)
			sum += temp
		}
		res <- sum
	}
	wg.Done()
}

func main() {
	start := time.Now()
	file, err := goquickdata.OpenFileMmap("./testdata/yellow_tripdata_2015-03.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	const chunkSize = 100_000

	cpus := runtime.NumCPU()
	workers := cpus - 2

	wg := &sync.WaitGroup{}
	chunkStream := make(chan []string, workers)
	resStream := make(chan int, 1)

	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go worker(chunkStream, resStream, wg)
	}

	go func() {
		file.Scan() // Skip first row
		buf := make([]string, chunkSize)
		for i := 0; file.Scan(); i++ {
			buf = append(buf, file.Text())
			if i == chunkSize {
				i = 0
				chunkStream <- buf
				buf = buf[:0]
			}
		}
		close(chunkStream)
		wg.Wait()
		close(resStream)
	}()

	sum := 0
	for res := range resStream {
		sum += res
	}

	fmt.Println(time.Since(start))
	fmt.Println(sum)
}
