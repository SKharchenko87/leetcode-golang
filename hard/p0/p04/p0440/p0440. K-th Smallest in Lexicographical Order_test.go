package p0440

import "testing"

func Test_findKthNumber(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 13, k: 2}, 10},
		{"Example 2", args{n: 1, k: 1}, 1},
		{"TestCase 41", args{n: 957747794, k: 424238336}, 481814499},
		{"My 1", args{n: 957747794, k: 957747794}, 99999999},
		{"My 2", args{n: 9899, k: 9889}, 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthNumber(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findKthNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
