package remove_duplicates_from_sorted_list

import (
	. "LeetCode-Go/common"
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
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
				head: BuildList([]int{1, 1, 2, 3, 3}),
			},
			want: BuildList([]int{1, 2, 3}),
		},
		{
			name: "case2",
			args: args{
				head: BuildList([]int{1, 1, 1, 3, 3}),
			},
			want: BuildList([]int{1, 3}),
		},
		{
			name: "case3",
			args: args{
				head: BuildList([]int{1, 1, 1, 1, 1}),
			},
			want: BuildList([]int{1}),
		},
		{
			name: "case1",
			args: args{
				head: BuildList([]int{1, 2, 3, 4, 5}),
			},
			want: BuildList([]int{1, 2, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}