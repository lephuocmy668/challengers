package ms

// https://leetcode.com/problems/maximal-square/submissions/
func maximalSquare(matrix [][]byte) int {
	ms := 0

	if len(matrix) == 0 {
		return ms
	}

	// init combinations
	combinations := [][]int{}
	for i := 0; i < len(matrix); i++ {
		combinations = append(combinations, make([]int, len(matrix[0])))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 || j == 0 {
				if string(matrix[i][j]) == "0" {
					combinations[i][j] = 0
				} else {
					combinations[i][j] = 1
				}
			} else if string(matrix[i][j]) == "1" {
				combinations[i][j] = min(min(combinations[i-1][j-1], combinations[i][j-1]), combinations[i-1][j]) + 1
			}
			ms = max(ms, combinations[i][j])
		}
	}

	return ms * ms
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
