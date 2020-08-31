package binary_tree_level_order_traversal

import (
	. "LeetCode-Go/common"
	"fmt"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 3,
	}
	res := levelOrder(root)
	fmt.Println(res[0])
	fmt.Println(res[1])
}