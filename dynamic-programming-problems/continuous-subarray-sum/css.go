package css

// https://leetcode.com/problems/continuous-subarray-sum/
// solve by momoization
func checkSubarraySum(nums []int, k int) bool {
	length := len(nums)
	combinations := [][]int{}

	for i := 1; i <= length; i++ {
		item := []int{}
		for j := 0; j < length; j++ {
			addItem := 1
			if i == 1 {
				addItem = nums[j]
			} else {
				if j < i-1 {
					addItem = combinations[i-2][j]

				} else {
					addItem = combinations[i-2][j-1] + nums[j]
				}
			}

			item = append(item, addItem)
			if i < 2 || j < 1 {
				continue
			}

			if (addItem == 0 && k == 0) || (k != 0 && addItem%k == 0) {
				return true
			}
		}
		combinations = append(combinations, item)
	}
	return false
}
