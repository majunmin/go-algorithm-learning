/**
  @Author: majm@ushareit.com
  @date: 2021/1/27
  @note:
**/
package common

import "errors"

var StackEmptyErr = errors.New("stack is empty")
var StackFullErr = errors.New("stack is full")

type Stack []interface{}

func (ss Stack) Len() int {
	return len(ss)
}

func (ss Stack) Cap() int {
	return cap(ss)
}

func (ss Stack) IsEmpty() bool {
	return len(ss) == 0
}

func (ss Stack) Top() (interface{}, error) {
	if ss.IsEmpty() {
		return nil, StackEmptyErr
	}
	return ss[0], nil
}

func (ss *Stack) Pop() (interface{}, error) {
	theStack := *ss
	if ss.IsEmpty() {
		return nil, StackEmptyErr
	}
	tail := theStack[len(*ss)-1]
	*ss = theStack[:len(theStack)-1]
	return tail, nil
}

func (ss *Stack) Push(item interface{}) error {
	if ss.Cap() == ss.Len() {
		return StackFullErr
	}
	*ss = append(*ss, item)
	return nil
}
