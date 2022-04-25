# 简单工厂（Simple Factory）
工厂方法模式是一种创建型设计模式， 其在父类中提供一个创建对象的方法， 允许子类决定实例化对象的类型。

Factory design pattern is a creational design pattern and it is also one of the most commonly used pattern. This pattern provides a way to hide the creation logic of the instances being created.
The client only interacts with a factory struct and tells the kind of instances that needs to be created. The factory class interacts with the corresponding concrete structs and returns the correct instance back.

## 实现

```go
type Course interface {
	GetID() int64
	GetName() string
}

// Chinese
type Chinese struct {
}

func (m *Chinese) GetID() int64 {
	return static.Chinese
}

func (m *Chinese) GetName() string {
	return "Chinese"
}

// English
type English struct {
}

func (m *English) GetID() int64 {
	return static.English
}

func (m *English) GetName() string {
	return "English"
}

// 简单工厂（Simple Factory）- 函数式使用方式
func New(ID int64) Course {
	if ID == static.Chinese {
		return &Chinese{}
	}

	if ID == static.English {
		return &English{}
	}

	return nil
}

// 简单工厂（Simple Factory）- 缓存式使用方式
var CourseMap = map[int64]Course{
	static.Chinese: &Chinese{},
	static.English: &English{},
}

func NewFromCache(ID int64) Course {
	// ok-idiom 模式
	if instance, ok := CourseMap[ID]; ok {
		return instance
	}
	return nil
}
```

## 用法

```go
var c Course

c = New(static.Chinese)
fmt.Println(c.GetName())
// Chinese

c = NewFromCache(static.English)
fmt.Println(c.GetName())
// English
```

## 应用场景
- 尽管简单工厂模式的代码实现中，有多处 if 分支判断逻辑，违背开闭原则，但权衡扩展性和可读性，这么实现在大多数情况下是没有问题的。
- 当对象的创建逻辑比较简单时，推荐使用简单工厂模式。
- 当对象的创建逻辑比较复杂时，推荐使用工厂方法模式。

