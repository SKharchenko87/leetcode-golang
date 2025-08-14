package p1823

import "testing"

func Test_findTheWinner(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{"Example 1", args{5, 2}, 3},
		{"Example 2", args{6, 5}, 1},
		{"Example 2", args{8, 8}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheWinner(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findTheWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
