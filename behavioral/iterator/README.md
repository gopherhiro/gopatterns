# 迭代器模式

迭代器模式（Iterator Design Pattern），也叫作游标模式（Cursor Design Pattern）。

迭代器模式将集合对象的遍历操作从集合类中拆分出来，放到迭代器类中，让两者的职责更加单一。

迭代器是一种行为设计模式，让你能在不暴露集合底层表现形式(列表、栈和树等)的情况下遍历集合中所有的元素。

Iterator is a behavioral design pattern that lets you traverse elements of a collection without exposing its underlying
representation (list, stack, tree, etc.).

![image](https://user-images.githubusercontent.com/65383410/165703393-790d4716-2153-4228-bda2-d8b72d65923d.png)

Iterators implement various traversal algorithms. Several iterator objects can traverse the same collection at the same
time.

![image](https://user-images.githubusercontent.com/65383410/165705777-3df019fa-32e1-4e62-823c-1c3941d7742c.png)

The main idea of the Iterator pattern is to extract the traversal behavior of a collection into a separate object called
an iterator.

## 实现

迭代器是用来遍历容器的，所以，一个完整的迭代器模式一般会涉及容器和容器迭代器两部分内容。

容器又包含容器接口、容器实现类，迭代器又包含迭代器接口、迭代器实现类。

容器对象通过依赖注入传递到迭代器类中。

![image](https://user-images.githubusercontent.com/65383410/165725528-11889e0b-6103-42c4-adc8-21a6392efa56.png)



```go
package iterator

// Collection
// 集合
type Collection interface {
	newIterator() Iterator
}

// UserCollection
// 具体集合
type user struct {
	name string
	age  int
}

type UserCollection struct {
	users []*user
}

// 一个集合可以对应多个迭代器
// 可以结合工厂模式来返回不同的迭代器
func (u *UserCollection) newIterator() Iterator {
	return &UserIterator{
		users: u.users,
	}
}

// Iterator
// 迭代器
type Iterator interface {
	hasMore() bool
	next() *user
}

// UserIterator
// 具体迭代器
type UserIterator struct {
	index int
	users []*user
}

func (u *UserIterator) hasMore() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *UserIterator) next() *user {
	if u.hasMore() {
		elem := u.users[u.index]
		u.index++
		return elem
	}
	return nil
}

```

## 用法

```go
package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	user1 := &user{
		name: "a",
		age:  1,
	}
	user2 := &user{
		name: "b",
		age:  2,
	}
	uc := &UserCollection{
		users: []*user{user1, user2},
	}

	iterator := uc.newIterator()
	for iterator.hasMore() {
		u := iterator.next()
		fmt.Printf("user:%+v\n", u)
	}
}

// Output
// user:&{name:a age:1}
// user:&{name:b age:2}
```

## 应用场景

### 为什么用迭代器来遍历容器呢？为什么要给容器设计对应的迭代器呢？
对于类似数组和链表这样的数据结构，遍历方式比较简单，直接使用 for 循环来遍历就足够了。但是，对于复杂的数据结构（比如树、图）来说，有各种复杂的遍历方式。比如，树有前中后序、按层遍历，图有深度优先、广度优先遍历等等。如果由客户端代码来实现这些遍历算法，势必增加开发成本，而且容易写错。如果将这部分遍历的逻辑写到容器类中，也会增加容器类代码的复杂性。

应对复杂性的方法就是拆分。我们可以将遍历操作拆分到迭代器类中。比如，针对图的遍历，我们就可以定义 DFSIterator、BFSIterator 两个迭代器类，让它们分别来实现深度优先遍历和广度优先遍历。

容器和迭代器都提供了抽象的接口，方便我们在开发的时候，基于接口编程。

当需要切换新的遍历算法的时候，客户端代码只需要切换迭代器即可，其他代码都不需要修改。

除此之外，添加新的遍历算法，我们只需要扩展新的迭代器类，也更符合开闭原则。


## 优点
- 迭代器模式封装集合内部的复杂数据结构，开发者不需要了解如何遍历，直接使用容器提供的迭代器即可。
- 迭代器模式将集合对象的遍历操作从集合类中拆分出来，放到迭代器类中，让两者的职责更加单一。
- 迭代器模式让添加新的遍历算法更加容易，更符合开闭原则。除此之外，因为迭代器都实现自相同的接口，在开发中，基于接口而非实现编程，替换迭代器也变得更加容易。

## 缺点

- 对于某些特殊集合，使用迭代器可能比直接遍历的效率低。
- 如果你的程序只与简单的集合进行交互，应用该模式可能会过度设计。
