package main

import "fmt"

func makeGreeting() func() string {
	return func() string {
		return "Hello world!"
	}
}

func main() {
	greet := makeGreeting()
	fmt.Println(greet())
}
