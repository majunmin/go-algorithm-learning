/**
  @Author: majm@ushareit.com
  @date: 2021/3/16
  @note:
**/
package sliding_window

import (
	"fmt"
	"testing"
)

func TestLongestSubString(t *testing.T) {
	lengthOfLongestSubstring("pwwkew")
}

func TestMaxSlidingWindow(t *testing.T) {

	fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3))
	fmt.Println(maxSlidingWindow([]int{1,-1}, 1))
}
