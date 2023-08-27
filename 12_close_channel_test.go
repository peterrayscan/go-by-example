package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_close_channel(t *testing.T) {
	type info struct {
		id     int
		detail string
	}
	jobs := make(chan info, 5)
	done := make(chan bool)

	go func() {
		for {
			got, open := <-jobs
			if open {
				fmt.Println("got job info:", got)
				fmt.Println("do sth.")
				time.Sleep(1 * time.Second)
				fmt.Println("sth is done.")
				fmt.Println()
			} else {
				fmt.Println("got all job info.")
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- info{
			id:     i,
			detail: fmt.Sprintf("this is job: %d", i),
		}
	}

	// Closing the channel will not affect the data already sent in the channel.
	// The close operation can be understood as adding a special close flag at the end of the data.
	close(jobs)
	<-done
	fmt.Println("all jobs are done.")
}
