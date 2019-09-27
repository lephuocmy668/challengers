package lcs

// https://leetcode.com/problems/longest-common-subsequence/submissions/
func longestCommonSubsequence(text1 string, text2 string) int {
	r1 := []rune(text1)
	r2 := []rune(text2)
	combinations := [][]int{}
	for i := 0; i < len(r1); i++ {
		combinations = append(combinations, make([]int, len(r2)))
	}

	for i := 0; i < len(r1); i++ {
		for j := 0; j < len(r2); j++ {
			if r2[j] == r1[i] {
				if i > 0 && j > 0 {
					combinations[i][j] = combinations[i-1][j-1] + 1
					continue
				}

				combinations[i][j] = 1
				continue
			}

			if i > 0 {
				if j > 0 {
					combinations[i][j] = max(combinations[i][j-1], combinations[i-1][j])
					continue
				}
				combinations[i][j] = max(0, combinations[i-1][j])
			} else {
				if j > 0 {
					combinations[i][j] = max(combinations[i][j-1], 0)
					continue
				}
				combinations[i][j] = 0
			}
		}
	}

	return combinations[len(r1)-1][len(r2)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
