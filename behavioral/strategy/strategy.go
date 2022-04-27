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
	// 上下文仅可通过策略接口同策略对象进行交互
	c.evictAlgo.evict(c)
	c.length--
}
