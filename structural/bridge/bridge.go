package bridge

import "fmt"

// 抽象:computer
type computer interface {
	print()
	setPrinter(printer)
}

// 精确抽象:mac
type mac struct {
	printer printer
}

func (m *mac) print() {
	fmt.Println("Print request for mac")
	m.printer.printFile()
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}

// 精确抽象:windows
type windows struct {
	printer printer
}

func (w *windows) print() {
	fmt.Println("Print request for windows")
	w.printer.printFile()
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}

// 实现:printer
type printer interface {
	printFile()
}

// 具体实现:epson
type epson struct {
}

func (p *epson) printFile() {
	fmt.Println("Printing by a EPSON Printer")
}

// 具体实现:hp
type hp struct {
}

func (p *hp) printFile() {
	fmt.Println("Printing by a HP Printer")
}
