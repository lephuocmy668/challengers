package lps

// https://leetcode.com/problems/longest-palindromic-substring
// Runtime: 124 ms, faster than 27.58% of Go online submissions for Longest Palindromic Substring.
// Memory Usage: 24.5 MB, less than 5.43% of Go online submissions for Longest Palindromic Substring.
func longestPalindrome(s string) string {
	r := []rune(s)
	longest := 1
	start := 0
	memoizations := [][]int{}
	isPalindrome := func(i int, j int) bool {
		if r[i] != r[j] {
			return false
		}
		if i-j == 1 || i-j == 2 {
			return true
		}
		if i < len(s) && j < len(s)-1 {
			if memoizations[i-1][j+1] == 1 {
				return true
			}
		}
		return false
	}

	if len(s) == 0 {
		return ""
	}

	for i := 0; i < len(s); i++ {
		row := make([]int, len(s))
		memoizations = append(memoizations, row)
		for j := 0; j < i; j++ {
			if j == i {
				memoizations[i][j] = 1
				break
			}
			if isPalindrome(i, j) {
				memoizations[i][j] = 1
				l := i - j + 1
				if l > longest {
					longest = l
					start = j
				}
			}
		}
	}

	return string(r[start : start+longest])
}
