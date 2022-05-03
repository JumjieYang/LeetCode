/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */

// @lc code=start
func longestValidParentheses(s string) int {
	left, right, longest := 0, 0, 0

	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	for _, val := range s {
		if val == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			longest = max(longest, 2*left)
		}

		if right > left {
			left = 0
			right = 0
		}
	}

	left, right = 0, 0

	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			longest = max(longest, 2*left)
		}

		if left > right {
			left = 0
			right = 0
		}
	}

	return longest
}

// @lc code=end

