package main

import (
	"fmt"
	"sort"
	"testing"
)

func Test_basic_sort(t *testing.T) {
	ints := []int{4, 6, 8, 6, 1, 2, 6, 987, 0, 7, 5, 22, 3}
	fmt.Println("before:", ints)
	fmt.Println("is sorted:", sort.IntsAreSorted(ints))
	sort.Ints(ints)
	fmt.Println("after sort:", ints)
	fmt.Println("is sorted:", sort.IntsAreSorted(ints))

	fmt.Println()

	strs := []string{"Q", "a", "A", "cat", "b", "B", "B", "c", "bug", "P", "banana"}
	fmt.Println("before:", strs)
	fmt.Println("is sorted:", sort.StringsAreSorted(strs))
	sort.Strings(strs)
	fmt.Println("after sort:", strs)
	fmt.Println("is sorted:", sort.StringsAreSorted(strs))
}
