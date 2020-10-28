package xuan_zhuan_shu_zu_de_zui_xiao_shu_zi_lcof

import "testing"

func Test_minArray(t *testing.T) {
	type args struct {
		numbers []int
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
				numbers: []int{1, 2, 3, 4, 5},
			},
			want: 1,
		},
		{
			name: "case2",
			args: args{
				numbers: []int{2, 2, 2, 0, 1},
			},
			want: 0,
		},
		{
			name: "case3",
			args: args{
				numbers: []int{3, 1, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minArray(tt.args.numbers); got != tt.want {
				t.Errorf("minArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
