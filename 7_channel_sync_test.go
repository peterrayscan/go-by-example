package main

import (
	"fmt"
	"testing"
	"time"
)

func sync(ch chan bool) {
	fmt.Println("working...")
	time.Sleep(1 * time.Second)
	fmt.Println("done.")
	ch <- true
}

func Test_channel_sync(t *testing.T) {
	do := make(chan bool)
	go sync(do)
	<-do
	fmt.Println("work has been finished.")
}
