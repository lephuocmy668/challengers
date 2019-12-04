package main

// https://leetcode.com/problems/n-queens/submissions/
// Runtime: 0 ms, faster than 100.00% of Go online submissions for N-Queens.
func solveNQueens(n int) [][]string {
	res := [][]string{}
	currentPosition := 0 // position is considering of queen at each column
	s := newStack()
	s.push(0)

	for !s.isEmpty() {
		row := s.size()
		if row == n {
			res = append(res, s.toResult())
		}

		// init valid position = -1
		validPosition := -1
	childLoop:
		for i := currentPosition; i < n; i++ {
			// if find out valid position, break child loop
			if isSafe(s, row, i) {
				validPosition = i
				break childLoop
			}
		}

		// if there is valid position, push to stack and continue
		if validPosition >= 0 {
			s.push(validPosition)
			currentPosition = 0
			continue
		}

		// if there is no valid position, return previous row and set next column()
		currentPosition = s.pop()
		if !s.isEmpty() {
			currentPosition++
			continue
		}

		// if return to first row and considering column is last column, return
		if currentPosition == n-1 {
			return res
		}
		currentPosition++
		s.push(currentPosition)
		currentPosition = 0
	}

	return res
}

func isSafe(board *stack, row, col int) bool {
	for i, e := range board.values {
		if e == col || i-row == e-col || i-row == col-e {
			return false
		}
	}

	return true
}

// stack is special stack that was implemented for this problem
type stack struct {
	values []int
}

func newStack() *stack {
	return &stack{
		values: []int{},
	}
}

// toResult convert stack elements to result of this problem
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

func (s *stack) size() int {
	return len(s.values)
}
