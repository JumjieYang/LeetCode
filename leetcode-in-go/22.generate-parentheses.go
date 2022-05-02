/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := make([]string, 0)

	backtracking(&res, n, n, "")

	return res
}

func backtracking(res *[]string, left, right int, cur string) {

	if left == right && left == 0 {
		*res = append(*res, cur)
		return
	}

	if left > 0 {
		backtracking(res, left-1, right, cur+"(")

	}

	if right > left {
		backtracking(res, left, right-1, cur+")")
	}

}

// @lc code=end

