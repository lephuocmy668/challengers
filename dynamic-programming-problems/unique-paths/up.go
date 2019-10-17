package up

// https://leetcode.com/problems/unique-paths/submissions/
func uniquePaths(m int, n int) int {
	combination := [][]int{}

	for i := 0; i < n; i++ {
		combination = append(combination, make([]int, m))

		for j := 0; j < m; j++ {
			if j == 0 || i == 0 {
				combination[i][j] = 1
				continue
			}
			combination[i][j] = combination[i-1][j] + combination[i][j-1]
		}
	}

	return combination[n-1][m-1]
}

// recursive solution
// func uniquePaths(m int, n int) int {
// 	if m == 1 && n == 1 {
// 		return 1
// 	} else {
// 		if m == 1 || n == 1 {
// 			return 1
// 		}
// 		return uniquePaths(m-1, n) + uniquePaths(m, n-1)
// 	}
// }
