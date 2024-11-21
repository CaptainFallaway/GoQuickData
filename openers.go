package goquickdata

import (
	"bufio"
	"io"
	"os"

	"github.com/go-mmap/mmap"
)

type Scannable interface {
	GetScanner() *bufio.Scanner
}

type DataFile struct {
	scanner *bufio.Scanner
	closer  io.Closer
}

func (df *DataFile) Close() error {
	return df.closer.Close()
}

func OpenFile(path string) (*DataFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	return &DataFile{scanner: scanner, closer: file}, nil
}

func OpenFileMmap(path string) (*DataFile, error) {
	mmapFile, err := mmap.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(mmapFile)
	return &DataFile{scanner: scanner, closer: mmapFile}, nil
}
