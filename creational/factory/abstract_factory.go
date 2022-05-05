package factory

import (
	"fmt"
	"gopatterns/pkg/static"
)

// 抽象工厂接口
type iSportsFactory interface {
	MakeShoe() iShoe
	MakeShirt() iShirt
}

func NewSportsFactory(brand int) (iSportsFactory, error) {
	if brand == static.Nike {
		return &adidas{}, nil
	}

	if brand == static.Adidas {
		return &nike{}, nil
	}
	return nil, fmt.Errorf("wrong brand type passed")
}

// 具体工厂:Nike
type nike struct {
}

func (n *nike) MakeShoe() iShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike shoe",
			size: 20,
		},
	}
}

func (n *nike) MakeShirt() iShirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike shirt",
			size: 21,
		},
	}
}

// 具体工厂:Adidas
type adidas struct {
}

func (a *adidas) MakeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas shoe",
			size: 30,
		},
	}
}

func (a *adidas) MakeShirt() iShirt {
	return &adidasShirt{
		shirt: shirt{
			logo: "adidas shirt",
			size: 31,
		},
	}
}

// 抽象产品:shoe
type iShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type shoe struct {
	logo string
	size int
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *shoe) getLogo() string {
	return s.logo
}

func (s *shoe) setSize(size int) {
	s.size = size
}

func (s *shoe) getSize() int {
	return s.size
}

// 抽象产品:shirt
type iShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type shirt struct {
	logo string
	size int
}

func (s *shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *shirt) getLogo() string {
	return s.logo
}

func (s *shirt) setSize(size int) {
	s.size = size
}

func (s *shirt) getSize() int {
	return s.size
}

// 具体产品：Nike shoe
type nikeShoe struct {
	shoe
}

// 具体产品：Nike shirt
type nikeShirt struct {
	shirt
}

// 具体产品：Adidas shoe
type adidasShoe struct {
	shoe
}

// 具体产品：Adidas shirt
type adidasShirt struct {
	shirt
}

func PrintShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func PrintShirtDetails(s iShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
