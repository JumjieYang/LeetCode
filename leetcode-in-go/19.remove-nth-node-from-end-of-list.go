/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	faster, slower := head, head

	for i := 0; i < n; i++ {
		faster = faster.Next
	}

	if faster == nil {
		return head.Next
	}

	for faster.Next != nil {
		faster = faster.Next
		slower = slower.Next
	}

	slower.Next = slower.Next.Next

	return head
}

// @lc code=end

