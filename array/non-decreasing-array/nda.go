package nda

// https://leetcode.com/problems/non-decreasing-array/
func checkPossibility(nums []int) bool {
	var violence bool

	i := 0
	for i < len(nums)-1 {
		if nums[i+1] < nums[i] {
			if violence {
				return false
			} else {
				if i > 0 && i <= len(nums)-3 && nums[i+2] < nums[i] && nums[i+1] < nums[i-1] {
					return false
				}
				violence = true
			}
		}
		i++
	}
	return true
}
