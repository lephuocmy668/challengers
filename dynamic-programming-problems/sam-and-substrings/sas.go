package sas

import (
	"strconv"
)

// https://www.hackerrank.com/challenges/sam-and-substrings/problem
// Complete the substrings function below.
func substrings(n string) int32 {
	r := []rune(n)
	length := len(r)
	combination := [][]string{}
	sum := int32(0)

	for i := 0; i < length; i++ {
		combination = append(combination, make([]string, length))
		if i == 0 {
			for j := 0; j < length; j++ {
				num, _ := strconv.ParseInt(string(r[j]), 10, 32)
				sum += int32(num)
				combination[i][j] = string(r[length-1-j])
			}
		} else {
			for j := i; j < length; j++ {
				baseStr := combination[0][j-i]
				str := combination[i-1][j]
				currentNum := str + baseStr

				combination[i][j] = currentNum
				num, _ := strconv.ParseInt(currentNum, 10, 32)
				sum += int32(num)
			}
		}

	}

	return sum
}
