package main

import (
	"fmt"
	"gopatterns/creational/builder"
)

func main() {
	normalBuilder := builder.Builder("normal")
	normalDirector := builder.NewDirector(normalBuilder)
	r := normalDirector.BuildHouse()
	fmt.Println(r)

	iglooBuilder := builder.Builder("igloo")
	iglooDirector := builder.NewDirector(iglooBuilder)
	r = iglooDirector.BuildHouse()
	fmt.Println(r)
}
