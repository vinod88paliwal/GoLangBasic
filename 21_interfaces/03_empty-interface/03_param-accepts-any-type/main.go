package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func specs(a interface{}) {
	fmt.Println(a)
}

func main() {
	fido := dog{animal{"Tiger"}, true}
	fifi := cat{animal{"Cow"}, true}
	specs(fido)
	specs(fifi)
}
