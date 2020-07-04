package add_two_number

import (
	. "LeetCode-Go/common"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummy := &ListNode{Val: -1}
	tmp := dummy
	var carry int
	for l1 != nil || l2 != nil {
		var l1val, l2val int
		if l1 != nil {
			l1val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2val = l2.Val
			l2 = l2.Next
		}
		sum := l1val + l2val + carry
		carry = sum / 10
		tmp.Next = &ListNode{Val: sum % 10}
		tmp = tmp.Next
	}
	if carry != 0 {
		tmp.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
