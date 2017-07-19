package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("Person Name & Age %s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []person based on
// the Age field.
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

//func (a ByAge) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	people := []person{
		{"Vinod", 31},
		{"Sujit", 42},
		{"Raj", 17},
		{"Sanjay", 26},
	}

	fmt.Println(people[0])
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

}

// https://golang.org/pkg/sort/#Sort
// https://golang.org/pkg/sort/#Interface

// String() string
// https://golang.org/doc/effective_go.html#printing
