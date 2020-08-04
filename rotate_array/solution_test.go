package rotate_array

import (
	"fmt"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k: 3,
			},
		},
		{
			name: "case2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k: 2,
			},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			fmt.Println(tt.args.nums)
		})
	}
}