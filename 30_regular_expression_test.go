package main

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"
)

func Test_regular_expression(t *testing.T) {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	fmt.Println()

	// r, _ := regexp.Compile("p([a-z]+)ch")
	r := regexp.MustCompile("p([a-z]+)ch")

	fmt.Println(r.Match([]byte("peach")))
	fmt.Println(r.MatchString("peach"))
	fmt.Println()

	fmt.Println(r.FindString("punch peach"), r.FindStringIndex("punch peach"))
	fmt.Println(r.FindStringSubmatch("punch peach"), r.FindStringSubmatchIndex("punch peach"))
	fmt.Println()

	fmt.Println(r.FindAllString("punch pinch peach", -1), r.FindAllStringIndex("punch pinch peach", -1))
	fmt.Println(r.FindAllStringSubmatch("punch pinch peach", -1), r.FindAllStringSubmatchIndex("punch pinch peach", -1))
	fmt.Println()

	fmt.Println(r.FindAllString("punch pinch peach", 2)) // only match 2 times
	fmt.Println()

	fmt.Println(r.ReplaceAllString("punch - pinch - peach", "<replaced>"))
	fmt.Println(string(r.ReplaceAllFunc([]byte("punch pinch peach"), bytes.ToUpper)))
	fmt.Println()

}
