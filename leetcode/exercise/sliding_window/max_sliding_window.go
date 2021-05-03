/**
  @Author: majm@ushareit.com
  @date: 2021/3/16
  @note:
**/
package sliding_window

import (
	"container/heap"
	"sort"
)

func maxSlidingWindow(nums []int, k int) []int {
	return priorityQueue(nums, k)
}

// 队列里面存的是索引
func Queue(nums []int, k int) []int {
	length := len(nums)
	if length < k {
		return nil
	}
	var queue []int
	result := make([]int, 0, length-k+1)

	push := func(v int) {
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[v] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, v)
	}

	for i := 0; i < k; i++ {
		push(i)
	}
	result = append(result, nums[queue[0]])
	for i := k; i < length; i++ {
		push(i)
		for queue[0] <= i-k {
			queue = queue[1:]
		}
		result = append(result, nums[queue[0]])
	}
	return result
}

var a []int

type hp struct {
	sort.IntSlice
}

// 实现大根堆 or 小根堆
func (h hp) Less(i, j int) bool {
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}

func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() interface{} {
	lastIdx := len(h.IntSlice) - 1
	val := h.IntSlice[lastIdx]
	h.IntSlice = h.IntSlice[:lastIdx]
	return val
}

func priorityQueue(nums []int, k int) []int {
	a = nums
	h := &hp{}
	for i := 0; i < k; i++ {
		h.Push(i)
	}
	heap.Init(h)

	result := make([]int, 1, len(nums)-k+1)
	result[0] = nums[h.IntSlice[0]]

	for i := k; i < len(nums); i++ {
		heap.Push(h, i)
		for h.IntSlice[0] <= i- k {
			heap.Pop(h)
		}
		result = append(result, nums[h.IntSlice[0]])
	}
	return result
}
