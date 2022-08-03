package roman_number

import "errors"

type ListNode struct {
	Val  int
	Next *ListNode
}
type Stack struct {
	top  *ListNode
	size int
}

func (s *Stack) Len() int {
	return s.size
}
func (s *Stack) Push(val int) {
	s.top = &ListNode{Val: val, Next: s.top}
	s.size++
}

func (s *Stack) Pop() (int, error) {
	if s.Len() > 0 {
		value := s.top.Val
		s.top = s.top.Next
		s.size--
		return value, nil
	}
	return 0, errors.New("no more items")
}

func isPalindrome(head *ListNode) bool {
	stack := new(Stack)
	for child := head; child != nil; child = child.Next {
		stack.Push(child.Val)
	}
	for child := head; child != nil; child = child.Next {
		childValue := child.Val
		stackValue, err := stack.Pop()
		if err != nil {
			return false
		}
		if childValue != stackValue {
			return false
		}
	}
	return true
}
