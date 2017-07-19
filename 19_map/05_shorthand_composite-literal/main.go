package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Vinod": "Good morning!",
		"Kitto": "Bijnor!",
	}

	fmt.Println(myGreeting["Vinod"])
}
