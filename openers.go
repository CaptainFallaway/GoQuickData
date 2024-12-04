package goquickdata

import (
	"io"
	"os"

	"github.com/go-mmap/mmap"
)

// opens the file with the os package, and returns a [DataScanner]
func OpenFile(path string) (io.ReadCloser, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// OpenFileMmap first of all does a syscall depending on os and memory maps the file specified.
// Memory mapping reads the whole file into memory (meaning that the process uses alot of ram),
// But this is beneficial because of the increased speeds.
func OpenFileMmap(path string) (io.ReadCloser, error) {
	mmapFile, err := mmap.Open(path)
	if err != nil {
		return nil, err
	}
	return mmapFile, nil
}
