package goquickdata

type Options struct {
	workerAmmount int
	chunkSize     int
}

type OptFunc func(*Options)

// NewOptions returns a new Options struct with default values
func NewOptions() *Options {
	return &Options{
		workerAmmount: 1,
		chunkSize:     10,
	}
}

func WorkerAmmount(ammount int) OptFunc {
	return func(qr *Options) {
		qr.workerAmmount = ammount
	}
}

func ChunkSize(size int) OptFunc {
	return func(qr *Options) {
		qr.chunkSize = size
	}
}
