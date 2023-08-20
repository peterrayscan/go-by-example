package main

import (
	"fmt"
	"testing"
)

func Test_channel_buffering(t *testing.T) {
	ch := make(chan string, 2)

	ch <- "first"
	ch <- "second"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
