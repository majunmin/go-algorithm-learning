/**
* Date: 2021/5/11 7:46 上午
* Author: majunmin
**/
package leetcode_0050

// x -> x^2  -> x^4 -> x^8*
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return myPow(1/x, n)
	}
	if n&1 == 1 {
		return myPow(x, n/2) * x
	} else {
		return myPow(x, n/2)
	}
	//if n >= 0 {
	//	return quickMul1(x, n)
	//}
	//return 1 / quickMul1(x, -n)
}

//时间复杂度 O(logN) --> 递归的层数
//空间复杂度 O(logN) --> 递归的层数, 递归的函数会使用栈空间
// 快速幂 + 递归
func quickMul1(x float64, n int) float64 {
	// recursion terminator
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	//
	y := quickMul1(x, n/2)
	if n&1 == 1 {
		return y * y * x
	}
	return y * y
}
