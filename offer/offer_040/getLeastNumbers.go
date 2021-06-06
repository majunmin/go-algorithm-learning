/**
* Date: 2021/6/1 10:13 上午
* Author: majunmin
**/
package offer_040

import "container/heap"

type IntHeap []int

func (ih IntHeap) Len() int {
	return len(ih)
}

// 小顶堆
func (ih IntHeap) Less(i, j int) bool {
	return ih[i] < ih[j]
}

func (ih IntHeap) Swap(i, j int) {
	ih[i], ih[j] = ih[j], ih[i]
}

func (ih *IntHeap) Push(x interface{}) {
	*ih = append(*ih, x.(int))
}

func (ih *IntHeap) Pop() interface{} {
	old := *ih
	n := len(old)
	item := old[n-1]
	*ih = (*ih)[0 : n-1]
	return item
}

// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
func getLeastNumbers(arr []int, k int) []int {
	if len(arr) < k {
		return arr
	}
	var result []int
	var intHeap IntHeap
	for _, item := range arr {
		intHeap = append(intHeap, item)
	}
	heap.Init(&intHeap)

	for i := 0; i < k; i++ {
		item := heap.Pop(&intHeap)
		result = append(result, item.(int))
	}
	return result

}
