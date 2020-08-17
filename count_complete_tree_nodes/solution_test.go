package count_complete_tree_nodes

import (
	"fmt"
	"testing"
)
import . "LeetCode-Go/common"
func Test_countNodes(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	res := countNodes(root)
	fmt.Println(res)
}