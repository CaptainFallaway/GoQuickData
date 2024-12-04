package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	goquickdata "github.com/CaptainFallaway/GoQuickData"
	"github.com/go-mmap/mmap"
)

// func worker(cs <-chan []string, res chan<- int, wg *sync.WaitGroup) {
// 	var (
// 		sum  = 0
// 		temp = 0
// 	)
// 	for chunk := range cs {
// 		sum = 0
// 		for _, row := range chunk {
// 			p := strings.Split(row, ",")[3]
// 			temp, _ = strconv.Atoi(p)
// 			sum += temp
// 		}
// 		res <- sum
// 	}
// 	wg.Done()
// }

func sumAdd(sums ...*Sum) *Sum {
	ret := new(Sum)
	for _, sum := range sums {
		ret.VendorID += sum.VendorID
		ret.passengerCount += sum.passengerCount
		ret.tripDistance += sum.tripDistance
		ret.pickupLon += sum.pickupLon
		ret.pickupLat += sum.pickupLat
		ret.rateCodeID += sum.rateCodeID
		ret.dropoffLon += sum.dropoffLon
		ret.dropoffLat += sum.dropoffLat
		ret.paymentType += sum.paymentType
		ret.fareAmmount += sum.fareAmmount
		ret.extra += sum.extra
		ret.mtaTax += sum.mtaTax
		ret.tipAmmount += sum.tipAmmount
		ret.tollsAmmount += sum.tollsAmmount
		ret.improvementSurcharge += sum.improvementSurcharge
		ret.totalAmmount += sum.totalAmmount
	}
	return ret
}

func processRow(row string) *Sum {
	ret := new(Sum)
	p := strings.Split(row, ",")
	ret.VendorID, _ = strconv.Atoi(p[0])
	ret.passengerCount, _ = strconv.Atoi(p[1])
	ret.tripDistance, _ = strconv.ParseFloat(p[2], 64)
	ret.pickupLon, _ = strconv.ParseFloat(p[3], 64)
	ret.pickupLat, _ = strconv.ParseFloat(p[4], 64)
	ret.rateCodeID, _ = strconv.Atoi(p[5])
	ret.dropoffLon, _ = strconv.ParseFloat(p[6], 64)
	ret.dropoffLat, _ = strconv.ParseFloat(p[7], 64)
	ret.paymentType, _ = strconv.Atoi(p[8])
	ret.fareAmmount, _ = strconv.ParseFloat(p[9], 64)
	ret.extra, _ = strconv.ParseFloat(p[10], 64)
	ret.mtaTax, _ = strconv.ParseFloat(p[11], 64)
	ret.tipAmmount, _ = strconv.ParseFloat(p[12], 64)
	ret.tollsAmmount, _ = strconv.ParseFloat(p[13], 64)
	ret.improvementSurcharge, _ = strconv.ParseFloat(p[14], 64)
	ret.totalAmmount, _ = strconv.ParseFloat(p[15], 64)
	return ret
}

func worker(chunk []byte) []*Sum {
	c := strings.Split(string(chunk), "\n")

	buf := make([]*Sum, 0, len(c))

	for _, row := range c {
		if len(row) == 0 {
			continue
		}
		buf = append(buf, processRow(row))
	}

	return buf
}

type Sum struct {
	VendorID             int
	passengerCount       int
	tripDistance         float64
	pickupLon            float64
	pickupLat            float64
	rateCodeID           int
	dropoffLon           float64
	dropoffLat           float64
	paymentType          int
	fareAmmount          float64
	extra                float64
	mtaTax               float64
	tipAmmount           float64
	tollsAmmount         float64
	improvementSurcharge float64
	totalAmmount         float64
}

// const chunkSize = 67_108_864

func main() {
	start := time.Now()
	file, err := mmap.Open("./testdata/yellow_tripdata_2015-03.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("File opened", time.Since(start))

	stat, _ := file.Stat()
	cpus := runtime.NumCPU()

	chunker := goquickdata.NewExpChunker(file, (int(stat.Size())/cpus)/4, "\n")
	fmt.Println("Chunker created", time.Since(start))

	wp := goquickdata.NewQuickData[[]*Sum](chunker, goquickdata.MaxWorkers())
	wp.Start(worker)
	fmt.Println("Worker pool started", time.Since(start))

	buf := make([]*Sum, 0)
	for res := range wp.Results() {
		buf = append(buf, res...)
	}
	fmt.Println(time.Since(start))
	sum := sumAdd(buf...)

	fmt.Println(time.Since(start))
	fmt.Println(sum)
}
