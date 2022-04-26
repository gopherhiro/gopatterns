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
