package reverse_integer

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
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
				x: 123,
			},
			want: 321,
		},
		{
			name: "case2",
			args: args{
				x: -123,
			},
			want: -321,
		},
		{
			name: "case3",
			args: args{
				x: 1092320099,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}