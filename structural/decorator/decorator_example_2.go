package decorator

import "fmt"

// 代理模式与装饰器模式对比
// 在代理模式中，附加的是跟原始类型无关的功能。
// 在装饰器模式中，装饰器附加的是跟原始类型相关的增强功能。

type IA interface {
	doSth()
}

type A struct {
}

func (a *A) doSth() {
	fmt.Println("I am A")
}

// AProxy 代理模式的代码结构
type AProxy struct {
	a A
}

func (ap *AProxy) doSth() {
	// 代理逻辑代码
	fmt.Println("I am proxy before")

	ap.a.doSth()

	// 代理逻辑代码
	fmt.Println("I am proxy after")
}

// ADecorator 装饰器模式的代码结构
type ADecorator struct {
	a A
}

func (ap *ADecorator) doSth() {
	// 功能增强代码
	fmt.Println("I am decorator before")

	ap.a.doSth()

	// 功能增强代码
	fmt.Println("I am decorator after")
}
