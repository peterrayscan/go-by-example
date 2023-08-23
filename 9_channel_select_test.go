package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_channel_select(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "done after 1 second."
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "done after 2 seconds."
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("got:", msg)
		case msg := <-ch2:
			fmt.Println("got:", msg)
		}
	}
}
