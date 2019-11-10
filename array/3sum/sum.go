package sum

import "sort"

// https://leetcode.com/problems/3sum/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums) - 1

		for j < k {
			if nums[i]+nums[j]+nums[k] > 0 {
				k--
				continue
			}
			if nums[i]+nums[j]+nums[k] < 0 {
				j++
				continue
			}

			result = append(result, []int{nums[i], nums[j], nums[k]})
			j++
			k--
			for j < k && nums[j] == nums[j-1] {
				j++
			}
			for k > j && nums[k] == nums[k+1] {
				k--
			}
		}
	}

	return result
}
