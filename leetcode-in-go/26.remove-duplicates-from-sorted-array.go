/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	slower := 0

	for faster := 1; faster < len(nums); faster++ {
		if nums[faster] != nums[slower] {
			slower++
			nums[slower] = nums[faster]
		}
	}

	return slower + 1

}

// @lc code=end

