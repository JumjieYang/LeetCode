/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	s = strings.Trim(s, " ")
	start := 0
	sign_factor := 1
	if len(s) == 0 {
		return 0
	}
	if s[0] == '+' {
		start = 1
	} else if s[0] == '-' {
		start = 1
		sign_factor = -1
	}

	result := 0
	for i := start; i < len(s); i++ {
		if s[i] < 48 || s[i] > 57 {
			return result * sign_factor
		}
		result = result*10 + int(s[i]-48)

		if result*sign_factor < math.MinInt32 {
			return math.MinInt32
		}

		if result*sign_factor > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return result * sign_factor
}

// @lc code=end
