package merge_two_sorted_lists

import (
	. "LeetCode-Go/common"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
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
				l1: BuildList([]int{1, 2, 4}),
				l2: BuildList([]int{1, 3, 4}),
			},
			want: BuildList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name: "case2",
			args: args{
				l1: BuildList([]int{1, 2, 4}),
				l2: nil,
			},
			want: BuildList([]int{1, 2, 4}),
		},
		{
			name: "case1",
			args: args{
				l1: nil,
				l2: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
