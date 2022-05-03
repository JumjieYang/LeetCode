/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

// @lc code=start
func rotate(matrix [][]int) {
	for i := range matrix {
		reverse(matrix, i)
	}

	r := len(matrix) - 1
	c := 0
	for r >= 0 && c < len(matrix[0]) {
		row := r - 1
		col := c + 1
		for row >= 0 && col < len(matrix[0]) {
			matrix[row][c], matrix[r][col] = matrix[r][col], matrix[row][c]
			row--
			col++
		}

		r--
		c++
	}
}

func reverse(matrix [][]int, rowIdx int) {
	row := matrix[rowIdx]
	l := 0
	r := len(row) - 1
	for l < r {
		row[l], row[r] = row[r], row[l]
		l++
		r--
	}
	matrix[rowIdx] = row
}

// @lc code=end

