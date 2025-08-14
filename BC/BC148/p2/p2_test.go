package p2

import "testing"

func Test_abs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{-1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.x); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCost(t *testing.T) {
	type args struct {
		arr []int
		brr []int
		k   int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{arr: []int{-7, 9, 5}, brr: []int{7, -2, -5}, k: 2}, 13},
		{"Example 2", args{arr: []int{2, 1}, brr: []int{2, 1}, k: 0}, 0},
		{"Example 3", args{arr: []int{3, -4}, brr: []int{4, 5}, k: 29}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.arr, tt.args.brr, tt.args.k); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
