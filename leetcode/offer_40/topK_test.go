/**
  @Author: majm@ushareit.com
  @date: 2021/1/16
  @note:
**/
package offer_40

import (
	"fmt"
	"testing"
)

func TestTopK(t *testing.T) {
	arr := []int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}

	fmt.Println(getLeastNumbers(arr, 8))
}
