package goquickdata

import "runtime"

type Options struct {
	MaxWorkers int
	ChunkSize  int
}

type Opt func(*Options)

func SetWorkers(maxWorkers int) Opt {
	return func(o *Options) {
		o.MaxWorkers = maxWorkers
	}
}

func MaxWorkers() Opt {
	return func(o *Options) {
		o.MaxWorkers = runtime.NumCPU()
	}
}

func ChunkSize(chunkSize int) Opt {
	return func(o *Options) {
		o.ChunkSize = chunkSize
	}
}
