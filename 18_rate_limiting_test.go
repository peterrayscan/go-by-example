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
		<-limiter.C
		fmt.Println(time.Now(), "handle req:", req)
	}
}

func Test_rate_limiting_v2(t *testing.T) {
	// Process a request every 1 second, continuously limiting speed.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.NewTicker(1 * time.Second)

	for req := range requests {
		<-limiter.C
		fmt.Println(time.Now(), "handle req:", req)
	}
}
