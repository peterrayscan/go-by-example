package main

import (
	"fmt"
	"sync"
	"testing"
)

type Account struct {
	mu      sync.Mutex
	balance map[string]int32
}

func (a *Account) add(to string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance[to] = a.balance[to] + 1
}

func (a *Account) AddBalance(wg *sync.WaitGroup, to string, amount int32) {
	for i := 0; i < int(amount); i++ {
		a.add(to)
	}
	wg.Done()
}

func Test_mutex(t *testing.T) {
	ac := Account{
		balance: map[string]int32{
			"ali_pay":    0,
			"wechat_pay": 0,
		},
	}

	var wg sync.WaitGroup

	wg.Add(3)
	go ac.AddBalance(&wg, "ali_pay", 1000)
	go ac.AddBalance(&wg, "wechat_pay", 1000)
	go ac.AddBalance(&wg, "wechat_pay", 1000)
	wg.Wait()

	fmt.Println(ac.balance)
}
