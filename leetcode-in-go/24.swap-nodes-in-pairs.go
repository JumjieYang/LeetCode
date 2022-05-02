/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := new(ListNode)
	dummyHead.Next = head
	pre := dummyHead

	for head != nil && head.Next != nil {
		cur, next := head, head.Next

		pre.Next = cur.Next
		cur.Next = next.Next
		next.Next = cur

		pre = cur
		head = head.Next
	}

	return dummyHead.Next
}

// @lc code=end

