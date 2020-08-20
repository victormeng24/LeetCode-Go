package remove_duplicates_from_sorted_list

import . "LeetCode-Go/common"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	scanner := head
	for scanner != nil && scanner.Next != nil {
		if scanner.Val == scanner.Next.Val {
			scanner.Next = scanner.Next.Next
		} else {
			scanner = scanner.Next
		}
	}
	return head
}