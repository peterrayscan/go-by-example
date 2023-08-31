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
	go ac.AddBalance(&wg, "ali_pay", 10000)
	go ac.AddBalance(&wg, "wechat_pay", 10000)
	go ac.AddBalance(&wg, "wechat_pay", 10000)
	wg.Wait()

	fmt.Println(ac.balance)
}

// ------------------------------------

// Using channels as a non capacity message queue, adding values without mutex lock.

type AccountV2 struct {
	ch      chan string
	balance map[string]int32
}

func (a *AccountV2) syncFromChan() {
	for to := range a.ch {
		a.balance[to] = a.balance[to] + 1
	}
}

func (a *AccountV2) AddBalanceV2(wg *sync.WaitGroup, to string, amount int32) {
	for i := 0; i < int(amount); i++ {
		a.ch <- to
	}
	wg.Done()
}

func Test_no_mutex_sync_by_channel(t *testing.T) {
	ac := AccountV2{
		ch: make(chan string),
		balance: map[string]int32{
			"ali_pay":    0,
			"wechat_pay": 0,
		},
	}

	go ac.syncFromChan()

	var wg sync.WaitGroup
	wg.Add(3)
	go ac.AddBalanceV2(&wg, "ali_pay", 10000)
	go ac.AddBalanceV2(&wg, "wechat_pay", 10000)
	go ac.AddBalanceV2(&wg, "wechat_pay", 10000)
	wg.Wait()

	close(ac.ch)

	fmt.Println(ac.balance)
}
