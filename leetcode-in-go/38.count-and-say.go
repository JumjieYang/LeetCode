/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	stringToSay := countAndSay(n - 1)

	answer := make([]byte, 0)
	for len(stringToSay) > 0 {
		val := stringToSay[0]
		num := 1
		for num < len(stringToSay) && stringToSay[num] == val {
			num++
		}
		stringToSay = stringToSay[num:]

		answer = append(answer, byte('0')+byte(num))
		answer = append(answer, byte(val))
	}

	return string(answer)
}

// @lc code=end

