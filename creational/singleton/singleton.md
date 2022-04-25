# Singleton Pattern

Singleton creational design pattern restricts the instantiation of a type to a single object.

Singleton is a creational design pattern that lets you ensure that a class has only one instance, while providing a global access point to this instance.

## Implementation

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

## Usage

```go
s := singleton.New()

s["this"] = "that"

s2 := singleton.New()

fmt.Println("This is ", s2["this"])
// This is that
```

## Rules of Thumb

- Singleton pattern represents a global state and most of the time reduces testability.