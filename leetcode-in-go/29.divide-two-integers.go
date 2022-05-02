/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	negative := (dividend < 0) != (divisor < 0)

	quotient, dividend, divisor := 0, int(math.Abs(float64(dividend))), int(math.Abs(float64(divisor)))

	for dividend >= divisor {
		sub := divisor
		add := 1
		for dividend >= sub<<1 {
			sub <<= 1
			add <<= 1
		}

		dividend -= sub
		quotient += add
	}

	if negative {
		return -quotient
	}

	return quotient
}

// @lc code=end

