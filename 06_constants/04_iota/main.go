package main

import "fmt"

//Within a constant declaration, the predeclared identifier iota represents successive untyped
//integer constants and once its initialized then it would be keep incremented even further constant
// not initialed explicitly with iota.

const (
	a = iota // 0
	b        // 1
	c        // 2
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
