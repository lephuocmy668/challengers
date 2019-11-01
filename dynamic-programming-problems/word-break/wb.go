package wb

// https://leetcode.com/problems/word-break/submissions/
// recursive solution
// func wordBreak(s string, wordDict []string) bool {
// 	ok := false

// 	for _, word := range wordDict {
// 		if strings.HasPrefix(s, word) {
// 			if s == word {
// 				return true
// 			}
// 			if wordBreak(strings.TrimPrefix(s, word), wordDict) {
// 				ok = true
// 			}
// 		}
// 	}

// 	return ok
// }

// memoization solution
func wordBreak(s string, wordDict []string) bool {
	runStr := []rune(s)
	combinations := make([]bool, len(s)+1)
	combinations[0] = true

	dictMap := make(map[string]bool)
	for _, word := range wordDict {
		dictMap[word] = true
	}

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if _, ok := dictMap[string(runStr[j:i])]; ok && combinations[j] {
				combinations[i] = true
			}
		}
	}

	return combinations[len(s)]
}
