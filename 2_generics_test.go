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

func Test_generics_v1(t *testing.T) {
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

type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	head, tail *element[T]
}

func (lst *List[T]) push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) getAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func Test_generics_v2(t *testing.T) {
	var lst List[int]
	lst.push(1)
	lst.push(12)
	lst.push(123)
	lst.push(1234)
	lst.push(12345)
	fmt.Println(lst.getAll())
}
