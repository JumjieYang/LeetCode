/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */

// @lc code=start
func nextPermutation(nums []int) {
	lastPeak := -1

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			lastPeak = i
			break
		}
	}

	if lastPeak == -1 {
		sort.Ints(nums)
		return
	}

	idx := lastPeak

	for i := lastPeak; i < len(nums); i++ {
		if nums[i] > nums[lastPeak-1] && nums[i] < nums[idx] {
			idx = i
		}
	}

	nums[idx], nums[lastPeak-1] = nums[lastPeak-1], nums[idx]

	sort.Ints(nums[lastPeak:])
}

// @lc code=end

