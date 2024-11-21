package goquickdata

type Chunker interface {
	GetChunk() [][]byte
}
