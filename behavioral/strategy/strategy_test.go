package strategy

import (
	"testing"
)

func TestStrategy(t *testing.T) {
	strategy := NewStrategy(3)
	cache := initCache(strategy)

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
