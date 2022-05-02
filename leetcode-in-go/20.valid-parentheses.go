/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	mp := make(map[rune]rune)
	stack := make([]rune, 0, 10)
	mp['('] = ')'
	mp['{'] = '}'
	mp['['] = ']'

	for _, val := range s {
		if v, ok := mp[val]; ok {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}

			ele := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if ele != val {
				return false
			}
		}
	}

	return len(stack) == 0

}

// @lc code=end

