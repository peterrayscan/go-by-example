package main

import (
	"errors"
	"fmt"
	"testing"
)

func build(i int) error {
	// do sth
	return errors.New("build fail with error: xxx")
}

type builder struct {
	param int
}

func (b *builder) Error() string {
	return fmt.Sprintf("build(param: %d) fail with err: xxx", b.param)
}

func (b *builder) String() string {
	return "string method"
}

func (b *builder) build() error {
	// do sth

	// cause the builder has implemented the Error() method
	// so builder could return as an error
	return b
}

func Test_error_handling(t *testing.T) {
	fmt.Println(build(1))

	bdr := builder{param: 2}
	err := bdr.build()
	fmt.Println(err)

	// error could also assert into it's original struct
	if ae, ok := err.(*builder); ok {
		fmt.Println("\noriginal struct's param:", ae.param)
		fmt.Printf("err after assert: \n%+v\n", ae)
	}
}
