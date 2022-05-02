/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func letterCombinations(digits string) []string {
	mp := make(map[byte][]string)

	mp['2'] = []string{"a", "b", "c"}
	mp['3'] = []string{"d", "e", "f"}
	mp['4'] = []string{"g", "h", "i"}
	mp['5'] = []string{"j", "k", "l"}
	mp['6'] = []string{"m", "n", "o"}
	mp['7'] = []string{"p", "q", "r", "s"}
	mp['8'] = []string{"t", "u", "v"}
	mp['9'] = []string{"w", "x", "y", "z"}

	res := make([]string, 0, 10)
	if len(digits) == 0 {
		return res
	}
	backtracking(&res, mp, digits, 0, "")

	return res
}

func backtracking(res *[]string, mp map[byte][]string, digits string, index int, cur string) {
	if index == len(digits) {
		*res = append(*res, cur)
		return
	}

	for _, val := range mp[digits[index]] {
		cur += val
		backtracking(res, mp, digits, index+1, cur)
		cur = cur[:len(cur)-1]
	}
}

// @lc code=end

