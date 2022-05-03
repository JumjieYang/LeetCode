/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	jumps, curEnd, farthest := 0, 0, 0
	max := func(x, y int) int {
		if x < y {
			return y
		}

		return x
	}
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])

		if i == curEnd {
			jumps++
			curEnd = farthest
		}
	}

	return jumps
}

// @lc code=end

