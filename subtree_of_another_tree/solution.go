package subtree_of_another_tree

import (
	. "LeetCode-Go/common"
	"fmt"
	"strings"
)

/**
给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。

示例 1:
给定的树 s:

     3
    / \
   4   5
  / \
 1   2
给定的树 t：

   4
  / \
 1   2
返回 true，因为 t 与 s 的一个子树拥有相同的结构和节点值。

示例 2:
给定的树 s：

     3
    / \
   4   5
  / \
 1   2
    /
   0
给定的树 t：

   4
  / \
 1   2
返回 false。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subtree-of-another-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	sb1 := strings.Builder{}
	preOrder(&sb1, s, "")
	sb2 := strings.Builder{}
	preOrder(&sb2, t, "")
	return strings.Contains(sb1.String(), sb2.String())
}

func preOrder(builder *strings.Builder, node *TreeNode, tag string)  {
	if node == nil {
		builder.WriteString(tag)
		return
	}
	builder.WriteString(fmt.Sprintf("(%d)", node.Val))
	preOrder(builder, node.Left, "L")
	preOrder(builder, node.Right, "R")
}