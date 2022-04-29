/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
import "math"

const max int = math.MaxInt32
const min int = math.MinInt32

func reverse(x int) int {

	result := 0
	negative := false
	if x < 0 {
		x = x * -1
		negative = true
	}

	for x > 0 {
		result = result*10 + x%10
		x = x / 10

		if (negative && result*-1 < min) || (!negative && result > max) {
			return 0
		}

	}

	if negative {
		return -1 * result
	}

	return result
}

// @lc code=end

