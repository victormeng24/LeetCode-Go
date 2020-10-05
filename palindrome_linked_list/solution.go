package palindrome_linked_list

import (
	. "LeetCode-Go/common"
)

/**
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	secondHalfHead := reverseLinkedList(slow)
	for secondHalfHead != nil {
		if secondHalfHead.Val != head.Val {
			return false
		}
		head, secondHalfHead = head.Next, secondHalfHead.Next
	}
	return true
}

func reverseLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, nxt := head, head.Next
	pre.Next = nil
	for nxt != nil {
		tmp := nxt.Next
		nxt.Next = pre
		pre, nxt = nxt, tmp
	}
	return pre
}
