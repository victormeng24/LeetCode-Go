package remove_linked_list_elements

import . "LeetCode-Go/common"

/**
删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Val: -1,
	}
	dummy.Next = head
	pre := dummy
	nxt := head
	for nxt != nil {
		if nxt.Val == val {
			pre.Next = nxt.Next
		} else {
			pre = pre.Next
		}
		nxt = nxt.Next
	}
	return dummy.Next
}
