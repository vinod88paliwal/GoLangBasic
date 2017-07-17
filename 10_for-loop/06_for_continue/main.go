package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			fmt.Println("Continue : ", i)
			continue
		}
		if i >= 50 {
			fmt.Println("Break : ", i)
			break
		}
	}
}
