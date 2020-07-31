package string_rotation_lcci

import "testing"

func Test_isFlipedString(t *testing.T) {
	type args struct {
		s1 string
		s2 string
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
				s1: "waterbottle",
				s2: "erbottlewat",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFlipedString(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isFlipedString() = %v, want %v", got, tt.want)
			}
		})
	}
}