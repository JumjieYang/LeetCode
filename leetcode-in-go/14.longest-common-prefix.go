/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if strs == nil {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	low, high := 1, 201

	for _, val := range strs {
		if len(val) < high {
			high = len(val)
		}
	}

	for low <= high {
		mid := (low + high) / 2

		if helper(strs, mid) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return strs[0][:(low+high)/2]
}

func helper(strs []string, length int) bool {
	prefix := strs[0][:length]

	for i := 1; i < len(strs); i++ {
		if strs[i][:length] != prefix {
			return false
		}
	}

	return true
}

// @lc code=end

