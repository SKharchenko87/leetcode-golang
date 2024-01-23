package p2

import (
	"reflect"
	"testing"
)

func Test_countOfPairs(t *testing.T) {
	type args struct {
		n int
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{3, 1, 3}, []int{6, 0, 0}},
		{"2", args{5, 2, 4}, []int{10, 8, 2, 0, 0}},
		{"3", args{4, 1, 1}, []int{6, 4, 2, 0}},
		{"3", args{10, 1, 1}, []int{6, 4, 2, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfPairs(tt.args.n, tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countOfPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
