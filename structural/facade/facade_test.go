package facade

import (
	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {
	d := &Day{}
	d.today()

	fmt.Println("")
	d.timeTravel()
}
