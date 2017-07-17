package main

import "fmt"

func main() {

	name := "Vinod"
	fmt.Println(&name) // 0x82023c080

	changeMe(&name)

	fmt.Println(&name) //0x82023c080
	fmt.Println(name)  //Paliwal
}

func changeMe(z *string) {
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // Vinod
	*z = "Paliwal"
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // Paliwal
}
