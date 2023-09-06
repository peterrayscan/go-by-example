package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_recover(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("create file with error")
		}
	}()
	createFile("/xxx/not_exist/test.txt")

	fmt.Println("will not output if panic before")
}

func createFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}
