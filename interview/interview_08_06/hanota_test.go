/**
* Date: 2021/5/7 4:00 下午
* Author: majunmin
**/
package interview_08_06

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFunc(t *testing.T) {
	a := int64(888)
	ratio := 10000

	fmt.Printf("%f \n", float64(a)/10*float64(ratio))

	round, _ := Round(float64(a)/10*float64(ratio), "2")
	fmt.Println(round)

}

// 四舍五入
func Round(n float64, pre string) (float64, error) {
	s := fmt.Sprintf("%."+pre+"f", n)
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func Test_hanota(t *testing.T) {
	A := []int{2, 1, 0}
	B := []int{}
	C := []int{}
	hanota(A, B, C)
}
