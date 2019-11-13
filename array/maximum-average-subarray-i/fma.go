package fma

// https://leetcode.com/problems/maximum-average-subarray-i/submissions/
func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < k {
		return 0
	}

	start := 0
	m := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		if i < k {
			m += nums[i]
			sum += nums[i]
			continue
		}
		sum = sum + nums[i] - nums[start]
		start++
		m = max(m, sum)
	}

	return float64(m) / float64(k)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
