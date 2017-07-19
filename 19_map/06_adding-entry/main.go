package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Vinod": "Good morning!",
		"Kitto": "Boinorr!",
	}

	myGreeting["Vinod"] = "Howdy"

	fmt.Println(myGreeting)
}
