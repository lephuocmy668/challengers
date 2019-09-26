package lis

// https://leetcode.com/problems/longest-increasing-subsequence/submissions/
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	maxLength := 0
	combinations := make([]int, len(nums))
	combinations[0] = 1

	for i := 1; i < len(combinations); i++ {
		temp := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				temp = max(temp, combinations[j]+1)
			}
		}
		combinations[i] = temp
		maxLength = max(maxLength, temp)
	}
	return maxLength
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
