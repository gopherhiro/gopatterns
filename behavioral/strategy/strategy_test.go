package strategy

import (
	"testing"
)

func TestStrategy(t *testing.T) {
	initStrategy := NewStrategy(3)
	cache := initCache(initStrategy)

	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	lru := &LRU{}
	cache.setEvictAlgo(lru)

	cache.add("d", "4")

	fifo := &FIFO{}
	cache.setEvictAlgo(fifo)

	cache.add("e", "5")

	runStrategy := NewStrategy(2)
	cache.setEvictAlgo(runStrategy)

	cache.add("e", "6")

}
