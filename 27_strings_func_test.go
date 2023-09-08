package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_strings_func(t *testing.T) {
	var p = fmt.Println
	p("Contains:  ", strings.Contains("test", "es"))        // contains sub string or not
	p("Count:     ", strings.Count("test test", "es"))      // count sub string's repeat times
	p("HasPrefix: ", strings.HasPrefix("test", "te"))       // has sub string as prefix
	p("HasSuffix: ", strings.HasSuffix("test", "st"))       // has sub string as suffix
	p("Index:     ", strings.Index("test", "e"))            // get sub string's first index
	p("Join:      ", strings.Join([]string{"a", "b"}, "-")) // return joined sub string between each string in slice
	p("Repeat:    ", strings.Repeat("a", 5))                // return sub string with repeate times
	p("Replace:   ", strings.Replace("foo", "o", "0", -1))  // replace old with new, -1 means replace all of old
	p("Replace:   ", strings.Replace("foo", "o", "0", 1))   // replace old with new, 1 means replace-location-index
	p("Split:     ", strings.Split("a-b-c-d-e", "-"))       // split string into a string-slice by sub string
	p("ToLower:   ", strings.ToLower("TEST"))               // trans into lower case
	p("ToUpper:   ", strings.ToUpper("test"))               // trans into upper case
	p()

	p("Len: ", len("hello世界")) // return length of strings's utf-8 encoded char-set
	p("Char:", "hello世界"[1])
	p("Char:", "hello世界"[10])
}
