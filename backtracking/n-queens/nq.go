package main

import "fmt"

func main() {
	solveNQueens(5)
}

func solveNQueens(n int) [][]string {
	res := [][]string{}
	currentPosition := 0
	s := newStack()
	s.push(0)

	for !s.isEmpty() {
		row := s.size()
		if row == n {
			fmt.Println(s.values)
			res = append(res, s.toResult())
		}

		validPosition := -1
		for i := currentPosition; i < n; i++ {
			if isSafe(s, row, i) {
				validPosition = i
			}
		}

		if validPosition >= 0 {
			s.push(validPosition)
			currentPosition = 0
		} else {
			currentPosition = s.pop()
			if s.isEmpty() {
				if currentPosition == n-1 {
					return res
				}
				currentPosition++
				s.push(currentPosition)
				currentPosition = 0
			} else {
				currentPosition++
			}
		}
	}

	return res
}

func isSafe(board *stack, row, col int) bool {
	if board.size() == row+1 {
		return false
	}

	for i, e := range board.values {
		if e == col || i-row == e-col || i-row == col-e {
			return false
		}
	}

	return true
}

type stack struct {
	values []int
}

func newStack() *stack {
	return &stack{
		values: []int{},
	}
}

func (s *stack) toResult() []string {
	res := []string{}
	for i := 0; i < s.size(); i++ {
		str := ""
		for j := 0; j < s.size(); j++ {
			if j == s.values[i] {
				str += "Q"
			} else {
				str += "."
			}
		}
		res = append(res, str)
	}
	return res
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

func (s *stack) size() int {
	return len(s.values)
}

func (s *stack) get(index int) int {
	return s.values[index]
}
