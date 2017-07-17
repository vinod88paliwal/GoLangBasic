package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Println("Len : ", len(sf))

	var total float64
	for _, v := range sf {
		total += v
	}
	fmt.Printf("%T \n", sf)
	return total / float64(len(sf))
}
