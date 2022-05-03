/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	first := findOccur(nums, target, true)

	if first == -1 {
		return []int{-1, -1}
	}

	second := findOccur(nums, target, false)

	return []int{first, second}
}

func findOccur(nums []int, target int, isFirst bool) int {

	start, end := 0, len(nums)-1

	for start <= end {
		mid := (start + end) / 2

		if nums[mid] == target {
			if isFirst {
				if start == mid || nums[mid-1] != target {
					return mid
				}

				end = mid - 1
			} else {
				if end == mid || nums[mid+1] != target {
					return mid
				}

				start = mid + 1
			}
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

// @lc code=end

