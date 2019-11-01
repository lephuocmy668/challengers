package rob

// https://leetcode.com/problems/house-robber/submissions/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	for i := 2; i < len(nums); i++ {
		if i == 2 {
			nums[i] += nums[0]
			continue
		}

		nums[i] += max(nums[i-3], nums[i-2])

	}
	return max(nums[len(nums)-2], nums[len(nums)-1])
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
