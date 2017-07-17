package main

import "fmt"

//Within a constant declaration, the predeclared identifier iota represents successive untyped
//integer constants. It is reset to 0 whenever the reserved word const appears in the source and
// increments after each ConstSpec. It can be used to construct a set of related constants:
const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
	d = iota
)

const (
	e = iota // 0
	f = iota // 1
	g = iota // 2
	h = iota
)

func main() {
	fmt.Println("A: ", a)
	fmt.Println("B: ", b)
	fmt.Println("C: ", c)
	fmt.Println("D: ", d)
	fmt.Println("Constants ********** ")
	fmt.Println("E: ", e)
	fmt.Println("F: ", f)
	fmt.Println("G: ", g)
	fmt.Println("H: ", h)

}
