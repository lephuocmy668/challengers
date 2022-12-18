// https://leetcode.com/problems/maximum-subarray/submissions/861514416/
// Runtime
// 96 ms
// Beats
// 94.90%
// Memory
// 9.6 MB
// Beats
// 45.13%
package ms

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var maxVal, currMax int = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currMax = max(nums[i], currMax+nums[i])
		maxVal = max(currMax, maxVal)
	}

	return maxVal
}
