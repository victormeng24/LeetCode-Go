package partition_array_into_three_parts_with_equal_sum

import "testing"

func Test_canThreePartsEqualSum(t *testing.T) {
	type args struct {
		A []int
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
				A: []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1},
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				A: []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canThreePartsEqualSum(tt.args.A); got != tt.want {
				t.Errorf("canThreePartsEqualSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
