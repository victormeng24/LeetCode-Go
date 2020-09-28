package distribute_candies_to_people

import (
	"reflect"
	"testing"
)

func Test_distributeCandies(t *testing.T) {
	type args struct {
		candies    int
		num_people int
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
				candies:    7,
				num_people: 4,
			},
			want: []int{1, 2, 3, 1},
		},
		{
			name: "case2",
			args: args{
				candies:    10,
				num_people: 3,
			},
			want: []int{5, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCandies(tt.args.candies, tt.args.num_people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
