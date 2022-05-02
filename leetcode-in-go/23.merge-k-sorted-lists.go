/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}

	mid := (start + end) / 2
	left := merge(lists, start, mid)
	right := merge(lists, mid+1, end)

	return merge2(left, right)
}

func merge2(left, right *ListNode) *ListNode {
	dummyHead := new(ListNode)
	cur := dummyHead

	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}

		cur = cur.Next
	}

	if left != nil {
		cur.Next = left
	}

	if right != nil {
		cur.Next = right
	}

	return dummyHead.Next
}

// @lc code=end

