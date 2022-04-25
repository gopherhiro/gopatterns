package singleton

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	s := New()

	s["this"] = "that"

	s2 := New()

	fmt.Println("This is ", s2["this"])
	// This is that
}
