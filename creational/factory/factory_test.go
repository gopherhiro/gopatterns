package factory

import (
	"fmt"
	"gopatterns/pkg/static"
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	var c Course

	c = NewCourse(static.Chinese)
	fmt.Println(c.GetName())
	// Chinese

	c = NewCourse(static.English)
	fmt.Println(c.GetName())
	// English
}

func TestFactoryMethod(t *testing.T) {
	chineseFactory := NewFactory(static.Chinese)
	c := chineseFactory.NewCourse()
	fmt.Println(c.GetName())
	// Chinese

	englishFactory := NewFactory(static.English)
	e := englishFactory.NewCourse()
	fmt.Println(e.GetName())
	// English
}

func TestAbstractFactory(t *testing.T) {
	nikeFactory, _ := NewSportsFactory(static.Nike)
	nikeShoe := nikeFactory.MakeShoe()
	PrintShoeDetails(nikeShoe)
	nikeShirt := nikeFactory.MakeShirt()
	PrintShirtDetails(nikeShirt)

	adidasFactory, _ := NewSportsFactory(static.Adidas)
	adidasShoe := adidasFactory.MakeShoe()
	PrintShoeDetails(adidasShoe)
	adidasShirt := adidasFactory.MakeShirt()
	PrintShirtDetails(adidasShirt)
}
