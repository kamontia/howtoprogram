package main

import (
	"./bar"
	"./foo"
)

func main() {
	foo.CallBar(bar.Bar{})
	bar.CallFoo(foo.Foo{})
}
