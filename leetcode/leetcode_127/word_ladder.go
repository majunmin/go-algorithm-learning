/**
  @Author: majm@ushareit.com
  @date: 2021/1/30
  @note:
**/
package leetcode_127

// 1. BFS
// 2. DFS
// 3. Two Ended BFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
	//return bfsSolution(beginWord, endWord, wordList)
	return twoEndBFSSolution(beginWord, endWord, wordList)
}

func twoEndBFSSolution(beginWord string, endWord string, wordList []string) int {
	n := len(wordList)
	if n == 0 {
		return 0
	}

	beginQueue := make([]string, 0)
	endQueue := make([]string, 0)
	beginQueue = append(beginQueue, beginWord)
	endQueue = append(endQueue, endWord)
	startUsed := make([]bool, len(wordList))
	endUsed := make([]bool, len(wordList))

	var idx int
	if idx = IdxOf(endWord, wordList); idx == -1 {
		return 0
	} else {
		endUsed[idx] = true
	}

	step := 1
	for len(beginQueue) > 0 {
		// 选择长度较小的 长度开始遍历，减少复杂度
		if len(beginQueue) > len(endQueue) {
			beginQueue, endQueue = endQueue, beginQueue
			startUsed, endUsed = endUsed, startUsed
		}
		size := len(beginQueue)
		for i := 0; i < size; i++ {
			curWord := beginQueue[i]
			// terminate
			// return step if curWord in endQueue
			if IdxOf(curWord, endQueue) != -1 {
				return step
			}
			for k, word := range wordList {
				if !startUsed[k] && IsDiffOne(curWord, word) {
					beginQueue = append(beginQueue, word)
					startUsed[k] = true
				}
			}
		}
		beginQueue = beginQueue[size:]
		step++
	}
	return 0
}

func bfsSolution(beginWord string, endWord string, wordList []string) int {
	n := len(wordList)
	queue := make([]string, 0)
	queue = append(queue, beginWord)
	//visited := make(map[string]struct{}, n)
	visited := make([]bool, n)

	// sterp init = 1, if beginWord == endWord
	step := 1
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curWord := queue[i]
			if curWord == endWord {
				return step
			}

			for j, word := range wordList {
				// 如果已经被访问过 那么抛弃
				if !visited[j] && IsDiffOne(curWord, word) {
					queue = append(queue, word)
					visited[j] = true
				}
			}
		}
		queue = queue[size:]
		step++
	}
	return 0
}

// 如果 单词很长，那么时间复杂度急速升高
func IsDiffOne(curWord string, word string) bool {
	diffCnt := 0
	curByts := []byte(curWord)
	byts := []byte(word)
	for i := 0; i < len(curByts); i++ {
		if curByts[i]-byts[i] != 0 {
			diffCnt++
		}
		if diffCnt > 1 {
			return false
		}
	}
	return diffCnt == 1
}

func IdxOf(word string, wordList []string) int {
	for i, s := range wordList {
		if s == word {
			return i
		}
	}
	return -1
}
