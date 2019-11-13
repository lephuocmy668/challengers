package ssm

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

// https://leetcode.com/problems/sum-of-subarray-minimums/
func sumSubarrayMins(A []int) int {
	result := 0
	mod := int(1e9) + 7

	s := newStack()
	for i := 0; i <= len(A); i++ {
		for !s.isEmpty() && (i == len(A) || A[i] < A[s.peek()]) {
			current := s.pop()

			prev := -1
			if !s.isEmpty() {
				prev = s.peek()
			}

			currentSum := A[current] * (current - prev) * (i - current)
			result = (result + (currentSum % mod)) % mod
		}
		if i < len(A) {
			s.push(i)
		}
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Simple approach
// // https://leetcode.com/problems/sum-of-subarray-minimums/
// func sumSubarrayMins(A []int) int {
// 	result := 0
// 	mod := int(1e9) + 7

// 	for i := 0; i < len(A); i++ {
// 		result += A[i]
// 		if i == len(A)-1 {
// 			break
// 		}

// 		m := A[i]
// 		for j := i + 1; j < len(A); j++ {
// 			m = min(m, A[j])
// 			result = (result + m) % mod
// 		}
// 	}
// 	return int(result)
// }

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }
