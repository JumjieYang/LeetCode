/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := 1000000

	for i := 0; i < len(nums)-2; i++ {

		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if int(math.Abs(float64(target-sum))) < int(math.Abs(float64(target-closest))) {
				closest = sum
			}

			if sum == target {
				return target
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return closest
}

// @lc code=end

