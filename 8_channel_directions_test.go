package main

import (
	"fmt"
	"testing"
)

func set(ch chan<- string, msg string) {
	ch <- msg
}

func get(ch <-chan string) string {
	return <-ch
}

func Test_channel_directions(t *testing.T) {
	ch := make(chan string, 1)
	set(ch, "this is a message.")
	fmt.Println("got:", get(ch))
}
