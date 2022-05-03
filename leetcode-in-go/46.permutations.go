/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
var output [][]int

func permute(nums []int) [][]int {
	output = [][]int{}
	backtracing(nums, []int{}, len(nums))

	return output
}

func backtracing(remaining []int, curr []int, length int) {
	if len(curr) == length {
		output = append(output, curr)
		return
	}

	for i, el := range remaining {
		if el > -11 {
			cpy := make([]int, len(remaining))
			copy(cpy, remaining)
			cpy[i] = -11
			backtracing(cpy, append(curr, el), length)
		}
	}
}

// @lc code=end

