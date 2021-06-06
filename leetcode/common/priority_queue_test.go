/**
* Date: 2021/6/1 10:51 上午
* Author: majunmin
**/
package common

import (
	"container/heap"
	"testing"
)

func Test_PriorityQueue(t *testing.T) {
	var pq PriorityQueue
	pq = append(pq, &Item{Value: 88, Priority: 10})
	pq = append(pq, &Item{Value: 89, Priority: 9})
	pq = append(pq, &Item{Value: 90, Priority: 8})
	pq = append(pq, &Item{Value: 91, Priority: 7})
	pq = append(pq, &Item{Value: 92, Priority: 6})
	pq = append(pq, &Item{Value: 93, Priority: 5})

	heap.Init(&pq)

	for i := 0; i < 8; i++ {
		item := heap.Pop(&pq)
		it := item.(*Item)
		t.Log(it.Value)
	}
}
