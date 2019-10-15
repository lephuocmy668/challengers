package lg

// https://leetcode.com/problems/longest-palindromic-substring/submissions/
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	r := []rune(s)
	length := len(r)

	combination := [][]bool{}
	maxLength := 1
	startAt := 0

	for i := 0; i < length; i++ {
		initial := make([]bool, length)
		initial[i] = true
		if i > 0 && r[i] == r[i-1] {
			initial[i-1] = true
			maxLength = 2
			startAt = i - 1
		}
		combination = append(combination, initial)
	}

	for l := 2; l <= length; l++ {
		for i := 0; i < length-l+1; i++ {
			j := i + l - 1
			if r[i] == r[j] && combination[i+1][j-1] {
				combination[i][j] = true
				if l > maxLength {
					maxLength = l
					startAt = i
				}
			}
		}
	}

	return string(r[startAt : startAt+maxLength])
}
