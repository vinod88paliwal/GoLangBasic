package main

import "fmt"

func main() {
	switch "VP" {
	case "Malik":
		fmt.Println("Malik Kumar")
	case "Sujit":
		fmt.Println("Sujit Barik")
	case "Raj":
		fmt.Println("Raj Agrawal")
	default:
		fmt.Println("Have you no friends?")
	}
}

/*
  no default fallthrough
  fallthrough is optional
  -- you can specify fallthrough by explicitly stating it
  -- break isn't needed like in other languages
*/
