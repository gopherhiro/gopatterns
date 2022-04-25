package main

import (
	"fmt"
	"gopatterns/creational/factory"
	"gopatterns/pkg/static"
)

func main() {
	var f factory.CourseFactory
	f = factory.NewInstance(static.Chinese)
	c := f.New()
	fmt.Println(c.GetName())
}
