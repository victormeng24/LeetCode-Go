package minimum_depth_of_binary_tree

import (
	. "LeetCode-Go/common"
	"math"
)
/**
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:

给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-depth-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	ret := math.MaxInt32
	if root.Left != nil {
		ret = min(ret, minDepth(root.Left) + 1)
	}
	if root.Right != nil {
		ret = min(ret, minDepth(root.Right) + 1)
	}
	return ret
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
