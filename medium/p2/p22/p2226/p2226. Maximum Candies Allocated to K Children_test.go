package p2226

import "testing"

func Test_maximumCandies(t *testing.T) {
	type args struct {
		candies []int
		k       int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{candies: []int{5, 8, 6}, k: 3}, 5},
		{"Example 2", args{candies: []int{2, 5}, k: 11}, 0},
		{"TestCase 27", args{candies: []int{4, 7, 5}, k: 16}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCandies(tt.args.candies, tt.args.k); got != tt.want {
				t.Errorf("maximumCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
