package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Vinod!"
	greeting[2] = "Paliwal!"
	greeting = append(greeting, "Kittttooss")
	greeting = append(greeting, "pino'Pari")
	greeting = append(greeting, "Raju Agrawal")
	greeting = append(greeting, "gudday")

	fmt.Println(greeting[6])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))
}
