package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	goquickdata "github.com/CaptainFallaway/GoQuickData"
)

func Equal(a, b, r int, val string) {
	if a != b {
		fmt.Printf("Expected %d, got %d. line %d val \n %s\n", a, b, r, val)
		os.Exit(1)
	}
}

func main() {
	file, err := goquickdata.OpenFileMmap("./testdata/Marvel_Movies_Dataset.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	const chunkSize = 10

	chunker := goquickdata.NewChunker(file, chunkSize)

	for {
		chunk, err := chunker.GetChunk()
		if err != nil {
			panic(err)
		}

		t := string(chunk)
		s := strings.Split(t, "\n")
		fmt.Println(len(s))

		for _, row := range s {
			fmt.Print("asdasd")
			fmt.Println(row)
		}

		time.Sleep(1 * time.Second)
	}
}
