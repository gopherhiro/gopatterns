# 观察者模式
亦称:事件订阅者、监听者、Event-Subscriber、Listener、Observer

观察者是一种行为设计模式，允许你定义一种订阅机制，可在对象事件发生时通知多个“观察”该对象的其他对象。

Observer is a behavioral design pattern that lets you define a subscription mechanism to notify multiple objects about any events that happen to the object they’re observing.

![image](https://user-images.githubusercontent.com/65383410/165441876-9b869dc7-cad3-4abd-bc9a-6329d4456446.png)

Magazine and newspaper subscriptions.

![image](https://user-images.githubusercontent.com/65383410/165442020-e056bc86-d23c-4e07-8d9d-15b15ac843f6.png)

## 实现

```go
package observer

import "fmt"

// 主体
type Subject interface {
	register(o Observer)
	deregister(o Observer)
	notifyObservers()
}

// 观察者
type Observer interface {
	update(msg string)
}

// 具体主体
type Item struct {
	observers []Observer
	name      string
	inStock   bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) haveNewEvent() {
	fmt.Printf("item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyObservers()
}

func (i *Item) register(o Observer) {
	i.observers = append(i.observers, o)
}

func (i *Item) deregister(o Observer) {
	i.observers = removeObserver(i.observers, o)
}

func (i *Item) notifyObservers() {
	for _, observer := range i.observers {
		observer.update(i.name)
	}
}

func removeObserver(obs []Observer, o Observer) []Observer {
	length := len(obs)
	for i, ob := range obs {
		if ob == o {
			// 把要删除的元素交换到最后一个位置，再执行删除。
			obs[length-1], obs[i] = obs[i], obs[length-1]
			return obs[:length-1]
		}
	}
	return obs
}

// 具体观察者:Customer
type Customer struct {
}

func (c *Customer) update(msg string) {
	// TODO:收到通知，执行自己的业务逻辑
	fmt.Println("Customer is notified")
	fmt.Println("Customer already know has new item:", msg)
}

// 具体观察者:Trader
type Trader struct {
}

func (t *Trader) update(msg string) {
	// TODO:收到通知，执行自己的业务逻辑
	fmt.Println("Trader is notified")
	fmt.Println("Trader already know has new item:", msg)
}

```

## 用法

```go
package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {
	item := NewItem("Nike Shirt")

	observer1 := &Customer{}
	observer2 := &Trader{}

	item.register(observer1)
	item.register(observer2)

	item.haveNewEvent()

	item.deregister(observer1)
	item.haveNewEvent()
}

```

## 应用场景
- 当一个对象状态的改变需要改变其他对象，或实际对象是事先未知的或动态变化的时，可使用观察者模式。
- 当应用中的一些对象必须观察其他对象时，可使用该模式。但仅能在有限时间内或特定情况下使用。
    - 订阅列表是动态的， 因此订阅者可随时加入或离开该列表。
    
## 优点
- 开闭原则。 你无需修改发布者代码就能引入新的订阅者类
- 你可以在运行时建立对象之间的联系。

## 缺点
- 订阅者的通知顺序是随机的。
