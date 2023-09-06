package main

import (
	"os"
	"testing"
)

func Test_panic_v1(t *testing.T) {
	panic("panic reason: xxx")
}

func Test_panic_v2(t *testing.T) {
	_, err := os.Create("/not_exist/temp.txt")
	if err != nil {
		panic(err)
	}
}
