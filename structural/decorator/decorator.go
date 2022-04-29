package decorator

// Pizza 基本组件和可选层次的通用方法。
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

// Example 2
type Component interface {
	Calc() int
}

type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 0
}

type MulDecorator struct {
	Component
	num int
}

func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func (d *MulDecorator) Calc() int {
	return d.Component.Calc() * d.num
}

type AddDecorator struct {
	Component
	num int
}

func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}
