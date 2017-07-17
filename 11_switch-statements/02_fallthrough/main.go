package main

import "fmt"

func main() {
	switch "VP" {
	case "Vinod":
		fmt.Println("Wassup Vinod")
	case "Peddy":
		fmt.Println("Wassup Peddy")
	case "VP":
		fmt.Println("Wassup VP")
		fallthrough
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case "KP":
		fmt.Println("Wassup KP")
	case "Sushant":
		fmt.Println("Wassup Sushant")
	}
}
