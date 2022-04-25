package factory

import (
	"fmt"
	"gopatterns/pkg/static"
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	var c Course

	c = New(static.Chinese)
	fmt.Println(c.GetName())
	// Chinese

	c = NewFromCache(static.English)
	fmt.Println(c.GetName())
	// English
}

func TestFactoryMethod(t *testing.T) {
	var f CourseFactory
	var c Course

	f = NewInstance(static.Chinese)
	c = f.New()
	fmt.Println(c.GetName())
	// Chinese

	f = NewInstance(static.English)
	c = f.New()
	fmt.Println(c.GetName())
	// English
}
