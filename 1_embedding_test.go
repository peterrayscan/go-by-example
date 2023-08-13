package main

import (
	"fmt"
	"testing"
)

type base struct {
	num int
}

func (b *base) desc() {
	fmt.Printf("base.num: %d\n", b.num)
}

type container struct {
	str string
	base
}

func Test_embedding(t *testing.T) {
	ctnr := container{
		base: base{num: 1024}, // can not use: 'num: 1024,'
		str:  "this is string",
	}

	fmt.Printf("container's base.num: %d, str: %s\n", ctnr.num, ctnr.str)
	fmt.Printf("container's base.num: %d, str: %s\n", ctnr.base.num, ctnr.str)

	ctnr.desc()
	ctnr.base.desc()

	type describer interface {
		desc()
	}

	var sth describer = &ctnr
	sth.desc()
}
