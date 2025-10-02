package p2894

import "testing"

func Test_differenceOfSums(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 10, m: 3}, 19},
		{"Example 2", args{n: 5, m: 6}, 15},
		{"Example 3", args{n: 5, m: 1}, -15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := differenceOfSums(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("differenceOfSums() = %v, want %v", got, tt.want)
			}
		})
	}
}
