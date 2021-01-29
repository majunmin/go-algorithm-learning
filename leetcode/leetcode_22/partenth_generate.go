/**
  @Author: majm@ushareit.com
  @date: 2021/1/27
  @note:
**/
package leetcode_22

import "container/list"

func generateParenthesis(n int) []string {
	var result []string

	//dfs(n, 0, 0, []rune{}, &result)
	bfs(n, &result)
	return result
}

// 递归
func dfs(n, left, right int, curPath []rune, result *[]string) {
	if len(curPath) == n*2 {
		*result = append(*result, string(curPath))
		return
	}

	if left < n {
		dfs(n, left+1, right, append(curPath, '('), result)
	}
	if right < left {
		dfs(n, left, right+1, append(curPath, ')'), result)
	}
}

type Node struct {
	path  string
	left  int
	right int
}

// 栈实现 递归
func bfs(n int, result *[]string) {
	queue := list.New()
	node := &Node{path: "("}

	queue.PushBack(node)
	node.left++

	for queue.Len() != 0 {
		ele := queue.Front()
		queue.Remove(ele)
		curNode := ele.Value.(*Node)
		if curNode.left == n && curNode.right == n {
			*result = append(*result, curNode.path)
			continue
		}

		if curNode.left < n {
			lNode := &Node{
				path:  curNode.path + "(",
				left:  curNode.left + 1,
				right: curNode.right,
			}
			queue.PushBack(lNode)
		}

		if curNode.right < curNode.left {
			rNode := &Node{
				path:  curNode.path + ")",
				left:  curNode.left,
				right: curNode.right + 1,
			}
			queue.PushBack(rNode)
		}
	}
}
