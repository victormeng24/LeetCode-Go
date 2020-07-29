package length_of_last_word

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
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
				s: "wqej wqe we1",
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				s: "a        ",
			},
			want: 1,
		},
		{
			name: "case3",
			args: args{
				s: "Hello World",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}