package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_timestamp(t *testing.T) {
	now := time.Now()
	fmt.Println(now)
	fmt.Println("Seconds from Uinx:\t\t", now.Unix())
	fmt.Println("MilliSeconds form Uinx:\t\t", now.UnixNano()/1000000)
	fmt.Println("MicroSeconds form Uinx:\t\t", now.UnixNano()/1000)
	fmt.Println("NanoSeconds form Uinx:\t\t", now.UnixNano())
	fmt.Println("")
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
