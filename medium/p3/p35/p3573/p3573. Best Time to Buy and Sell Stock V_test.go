package p3573

import "testing"

func Test_maximumProfit(t *testing.T) {
	type args struct {
		prices []int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{prices: []int{1, 7, 9, 8, 2}, k: 2}, 14},
		{"Example 2", args{prices: []int{12, 16, 19, 19, 8, 1, 19, 13, 9}, k: 3}, 36},
		{"TestCase 420", args{prices: []int{8, 4, 15, 7, 4, 7, 2, 14, 15}, k: 3}, 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProfit(tt.args.prices, tt.args.k); got != tt.want {
				t.Errorf("maximumProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
