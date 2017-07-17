package main

import "fmt"

func main() {
	m := make([]string, 2, 25)
	fmt.Println(m) // [ ]
	changeMe(m)
	fmt.Println(m) // [Vinod Paliwal]
}

func changeMe(z []string) {
	z[0] = "Vinod"
	z[1] = "Paliwal"
	fmt.Println(z[0]) // [Vinod]
}
