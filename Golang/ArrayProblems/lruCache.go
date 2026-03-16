package main

import (
	"container/list"
	"fmt"
)

func main() {
	lrucache := NewLruCache(2)
	lrucache.Put(1, 2)
	lrucache.Put(3, 6)
	lrucache.Put(4, 8)
	fmt.Println(lrucache.Get(1))
}

type entry struct {
	key   int
	value int
}

// LRU cache Q front is most recently used. queue back is stale element.
type LruCache struct {
	capacity int
	cache    map[int]*list.Element
	queue    *list.List
}

func NewLruCache(capacity int) *LruCache {
	return &LruCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		queue:    list.New(),
	}
}

func (lru *LruCache) Put(key int, value int) {
	if entry1, ok := lru.cache[key]; ok {
		entry1.Value.(*entry).value = value
		lru.cache[key] = entry1
		// move to most recently accessed element
		lru.queue.MoveToFront(entry1)
	} else {
		if len(lru.cache) == lru.capacity {
			evict := lru.queue.Back()
			delete(lru.cache, evict.Value.(*entry).key)
			lru.queue.Remove(evict)
		}
		// insert new entry in cache and queue
		newElement := &entry{key, value}
		ele := lru.queue.PushFront(newElement)
		lru.cache[key] = ele
	}

}

func (lru *LruCache) Get(key int) (value int) {
	if element, ok := lru.cache[key]; ok {
		// mark currently accessed element as most recently used.
		lru.queue.MoveToFront(element)
		return element.Value.(*entry).value
	}
	return -1

}

func (lru *LruCache) Remove(key int) {
	element := lru.cache[key]
	delete(lru.cache, key)
	lru.queue.Remove(element)
}

func (lru *LruCache) Len() int {
	return len(lru.cache)
}
