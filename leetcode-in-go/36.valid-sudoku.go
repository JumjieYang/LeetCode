/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]int, 9)
	cols := make([]map[byte]int, 9)
	boxes := make([]map[byte]int, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]int, 9)
		cols[i] = make(map[byte]int, 9)
		boxes[i] = make(map[byte]int, 9)

	}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			val := board[r][c]

			if val == '.' {
				continue
			}

			if ok := rows[r][val]; ok != 0 {
				return false
			}

			rows[r][val] = 1

			if ok := cols[c][val]; ok != 0 {
				return false
			}
			cols[c][val] = 1

			idx := (r/3)*3 + c/3

			if ok := boxes[idx][val]; ok != 0 {
				return false
			}
			boxes[idx][val] = 1
		}
	}

	return true

}

// @lc code=end

