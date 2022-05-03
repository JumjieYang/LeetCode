/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func helper(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	res := helper(x, n/2)
	res = res * res

	if n%2 != 0 {
		res = res * x
	}

	return res
}

func myPow(x float64, n int) float64 {
	var res = helper(x, int(math.Abs(float64(n))))

	if n < 0 {
		return 1 / res
	}
	return res
}

// @lc code=end

