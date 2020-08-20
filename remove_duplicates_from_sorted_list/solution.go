package remove_duplicates_from_sorted_list

import . "LeetCode-Go/common"

/**
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
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