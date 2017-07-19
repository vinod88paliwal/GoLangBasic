package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := &person{"Vinod", 30}
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)

	p1.name = "Hello"
	p1.age = 40

	fmt.Println(p1.name)
	fmt.Println(p1.age)
}
