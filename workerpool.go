package goquickdata

import "sync"

type Task[T any, R any] struct {
	Input  T
	Result chan R
}

type WorkerPool[T any, R any] struct {
	maxWorkers  int
	taskQueue   chan Task[T, R]
	resultQueue chan R
	wg          *sync.WaitGroup

	started bool
	stopped bool
}

func NewWorkerPool[T any, R any](maxWorkers int) *WorkerPool[T, R] {
	return &WorkerPool[T, R]{
		maxWorkers:  maxWorkers,
		taskQueue:   make(chan Task[T, R], maxWorkers),
		resultQueue: make(chan R, maxWorkers),
		wg:          &sync.WaitGroup{},
	}
}

func (wp *WorkerPool[T, R]) Start(procFunc func(T) R) {
	wp.started = true
	wp.wg.Add(wp.maxWorkers)
	for i := 0; i < wp.maxWorkers; i++ {
		go wp.worker(procFunc)
	}
}

func (wp *WorkerPool[T, R]) worker(procFunc func(T) R) {
	for task := range wp.taskQueue {
		res := procFunc(task.Input)
		wp.resultQueue <- res
		task.Result <- res
	}
	wp.wg.Done()
}

// Submit also Returns optional channel if you want instant result for this task
func (wp *WorkerPool[T, R]) Submit(input T) <-chan R {
	if !wp.started {
		panic("workerpool never started")
	} else if wp.stopped {
		panic("submitted input to closed workerpool")
	}

	resultChan := make(chan R, 1)
	wp.taskQueue <- Task[T, R]{input, resultChan}
	return resultChan
}

func (wp *WorkerPool[T, R]) Results() <-chan R {
	return wp.resultQueue
}

func (wp *WorkerPool[T, R]) Stop() {
	close(wp.taskQueue)
}
