package main

import (
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "Vinod",
			Last:  "Paliwal",
			Age:   30,
		},
		First:         "Nine Nine Nine",
		LicenseToKill: true,
	}

	p2 := doubleZero{
		person: person{
			First: "Malik",
			Last:  "Paliwal",
			Age:   26,
		},
		First:         "If looks could change",
		LicenseToKill: false,
	}

	// fields and methods of the inner-type are promoted to the outer-type
	fmt.Println(p1.First, p1.person.First)
	fmt.Println(p2.First, p2.person.First)
}
