package remove_duplicate_node_icci

import . "LeetCode-Go/common"

/**
编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。

示例1:

 输入：[1, 2, 3, 3, 2, 1]
 输出：[1, 2, 3]
示例2:

 输入：[1, 1, 1, 1, 2]
 输出：[1, 2]
提示：

链表长度在[0, 20000]范围内。
链表元素在[0, 20000]范围内。
智障进阶，忽略
进阶：

如果不得使用临时缓冲区，该怎么解决？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicate-node-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
	m := make(map[int]bool)
	dummy := &ListNode{
		Val: -1,
	}
	dummy.Next = head
	pre, nxt := dummy, head
	for nxt != nil {
		if m[nxt.Val] {
			pre.Next = pre.Next.Next
		} else {
			m[nxt.Val] = true
			pre = pre.Next
		}
		nxt = nxt.Next
	}
	return dummy.Next
}
