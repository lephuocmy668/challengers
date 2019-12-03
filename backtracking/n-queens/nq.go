package main

import (
	"fmt"
	"time"
)

func main() {
	solveNQueens(4)
	// s := newStack()
	// s.push(0)
	// s.push(2)
	// fmt.Println("=============", isSafe(s, row, i))
}

func solveNQueens(n int) [][]string {
	currentPosition := 0
	s := newStack()
	s.push(0)

	for !s.isEmpty() {
		row := s.size()

		validPosition := -1
	l:
		for i := currentPosition; i < n; i++ {
			if isSafe(s, row, i) {
				validPosition = i
				break l
			}
		}

		fmt.Println("==========-1========", validPosition, s.values)

		if validPosition > 0 {
			s.push(validPosition)
			currentPosition = 0
			fmt.Println("=======1===========", s.values)
		} else {
			currentPosition = s.pop()
			fmt.Println("=======2===========", currentPosition)
			if s.isEmpty() {
				if currentPosition == n {
					return nil
				}
				s.push(currentPosition)
			} else {
				currentPosition++
			}
		}
		time.Sleep(1 * time.Second)
	}

	return nil
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

// Create empty stack and set current position to 0

// Repeat {

//    loop from current position to the last position until valid position found //current row

//    if there is a valid position {
//       push the position to stack, set current position to 0 // move to next row
//    }

//    if there is no valid position {
//       if stack is empty, break // stop search
//       else pop stack, set current position to next position // backtrack to previous row
//    }

//    if stack has size N { // a solution is found
//       pop stack, set current position to next position // backtrack to find next solution
//    }

// }
