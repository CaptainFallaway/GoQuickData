package goquickdata

type Options struct {
	workerAmmount int
	chunkSize     int
}

// NewOptions returns a new Options struct with default values
func NewOptions() *Options {
	return &Options{
		workerAmmount: 1,
		chunkSize:     10,
	}
}

type OptFunc func(*Options)

func WorkerAmmount[T any](ammount int) OptFunc {
	return func(qr *Options) {
		qr.workerAmmount = ammount
	}
}

func ChunkSize[T any](size int) OptFunc {
	return func(qr *Options) {
		qr.chunkSize = size
	}
}
