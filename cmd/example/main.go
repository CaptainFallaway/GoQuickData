package main

import (
	"bufio"
	"os"

	goquickdata "github.com/CaptainFallaway/GoQuickData"
)

func main() {
	file, err := os.Open("./testdata/yellow_tripdata_2015-03.csv")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	goquickdata.NewQuickReducer[any](scanner, goquickdata.WorkerAmmount(4), goquickdata.ChunkSize(100))
}
