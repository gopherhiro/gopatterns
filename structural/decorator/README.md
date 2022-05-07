# 装饰器模式

Attach additional responsibilities to an object dynamically. Decorators provide a flexible alternative to subclassing
for extending functionality.

装饰模式是一种结构型设计模式， 允许你通过将对象放入包含行为的特殊封装对象中来为原对象绑定新的行为。

#### 真实世界类比

穿上多件衣服将获得组合性的效果。类似一种打包嵌套，可以不断的给对象叠加新的行为。

![image](https://user-images.githubusercontent.com/65383410/165287892-eaf53215-3545-4e1d-87d4-4f40518d6826.png)

穿衣服是使用装饰的一个例子。 觉得冷时， 你可以穿一件毛衣。 如果穿毛衣还觉得冷， 你可以再套上一件夹克。 如果遇到下雨， 你还可以再穿一件雨衣。 所有这些衣物都 “扩展” 了你的基本行为， 但它们并不是你的一部分，
如果你不再需要某件衣物， 可以方便地随时脱掉。

## 装饰器的特点

- 装饰器类和原始类继承同样的父类。这样就可以对原始类“嵌套”多个装饰器类。
- 装饰器类是对原始类功能的增强，这也是装饰器模式应用场景的一个重要特点。

## 实现

```go
package decorator

// 基本组件和可选层次的通用方法。
type Pizza interface {
	price() int
}

// 创建一个具体组件类， 并定义其基础行为。
type vegetable struct {
	p int
}

func (v *vegetable) price() int {
	v.p = 15
	return v.p
}

// 具体装饰
type withTomato struct {
	pizza Pizza
}

func (t *withTomato) price() int {
	return t.pizza.price() + 7
}

// 具体装饰
type withCheese struct {
	pizza Pizza
}

func (c *withCheese) price() int {
	return c.pizza.price() + 10
}

```

## 用法

```go
package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	pizza := &vegetable{}
	fmt.Println("price of vegetable Pizza:", pizza.price())

	withTomato := &withTomato{
		pizza: pizza,
	}
	fmt.Println("price of vegetable Pizza with tomato:", withTomato.price())

	withCheese := withCheese{
		pizza: pizza,
	}
	fmt.Println("price of vegetable Pizza with tomato and cheese:", withCheese.price())
}

Output:
price of vegetable Pizza: 15
price of vegetable Pizza with tomato: 22
price of vegetable Pizza with tomato and cheese: 25

```

## 应用场景

- 如果你希望在无需修改代码的情况下即可使用对象， 且希望在运行时为对象新增额外的行为， 可以使用装饰模式。
- 如果用继承来扩展对象行为的方案难以实现或者根本不可行， 你可以使用该模式。

## 优点

- 你无需创建新子类即可扩展对象的行为。
- 你可以在运行时添加或删除对象的功能。
- 你可以用多个装饰封装对象来组合几种行为。
- 单一职责原则。 你可以将实现了许多不同行为的一个大类拆分为多个较小的类。

## 缺点

- 在封装器栈中删除特定封装器比较困难。
- 实现行为受装饰栈顺序影响。
- 各层的初始化配置代码看上去可能会很糟糕。
