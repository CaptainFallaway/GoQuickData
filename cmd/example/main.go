package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("./testdata/yellow_tripdata_2015-03.csv")
	if err != nil {
		panic(err)
	}

	bufio.NewScanner(file)
}
