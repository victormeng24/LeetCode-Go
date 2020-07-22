package add_two_number

import (
	"reflect"
	"testing"
)
import . "LeetCode-Go/common"

func Test_addTwoNumbers(t *testing.T) {
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
				l1: BuildList([]int{2, 4, 3}),
				l2: BuildList([]int{5, 6, 4}),
			},
			want: BuildList([]int{7, 0, 8}),
		},
		{
			name: "case2",
			args: args{
				l1: BuildList([]int{0}),
				l2: BuildList([]int{5, 6, 4}),
			},
			want: BuildList([]int{5, 6, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
