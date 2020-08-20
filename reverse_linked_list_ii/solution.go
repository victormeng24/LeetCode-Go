package reverse_linked_list_ii

import . "LeetCode-Go/common"

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := ListNode{
		Val: -1,
	}
	dummy.Next = head
	scanner := &dummy
	scannerPre := &dummy
	for m > 0 {
		scannerPre = scanner
		scanner = scanner.Next
		m--
		n--
	}
	pre := scanner
	nxt := scanner.Next
	pre.Next = nil
	for nxt != nil && n > 0 {
		tmp := nxt.Next
		nxt.Next = pre
		pre, nxt = nxt, tmp
		n--
	}
	scannerPre.Next = pre
	scanner.Next = nxt
	return dummy.Next
}