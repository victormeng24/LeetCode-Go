package zuo_xuan_zhuan_zi_fu_chuan_lcof

import "testing"

func Test_reverseLeftWords(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				s: "abcdefg",
				n: 2,
			},
			want: "cdefgab",
		},
		{
			name: "case2",
			args: args{
				s: "lrloseumgh",
				n: 6,
			},
			want: "umghlrlose",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLeftWords(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("reverseLeftWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
