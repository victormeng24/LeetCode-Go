package sum_of_left_leaves

import (
	. "LeetCode-Go/common"
)
/**
计算给定二叉树的所有左叶子之和。

示例：

    3
   / \
  9  20
    /  \
   15   7

在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-left-leaves
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

var res int
func sumOfLeftLeaves(root *TreeNode) int {
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode)  {
	if root == nil {
		return
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		res += root.Left.Val
	}
	dfs(root.Left)
	dfs(root.Right)
}