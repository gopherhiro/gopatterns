# 迭代器模式

迭代器是一种行为设计模式，让你能在不暴露集合底层表现形式(列表、栈和树等)的情况下遍历集合中所有的元素。

Iterator is a behavioral design pattern that lets you traverse elements of a collection without exposing its underlying
representation (list, stack, tree, etc.).

![image](https://user-images.githubusercontent.com/65383410/165703393-790d4716-2153-4228-bda2-d8b72d65923d.png)

Iterators implement various traversal algorithms. Several iterator objects can traverse the same collection at the same
time.

![image](https://user-images.githubusercontent.com/65383410/165705777-3df019fa-32e1-4e62-823c-1c3941d7742c.png)

The main idea of the Iterator pattern is to extract the traversal behavior of a collection into a separate object called
an iterator.

迭代器模式的主要思想是将集合的遍历行为抽取为单独的迭 代器对象。

## 实现

```go
package iterator

// Collection
// 集合
type Collection interface {
	newIterator() Iterator
}

// UserCollection
// Concrete collection
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
// 迭代接口
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

- 迭代器封装了与复杂数据结构进行交互的细节，为客户端提供多个访问集合元素的简单方法。这种方式不仅对客户端来说非常方便，而且能避免客户端在直接与集合交互时执行错误或有害的操作，从而起到保护集合的作用。
- 重要迭代算法的代码往往体积非常庞大。当这些代码被放置在程序业务逻辑中时，它会让原始代码的职责模糊不清，降低其可维护性。因此，将遍历代码移到特定的迭代器中可使程序代码更加精炼和简洁。

## 优点

- 使用该模式可以减少程序中重复的遍历代码。
- 你可以并行遍历同一集合， 因为每个迭代器对象都包含其自身的遍历状态。

## 缺点

- 对于某些特殊集合，使用迭代器可能比直接遍历的效率低。
- 如果你的程序只与简单的集合进行交互，应用该模式可能会过度设计。
