package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func Test_non_atomic_counter(t *testing.T) {
	var num uint32
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			num = num + 1
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(num)
}

func Test_atomic_counter(t *testing.T) {
	var num uint32
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddUint32(&num, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(num)
}
