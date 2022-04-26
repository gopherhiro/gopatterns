package bridge

import "fmt"

// 抽象:computer
// 提供高层控制逻辑， 依赖于完成底层实际工作的实现对象。
// 如果你的高层逻辑有多个变体， 则可通过扩展抽象基类为每个变体创建一个精确抽象。
type computer interface {
	print()
	setPrinter(printer)
}

// 精确抽象:mac
// 提供控制逻辑的变体
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
// 提供控制逻辑的变体
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
// 为所有具体实现声明通用接口。
type printer interface {
	printFile()
}

// 具体实现:epson
// 包含特定于平台的代码。
type epson struct {
}

func (p *epson) printFile() {
	fmt.Println("Printing by a EPSON Printer")
}

// 具体实现:hp
// 包含特定于平台的代码。
type hp struct {
}

func (p *hp) printFile() {
	fmt.Println("Printing by a HP Printer")
}
