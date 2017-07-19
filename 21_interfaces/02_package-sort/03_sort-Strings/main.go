package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Vinod", "Sujit", "Sanjay", "Raj"}
	sort.Strings(s)
	fmt.Println(s)
}
