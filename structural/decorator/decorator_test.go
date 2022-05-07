package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	pizza := &vegetable{}
	fmt.Println("price of vegetable Pizza:", pizza.price())

	withTomato := &withTomato{
		pizza: pizza,
	}
	fmt.Println("price of vegetable Pizza with tomato:", withTomato.price())

	withCheese := withCheese{
		pizza: pizza,
	}
	fmt.Println("price of vegetable Pizza with tomato and cheese:", withCheese.price())
}

func TestExample1Decorator(t *testing.T) {
	var c Component = &ConcreteComponent{}
	c = WarpAddDecorator(c, 10)
	c = WarpMulDecorator(c, 8)
	res := c.Calc()

	fmt.Printf("res %d\n", res)
	// Output:
	// res 80
}

func TestExample2Decorator(t *testing.T) {
	var ia IA

	// proxy pattern
	ia = &AProxy{}
	ia.doSth()

	// decorator pattern
	ia = &ADecorator{}
	ia.doSth()

	// Output:
	// I am proxy before
	// I am A
	// I am proxy after

	// I am decorator before
	// I am A
	// I am decorator after
}
