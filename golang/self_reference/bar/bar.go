package bar

import (
	"fmt"

	"../foo"
)

type Bar struct{}

func (b Bar) P() {
	fmt.Println("bar")
}

func CallFoo() {
	foo.Foo{}.P()
}
