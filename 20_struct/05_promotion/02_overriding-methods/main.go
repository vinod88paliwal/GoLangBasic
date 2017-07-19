package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func (p person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

func (dz doubleZero) Greeting() {
	fmt.Println("Miss Paliwal, so good to see you.")
}

func main() {
	p1 := person{
		Name: "Vinod Paliwal",
		Age:  30,
	}

	p2 := doubleZero{
		person: person{
			Name: "Malik Paliwal",
			Age:  23,
		},
		LicenseToKill: true,
	}
	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting()
	fmt.Println("Name: ", p1.Name)
}
