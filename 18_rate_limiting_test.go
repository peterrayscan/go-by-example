package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_rate_limiting_v1(t *testing.T) {
	// Process a request every 1 second, continuously limiting speed.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.NewTicker(1 * time.Second)

	for req := range requests {
		fmt.Println(<-limiter.C, "handle req:", req)
	}
}

func Test_rate_limiting_v2(t *testing.T) {
	//Supports simultaneous serial processing of three requests in a queue.
	// By default, one request is placed in the queue every second.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.NewTicker(1 * time.Second)

	queue := make(chan time.Time, 3)
	queue <- time.Now()
	queue <- time.Now()
	queue <- time.Now()

	go func() {
		for t := range limiter.C {
			queue <- t
		}
	}()

	for req := range requests {
		fmt.Println(<-queue, "handle req:", req)
	}
}
