# 职责链模式

职责链是一种行为设计模式，允许你将请求沿着处理者链进行发送。收到请求后，每个处理者均可对请求进行处理，或将其传递给链上的下个处理者。

Chain of Responsibility is a behavioral design pattern that lets you pass requests along a chain of handlers. Upon
receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.、

![image](https://user-images.githubusercontent.com/65383410/165558246-2d36cb45-e609-4129-85f5-7964d2532122.png)

Handlers are lined up one by one, forming a chain.

![image](https://user-images.githubusercontent.com/65383410/165558363-9af154e4-9ad2-4410-93d7-d5b0cbfaef56.png)

在职责链模式中，多个处理器依次处理同一个请求。一个请求先经过 A 处理器处理，然后再把请求传递给 B 处理器，B 处理器处理完后再传递给 C 处理器，以此类推，形成一个链条。链条上的每个处理器各自承担各自的处理职责，所以叫作职责链模式。

## 演变

- 请求沿着职责链传递，直到链上的某个接收对象能够处理它为止。（原生）
- 请求沿着职责链传递，链上的每一个接收对象都会处理它。不存在中途停止的情况。（变体）

## 作用

- 复用
- 拓展

在实际的项目开发中比较常用，特别是框架开发中，我们可以利用它们来提供框架的扩展点，能够让框架的使用者在不修改框架源码的情况下，基于扩展点定制化框架的功能。

## 实现

职责链模式有两种常用的实现。

- 一种是使用链表来存储处理器。
- 一种是使用数组来存储处理器，数组实现方式更加简单。

```go
package responsibility

import "fmt"

// 职责链为：A -> B -> C

// Handler 处理者(Handler)声明了所有具体处理者的通用接口。
type Handler interface {
	handle()
	setSuccessor(s Handler)
}

// HandlerA ConcreteHandler
type HandlerA struct {
	successor Handler
}

func (a *HandlerA) handle() {
	handled := false
	fmt.Println("HandlerA handled")
	if !handled && a.successor != nil {
		a.successor.handle()
	}
}

func (a *HandlerA) setSuccessor(s Handler) {
	a.successor = s
}

// HandlerB ConcreteHandler
type HandlerB struct {
	successor Handler
}

func (b *HandlerB) setSuccessor(s Handler) {
	b.successor = s
}

func (b *HandlerB) handle() {
	handled := false
	fmt.Println("HandlerB handled")
	if !handled && b.successor != nil {
		b.successor.handle()
	}
}

// HandlerC ConcreteHandler
// 拓展一个新的处理器
type HandlerC struct {
	successor Handler
}

func (c *HandlerC) setSuccessor(s Handler) {
	c.successor = s
}

func (c *HandlerC) handle() {
	handled := false
	fmt.Println("HandlerC handled")
	if !handled && c.successor != nil {
		c.successor.handle()
	}
}

// HandlerChain 处理器链（职责链）
// 就是一个记录了链头、链尾的链表。
type HandlerChain struct {
	head Handler // 记录链头是为了方便启动处理
	tail Handler // 记录链尾是为了方便添加处理器
}

// 尾插法建立单链表
func (hc *HandlerChain) addHandler(h Handler) {
	// 清空 h 后续节点
	h.setSuccessor(nil)
	if hc.head == nil {
		hc.head = h
		hc.tail = h
	}
	hc.tail.setSuccessor(h)
	hc.tail = h
}

func (hc *HandlerChain) handle() {
	if hc.head != nil {
		hc.head.handle()
	}
}

```

## 用法

```go
package responsibility

import (
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {
	chain := &HandlerChain{}
	chain.addHandler(&HandlerA{})
	chain.addHandler(&HandlerB{})
	chain.addHandler(&HandlerC{})
	chain.handle()
}

// Output:
// HandlerA handled
// HandlerB handled
// HandlerC handled
```

## 总结

在职责链模式中，多个处理器依次处理同一个请求。一个请求先经过 A 处理器处理，然后再把请求传递给 B 处理器，B 处理器处理完后再传递给 C 处理器，以此类推，形成一个链条。链条上的每个处理器各自承担各自的处理职责，所以叫作职责链模式。

在 GoF 的定义中，一旦某个处理器能处理这个请求，就不会继续将请求传递给后续的处理器了。当然，在实际的开发中，也存在对这个模式的变体，那就是请求不会中途终止传递，而是会被所有的处理器都处理一遍。

职责链模式有两种常用的实现。一种是使用链表来存储处理器，另一种是使用数组来存储处理器，后面一种实现方式更加简单。

## 应用场景

- 当程序需要使用不同方式处理不同种类请求，而且请求类型和顺序预先未知时，可以使用该模式。
- 当必须按顺序执行多个处理者时，可以使用该模式。
- 如果所需处理者及其顺序必须在运行时进行改变，可以使用该模式。

## 优点

- 应对代码的复杂性
- 提高代码的拓展性
- 你可以控制请求处理的顺序。
- 开闭原则。你可以在不更改现有代码的情况下在程序中新增处理者。

## 缺点

- 部分请求可能未被处理。
