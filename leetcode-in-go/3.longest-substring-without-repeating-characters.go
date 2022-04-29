/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	left, longest, cache := 0, 0, make([]int, 128)
	maxFn := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i, val := range s {
		left = maxFn(left, cache[val]+1)
		longest = maxFn(longest, i-left+1)
		cache[val] = i
	}
	return longest
}

// @lc code=end

