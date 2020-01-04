package maxprofit

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	min := int(^uint(0) >> 1)
	profix := 0

	for _, p := range prices {
		if p < min {
			min = p
			continue
		}
		profix = max(p-min, profix)
	}

	return profix
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
