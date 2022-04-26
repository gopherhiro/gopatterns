# 桥接模式
桥接模式是一种结构型设计模式， 可将一个大类或一系列紧密相关的类拆分为抽象和实现两个独立的层次结构， 从而能在开发时分别使用。

![image](https://user-images.githubusercontent.com/65383410/165283963-25afef4c-4ae1-4d5a-ae22-e5da615a0daa.png)

## 问题
假如你有一个几何 形状Shape类， 从它能扩展出两个子类： ​ 圆形Circle和 方形Square 。 你希望对这样的类层次结构进行扩展以使其包含颜色， 所以你打算创建名为 红色Red和 蓝色Blue的形状子类。 但是， 由于你已有两个子类， 所以总共需要创建四个类才能覆盖所有组合， 例如 蓝色圆形Blue­Circle和 红色方形Red­Square 。


所有组合类的数量将以几何级数增长。

![image](https://user-images.githubusercontent.com/65383410/165284572-1ae4fcb6-b649-4059-a5d5-ff983b8b6820.png)

在层次结构中新增形状和颜色将导致代码复杂程度指数增长。 例如添加三角形状， 你需要新增两个子类， 也就是每种颜色一个； 此后新增一种新颜色需要新增三个子类， 即每种形状一个。 如此以往， 情况会越来越糟糕。

## 解决方案
问题的根本原因是我们试图在两个独立的维度——形状与颜色——上扩展形状类。

桥接模式通过将继承改为组合的方式来解决这个问题。 具体来说， 就是抽取其中一个维度并使之成为独立的类层次， 这样就可以在初始类中引用这个新层次的对象， 从而使得一个类不必拥有所有的状态和行为。

将一个类层次转化为多个相关的类层次， 避免单个类层次的失控。

![image](https://user-images.githubusercontent.com/65383410/165284833-d6f07173-d950-4673-a56f-24f78a8462ab.png)

根据该方法， 我们可以将颜色相关的代码抽取到拥有 红色和 蓝色两个子类的颜色类中， 然后在 形状类中添加一个指向某一颜色对象的引用成员变量。 现在， 形状类可以将所有与颜色相关的工作委派给连入的颜色对象。 这样的引用就成为了 形状和 颜色之间的桥梁。 此后， 新增颜色将不再需要修改形状的类层次， 反之亦然。

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
    

