package p0762

import "testing"

func Test_countPrimeSetBits(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{left: 6, right: 10}, 4},
		{"Example 2", args{left: 10, right: 15}, 5},
		{"TestCase 172", args{left: 248657, right: 257888}, 3381},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimeSetBits(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("countPrimeSetBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
