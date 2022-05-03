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

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := (left + right) / 2
	l := merge(lists, left, mid)
	r := merge(lists, mid+1, right)

	return merge2(l, r)
}

func merge2(l, r *ListNode) *ListNode {
	dummyHead := new(ListNode)

	cur := dummyHead

	for l != nil && r != nil {
		if l.Val < r.Val {
			cur.Next = l
			l = l.Next
		} else {
			cur.Next = r
			r = r.Next
		}

		cur = cur.Next
	}

	if l == nil {
		cur.Next = r
	}

	if r == nil {
		cur.Next = l
	}

	return dummyHead.Next
}

// @lc code=end

