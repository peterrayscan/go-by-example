package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_defer(t *testing.T) {
	file, err := os.Create("/workspaces/go-by-example/temp.txt")
	if err != nil {
		panic(err)
	}
	defer closeFile(file)

	fmt.Fprintln(file, "input data")
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		os.Exit(1)
	}
}
