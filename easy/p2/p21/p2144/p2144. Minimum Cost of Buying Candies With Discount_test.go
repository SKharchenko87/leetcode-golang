package p2144

import "testing"

func Test_minimumCost(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{cost: []int{1, 2, 3}}, 5},
		{"Example 2", args{cost: []int{6, 5, 7, 9, 2, 2}}, 23},
		{"Example 3", args{cost: []int{5, 5}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCost(tt.args.cost); got != tt.want {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
