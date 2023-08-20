package main

import (
	"fmt"
	"testing"
)

func Test_channels(t *testing.T) {
	ch := make(chan string)

	go func() {
		ch <- "message from goroutine"
	}()

	fmt.Println(<-ch)
}
