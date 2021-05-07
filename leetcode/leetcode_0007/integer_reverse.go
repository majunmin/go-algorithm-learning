/**
* Date: 2021/5/3 5:23 下午
* Author: majunmin
**/
package leetcode_0007

import "math"

func reverse(x int) int {

	res := 0
	for x != 0 {
		res = res*10 + x%10
		x = x / 10
		if res < math.MinInt32 || res > math.MaxInt32 {
			return 0
		}
	}
	return res
}
