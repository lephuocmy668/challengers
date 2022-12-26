// Best Time to Buy and Sell Stock II
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
package solution

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) (result int) {
	if len(prices) < 2 {
		return
	}

	for i := 1; i < len(prices); i++ {
		result += max(0, prices[i]-prices[i-1])
	}

	return
}
