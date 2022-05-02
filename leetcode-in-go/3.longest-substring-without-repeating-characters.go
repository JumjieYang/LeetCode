/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	left, longest, cache := 0, 0, make([]int, 128)

	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	for i, val := range s {
		left = max(left, cache[val])
		longest = max(longest, i-left+1)

		cache[val] = i + 1
	}

	return longest
}

// @lc code=end

