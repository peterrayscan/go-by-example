package main

import (
	"fmt"
	"testing"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- string) {
	for jobID := range jobs {
		fmt.Println("worker:", id, "start job:", jobID)
		time.Sleep(1 * time.Second)
		result := fmt.Sprintf("worker: %d finished job: %d", id, jobID)
		fmt.Println(result)
		results <- result
	}
}

func Test_work_pool(t *testing.T) {
	const jobNum = 5
	jobs := make(chan int, jobNum)
	results := make(chan string, jobNum)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// send jobs to work
	for j := 1; j <= jobNum; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= jobNum; a++ {
		<-results
	}
}
