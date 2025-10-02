package p0837

import "testing"

func Test_new21Game(t *testing.T) {
	type args struct {
		n      int
		k      int
		maxPts int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Example 1", args{n: 10, k: 1, maxPts: 10}, 1.00000},
		{"Example 2", args{n: 6, k: 1, maxPts: 10}, 0.60000},
		{"Example 3", args{n: 21, k: 17, maxPts: 10}, 0.73278},
		{"TestCase 129", args{n: 1, k: 0, maxPts: 1}, 1.00000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := new21Game(tt.args.n, tt.args.k, tt.args.maxPts); got != tt.want {
				t.Errorf("new21Game() = %v, want %v", got, tt.want)
			}
		})
	}
}
