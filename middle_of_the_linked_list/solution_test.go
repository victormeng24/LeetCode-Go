package middle_of_the_linked_list

import (
	. "LeetCode-Go/common"
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
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
				head: BuildList([]int{1, 2, 3, 4, 5}),
			},
			want: BuildList([]int{3, 4, 5}),
		},
		{
			name: "case2",
			args: args{
				head: BuildList([]int{1, 2, 3, 4, 5, 6}),
			},
			want: BuildList([]int{4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
