package shuffle_the_array

import (
	"reflect"
	"testing"
)

func Test_shuffle(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				nums: []int{2, 5, 1, 3, 4, 7},
				n:    3,
			},
			want: []int{2, 3, 5, 4, 1, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shuffle(tt.args.nums, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}
