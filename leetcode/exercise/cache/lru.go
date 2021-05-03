/**
  @Author: majm@ushareit.com
  @date: 2021/3/9
  @note:
**/
package cache

import "container/list"

type Key interface{}

type entry struct {
	key   Key         // entry中保存 key是为了后面方便删除 key
	Value interface{} // 具体值
}

type Cache struct {
	capacity int
	ll       *list.List
	OnEvict  func(key Key, value interface{})
	cache    map[Key]*list.Element
}

func NewCache(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		ll:       list.New(),
		cache:    make(map[Key]*list.Element),
	}
}

func (c *Cache) Get(key Key) interface{} {
	if c.cache == nil {
		c.cache = make(map[Key]*list.Element)
		c.ll = list.New()
	}

	if ee, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ee)
		return ee.Value
	}
	return nil
}

func (c *Cache) Add(key Key, value interface{}) {

	if c.cache == nil {
		c.cache = make(map[Key]*list.Element)
		c.ll = list.New()
	}
	if ee, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ee)
		ee.Value.(*entry).Value = value
	}
	ee := c.ll.PushFront(&entry{
		key:   key,
		Value: value,
	})
	c.cache[key] = ee
	// c.capacity =0  表示不限制
	if c.capacity != 0 && c.Len() > c.capacity {
		c.RemoveOldest()
	}

}

func (c Cache) Len() int {
	if c.ll == nil {
		return 0
	}
	return c.ll.Len()
}

func (c *Cache) RemoveKey(key Key) {
	if c.cache == nil {
		return
	}

	if ee, ok := c.cache[key]; ok {
		c.removeElement(ee)
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

// 内部方法 可以抱枕  c.ll 不为nil
func (c *Cache) removeElement(ee *list.Element) {
	c.ll.Remove(ee)
	entry := ee.Value.(*entry)
	delete(c.cache, entry.key)

	if c.OnEvict != nil {
		c.OnEvict(entry.key, entry.Value)
	}

}

func (c *Cache) Clear() {
	if c.OnEvict != nil {
		for _, ee := range c.cache {
			entry := ee.Value.(*entry)
			c.OnEvict(entry.key, entry.Value)
		}
	}

	c.ll = nil
	c.cache = nil
}
