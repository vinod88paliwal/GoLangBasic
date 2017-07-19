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
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "Vinod",
			Last:  "Paliwal",
			Age:   30,
		},
		LicenseToKill: true,
	}

	p2 := doubleZero{
		person: person{
			First: "Malik",
			Last:  "Paliwal",
			Age:   26,
		},
		LicenseToKill: false,
	}

	fmt.Println(p1.First, p1.Last, p1.Age, p1.LicenseToKill)
	fmt.Println(p2.First, p2.Last, p2.Age, p2.LicenseToKill)
}
