package p2787

import "testing"

func Test_numberOfWays(t *testing.T) {
	type args struct {
		n int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 10, x: 2}, 1},
		{"Example 2", args{n: 4, x: 1}, 2},
		{"TestCase 568", args{n: 8, x: 3}, 1},
		{"TestCase 572", args{n: 115, x: 1}, 1490528},
		{"TestCase 1167", args{n: 234, x: 1}, 520706283},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWays(tt.args.n, tt.args.x); got != tt.want {
				t.Errorf("numberOfWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
