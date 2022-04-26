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
