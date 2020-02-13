package ts

// https://leetcode.com/problems/target-sum/
// Runtime: 680 ms, faster than 25.93% of Go online submissions for Target Sum.
// Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Target Sum.
// Next challenges:
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

func findTargetSumWays(nums []int, target int) int {
	return dp(nums, 0, 0, target)
}
