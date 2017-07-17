package main

import "fmt"

func main() {
	greet("Vinod", "Paliwal")
}

func greet(fname, lname string) {
	fmt.Println(fname, lname)
}
