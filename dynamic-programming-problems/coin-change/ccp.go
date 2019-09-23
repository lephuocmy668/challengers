package ccp

// https://www.hackerrank.com/challenges/coin-change/problem

/*
 * Complete the 'getWays' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. LONG_INTEGER_ARRAY c
 */
func getWays(n int32, c []int64) int64 {
	// init combinations
	combinations := make([]int64, n+1)
	combinations[0] = 1

	for _, coin := range c {
		for i := coin; i < int64(len(combinations)); i++ {
			if i >= coin {
				combinations[i] += combinations[i-coin]
			}
		}
	}

	return combinations[n]
}
