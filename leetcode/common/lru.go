/**
  @Author: majm@ushareit.com
  @date: 2021/1/25
  @note:
**/
package common

import "container/list"

type Cache struct {
	MaxEntries int
	// 保存 lru的 顺序,进行lru的操作
	ll    *list.List
	cache map[interface{}]*list.Element
	// callback on ele remove
	OnEvicted func(key Key, value interface{})
}

type Key interface{}

type entry struct {
	key   Key
	value interface{}
}

// lru cache
func (c *Cache) Add(key Key, val interface{}) {
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}

	if ee, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ee)
		ee.Value.(*entry).value = val
		return
	}
	entry := &entry{key: key, value: val}
	ele := c.ll.PushFront(entry)
	c.cache[key] = ele

	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		c.RemoveOldest()
	}
}

func (c *Cache) Get(key Key) (interface{}, bool) {
	if c.cache == nil {
		return nil, false
	}
	if ele, hit := c.cache[key]; hit {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(entry)
		return kv.value, true
	}
	return nil, false
}

func (c *Cache) Remove(key Key) interface{} {
	if c.cache == nil {
		return nil
	}
	if ele, hit := c.cache[key]; hit {
		c.removeElement(ele)
		return ele.Value
	}
	return nil
}

func (c *Cache) removeElement(e *list.Element) {
	c.ll.Remove(e)
	kv := e.Value.(entry)
	delete(c.cache, kv.key)
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}

func (c *Cache) RemoveOldest() {
	if c.cache == nil {
		return
	}
	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}
