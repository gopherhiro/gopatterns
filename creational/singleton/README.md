# 单例模式

Singleton creational design pattern restricts the instantiation of a type to a single object.

Singleton is a creational design pattern that lets you ensure that a class has only one instance, while providing a global access point to this instance.

![image](https://user-images.githubusercontent.com/65383410/165093763-471b6849-e56b-49d1-a610-3acaaabc17ac.png)


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
- DB instance
- 生产唯一序号
- Web计数器
- 缓存
