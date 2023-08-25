package main

import (
	"fmt"
	"testing"
)

func Test_non_blocking_channel(t *testing.T) {
	ch := make(chan bool)
	select {
	case ch <- true:
		fmt.Println("send a message to ch.")
	case <-ch:
		fmt.Println("get a message from ch.")
	default:
		fmt.Println("no action.")
	}
}
