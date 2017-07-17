package main

import "fmt"

type customer struct {
	name string
	age  int
}

func main() {
	c1 := customer{"Vinod", 30}
	fmt.Println(&c1.name)                // 0x8201e4120
	fmt.Println("c1.name ==> ", c1.name) //Vinod

	changeMe(&c1)

	fmt.Println(c1)       // {Paliwal 44}
	fmt.Println(&c1.name) // 0x8201e4120
}

func changeMe(z *customer) {
	fmt.Println(z)       // &{Vinod 30}
	fmt.Println(&z.name) // 0x8201e4120
	z.name = "Paliwal"
	fmt.Println(z)       // &{Paliwal 30}
	fmt.Println(&z.name) // 0x8201e4120

}
