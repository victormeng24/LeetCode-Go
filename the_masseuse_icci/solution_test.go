package the_masseuse_icci

import "testing"

func Test_massage(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: 4,
		},
		{
			name: "case2",
			args: args{
				nums: []int{2, 7, 9, 3, 1},
			},
			want: 12,
		},
		{
			name: "case3",
			args: args{
				nums: []int{2, 1, 4, 5, 3, 1, 1, 3},
			},
			want: 12,
		},
		{
			name: "case1",
			args: args{
				nums: []int{2, 1, 1, 2},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := massage(tt.args.nums); got != tt.want {
				t.Errorf("massage() = %v, want %v", got, tt.want)
			}
		})
	}
}
