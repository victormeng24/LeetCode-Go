package implement_strstr

import "testing"

func Test_kmp(t *testing.T) {
	type args struct {
		s string
		p string
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
				s: "hello",
				p: "ll",
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				s: "aaaaa",
				p: "bba",
			},
			want: -1,
		},
		{
			name: "case3",
			args: args{
				"mississippi",
				"issip",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kmp(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("kmp() = %v, want %v", got, tt.want)
			}
		})
	}
}