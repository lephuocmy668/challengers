package mps

// https://leetcode.com/problems/minimum-path-sum/submissions/
func minPathSum(grid [][]int) (result int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 {
				if j > 0 {
					grid[i][j] += grid[i][j-1]
				}
			} else {
				if j == 0 {
					grid[i][j] += grid[i-1][j]
				} else {
					grid[i][j] += min(grid[i-1][j], grid[i][j-1])
				}
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
