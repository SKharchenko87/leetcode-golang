package p3100

import "testing"

func Test_maxBottlesDrunk(t *testing.T) {
	type args struct {
		numBottles  int
		numExchange int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{numBottles: 13, numExchange: 6}, 15},
		{"Example 2", args{numBottles: 10, numExchange: 3}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxBottlesDrunk(tt.args.numBottles, tt.args.numExchange); got != tt.want {
				t.Errorf("maxBottlesDrunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
