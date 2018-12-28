package foo

import (
	"fmt"

	"../common"
)

type Foo struct{}

func (f Foo) P() {
	fmt.Println("foo")
}

func CallBar(p common.Printer) {
	p.P()
}
