/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	rev := 0
	temp := x
	for x != 0 {
		rev = rev*10 + x%10
		x /= 10
	}

	return temp == rev
}

// @lc code=end

