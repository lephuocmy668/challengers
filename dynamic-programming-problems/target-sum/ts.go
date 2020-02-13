package ts

func dp(nums []int, currentPosition int, currentSum int, target int) int {
	if currentPosition > len(nums) {
		return 0
	}

	if currentPosition == len(nums) {
		if currentSum == target {
			return 1
		}
		return 0
	}

	count := dp(nums, currentPosition+1, currentSum+nums[currentPosition], target)
	count += dp(nums, currentPosition+1, currentSum-nums[currentPosition], target)
	return count
}

// https://leetcode.com/problems/target-sum/
func findTargetSumWays(nums []int, target int) int {
	return dp(nums, 0, 0, target)
}
