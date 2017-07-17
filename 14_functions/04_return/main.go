package main

import "fmt"

func main() {
	fmt.Println(greet("Vinod ", "Paliwal"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
