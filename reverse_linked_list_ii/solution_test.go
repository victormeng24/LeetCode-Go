package reverse_linked_list_ii

import (
	. "LeetCode-Go/common"
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head *ListNode
		m    int
		n    int
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
				m: 2,
				n: 4,
			},
			want: BuildList([]int{1, 4, 3, 2, 5}),
		},
		{
			name: "case2",
			args: args{
				head: BuildList([]int{1, 2, 3, 4, 5}),
				m: 1,
				n: 5,
			},
			want: BuildList([]int{5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}