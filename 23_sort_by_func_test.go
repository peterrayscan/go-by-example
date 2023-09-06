package main

import (
	"fmt"
	"sort"
	"testing"
)

type Student struct {
	Name   string
	Age    int
	Height int
}

func Test_sort_by_func_v1(t *testing.T) {
	students := []Student{
		{"Alice", 23, 175},
		{"John", 18, 175},
		{"Eve", 18, 165},
		{"David", 18, 185},
		{"Bob", 25, 170},
	}

	// rule: age asc

	sort.SliceStable(students, func(i, j int) bool {
		return students[i].Age < students[j].Age
	})
	fmt.Println(students)
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Test_sort_by_func_v2(t *testing.T) {
	students := []Student{
		{"Alice", 23, 175},
		{"John", 18, 175},
		{"Eve", 18, 165},
		{"David", 18, 185},
		{"Bob", 25, 170},
	}
	sort.Stable(Students(students))
	fmt.Println(students)
}
