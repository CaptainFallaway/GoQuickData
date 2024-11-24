package main

import (
	"bufio"
	"fmt"
	"time"

	goquickdata "github.com/CaptainFallaway/GoQuickData"
)

func main() {
	wholeTestStart := time.Now()
	file, err := goquickdata.OpenFileMmap("./testdata/yellow_tripdata_2015-03.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileOpenTime := fmt.Sprintf("Took %v To Open File", time.Since(wholeTestStart))

	actual := time.Now()

	scanner := file.Scanner

	scanner.Split(bufio.ScanLines)
	scanner.Text()

	i := 0

	for scanner.Scan() {
		scanner.Text()
		i++
	}

	fmt.Printf("%v Lines\n%v\n", i, fileOpenTime)
	fmt.Printf("Reading Lines Took %v\n", time.Since(actual))
	fmt.Printf("Whole Test Took %v", time.Since(wholeTestStart))

	for {
	}
}
