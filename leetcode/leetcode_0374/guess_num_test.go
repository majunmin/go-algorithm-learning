/**
  @Author: majm@ushareit.com
  @date: 2021/2/9
  @note:
**/
package leetcode_0374

import (
	"fmt"
	"math"
	"testing"
)

func TestDmo(t *testing.T) {
	num := int8(math.MinInt8)
	fmt.Println(num)
	fmt.Println(num >> 3)
	fmt.Printf("%X\n", num)
	fmt.Printf("%b  , %d \n", num+1, num+1)
	fmt.Printf("%b  , %d \n", 127, 127)

	fmt.Printf("%X\n", num+2)
	fmt.Printf("%X\n", num>>1)
	fmt.Printf("%X\n", num>>3)

	//num2 := uint32(0xFFFFFFFF)
	//fmt.Println(num2)
	//fmt.Println(num2>>3)
	//fmt.Printf("%X\n", num2>>3)

}

func TestDemo2(t *testing.T) {

	fmt.Println(guessNumber(1))

}
