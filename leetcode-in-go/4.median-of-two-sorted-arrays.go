/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	return float64(findKth(nums1, nums2, 0, m-1, 0, n-1, (m+n+1)/2)+findKth(nums1, nums2, 0, m-1, 0, n-1, (m+n+2)/2)) / float64(2)
}

func findKth(nums1, nums2 []int, start1, end1, start2, end2, k int) int {
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1

	if len1 > len2 {
		return findKth(nums2, nums1, start2, end2, start1, end1, k)
	}

	if len1 == 0 {
		return nums2[start2+k-1]
	}

	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}

	i := start1 + min(len1, k/2) - 1
	j := start2 + min(len2, k/2) - 1

	if nums1[i] < nums2[j] {
		return findKth(nums1, nums2, i+1, end1, start2, end2, k-i+start1-1)
	}

	return findKth(nums1, nums2, start1, end1, j+1, end2, k-j+start2-1)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

