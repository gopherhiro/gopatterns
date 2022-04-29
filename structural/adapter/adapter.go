package adapter

import "fmt"

/**
ITarget 表示要转化成的接口定义。
A 是一组不兼容 ITarget 接口定义的接口，Adaptor 将 A 转化成一组符合 ITarget 接口定义的接口。
*/

type iTarget interface {
	f1()
	f2()
	f3()
}

// A ...
type A struct {
}

func (a *A) fa() {
	fmt.Println("I am A fa")
}

func (a *A) fb() {
	fmt.Println("I am A fb")
}

func (a *A) fc() {
	fmt.Println("I am A fc")
}

// Adaptor ...
type Adaptor struct {
	a *A
}

func (ad *Adaptor) f1() {
	ad.a.fa()
}

func (ad *Adaptor) f2() {
	// 重新实现
	ad.a.fb()
	fmt.Println("I reimplement sth")
}

func (ad *Adaptor) f3() {
	ad.a.fc()
}
