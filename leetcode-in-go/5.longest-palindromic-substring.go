/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	var longest string

	for i, _ := range s {
		tempLong := spread(s, i, i)
		even := spread(s, i, i+1)

		if len(tempLong) < len(even) {
			tempLong = even
		}

		if len(longest) < len(tempLong) {
			longest = tempLong
		}

	}

	return longest
}

func spread(s string, i, j int) string {
	for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {
	}

	return s[i+1 : j]
}

// @lc code=end

