package quickreducer

import "io/fs"

type Config struct {
	Mmap        bool
	WorkerCount int
	File        fs.File
}
