package _6

// ListNode 节点
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 从尾到头打印链表
 * https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/description/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	if head == nil {
		res := make([]int, 0)
		return res
	}
	return append(reversePrint(head.Next), head.Val)
}
