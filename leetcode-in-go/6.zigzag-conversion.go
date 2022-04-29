/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] Zigzag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
	cur := 0
	dir := -1

	if numRows == 1 {
		return s
	}
	lines := make([][]rune, numRows)

	for _, s := range s {
		lines[cur] = append(lines[cur], s)
		if cur == numRows-1 || cur == 0 {
			dir = -1 * dir
		}

		cur += dir
	}

	var result string
	for _, line := range lines {
		result += string(line)
	}

	return result
}

// @lc code=end

