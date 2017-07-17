package main

import "fmt"

func main() {
	switch "Vinod" {
	case "VP", "Vinod":
		fmt.Println("Wassup VP, or, err, Vinod")
	case "Mery", "Medhi":
		fmt.Println("Both of your names start with M")
	case "Kumar", "Sushant":
		fmt.Println("Wassup Kumar / Sushant")
	}
}
