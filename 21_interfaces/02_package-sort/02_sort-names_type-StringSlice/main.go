package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Vinod", "Sujit", "Sanjay", "Raj"}
	fmt.Println(s)
	//	sort.StringSlice(s).Sort()
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)

}

// https://golang.org/pkg/sort/#Sort
