package nthdigit

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/nth-digit
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Nth Digit.
// Memory Usage: 2 MB, less than 16.67% of Go online submissions for Nth Digit.
func findNthDigit(n int) int {
	var m, start, length, count int64 = int64(n), 1, 1, 9

	// get the number that contain digit we finding
	for m > length*count {
		m -= length * count
		count *= 10
		start *= 10
		length++
	}
	num := start + (m-1)/length
	numStr := fmt.Sprintf("%d", num)

	// in numStr, get index of digit we finding
	index := (m - 1) % length

	char := string([]rune(numStr)[int(index)])
	if result, err := strconv.Atoi(char); err == nil {
		return result
	}
	return 0
}
