package remove_duplicate_node_icci

import (
	. "LeetCode-Go/common"
	"reflect"
	"testing"
)

func Test_removeDuplicateNodes(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				head: BuildList([]int{1, 2, 3, 3, 2, 1}),
			},
			want: BuildList([]int{1, 2, 3}),
		},
		{
			name: "case2",
			args: args{
				head: BuildList([]int{1, 1, 1, 1, 1, 2}),
			},
			want: BuildList([]int{1, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicateNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
