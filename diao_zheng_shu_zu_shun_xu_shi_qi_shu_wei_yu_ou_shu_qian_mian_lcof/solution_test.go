package diao_zheng_shu_zu_shun_xu_shi_qi_shu_wei_yu_ou_shu_qian_mian_lcof

import (
	"reflect"
	"testing"
)

func Test_exchange(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{1, 2, 3, 4},
			},
			want: []int{1, 3, 4, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exchange(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exchange() = %v, want %v", got, tt.want)
			}
		})
	}
}
