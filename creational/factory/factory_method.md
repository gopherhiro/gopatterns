# 工厂方法模式
![image](https://user-images.githubusercontent.com/65383410/165209967-e71b9e32-7890-4f22-8552-5b6481622aa4.png)

Factory design pattern is a creational design pattern and it is also one of the most commonly used pattern. This pattern provides a way to hide the creation logic of the instances being created.
The client only interacts with a factory struct and tells the kind of instances that needs to be created. The factory class interacts with the corresponding concrete structs and returns the correct instance back.

工厂方法模式：你告诉工厂需要的产品，工厂先拉投资，创建一个生产该产品工厂，然后再通过这个工厂生产你所需要的产品。

## 实现

```go
// 为工厂类再创建一个简单工厂，也就是工厂的工厂，用来创建工厂类对象。
func NewFactory(ID int) CourseFactory {
	if ID == static.Chinese {
		return newChineseFactory()
	}
	if ID == static.English {
		return newEnglishFactory()
	}
	return nil
}
```

## 用法

```go
var f CourseFactory
var c Course

f = NewFactory(static.Chinese)
c = f.NewCourse()
fmt.Println(c.GetName())
// Chinese

f = NewFactory(static.English)
c = f.NewCourse()
fmt.Println(c.GetName())
// English
```


## 应用场景

### 什么时候该用工厂模式？相对于直接 new 来创建对象，用工厂模式来创建究竟有什么好处呢？
当创建逻辑比较复杂，是一个“大工程”的时候，我们就考虑使用工厂模式，封装对象的创建过程，将对象的创建和使用相分离。

### 何为创建逻辑比较复杂呢？
有以下两种情况：
- 第一种情况：代码中存在 if-else 分支判断，动态地根据不同的类型创建不同的对象。针对这种情况，我们就考虑使用工厂模式，将这一大坨 if-else 创建对象的代码抽离出来，放到工厂类中。
- 第二种情况：尽管我们不需要根据不同的类型创建不同的对象，但是，单个对象本身的创建过程比较复杂，比如要组合其他类对象，做各种初始化操作。在这种情况下，我们也可以考虑使用工厂模式，将对象的创建过程封装到工厂类中。

对于第一种情况，当每个对象的创建逻辑都比较简单的时候，推荐使用简单工厂模式，将多个对象的创建逻辑放到一个工厂类中。当每个对象的创建逻辑都比较复杂的时候，为了避免设计一个过于庞大的简单工厂类，推荐使用工厂方法模式，将创建逻辑拆分得更细，每个对象的创建逻辑独立到各自的工厂类中。同理，对于第二种情况，因为单个对象本身的创建逻辑就比较复杂，所以，建议使用工厂方法模式。

上升一个思维层面来看工厂模式，它的作用无外乎下面这四个，这也是判断要不要使用工厂模式的最本质的参考标准。
- 封装变化：创建逻辑有可能变化，封装成工厂类之后，创建逻辑的变更对调用者透明。
- 代码复用：创建代码抽离到独立的工厂类之后可以复用。
- 隔离复杂性：封装复杂的创建逻辑，调用者无需了解如何创建对象。
- 控制复杂度：将创建代码抽离出来，让原本的函数或类职责更单一，代码更简洁。

