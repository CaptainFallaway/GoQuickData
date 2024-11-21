package goquickdata

import (
	"bufio"
)

type QuickReducer[T any] struct {
	scanner *bufio.Scanner

	opts *Options

	stopChan   chan struct{}
	chunkChan  chan [][]byte
	resultChan chan T
}

func NewQuickReducer[T any](scanner *bufio.Scanner, opts ...OptFunc) *QuickReducer[T] {
	options := NewOptions()

	for _, cf := range opts {
		cf(options)
	}

	chunkChanSize := options.workerAmmount * options.chunkSize
	resultChanSize := options.workerAmmount * chunkChanSize

	return &QuickReducer[T]{
		chunkChan:  make(chan [][]byte, chunkChanSize),
		resultChan: make(chan T, resultChanSize),
		stopChan:   make(chan struct{}),
		scanner:    scanner,
		opts:       options,
	}
}

func (qr *QuickReducer[T]) StartProcessing() {

}

func (qr *QuickReducer[T]) StopProcessing() {

}

func (qr *QuickReducer[T]) GetResultChan() chan T {
	return nil
}
