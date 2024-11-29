package goquickdata

import (
	"bufio"
	"io"
	"os"

	"github.com/go-mmap/mmap"
)

// DataScanner is a wrapper around bufio.Scanner that also implements the io.Closer interface.
// This is useful for opening files with either implementation and closing them when we're done.
type DataScanner struct {
	*bufio.Scanner
	closer io.Closer
}

func newDataScanner(scanner *bufio.Scanner, closer io.Closer) *DataScanner {
	return &DataScanner{Scanner: scanner, closer: closer}
}

// Close closes the file and retuns a error if something happend
func (df *DataScanner) Close() error {
	return df.closer.Close()
}

// OpenFile opens the file with the os package, and returns a [DataScanner]
func OpenFile(path string) (*DataScanner, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	return newDataScanner(scanner, file), nil
}

// OpenFileMmap first of all does a syscall depending on os and memory maps the file specified.
// Memory mapping reads the whole file into memory (meaning that the process uses alot of ram),
// But this is beneficial because of the increased speeds.
func OpenFileMmap(path string) (*DataScanner, error) {
	mmapFile, err := mmap.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(mmapFile)
	return newDataScanner(scanner, mmapFile), nil
}
