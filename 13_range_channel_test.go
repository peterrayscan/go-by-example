package main

import (
	"fmt"
	"testing"
)

func Test_range_channel(t *testing.T) {
	queue := make(chan string, 2)
	queue <- "first"
	queue <- "two"
	close(queue)

	// range over a channel, only has one iteration variable(channel data).
	for elem := range queue {
		fmt.Println(elem)
	}
}
