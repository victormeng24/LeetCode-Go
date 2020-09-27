package compress_string_lcci

import (
	"testing"
)

func Test_compressString(t *testing.T) {
	type args struct {
		S string
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
				S: "aabbcccccdd",
			},
			want: "a2b2c5d2",
		},
		{
			name: "case2",
			args: args{
				S: "abbcd",
			},
			want: "abbcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressString(tt.args.S); got != tt.want {
				t.Errorf("compressString() = %v, want %v", got, tt.want)
			}
		})
	}
}
