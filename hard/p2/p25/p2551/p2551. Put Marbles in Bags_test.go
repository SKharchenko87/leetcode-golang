package p2551

import "testing"

func Test_putMarbles(t *testing.T) {
	type args struct {
		weights []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{weights: []int{1, 3, 5, 1}, k: 2}, 4},
		{"Example 2", args{weights: []int{1, 3}, k: 2}, 0},
		{"Test 0", args{weights: []int{100, 300, 50, 1, 1}, k: 2}, 398},
		{"Test 1", args{weights: []int{1, 3, 5, 1}, k: 1}, 0},
		{"Test 2", args{weights: []int{1000000000, 1000000000, 1000000000, 1000000000}, k: 3}, 0},
		{"Test 3", args{weights: []int{1, 4, 2, 5, 2}, k: 3}, 3},
		{"Test 4", args{weights: []int{24, 16, 62, 27, 8, 3, 70, 55, 13, 34, 9, 29, 10}, k: 11}, 168},
		{"Test 5", args{weights: []int{68, 65, 5, 74, 12, 67, 10, 55, 27, 38, 69, 54, 62, 50, 30, 3, 1, 24, 39, 65, 73, 33, 43, 9, 64}, k: 9}, 562},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := putMarbles(tt.args.weights, tt.args.k); got != tt.want {
				t.Errorf("putMarbles() = %v, want %v", got, tt.want)
			}
		})
	}
}
