package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	var target iTarget

	a := &A{}
	target = &Adaptor{a: a}
	target.f1()
	target.f2()
	target.f3()
}
