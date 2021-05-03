/**
  @Author: majm@ushareit.com
  @date: 2021/3/8
  @note:
**/
package leetcode_0322

import "math"

var cache = make(map[int]int)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
	}

	for i := 0; i < amount+1; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i-coins[j]], dp[i]) + 1
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]

}

func extChange(coins []int, amount int) int {
	if val, ok := cache[amount]; ok {
		return val
	}

	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	var result = math.MaxInt32
	for _, coin := range coins {
		subMin := extChange(coins, amount-coin)
		if subMin == -1 {
			continue
		}
		result = min(subMin+1, result)
	}
	// 没有符合问题的解
	if result == math.MaxInt32 {
		result = -1
	}
	cache[amount] = result
	return result
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
