# 桥接模式
桥接模式是一种结构型设计模式， 可将一个大类或一系列紧密相关的类拆分为抽象和实现两个独立的层次结构， 从而能在开发时分别使用。

## 实现
### 具体步骤
1.明确类中独立的维度。 独立的概念可能是： 抽象/平台， 域/基础设施， 前端/后端或接口/实现。

2.了解客户端的业务需求， 并在抽象基类中定义它们。

3.确定在所有平台上都可执行的业务。 并在通用实现接口中声明抽象部分所需的业务。

4.为你域内的所有平台创建实现类， 但需确保它们遵循实现部分的接口。

5.在抽象类中添加指向实现类型的引用成员变量。 抽象部分会将大部分工作委派给该成员变量所指向的实现对象。

6.如果你的高层逻辑有多个变体， 则可通过扩展抽象基类为每个变体创建一个精确抽象。

7.客户端代码必须将实现对象传递给抽象部分的构造函数才能使其能够相互关联。 此后， 客户端只需与抽象对象进行交互， 无需和实现对象打交道。
### 概念示例
假设你有两台电脑： 一台 Mac 和一台 Windows。 还有两台打印机： 爱普生和惠普。 这两台电脑和打印机可能会任意组合使用。 客户端不应去担心如何将打印机连接至计算机的细节问题。

如果引入新的打印机， 我们也不会希望代码量成倍增长。 所以， 我们创建了两个层次结构， 而不是 2x2 组合的四个结构体：
- 抽象层： 代表计算机
- 实施层： 代表打印机

这两个层次可通过桥接进行沟通， 其中抽象层 （计算机） 包含对于实施层 （打印机） 的引用。 抽象层和实施层均可独立开发， 不会相互影响。

```go
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

```

## 用法

```go
package bridge

import "testing"

func TestBridge(t *testing.T) {
	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	macComputer := &mac{}
	winComputer := &windows{}

	// mac
	macComputer.setPrinter(hpPrinter)
	macComputer.print()

	macComputer.setPrinter(epsonPrinter)
	macComputer.print()

	// windows
	winComputer.setPrinter(hpPrinter)
	winComputer.print()

	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
}
```

## 应用场景
- 如果你想要拆分或重组一个具有多重功能的庞杂类 （例如能与多个数据库服务器进行交互的类）， 可以使用桥接模式。
    - 类的代码行数越多， 弄清其运作方式就越困难， 对其进行修改所花费的时间就越长。 一个功能上的变化可能需要在整个类范围内进行修改， 而且常常会产生错误， 甚至还会有一些严重的副作用。
    - 桥接模式可以将庞杂类拆分为几个类层次结构。 此后， 你可以修改任意一个类层次结构而不会影响到其他类层次结构。 这种方法可以简化代码的维护工作， 并将修改已有代码的风险降到最低。
- 如果你希望在几个独立维度上扩展一个类， 可使用桥接模式。
- 如果你需要在运行时切换不同实现方法， 可使用桥接模式。

## 优点
- 你可以创建与平台无关的类和程序。
- 客户端代码仅与高层抽象部分进行互动，不会接触到平台的详细信息。
- 开闭原则。你可以新增抽象部分和实现部分，且它们之间不会相互影响。
- 单一职责原则。抽象部分专注于处理高层逻辑，实现部分处理平台细节。
## 缺点
- 对高内聚的类使用该模式可能会让代码更加复杂。
    

