/**
  @Author: majm@ushareit.com
  @date: 2021/2/6
  @note:
**/
package leetcode_0146

import "container/list"

type Item struct {
	Key   int
	Value int
}

type LRUCache struct {
	cap      int
	cacheMap map[int]*list.Element
	ll       *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:      capacity,
		cacheMap: make(map[int]*list.Element, capacity+1),
		ll:       list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if this.IsEmpty() {
		return -1
	}
	if ele, ok := this.cacheMap[key]; ok {
		item := ele.Value.(*Item)
		// 刚访问的元素 移到队头
		this.ll.MoveToFront(ele)
		return item.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.cacheMap == nil {
		this.cacheMap = make(map[int]*list.Element, this.cap+1)
		this.ll = list.New()
	}
	if ele, ok := this.cacheMap[key]; ok {
		item := ele.Value.(*Item)
		item.Value = value
		this.ll.MoveToFront(ele)
		return
	}

	item := &Item{
		Key:   key,
		Value: value,
	}
	ele := this.ll.PushFront(item)
	this.cacheMap[key] = ele

	if this.cap != 0 && this.cap < this.ll.Len() {
		this.removeOldest()
	}
}

func (this *LRUCache) IsEmpty() bool {
	return this.ll.Len() == 0
}

func (this *LRUCache) removeOldest() {
	if this.cacheMap == nil {
		return
	}
	ele := this.ll.Back()
	if ele != nil {
		this.removeElement(ele)
	}
}

func (this *LRUCache) removeElement(ele *list.Element) {
	kv := ele.Value.(*Item)
	delete(this.cacheMap, kv.Key)
	this.ll.Remove(ele)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
