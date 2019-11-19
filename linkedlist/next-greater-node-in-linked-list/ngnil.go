package ngnil

import (
	"github.com/lephuocmy668/challengers/linkedlist"
)

// https://leetcode.com/problems/next-greater-node-in-linked-list/
func nextLargerNodes(head *linkedlist.ListNode) []int {
	temp := 0
	result := []int{}
	arr := []int{}
	stack := newStack()
	node := head

	for node != nil {
		for !stack.isEmpty() && node.Val > arr[stack.peek()] {
			index := stack.pop()
			result[index] = node.Val
		}

		result = append(result, 0)
		arr = append(arr, node.Val)
		stack.push(temp)
		temp++

		node = node.Next
	}

	for !stack.isEmpty() && arr[len(arr)-1] > arr[stack.peek()] {
		index := stack.pop()
		result[index] = arr[len(arr)-1]
	}

	return result
}

type stack struct {
	values []int
}

func newStack() *stack {
	return &stack{
		values: []int{},
	}
}

func (s *stack) isEmpty() bool {
	return len(s.values) == 0
}

func (s *stack) push(num int) {
	s.values = append(s.values, num)
}

func (s *stack) pop() int {
	result := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return result
}

func (s *stack) peek() int {
	return s.values[len(s.values)-1]
}
