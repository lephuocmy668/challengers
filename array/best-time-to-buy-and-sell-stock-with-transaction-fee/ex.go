package maxprofit

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee
func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}

	min := prices[0]
	max := prices[0]
	res := 0

	for _, p := range prices {
		if p < min || p+fee < max {
			profix := max - min - fee
			if profix > 0 {
				res += profix
			}
			max = p
			min = p
			continue
		}
		if p > max {
			max = p
		}
	}

	return res + (getMax(0, max-min-fee))
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
