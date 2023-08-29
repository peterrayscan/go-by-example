package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func workerV2(id int) {
	fmt.Printf("worker: %d, starting.\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker: %d, done.\n", id)
}

func Test_wait_group(t *testing.T) {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		j := i

		go func() {
			defer wg.Done()
			workerV2(j)
		}()
	}
	wg.Wait()
}
