package palindrome_linked_list

import (
	. "LeetCode-Go/common"
	"reflect"
	"testing"
)

func Test_reverseLinkedList(t *testing.T) {
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
				head: BuildList([]int{1, 2, 3, 4}),
			},
			want: BuildList([]int{4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLinkedList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				head: BuildList([]int{1, 2, 3, 2, 1}),
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				head: BuildList([]int{1, 2, 2, 1}),
			},
			want: true,
		},
		{
			name: "case3",
			args: args{
				BuildList([]int{1, 2}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
