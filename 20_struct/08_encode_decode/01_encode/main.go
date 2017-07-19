package main

import (
	"encoding/json"
	"os"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := person{"Vinod", "Paliwal", 30, 999}
	json.NewEncoder(os.Stdout).Encode(p1)
}
