package rmdfsa

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	count := 1
	currentNum := nums[0]
	currentNumCount := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == currentNum {
			currentNumCount += 1
			if currentNumCount < 3 {
				count++
			}
		} else {
			currentNum = nums[i]
			currentNumCount = 1
			count++
		}
	}

	return count
}
