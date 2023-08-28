package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_timer(t *testing.T) {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("timer1 is fired.")

	timer2 := time.NewTimer(1 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 is fired.")
	}()

	// in some conditiion, you may want to stop the timer before it fired.
	if true {
		success := timer2.Stop()
		if success {
			fmt.Println("timer2 is stopped.")
		}
	}
	time.Sleep(2 * time.Second)
}
