package strategy

import "fmt"

// 以缓存淘汰算法举例，来研究策略模式

// EvictAlgo Strategy interface
type EvictAlgo interface {
	evict(c *Cache)
}

// 不同的缓存淘汰策略算法

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
