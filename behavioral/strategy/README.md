# 策略模式

策略模式，英文全称是 Strategy Design Pattern。

Define a family of algorithms, encapsulate each one, and make them interchangeable. Strategy lets the algorithm vary
independently from clients that use it.

策略是一种行为设计模式，它能让你定义一系列算法，并将每种算法分别放入独立的类中，以使算法的对象能够相互替换。

不同客户端可通过一个简单接口执行算法，并能在运行时进行切换。

![image](https://user-images.githubusercontent.com/65383410/165496250-843144a1-5684-4196-87dd-b99f154a3a64.png)

## Real-World Analogy

Various strategies for getting to the airport.

![image](https://user-images.githubusercontent.com/65383410/165477697-faf51166-a800-4206-bca5-4ba339ba2b5d.png)

Imagine that you have to get to the airport. You can catch a bus, order a cab, or get on your bicycle. These are your
transportation strategies. You can pick one of the strategies depending on factors such as budget or time constraints.

## 作用

策略模式解耦策略的定义、创建、使用三个部分。

##### 策略的定义

策略类的定义比较简单，包含一个策略接口和一组实现这个接口的策略类。因为所有的策略类都实现相同的接口，所以，客户端代码基于接口编程，可以灵活地替换不同的策略。

##### 策略的创建

因为策略模式会包含一组策略，在使用它们的时候，一般会通过类型（type）来判断创建哪个策略来使用。为了封装创建逻辑，我们需要对客户端代码屏蔽创建细节。我们可以把根据 type 创建策略的逻辑抽离出来，放到工厂类中。

一般来讲，如果策略类是无状态的，不包含成员变量，只是纯粹的算法实现，这样的策略对象是可以被共享使用的，不需要在每次调用的时候，都创建一个新的策略对象。针对这种情况，我们事先创建好每个策略对象，缓存到工厂类中，用的时候直接返回。

##### 策略的使用
我们知道，策略模式包含一组可选策略，客户端代码一般如何确定使用哪个策略呢？最常见的是**运行时动态确定使用哪种策略**，这也是策略模式最典型的应用场景。

这里的**运行时动态**指的是，我们事先并不知道会使用哪个策略，而是在程序运行期间，根据配置、用户输入、计算结果等这些不确定因素，动态决定使用哪种策略。

## 实现

```go
package strategy

import "fmt"

// 以缓存淘汰算法举例，来研究策略模式

// EvictAlgo Strategy interface
// 策略(Strategy)接口是所有具体策略的通用接口，它声明 了一个上下文用于执行策略的方法。
type EvictAlgo interface {
	evict(c *Cache)
}

// 不同的缓存淘汰策略算法
// 具体策略(Concrete Strategies)实现了上下文所用算法的各 种不同变体。

// FIFO Concrete strategy
type FIFO struct {
}

func (f *FIFO) evict(c *Cache) {
	fmt.Println("Evicting by FIFO strategy")
}

// LRU Concrete strategy
type LRU struct {
}

func (l *LRU) evict(c *Cache) {
	fmt.Println("Evicting by LRU strategy")
}

// LFU Concrete strategy
type LFU struct {
}

func (l *LFU) evict(c *Cache) {
	fmt.Println("Evicting by LFU strategy")
}

// Cache Context
// 上下文(Context)维护指向具体策略的引用
type Cache struct {
	storage   map[string]string
	evictAlgo EvictAlgo
	length    int
	capacity  int
}

func initCache(e EvictAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:   storage,
		evictAlgo: e,
		length:    0,
		capacity:  2,
	}
}

// 上下文则会提供一个设置器以便客户端在运行时替换相关联的策略。
func (c *Cache) setEvictAlgo(e EvictAlgo) {
	c.evictAlgo = e
}

func (c *Cache) add(key, value string) {
	if c.length == c.capacity {
		c.evict()
	}
	c.storage[key] = value
	c.length++
}

func (c *Cache) evict() {
	c.evictAlgo.evict(c)
	c.length--
}

```

## 用法

```go
package strategy

import (
	"testing"
)

func TestStrategy(t *testing.T) {
	lfu := &LFU{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	lru := &LRU{}
	cache.setEvictAlgo(lru)

	cache.add("d", "4")

	fifo := &FIFO{}
	cache.setEvictAlgo(fifo)

	cache.add("e", "5")
}

// Output:
// Evicting by LFU strategy
// Evicting by LRU strategy
// Evicting by FIFO strategy
```

## 应用场景

- 最常见的应用场景是，利用它来避免冗长的 if-else 或 switch 分支判断。
- When an object needs to support different behavior and you want to change the behavior at run time.
- When you want to avoid a lot of conditionals of choosing the runtime behavior.
- When you have different algorithms that are similar and they only differ in the way they execute some behavior.

## 优点

- 你可以在运行时切换对象内的策略算法。
- 你可以将算法的实现和使用算法的代码隔离开来。
- 开闭原则。你无需对上下文进行修改就能够引入新的策略。

## 缺点

- 如果你的策略极少发生改变，那么没有任何理由引入新的类和接口。使用该模式只会让程序过于复杂。
- 客户端必须知晓策略间的不同——它需要选择合适的策略。
