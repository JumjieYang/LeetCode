/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	cur, max, left, right := 0, 0, 0, len(height)-1

	minFn := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	maxFn := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for {
		if left >= right {
			break
		}

		cur = minFn(height[left], height[right]) * (right - left)
		max = maxFn(cur, max)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}

// @lc code=end

