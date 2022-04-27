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
