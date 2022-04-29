/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// @lc code=start
func isMatch(s string, p string) bool {
	ls := len(s) + 1
	lp := len(p) + 1

	dp := make([][]bool, ls)
	for i := 0; i < ls; i++ {
		dp[i] = make([]bool, lp)
	}
	dp[0][0] = true

	for j := 1; j < lp; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i < ls; i++ {
		for j := 1; j < lp; j++ {
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if dp[i][j-2] {
					dp[i][j] = true
				} else if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[ls-1][lp-1]
}

// @lc code=end

