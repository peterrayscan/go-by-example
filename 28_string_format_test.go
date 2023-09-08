package main

import (
	"fmt"
	"os"
	"testing"
)

type Point struct {
	x, y int
}

func Test_string_format(t *testing.T) {
	var p = fmt.Printf
	point := Point{1, 2}

	p("struct1: %v\n", point)
	p("struct2: %+v\n", point)
	p("struct3: %#v\n", point)

	p("type: %T\n", point)
	p("bool: %t\n", true)
	p("int: %d\n", 123)
	p("bin: %b\n", 14)
	p("char: %c\n", 33)

	p("hex: %x\n", 456)

	p("float1: %f\n", 78.9)
	p("float2: %e\n", 123400000.0)
	p("float3: %E\n", 123400000.0)

	p("str1: %s\n", "\"string\"")
	p("str2: %q\n", "\"string\"")
	p("str3: %x\n", "hex this")

	p("pointer: %p\n", &point)

	p("width1: |%6d|%6d|\n", 12, 345)
	p("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	p("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	p("width4: |%6s|%6s|\n", "foo", "b")
	p("width5: |%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
