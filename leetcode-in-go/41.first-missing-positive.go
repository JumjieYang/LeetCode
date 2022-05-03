/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	one := false
	abs := func(x int) int {
		return int(math.Abs(float64(x)))
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			one = true
			break
		}
	}

	if !one {
		return 1
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 || nums[i] > len(nums) {
			nums[i] = 1
		}
	}

	for i := 0; i < len(nums); i++ {
		a := abs(nums[i])

		if a == len(nums) {
			nums[0] = -abs(nums[0])
		} else {
			nums[a] = -abs(nums[a])
		}
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] > 0 {
		return len(nums)
	}

	return len(nums) + 1
}

// @lc code=end

