package p0322

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{coins: []int{1, 2, 5}, amount: 11}, 3},
		{"Example 2", args{coins: []int{2}, amount: 3}, -1},
		{"Example 3", args{coins: []int{1}, amount: 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
