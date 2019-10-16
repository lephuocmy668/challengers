package dw

import (
	"strconv"
	"strings"
)

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	if strings.HasPrefix(s, "0") {
		return 0
	}

	r := []rune(s)
	combination := make([]int, len(r))
	combination[0] = 1

	for i := 1; i < len(r); i++ {
		lastNum, _ := strconv.Atoi(string(r[i]))
		lastTwoNum, _ := strconv.Atoi(string(r[i-1]) + string(r[i]))
		if lastNum > 0 {
			combination[i] += combination[i-1]
		}
		if lastTwoNum > 9 && lastTwoNum < 27 {
			if i > 1 {
				combination[i] += combination[i-2]
			} else {
				combination[i]++
			}
		}
	}

	return combination[len(r)-1]
}
