package p2530

import "testing"

func Test_maxKelements(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{nums: []int{10, 10, 10, 10, 10}, k: 5}, 50},
		{"Example 2", args{nums: []int{1, 10, 3, 3, 3}, k: 3}, 17},
		{"TestCase 6", args{nums: []int{672579538, 806947365, 854095676, 815137524}, k: 3}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxKelements(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxKelements() = %v, want %v", got, tt.want)
			}
		})
	}
}
