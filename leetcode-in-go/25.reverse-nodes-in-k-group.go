/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	count := 0
	cur := head

	for count < k && cur != nil {
		cur = cur.Next
		count++
	}

	if count == k {
		newHead := reverse(head, k)

		head.Next = reverseKGroup(cur, k)

		return newHead
	}

	return head
}

func reverse(head *ListNode, k int) *ListNode {
	var dummyHead *ListNode
	cur := head
	for k > 0 {
		next := cur.Next
		cur.Next = dummyHead
		dummyHead = cur

		cur = next

		k--
	}

	return dummyHead
}

// @lc code=end

