# 门面模式
门面模式，也叫外观模式，英文全称是 Facade Design Pattern。

Provide a unified interface to a set of interfaces in a subsystem. Facade defines a higher-level interface that makes the subsystem easier to use.

门面模式为子系统提供一组统一的接口，定义一组高层接口「让子系统更易用」。

## 举例理解
假设有一个系统 A，提供了 a、b、c、d 四个接口。系统 B 完成某个业务功能，需要调用 A 系统的 a、b、d 接口。利用门面模式，我们提供一个包裹 a、b、d 接口调用的门面接口 x，给系统 B 直接使用。

## 实现

```go
package facade

import "fmt"

type T interface {
	yesterday()
	today()
	tomorrow()
}

type Day struct {
}

func (d *Day) yesterday() {
	fmt.Println("I am yesterday")
}

func (d *Day) today() {
	fmt.Println("I am today")
}

func (d *Day) tomorrow() {
	fmt.Println("I am tomorrow")
}

// 门面方法：可以进行时光穿梭
func (d *Day) timeTravel() {
	d.yesterday()
	d.today()
	d.tomorrow()
}

```

## 用法

```go
package facade

import (
	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {
	d := &Day{}
	d.today()

	fmt.Println("")
	d.timeTravel()
}

// Output:
// I am today

// I am yesterday
// I am today
// I am tomorrow
```

## 应用场景
门面模式应用场景比较明确，主要在接口（粒度）设计方面使用。
1. 解决易用性问题
2. 解决性能问题
3. 解决分布式事务问题

##### 解决易用性问题
门面模式可以用来封装系统的底层实现，隐藏系统的复杂性，提供一组更加简单易用、更高层的接口。

Linux 系统调用函数就可以看作一种“门面”。它是 Linux 操作系统暴露给开发者的一组“特殊”的编程接口，它封装了底层更基础的 Linux 内核调用。再比如，Linux 的 Shell 命令，实际上也可以看作一种门面模式的应用。它继续封装系统调用，提供更加友好、简单的命令，让我们可以直接通过执行命令来跟操作系统交互。

类似封装、抽象的设计思想：“提供更抽象的接口，封装底层实现细节。”

##### 解决性能问题
我们通过将多个接口调用替换为一个门面接口调用，减少网络通信成本，提高 App 客户端的响应速度。

##### 解决分布式事务问题
设计一个新接口，让新接口在一个事务中执行具有事务要求业务需求。
只是这么做，可能会造成不相关业务耦合，不便于后期维护。

## 优点
- 可以让代码独立于复杂子系统。

## 缺点
- 可能造成不必要的耦合。

## 总结
接口粒度设计得太大，太小都不好。太大会导致接口不可复用，太小会导致接口不易用。在实际的开发中，接口的可复用性和易用性需要“微妙”的权衡。
针对这个问题，一个基本的处理原则是：**尽量保持接口的可复用性，但针对特殊情况，允许提供冗余的门面接口，来提供更易用的接口**。

### 闪记
适配器和门面模式的区别：
- 适配器是做接口转换，解决的是原接口和目标接口不匹配的问题。
- 门面模式做接口整合，解决的是多接口调用带来的问题。
