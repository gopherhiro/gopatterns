# 单例模式

单例设计模式（Singleton Design Pattern）理解起来非常简单：一个类只允许创建一个对象（或者实例）。

Singleton creational design pattern restricts the instantiation of a type to a single object.

Singleton is a creational design pattern that lets you ensure that a class has only one instance, while providing a
global access point to this instance.

![image](https://user-images.githubusercontent.com/65383410/165093763-471b6849-e56b-49d1-a610-3acaaabc17ac.png)

##### 弄清楚这样几个问题

- 为什么要使用单例？它能解决哪些问题？
- 单例存在哪些问题？有何替代的解决方案？

##### 单例模式解决的问题

- 处理资源访问冲突
    - 日志记录对象
- 表示全局唯一类
    - 从业务概念上，如果有些数据在系统中只应保存一份，那就比较适合设计为单例类。比如：
        1. 配置信息类。
        2. 唯一递增 ID 号码生成器。

###### 单例模式存在的问题

1. 单例对 OOP 特性的支持不友好
2. 单例会隐藏类之间的依赖关系
3. 单例对代码的扩展性不友好。
    - 硬编码、耦合性强，改动范围大。
    - 所以，数据库连接池、线程池这类的资源池，最好还是不要设计成单例类。
4. 单例对代码的可测试性不友好


## 实现

```go
package singleton

type singleton map[string]string

var (
	once sync.Once

	instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}
```

## 使用

```go
s := singleton.New()

s["this"] = "that"

s2 := singleton.New()

fmt.Println("This is ", s2["this"])
// This is that
```

## 运用场景

- Logger instance
- 生产唯一序号
- Web计数器
- 配置信息类
