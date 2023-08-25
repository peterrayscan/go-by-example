package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_timeout_handling(t *testing.T) {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "c1 has done."
	}()

	select {
	case got := <-c1:
		fmt.Println(got)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout in c1.")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "c2 has done."
	}()

	select {
	case got := <-c2:
		fmt.Println(got)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout in c2.")
	}
}
