/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	premute := make([]int, 0, len(nums))
	res := make([][]int, 0)

	res = traverse(nums, premute, res, len(nums))

	return res
}

func traverse(nums, premute []int, res [][]int, ln int) [][]int {
	if len(premute) == ln {
		premuteCopy := append([]int(nil), premute...)
		res = append(res, premuteCopy)
		return res
	}

	existing := make(map[int]bool, 0)

	for idx := range nums {
		if !existing[nums[idx]] {
			newNums := append([]int(nil), nums[:idx]...)
			newNums = append(newNums, nums[idx+1:]...)

			res = traverse(newNums, append(premute, nums[idx]), res, ln)

			existing[nums[idx]] = true
		}
	}

	return res
}

// @lc code=end

