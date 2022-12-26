package _24

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 反转链表
 * https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/description/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// 迭代
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// 递归

	return prev
}
