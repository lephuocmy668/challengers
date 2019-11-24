package sr

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/submissions/
// Runtime: 4 ms, faster than 99.56% of Go online submissions for
// Find First and Last Position of Element in Sorted Array.
func searchRange(nums []int, target int) []int {
	start := -1
	end := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			if start == -1 {
				start = i
				if i == len(nums)-1 || (i < len(nums)-1 && nums[i+1] != target) {
					end = i
				}
				continue
			}

			if i == len(nums)-1 {
				if end == -1 {
					end = i
				}
				continue
			}

			if nums[i+1] != target {
				end = i
			}
		}
	}

	return []int{start, end}
}
