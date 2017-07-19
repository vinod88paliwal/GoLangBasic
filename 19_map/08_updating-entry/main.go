package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Vinod": "Good morning!",
		"Kitto": "Boinorr!",
	}

	myGreeting["Vinod"] = "Howz"
	fmt.Println(myGreeting)
	myGreeting["Kitto"] = "Guddy"
	fmt.Println(myGreeting)
}
