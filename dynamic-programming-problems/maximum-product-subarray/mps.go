package mps

import (
	"math"
)

// https://leetcode.com/problems/maximum-product-subarray/
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := math.MinInt16
	var am, duong, tempAm, tempDuong int

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			am = nums[i]
			duong = nums[i]
			result = max(result, duong)
			continue
		}

		tempAm = am
		tempDuong = duong
		duong = max(nums[i], max(tempAm*nums[i], tempDuong*nums[i]))
		am = min(nums[i], min(tempAm*nums[i], tempDuong*nums[i]))
		result = max(result, duong)
	}

	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
