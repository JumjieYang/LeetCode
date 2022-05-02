/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	mp := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	num, lastNum := 0, 1001

	for _, val := range s {
		if mp[val] > lastNum {
			num -= 2 * lastNum
		}

		num += mp[val]
		lastNum = mp[val]
	}

	return num

}

// @lc code=end

