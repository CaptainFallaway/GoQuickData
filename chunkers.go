package goquickdata

import (
	"bufio"
	"bytes"
	"io"
)

type Chunkable interface {
	GetChunk() ([]byte, error)
}

type ExpChunker struct {
	reader io.Reader

	buf      []byte
	leftover []byte

	split string
}

func NewExpChunker(reader io.Reader, size int, split string) *ExpChunker {
	return &ExpChunker{
		reader:   reader,
		buf:      make([]byte, size),
		leftover: make([]byte, 0, size),
		split:    split,
	}
}

func (c *ExpChunker) GetChunk() ([]byte, error) {
	readTotal, err := c.reader.Read(c.buf)
	if err != nil {
		return nil, err
	}

	c.buf = c.buf[:readTotal]

	ret := make([]byte, readTotal)
	copy(ret, c.buf)

	lastNewLineIndex := bytes.LastIndex(c.buf, []byte(c.split))

	ret = append(c.leftover, c.buf[:lastNewLineIndex]...)
	c.leftover = make([]byte, len(c.buf[lastNewLineIndex:]))
	copy(c.leftover, c.buf[lastNewLineIndex:])

	return ret, nil
}

type Chunker struct {
	scanner *bufio.Scanner
	size    int
}

func NewChunker(reader io.Reader, size int) *Chunker {
	return &Chunker{
		scanner: bufio.NewScanner(reader),
		size:    size,
	}
}

func (c *Chunker) GetChunk() ([]byte, error) {
	buf := make([]byte, 0)
	for i := 0; i < c.size && c.scanner.Scan(); i++ {
		buf = append(buf, c.scanner.Bytes()...)
		if i != c.size-1 {
			buf = append(buf, '\n')
		}
	}
	if err := c.scanner.Err(); err != nil {
		return nil, err
	}
	if len(buf) == 0 {
		return nil, io.EOF
	}
	return buf, nil
}
