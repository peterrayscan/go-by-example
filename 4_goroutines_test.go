package main

import (
	"fmt"
	"testing"
	"time"
)

func doSth() {
	// do sth here
	time.Sleep(1 * time.Second)

	// print result
	fmt.Println("now:", time.Now().Format("2006-01-02 15:04:05.000000000"))
}

func Test_goroutines(t *testing.T) {
	fmt.Println("without goroutines")
	for i := 0; i < 3; i++ {
		doSth()
	}

	fmt.Println()

	fmt.Println("with goroutines")
	for i := 0; i < 3; i++ {
		go doSth()
	}

	// wait until goroutine is done
	time.Sleep(2 * time.Second)
}
