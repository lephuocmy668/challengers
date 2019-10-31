package up

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	for i := 0; i < len(obstacleGrid); i++ {
		item := obstacleGrid[i]
		for j := 0; j < len(item); j++ {
			if item[j] == 1 {
				item[j] = -1
				continue
			}
			if i == 0 && j == 0 {
				item[j] = 1
				if i == len(obstacleGrid)-1 && j == len(item)-1 {
					return returnResult(item[j])
				}
				continue
			}

			if i == 0 {
				if item[j-1] == -1 {
					item[j] = -1
				} else {
					item[j] = 1
				}
				if i == len(obstacleGrid)-1 && j == len(item)-1 {
					return returnResult(item[j])
				}
				continue
			}

			if j == 0 {
				if obstacleGrid[i-1][j] == -1 {
					item[j] = -1
				} else {
					item[j] = 1
				}
				if i == len(obstacleGrid)-1 && j == len(item)-1 {
					return returnResult(item[j])
				}
				continue
			}

			if item[j-1] != -1 {
				item[j] += item[j-1]
			}

			if obstacleGrid[i-1][j] != -1 {
				item[j] += obstacleGrid[i-1][j]
			}

			if i == len(obstacleGrid)-1 && j == len(item)-1 {
				return returnResult(item[j])
			}
		}
	}
	return 0
}

func returnResult(res int) int {
	if res == -1 {
		return 0
	}
	return res
}
