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
	//return priorityQueue(nums, k)
	return queueSolution(nums, k)
}

func queueSolution(nums []int, k int) []int {

	length := len(nums)
	if k == 1 {
		return nums
	}
	// store  index
	var queue []int
	push := func(i int) {
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	var ans []int
	ans = append(ans, nums[queue[0]])
	for i := k; i < length; i++ {
		push(i)
		if queue[0] <= i-k {
			queue = queue[1:]
		}
		ans = append(ans, nums[queue[0]])
	}
	return ans
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

func (hp *hp) Push(x interface{}) {
	hp.IntSlice = append(hp.IntSlice, x.(int))
}

func (hp hp) Less(i, j int) bool {
	return a[hp.IntSlice[i]] > a[hp.IntSlice[j]]
}

func (hp *hp) Pop() interface{} {
	length := len(hp.IntSlice)
	v := hp.IntSlice[length-1]
	hp.IntSlice = hp.IntSlice[:length-1]
	return v
}

func priorityQueue(nums []int, k int) []int {

	window := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		window.IntSlice[i] = i
	}
	heap.Init(window)

	length := len(nums)
	ans := make([]int, 1, length-k+1)
	ans[0] = nums[window.IntSlice[0]]

	for i := k; i < length; i++ {
		heap.Push(window, i)
		for window.IntSlice[0] <= i-k {
			heap.Pop(window)
		}
		ans = append(ans, nums[window.IntSlice[0]])
	}
	return ans
}
