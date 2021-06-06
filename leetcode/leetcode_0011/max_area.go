/**
* Date: 2021/5/29 7:55 上午
* Author: majunmin
**/
package leetcode_0011

// https://leetcode-cn.com/problems/container-with-most-water/
func maxArea(height []int) int {

	if len(height) < 1 {
		return 0
	}
	left, right := 0, len(height)-1
	var maxArea int
	for left < right {
		maxArea = maxInt(maxArea, minInt(height[left], height[right])*(right-left))

		l, r := left, right
		if height[left] > height[right] {
			for right--; left < right && height[right] <= height[r]; {
				right--
			}
		} else {
			for left++; left < right && height[left] <= height[l]; {
				left++
			}
		}
	}
	return maxArea
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
