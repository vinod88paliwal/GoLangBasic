package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bijnor!",
		2: "Bijnor dias!",
		3: "Bognsh!",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, 7)
	fmt.Println(myGreeting)
}
