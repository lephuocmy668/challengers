package lds

import (
	"sort"
)

// https://leetcode.com/problems/largest-divisible-subset
// Runtime: 44 ms, faster than 36.00% of Go online submissions for Largest Divisible Subset.
// Memory Usage: 3.9 MB, less than 100.00% of Go online submissions for Largest Divisible Subset.
func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	if len(nums) == 1 {
		return nums
	}

	var maxLength int
	result := []int{nums[0]}
	combination := [][]int{[]int{nums[0]}}

	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		length := 1
		combination = append(combination, []int{nums[i]})

		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				if len(combination[j])+1 > length {
					length = len(combination[j]) + 1
					combination[i] = append(combination[j], nums[i])
					if len(combination[i]) > maxLength {
						maxLength = len(combination[i])
						result = combination[i]
					}
				}
			}
		}
	}

	return result
}
