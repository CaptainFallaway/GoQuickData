package goquickdata

import (
	"errors"
	"io"
	"sync"
)

type QuickData[R any] struct {
	options *Options

	data Chunkable

	taskQueue   chan []byte
	resultQueue chan R
	shutdownWg  *sync.WaitGroup

	// err stores the first error that occurs during processing.
	err error
}

func NewQuickData[R any](data Chunkable, opts ...Opt) *QuickData[R] {
	options := new(Options)

	for _, opt := range opts {
		opt(options)
	}

	return &QuickData[R]{
		options:     options,
		data:        data,
		taskQueue:   make(chan []byte, options.ChunkSize),
		resultQueue: make(chan R, options.ChunkSize),
		shutdownWg:  new(sync.WaitGroup),
	}
}

func (wp *QuickData[R]) worker(procFunc func([]byte) R) {
	for task := range wp.taskQueue {
		res := procFunc(task)
		wp.resultQueue <- res
	}

	wp.shutdownWg.Done()
}

// Setup initializes the worker pool with the given processing function.
func (wp *QuickData[R]) Start(procFunc func([]byte) R) {
	wp.shutdownWg.Add(wp.options.MaxWorkers)

	for i := 0; i < wp.options.MaxWorkers; i++ {
		go wp.worker(procFunc)
	}

	go func() {
		for {
			chunk, err := wp.data.GetChunk()
			if err != nil {
				if !errors.Is(err, io.EOF) {
					wp.err = err
				}
				break
			}

			wp.taskQueue <- chunk
		}
		wp.Stop()
	}()
}

func (wp *QuickData[R]) Err() error {
	return wp.err
}

func (wp *QuickData[R]) Results() <-chan R {
	return wp.resultQueue
}

func (wp *QuickData[R]) Stop() {
	close(wp.taskQueue)
	wp.shutdownWg.Wait()
	close(wp.resultQueue)
}
