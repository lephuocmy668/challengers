package mcs

// https://www.hackerrank.com/challenges/maxsubarray/problem
// Complete the maxSubarray function below.
func maxSubarray(arr []int32) []int32 {
	if len(arr) == 0 {
		return []int32{0, 0}
	}

	m1 := arr[0]
	m2 := arr[0]
	lastM1 := arr[0]

	for i := 1; i < len(arr); i++ {
		var maximumSumsOfSubsequences = arr[i]
		if lastM1 > 0 {
			maximumSumsOfSubsequences = arr[i] + lastM1
		}

		m1 = max(m1, maximumSumsOfSubsequences)
		lastM1 = maximumSumsOfSubsequences

		if arr[i] > 0 {
			m2 += arr[i]
		} else {
			m2 = max(m2, arr[i])
		}
	}

	return []int32{m1, m2}
}

func max(a int32, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
