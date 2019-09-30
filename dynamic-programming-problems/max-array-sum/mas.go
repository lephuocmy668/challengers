package mas

func max(a int32, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

// Complete the maxSubsetSum function below.
// https://www.hackerrank.com/challenges/max-array-sum/problem
func maxSubsetSum(arr []int32) int32 {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	if len(arr) == 2 {
		return max(arr[0], arr[1])
	}

	if len(arr) == 3 {
		return max(max(arr[0], arr[1]), arr[2])
	}

	combination := make([]int32, len(arr))
	combination[0] = arr[0]
	combination[1] = arr[1]
	combination[2] = max(max(arr[0], arr[2]), arr[0]+arr[2])

	for i := 3; i < len(arr); i++ {
		lastMax := max(combination[i-3], combination[i-2])
		combination[i] = max(max(lastMax, arr[i]), arr[i]+lastMax)
	}

	return max(combination[len(arr)-1], combination[len(arr)-2])
}
