# 适配器模式

Convert the interface of a class into another interface clients expect. Adapter lets classes work together that couldn't
otherwise because of incompatible interfaces.

适配器模式的英文翻译是 Adapter Design Pattern。

顾名思义，这个模式就是用来做适配的，它将不兼容的接口转换为可兼容的接口。

适配器模式是一种结构型设计模式， 它能使接口不兼容的对象能够相互合作，可担任两个对象间的转换器。

![image](https://user-images.githubusercontent.com/65383410/165320220-417d50a8-dc92-43d4-878b-6eb333deaf1d.png)

#### 举例

![image](https://user-images.githubusercontent.com/65383410/165323721-5050814e-e14d-4b9f-9f6b-b4341f47d8ec.png)

你可以创建一个适配器。 这是一个特殊的对象， 能够转换对象接口， 使其能与其他对象进行交互。

![image](https://user-images.githubusercontent.com/65383410/165323740-b4a6236c-2684-410f-8c51-110f8ad567c1.png)

## 实现

```go
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

```

## 用法

```go
package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	var target iTarget

	// 通过适配器 Adapter，A 转化成实现 iTarget 接口定义的类型
	a := &A{}
	target = &Adaptor{a: a}
	target.f1()
	target.f2()
	target.f3()
}

// Output:

// I am A fa

// I am A fb
// I reimplement sth

// I am A fc
```

## 应用场景

## 优点

## 缺点
