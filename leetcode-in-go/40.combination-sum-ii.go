/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	ans := [][]int{}
	dpH(candidates, target, 0, nil, &ans)

	return ans
}

func dpH(candidates []int, target, idx int, temp []int, ans *[][]int) {
	if target == 0 {
		t := append([]int{}, temp...)
		*ans = append(*ans, t)
		return
	}

	if target < 0 {
		return
	}

	for i := idx; i < len(candidates); i++ {
		if i > idx && candidates[i] == candidates[i-1] {
			continue
		}

		dpH(candidates, target-candidates[i], i+1, append(temp, candidates[i]), ans)
	}
}

// @lc code=end

