package ssek

// https://leetcode.com/problems/subarray-sum-equals-k
func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	m[0] = 1
	sum := 0
	count := 0

	for _, v := range nums {
		sum += v
		if m[sum-k] > 0 {
			count += m[sum-k]
		}
		m[sum]++
	}

	return count
}
