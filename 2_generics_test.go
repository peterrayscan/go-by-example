package main

import (
	"fmt"
	"testing"
)

func getKeys[K comparable, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Test_generics(t *testing.T) {
	m1 := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Println(getKeys(m1))
	fmt.Println(getKeys(m2))
	fmt.Println(getKeys[int, string](m1))
	fmt.Println(getKeys[string, int](m2))
}
