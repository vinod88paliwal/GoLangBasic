package main

import "fmt"

func main() {

	name := "Vinod"
	fmt.Println(name) // Vinod

	changeMe(name)

	fmt.Println(name) // Vinod
}

func changeMe(z string) {
	fmt.Println(z) // Vinod
	z = "Paliwal"
	fmt.Println(z) // Paliwal
}
