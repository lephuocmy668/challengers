package rmdfsa

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
// Runtime: 4 ms, faster than 89.32% of Go online submissions for Remove Duplicates from Sorted Array II.
// Memory Usage: 3 MB, less than 100.00% of Go online submissions for Remove Duplicates from Sorted Array II.
func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	index := 0
	count := 0
	currentNum := 0
	currentNumCount := 0

	for _, num := range nums {
		if num == currentNum {
			if currentNumCount < 2 {
				count++
				currentNumCount++
				nums[index] = currentNum
				index++
			}
		} else {
			count++
			currentNum = num
			currentNumCount = 1
			nums[index] = currentNum
			index++
		}
	}

	return count
}
