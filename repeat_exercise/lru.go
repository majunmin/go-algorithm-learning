/**
* Date: 2021/5/24 9:06 上午
* Author: majunmin
**/
package repeat_exercise

import (
	"container/list"
)

const (
	defaultEntrySize = 1 << 5
)

type Key interface{}

// key 格式, list.List 的节点
type entry struct {
	key   Key
	value interface{}
}

type Cache struct {
	MaxEntry int
	ll       *list.List
	cache    map[Key]*list.Element

	OnEvict func(key Key, value interface{})
}

func NewCache(maxEntry int) *Cache {
	return &Cache{
		ll:       list.New(),
		MaxEntry: maxEntry,
		cache:    make(map[Key]*list.Element, maxEntry),
	}
}

func (c *Cache) Add(key Key, value interface{}) {
	if c.cache == nil {
		c.cache = make(map[Key]*list.Element, 0)
		c.ll = list.New()
	}
	if ee, exist := c.cache[key]; exist {
		c.ll.MoveToFront(ee)
		ee.Value.(*entry).value = value
		return
	}
	ee := c.ll.PushFront(&entry{key: key, value: value})
	c.cache[key] = ee
	if c.MaxEntry != 0 && c.ll.Len() > c.MaxEntry {
		c.RemoveOldest()
	}
}

func (c *Cache) Remove(key Key) {
	if c.cache == nil {
		return
	}

	if ee, hit := c.cache[key]; hit {
		c.removeElement(ee)
	}
}

func (c *Cache) removeElement(ee *list.Element) {
	if c.cache == nil {
		return
	}
	kv := ee.Value.(*entry)
	c.ll.Remove(ee)
	delete(c.cache, kv.key)
}

func (c *Cache) Get(key Key) (interface{}, bool) {
	if c.cache == nil {
		return nil, false
	}
	if ee, hit := c.cache[key]; hit {
		c.ll.MoveToFront(ee)
		entry := ee.Value.(*entry)
		return entry.value, true
	}
	return nil, false
}

func (c *Cache) Len() int {
	if c.cache == nil {
		return 0
	}
	return c.ll.Len()
}

// RemoveOldest remove the oldest item from the cache
func (c *Cache) RemoveOldest() {
	if c.cache == nil {
		return
	}
	ee := c.ll.Back()
	c.ll.Remove(ee)
	delete(c.cache, ee.Value.(*entry).key)
}
