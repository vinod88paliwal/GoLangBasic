package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 person
	fmt.Println("First1 :", p1.First)
	fmt.Println("Last1 :", p1.Last)
	fmt.Println("Age1 :", p1.Age)

	bs := []byte(`{"First":"Vinod", "Last":"Paliwal", "Age":30}`)
	json.Unmarshal(bs, &p1)

	fmt.Println("--------------")
	fmt.Println("First2 :", p1.First)
	fmt.Println("Last2 :", p1.Last)
	fmt.Println("Age2 :", p1.Age)
	fmt.Printf("%T \n", p1)
}
