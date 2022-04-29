/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)

	for i, val := range nums {
		_, ok := mp[target-val]

		if ok {
			return []int{i, mp[target-val]}
		}

		mp[val] = i
	}

	return nil
}

// @lc code=end

