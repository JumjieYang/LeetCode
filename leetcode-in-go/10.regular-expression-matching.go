/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// @lc code=start
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)

	for i := 0; i <= len(s); i++ {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true

	for i := 2; i <= len(p); i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
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

	return dp[len(s)][len(p)]
}

// @lc code=end

