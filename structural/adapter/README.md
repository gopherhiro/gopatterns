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
**解决接口不兼容问题。**
1. 封装有缺陷的接口设计。
2. 统一多个类的接口设计。
3. 替换依赖的外部系统。
4. 兼容老版本接口。
5. 适配不同格式的数据。

## 优点

## 缺点

## 总结：
**代理模式**：代理模式在不改变原始类接口的条件下，为原始类定义一个代理类，主要目的是控制访问，而非加强功能，这是它跟装饰器模式最大的不同。
**桥接模式**：桥接模式的目的是将接口部分和实现部分分离，从而让它们可以较为容易、也相对独立地加以改变。
**装饰器模式**：装饰者模式在不改变原始类接口的情况下，对原始类功能进行增强，并且支持多个装饰器的嵌套使用。
**适配器模式**：适配器模式是一种事后的补救策略。适配器提供跟原始类不同的接口，而代理模式、装饰器模式提供的都是跟原始类相同的接口。
