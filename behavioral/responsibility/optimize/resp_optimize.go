package optimize

import "fmt"

/*
由于处理器类的 handle() 函数，不仅包含自己的业务逻辑，还包含对下一个处理器的调用。
一个不熟悉这种代码结构的同学，在添加新的处理器类的时候，
很有可能忘记在 handle() 函数中调用 successor.handle()，这就会导致代码出现 bug。
针对这个问题，我们对代码进行重构，利用模板模式，将调用 successor.handle() 的逻辑从具体的处理器类中剥离出来，放到基础处理者中。
这样具体的处理器类只需要实现自己的业务逻辑就可以了。
*/

// 职责链为：A -> B -> C

// Handler 处理者(Handler)声明了所有具体处理者的通用接口。
type Handler interface {
	doHandle() bool
	setSuccessor(h Handler)
	getSuccessor() Handler
}

// BaseHandler 基础处理者(Base Handler)是一个可选的类，你可以将所有处理者共用的样本代码放置在这里
// 也就是模板模式
type BaseHandler struct {
	handler Handler
}

func (bh *BaseHandler) handle() {
	var handled bool
	var successor Handler
	handled = bh.handler.doHandle()
	successor = bh.handler.getSuccessor()
	for !handled && successor != nil {
		handled = successor.doHandle()
		successor = successor.getSuccessor()
	}

	/***变体：所有的处理器都会处理请求***/
	/*	var successor Handler
		bh.handler.doHandle()
		successor = bh.handler.getSuccessor()
		for successor != nil {
			successor.doHandle()
			successor = successor.getSuccessor()
		}*/
}

// HandlerA ConcreteHandler
type HandlerA struct {
	successor Handler
}

func (a *HandlerA) doHandle() bool {
	handled := true
	fmt.Println("HandlerA handled")
	return handled
}

func (a *HandlerA) setSuccessor(s Handler) {
	a.successor = s
}

func (a *HandlerA) getSuccessor() Handler {
	return a.successor
}

// HandlerB ConcreteHandler
type HandlerB struct {
	successor Handler
}

func (b *HandlerB) doHandle() bool {
	handled := false
	fmt.Println("HandlerB handled")
	return handled
}

func (b *HandlerB) setSuccessor(s Handler) {
	b.successor = s
}

func (b *HandlerB) getSuccessor() Handler {
	return b.successor
}

// HandlerC ConcreteHandler
type HandlerC struct {
	successor Handler
}

func (c *HandlerC) doHandle() bool {
	handled := false
	fmt.Println("HandlerC handled")
	return handled
}

func (c *HandlerC) setSuccessor(s Handler) {
	c.successor = s
}

func (c *HandlerC) getSuccessor() Handler {
	return c.successor
}

/*******链表实现方式********/

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
		return
	}
	hc.tail.setSuccessor(h)
	hc.tail = h
}

func (hc *HandlerChain) handle() {
	if hc.head != nil {
		bh := &BaseHandler{hc.head}
		bh.handle()
	}
}

/*******数组实现方式********/
// 这个方式更简单，不用考虑successor

// AHandlerChain 处理器链（职责链）
type AHandlerChain struct {
	handlers []Handler
}

func (hc *AHandlerChain) addHandler(h Handler) {
	hc.handlers = append(hc.handlers, h)
}

func (hc *AHandlerChain) handle() {
	for _, handler := range hc.handlers {
		handled := handler.doHandle()
		if handled {
			fmt.Println("request handled")
			break
		}
	}
}
