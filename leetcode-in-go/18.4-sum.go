/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0, 10)
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)

	kSum(&res, nums, []int{}, 4, target)

	return res
}

func kSum(res *[][]int, nums, curRes []int, k, target int) {
	if k == 2 {
		left, right := 0, len(nums)-1

		for left < right {
			sum := nums[left] + nums[right]
			newRes := curRes
			if sum == target {
				newRes = append(newRes, nums[left], nums[right])
				*res = append(*res, newRes)

				for left+1 < len(nums) && nums[left] == nums[left+1] {
					left++
				}

				for right-1 >= 0 && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}

		return
	}

	lastVisited := 1000000001
	for i := 0; i < len(nums)-(k-1); i++ {
		if nums[i] == lastVisited {
			continue
		}
		curRes = append(curRes, nums[i])
		kSum(res, nums[i+1:], curRes, k-1, target-nums[i])
		curRes = curRes[:len(curRes)-1]
		lastVisited = nums[i]
	}

	return
}

// @lc code=end

