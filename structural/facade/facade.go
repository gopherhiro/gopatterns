package facade

import "fmt"

type T interface {
	yesterday()
	today()
	tomorrow()
}

type Day struct {
}

func (d *Day) yesterday() {
	fmt.Println("I am yesterday")
}

func (d *Day) today() {
	fmt.Println("I am today")
}

func (d *Day) tomorrow() {
	fmt.Println("I am tomorrow")
}

// 门面方法：可以进行时光穿梭
func (d *Day) timeTravel() {
	d.yesterday()
	d.today()
	d.tomorrow()
}
