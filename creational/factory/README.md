# 简单工厂（Simple Factory）
![image](https://user-images.githubusercontent.com/65383410/165209904-eb648b80-f577-4603-9452-e7573619a8d8.png)

Factory design pattern is a creational design pattern and it is also one of the most commonly used pattern. This pattern provides a way to hide the creation logic of the instances being created.
The client only interacts with a factory struct and tells the kind of instances that needs to be created. The factory class interacts with the corresponding concrete structs and returns the correct instance back.

简单工厂模式：工厂能够生产所有产品，你只要告诉工厂，你需要什么产品，工厂就会生产并返回你需要的产品。

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

