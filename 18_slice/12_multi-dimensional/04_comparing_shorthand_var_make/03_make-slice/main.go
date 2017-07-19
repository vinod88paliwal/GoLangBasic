package main

import (
	"fmt"
)

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	student[0] = "Vinod"
	// student = append(student, "Vinod")
	fmt.Println(student)
	fmt.Println(students)
}
