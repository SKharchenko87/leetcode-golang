package p2110

import "testing"

func Test_getDescentPeriods(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{prices: []int{3, 2, 1, 4}}, 7},
		{"Example 2", args{prices: []int{8, 6, 7, 7}}, 4},
		{"Example 3", args{prices: []int{1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDescentPeriods(tt.args.prices); got != tt.want {
				t.Errorf("getDescentPeriods() = %v, want %v", got, tt.want)
			}
		})
	}
}
