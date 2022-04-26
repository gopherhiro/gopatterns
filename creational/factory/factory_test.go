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
	var f CourseFactory
	var c Course

	f = NewFactory(static.Chinese)
	c = f.NewCourse()
	fmt.Println(c.GetName())
	// Chinese

	f = NewFactory(static.English)
	c = f.NewCourse()
	fmt.Println(c.GetName())
	// English
}
