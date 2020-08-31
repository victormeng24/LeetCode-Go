package binary_tree_level_order_traversal

import . "LeetCode-Go/common"
/**
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

 

示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
var res [][]int
func levelOrder(root *TreeNode) [][]int {
	res = make([][]int, 0)
	if root != nil {
		dfs( 0, root)
	}
	return res
}

func dfs(depth int, node *TreeNode)  {
	if depth == len(res) {
		res = append(res, []int{})
	}
	res[depth] = append(res[depth], node.Val)
	if node.Left != nil {
		dfs(depth + 1, node.Left)
	}
	if node.Right != nil {
		dfs(depth + 1, node.Right)
	}
}