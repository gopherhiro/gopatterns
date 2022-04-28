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
