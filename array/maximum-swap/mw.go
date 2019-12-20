package mw

import (
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/maximum-swap/submissions/
func maximumSwap(num int) int {
	sortedDigits := strings.Split(strconv.Itoa(num), "")
	digits := strings.Split(strconv.Itoa(num), "")

	sort.SliceStable(sortedDigits, func(i, j int) bool {
		a, _ := strconv.Atoi(sortedDigits[i])
		b, _ := strconv.Atoi(sortedDigits[j])
		return a > b
	})

parentLoop:
	for i := 0; i < len(sortedDigits); i++ {
		if digits[i] != sortedDigits[i] {
			findDigit := digits[i]
			digits[i] = sortedDigits[i]

			for j := len(sortedDigits) - 1; j > i; j-- {
				if digits[j] == sortedDigits[i] {
					digits[j] = findDigit
					break parentLoop
				}
			}
		}
	}

	stringResult := strings.Join(digits, "")
	result, _ := strconv.Atoi(stringResult)
	return result
}
