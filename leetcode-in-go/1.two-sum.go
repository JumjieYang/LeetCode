/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)

	for i, val := range nums {
		if idx, ok := mp[target-val]; ok {
			return []int{idx, i}
		}

		mp[val] = i
	}

	return nil
}

// @lc code=end

