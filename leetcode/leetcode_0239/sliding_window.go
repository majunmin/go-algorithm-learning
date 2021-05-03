/**
  @Author: majm@ushareit.com
  @date: 2021/2/20
  @note:
**/
package leetcode_0239

import (
	"container/heap"
	"sort"
)

// https://leetcode-cn.com/problems/sliding-window-maximum/
func maxSlidingWindow(nums []int, k int) []int {
	length := len(nums)
	if length == 0 {
		return nil
	}
	if k == 0 || length < k {
		return nil
	}

	result := make([]int, 1, length-k+1)
	queue := []int{}

	push := func(i int) {
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	result[0] = nums[queue[0]]
	for i := k; i < length; i++ {
		push(i)
		// 移除 滑动窗口外的的元素
		for queue[0] < i-k+1 {
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

func (h hp) Less(i, j int) bool {
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h hp) Pop() interface{} {
	a = h.IntSlice
	v := h.IntSlice[h.Len()-1]
	h.IntSlice = a[:h.Len()-1]
	return v
}

// heap 中 存放索引, 索引的排序
func priorityQueue(nums []int, k int) []int {
	a = nums
	length := len(a)
	result := make([]int, 1, length-k+1)

	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	result[0] = nums[q.IntSlice[0]]
	for i := k; i < length; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		result = append(result, nums[q.IntSlice[0]])
	}

	return result
}
