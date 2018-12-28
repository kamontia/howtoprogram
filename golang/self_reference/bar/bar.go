package bar

import (
	"fmt"

	"../common"
)

type Bar struct{}

func (b Bar) P() {
	fmt.Println("bar")
}

func CallFoo(p common.Printer) {
	p.P()
}
