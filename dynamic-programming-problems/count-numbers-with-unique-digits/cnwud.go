package cnwud

func countByNumDigit(n int) int {
	if n == 1 {
		return 10
	}

	count := 9
	digits := 9
	for i := 2; i <= n; i++ {
		count = count * digits
		digits--
	}

	return count
}

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}

	result := 0
	for i := 1; i <= n; i++ {
		result += countByNumDigit(i)
	}
	return result
}
