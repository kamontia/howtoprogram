package foo

import (
	"fmt"

	"../bar"
)

type Foo struct{}

func (f Foo) P() {
	fmt.Println("foo")
}

func CallBar() {
	bar.Bar{}.P()
}
