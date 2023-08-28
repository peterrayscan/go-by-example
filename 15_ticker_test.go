package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_ticker(t *testing.T) {
	ticker := time.NewTicker(800 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("now:", time.Now())
			case <-done:
				return
			}
		}
	}()

	// // wait for 3 times ticker's ticking

	// time.Sleep(2500 * time.Millisecond)
	timer := time.NewTimer(2500 * time.Millisecond)
	<-timer.C

	// ticker must be stopped, otherwise it will keeping ticking.
	ticker.Stop()
	fmt.Println("ticker stopped.")

	// also need to close goroutine's for loop.
	done <- true
	fmt.Println("goroutine killed.")
}
