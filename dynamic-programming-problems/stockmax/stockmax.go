package stockmax

// https://www.hackerrank.com/challenges/stockmax

/*
 * Complete the 'stockmax' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER_ARRAY prices as parameter.
 */

func stockmax(prices []int32) int64 {
	// Write your code here
	profit := int64(0)
	maxSoFar := int32(0)

	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] >= maxSoFar {
			maxSoFar = prices[i]
		}
		profit += int64(maxSoFar - prices[i])
	}
	return profit

}
