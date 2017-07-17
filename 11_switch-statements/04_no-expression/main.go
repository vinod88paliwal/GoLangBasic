package main

import "fmt"

func main() {

	myFriendsName := "VKP"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Wassup my friend with name of length 2")
	case myFriendsName == "VP":
		fmt.Println("Wassup VP")
	case myFriendsName == "PD":
		fmt.Println("Wassup PD")
	case myFriendsName == "VP", myFriendsName == "KP":
		fmt.Println("Your name is either VP or KP")
	case myFriendsName == "Raj":
		fmt.Println("Wassup Raj")
	case myFriendsName == "Sushant":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("nothing matched; this is the default")
	}
}

/*
  expression not needed
  -- if no expression provided, go checks for the first case that evals to true
  -- makes the switch operate like if/if else/else
  cases can be expressions
*/
