/**
* Date: 2021/5/7 1:32 下午
* Author: majunmin
**/
package interview_08_06

// https://leetcode-cn.com/problems/hanota-lcci/
//
// 输入：A = [2, 1, 0], B = [], C = []
// 输出：C = [2, 1, 0]
//
func hanota(A []int, B []int, C []int) []int {
	n := len(A)
	if n == 0 {
		return nil
	}
	move(n, &A, &B, &C)
	return C
}

/**
1.定义问题的递归函数，明确函数的功能,我们定义这个函数的功能为：把 A 上面的 n 个圆盘经由 B 移到 C
2.关系公式：move(n from A to C) = move(n-1 from A to B) + move(A to C) + move(n-1 from B to C`)
分析可得，我们分三步走。
(1)先把A上的n-1个圆盘，通过和C的操作移动到B上。
(2)把A剩下的最大的一个盘移动到C。
(3)再把B上的n-1个盘，通过和A的操作移动到C上。
不管怎么移动问题和子问题的关系都可以是这三步，下面就可以开始递归。

链接：https://leetcode-cn.com/problems/hanota-lcci/solution/zhi-qian-yi-zhi-tiao-guo-yi-nuo-ta-zhe-ci-zhong-yu/
*/
func move(n int, A *[]int, B *[]int, C *[]int) {
	// terminate condition
	if n == 1 {
		// move A -> C
		*C = append(*C, (*A)[len(*A)-1])
		*A = (*A)[:len(*A)-1]
		return
	}
	move(n-1, A, C, B)
	move(1, A, B, C)
	move(n-1, B, A, C)
}
