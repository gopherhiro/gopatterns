package builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	normalBuilder := Builder("normal")
	normalDirector := NewDirector(normalBuilder)
	r := normalDirector.BuildHouse()
	fmt.Println(r)

	iglooBuilder := Builder("igloo")
	iglooDirector := NewDirector(iglooBuilder)
	r = iglooDirector.BuildHouse()
	fmt.Println(r)
}
