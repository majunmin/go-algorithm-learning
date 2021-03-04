/**
  @Author: majm@ushareit.com
  @date: 2021/3/1
  @note:
**/
package leetcode_0303

// 本题考察的是前缀和
// 缓存前面的计算
type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	numsArr := NumArray{
		preSum: make([]int, len(nums)+1),
	}

	for i := 0; i < len(nums); i++ {
		numsArr.preSum[i+1] = numsArr.preSum[i] + nums[i]
	}
	return numsArr
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.preSum[j+1] - this.preSum[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
