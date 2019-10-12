package tmn

import (
	"math"
)

func thirdMax(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	firstMaxNum := math.MinInt64
	sencondMaxNum := math.MinInt64
	thirdMaxNum := math.MinInt64

	for _, num := range nums {
		if num > firstMaxNum {
			thirdMaxNum = sencondMaxNum
			sencondMaxNum = firstMaxNum
			firstMaxNum = num
			continue
		}
		if num == firstMaxNum {
			continue
		}

		if num > sencondMaxNum {
			thirdMaxNum = sencondMaxNum
			sencondMaxNum = num
			continue
		}
		if num == sencondMaxNum {
			continue
		}

		if num > thirdMaxNum {
			thirdMaxNum = num
		}
	}

	if thirdMaxNum == math.MinInt64 {
		return firstMaxNum
	}
	return thirdMaxNum
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
