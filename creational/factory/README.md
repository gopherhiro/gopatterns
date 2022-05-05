# 简单工厂（Simple Factory）

![image](https://user-images.githubusercontent.com/65383410/165209904-eb648b80-f577-4603-9452-e7573619a8d8.png)

Factory design pattern is a creational design pattern and it is also one of the most commonly used pattern. This pattern
provides a way to hide the creation logic of the instances being created. The client only interacts with a factory
struct and tells the kind of instances that needs to be created. The factory class interacts with the corresponding
concrete structs and returns the correct instance back.

工厂能够生产所有产品，你只要告诉工厂，你需要什么产品，工厂就会生产并返回你需要的产品。

简单工厂：封装创建对象过程的类，即是简单工厂类。

工厂方法：多态设计思想结合「简单工厂」的使用。简称之，工厂的工厂。

抽象工厂：接口抽象，让一个工厂可以创建多个不同类型的对象。

## 问答

问：什么时候该用工厂模式？

答：当创建逻辑比较复杂，是一个“大工程”的时候，我们就考虑使用工厂模式，封装对象的创建过程，将对象的创建和使用相分离。

问：相对于直接 new 来创建对象，用工厂模式来创建究竟有什么好处呢？

答：封装变化、代码复用、隔离复杂性、控制复杂度。

- 封装变化：创建逻辑有可能变化，封装成工厂类之后，创建逻辑的变更对调用者透明。
- 代码复用：创建代码抽离到独立的工厂类之后可以复用。
- 隔离复杂性：封装复杂的创建逻辑，调用者无需了解如何创建对象。
- 控制复杂度：将创建代码抽离出来，让原本的函数或类职责更单一，代码更简洁。

## 实现

```go
// 简单工厂（Simple Factory）- 函数式使用方式
func NewCourse(ID int64) Course {
if ID == static.Chinese {
return newChinese()
}

if ID == static.English {
return newEnglish()
}

return nil
}
```

## 用法

```go
var c Course

c = NewCourse(static.Chinese)
fmt.Println(c.GetName())
// Chinese

c = NewCourse(static.English)
fmt.Println(c.GetName())
// English
```

## 应用场景

- 尽管简单工厂模式的代码实现中，有多处 if 分支判断逻辑，违背开闭原则，但权衡扩展性和可读性，这么实现在大多数情况下是没有问题的。
- 当对象的创建逻辑比较简单时，推荐使用简单工厂模式。
- 当对象的创建逻辑比较复杂时，推荐使用工厂方法模式。
- 当需同时创建多个不同对象时，推荐使用抽象工厂模式。

