package p3347

import "testing"

func Test_maxFrequency(t *testing.T) {
	type args struct {
		nums          []int
		k             int
		numOperations int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 4, 5}, k: 1, numOperations: 2}, 2},
		{"Example 2", args{nums: []int{5, 11, 20, 20}, k: 5, numOperations: 1}, 2},
		{"TestCase 179", args{nums: []int{7, 2}, k: 79, numOperations: 1}, 2},
		{"TestCase 418", args{nums: []int{999999997, 999999999, 999999999}, k: 999999999, numOperations: 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequency(tt.args.nums, tt.args.k, tt.args.numOperations); got != tt.want {
				t.Errorf("maxFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
