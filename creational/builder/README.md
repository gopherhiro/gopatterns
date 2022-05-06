# 建造者模式

Builder 模式，也称为建造者模式 or 生成器模式。 是一种创建型设计模式，使你能够分步骤创建复杂对象。该模式允许你使用相同的创建代码生成不同类型和形式的对象。

![image](https://user-images.githubusercontent.com/65383410/165256671-ec106afd-4036-4647-99c9-269eb2955836.png)

生成器模式建议将对象构造代码从产品类中抽取出来， 并将其放在一个名为生成器的独立对象中。
![image](https://user-images.githubusercontent.com/65383410/165235715-43884ec1-5c98-4949-80d0-25b89a797bf5.png)

该模式会将对象构造过程划分为一组步骤。每次创建对象时， 你都需要通过生成器对象执行一系列步骤。 重点在于你无需调用所有步骤， 而只需调用创建特定对象配置所需的那些步骤即可。

当你需要创建不同形式的产品时，其中的一些构造步骤可能需要不同的实现。 例如，木屋的房门可能需要使用木头制造，而城堡的房门则必须使用石头制造。

在这种情况下，你可以创建多个不同的生成器，用不同方式实现一组相同的创建步骤。然后你就可以在创建过程中使用这些生成器来生成不同类型的对象。

不同生成器以不同方式执行相同的任务。

![image](https://user-images.githubusercontent.com/65383410/165236788-7f51918d-de7d-4541-89cd-7216fa582bc0.png)

假设第一个建造者使用木头和玻璃制造房屋，第二个建造者使用石头和钢铁，而第三个建造者使用黄金和钻石。

在调用同一组步骤后，第一个建造者会给你一栋普通房屋，第二个会给你一座小城堡，而第三个则会给你一座宫殿。

但是，只有在调用构造步骤的客户端代码可以通过通用接口与建造者进行交互时，这样的调用才能返回需要的房屋。

## 为什么需要建造者模式？

建造者模式用来创建一种类型复杂的对象。 可以根据不同的选择，「定制化」地创建不同的对象。

## 实现

### 概念示例

当所需产品较为复杂且需要多个步骤才能完成时， 也可以使用生成器模式。 在这种情况下， 使用多个构造方法比仅仅使用一个复杂可怕的构造函数更简单。 分为多个步骤进行构建的潜在问题是， 构建不完整的和不稳定的产品可能会被暴露给客户端。
生成器模式能够在产品完成构建之前使其处于私密状态。

在下方的代码中， 我们可以看到 igloo­Builder冰屋生成器与 normal­Builder普通房屋生成器可建造不同类型房屋， 即 igloo冰屋和 normal­House普通房屋 。 每种房屋类型的建造步骤都是相同的。 主管
（可选） 结构体可对建造过程进行组织。

```go
package builder

// 要建造的对象
type house struct {
	windowType string
	doorType   string
	floor      int
}

// 建造者模式
type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func Builder(builder string) iBuilder {
	if builder == "normal" {
		return newNormalBuilder()
	}

	if builder == "igloo" {
		return newIglooBuilder()
	}
	return nil
}

// 指导者
type director struct {
	builder iBuilder
}

func NewDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) BuildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

// 普通建造者，建造普通房屋
type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *normalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *normalBuilder) setNumFloor() {
	b.floor = 1
}

func (b *normalBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// 冰屋建造者，建造冰屋
type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *iglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *iglooBuilder) setNumFloor() {
	b.floor = 2
}

func (b *iglooBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

```

## 用法

```go
normalBuilder := builder.Builder("normal")
normalDirector := builder.NewDirector(normalBuilder)
r := normalDirector.BuildHouse()
fmt.Println(r)

iglooBuilder := builder.Builder("igloo")
iglooDirector := builder.NewDirector(iglooBuilder)
r = iglooDirector.BuildHouse()
fmt.Println(r)
```

## 应用场景

- 如果你需要创建的各种形式的产品，它们的制造过程相似且仅有细节上的差异，此时可使用生成器模式。
- 生成器模式让你能分步骤构造产品。你可以延迟执行某些步骤而不会影响最终产品。
- 生成器在执行制造步骤时，不能对外发布未完成的产品。生成器模式能够在产品完成构建之前使其处于私密状态。这可以避免客户端代码获取到不完整结果对象的情况。
