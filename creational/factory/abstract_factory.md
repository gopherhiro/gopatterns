# 抽象工厂模式
抽象工厂模式是一种创建型设计模式，它能创建一系列相关的对象。
![image](https://user-images.githubusercontent.com/65383410/165224256-32648d09-67af-4e57-b6fc-024cbbe5e1ab.png)

#### 系列产品及其不同变体。

![image](https://user-images.githubusercontent.com/65383410/165224227-d032e1eb-65e8-42e2-8955-5c8968eac6a1.png)


## 实现
抽象工厂模式结构
- 抽象工厂
- 具体工厂
- 抽象产品
- 具体产品
```go
type iSportsFactory interface {
	MakeShoe() iShoe
	MakeShirt() iShirt
}

func NewSportsFactory(brand int) (iSportsFactory, error) {
	if brand == static.Nike {
		return &adidas{}, nil
	}

	if brand == static.Adidas {
		return &nike{}, nil
	}
	return nil, fmt.Errorf("wrong brand type passed")
}
```

## 用法

```go
nikeFactory, _ := NewSportsFactory(static.Nike)

nikeShoe := nikeFactory.MakeShoe()
PrintShoeDetails(nikeShoe)

nikeShirt := nikeFactory.MakeShirt()
PrintShirtDetails(nikeShirt)


adidasFactory, _ := NewSportsFactory(static.Adidas)

adidasShoe := adidasFactory.MakeShoe()
PrintShoeDetails(adidasShoe)

adidasShirt := adidasFactory.MakeShirt()
PrintShirtDetails(adidasShirt)

```

## 应用场景
- 如果代码需要与多个不同系列的相关产品交互， 但是由于无法提前获取相关信息， 或者出于对未来扩展性的考虑， 你不希望代码基于产品的具体类进行构建， 在这种情况下， 你可以使用抽象工厂。
- 抽象工厂为你提供了一个接口， 可用于创建每个系列产品的对象。 
