package kcms

// https://leetcode.com/problems/k-concatenation-maximum-sum/
func kConcatenationMaxSum(arr []int, k int) int {
	totalSum := 0

	localMaxSum := 0
	localCurSum := 0

	maxSumLeft := 0
	curSumLeft := 0

	maxSumRight := 0
	curSumRight := 0

	for i := 0; i < len(arr); i++ {
		curSumLeft += arr[i]
		maxSumLeft = max(curSumLeft, maxSumLeft)

		curSumRight += arr[len(arr)-i-1]
		maxSumRight = max(curSumRight, maxSumRight)

		totalSum += arr[i]

		localCurSum = max(localCurSum+arr[i], 0)
		localMaxSum = max(localCurSum, localMaxSum)
	}

	return max(max(max(maxSumLeft+maxSumRight, (totalSum*max(k-2, 0))+maxSumLeft+maxSumRight), localMaxSum), totalSum*k)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
