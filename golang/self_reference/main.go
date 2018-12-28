package main

import (
	"./bar"
	"./foo"
)

func main() {
	bar.Bar{}.P()
	foo.Foo{}.P()
}
