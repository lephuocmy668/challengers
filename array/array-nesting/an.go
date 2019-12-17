package an

// https://leetcode.com/problems/array-nesting/submissions/
// Runtime: 16 ms, faster than 75.76% of Go online submissions for Array Nesting.
func arrayNesting(nums []int) int {
	lenth := len(nums)
	m := 0
	trackingArr := make([]bool, lenth)

	for i := 0; i < lenth; i++ {
		m = max(m, find(trackingArr, nums, i))
	}

	return m
}

func find(trackingArr []bool, nums []int, index int) int {
	count := 0
	for index < len(nums) && index > -1 && !trackingArr[index] {
		count++
		trackingArr[index] = true
		index = nums[index]
	}
	return count
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
