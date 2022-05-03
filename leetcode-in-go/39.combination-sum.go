/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	backtrack(candidates, target, 0, 0, []int{}, &res)
	return res
}

func backtrack(candidates []int, target int, start int, currSum int, curr []int, res *[][]int) {
	if currSum == target {
		curr2 := make([]int, len(curr))
		copy(curr2, curr)
		*res = append(*res, curr2)
		return
	} else if currSum > target {
		return
	}
	for i := start; i < len(candidates); i++ {
		currSum += candidates[i]
		curr = append(curr, candidates[i])
		backtrack(candidates, target, i, currSum, curr, res)
		currSum -= candidates[i]
		curr = curr[:len(curr)-1]
	}
}

// @lc code=end

