/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */

// @lc code=start
import "bytes"

func intToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var str bytes.Buffer
	for i := 0; i < len(vals) && num != 0; i++ {
		for {
			if num-vals[i] < 0 {
				break
			}
			num -= vals[i]
			str.WriteString(symbols[i])
		}
	}

	return str.String()
}

// @lc code=end

