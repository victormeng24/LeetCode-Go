package binary_tree_level_order_traversal_ii

import (
	. "LeetCode-Go/common"
	"reflect"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				root: func() *TreeNode {
					root := &TreeNode{
						Val: 3,
					}
					root.Left = &TreeNode{
						Val: 9,
					}
					root.Right = &TreeNode{
						Val: 20,
					}
					return root
				}(),
			},
			want: [][]int{
				{9, 20},
				{3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
