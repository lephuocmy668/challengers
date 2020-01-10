package ms

// https://leetcode.com/problems/minimum-swaps-to-make-sequences-increasing
// Runtime: 12 ms, faster than 20.83% of Go online submissions for Minimum Swaps To Make Sequences Increasing.
// Memory Usage: 4.3 MB, less than 100.00% of Go online submissions for Minimum Swaps To Make Sequences Increasing.
func minSwap(A []int, B []int) int {
	const MaxInt = int(^uint(0) >> 1)

	if len(A) < 2 {
		return 0
	}

	n := 0
	k := 1

	for i := 1; i < len(A); i++ {
		tempN := MaxInt
		tempS := MaxInt

		if A[i-1] < A[i] && B[i-1] < B[i] {
			tempN = min(tempN, n)
			tempS = min(tempS, k+1)
		}

		if A[i-1] < B[i] && B[i-1] < A[i] {
			tempN = min(tempN, k)
			tempS = min(tempS, n+1)
		}

		n = tempN
		k = tempS
	}

	return min(n, k)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
