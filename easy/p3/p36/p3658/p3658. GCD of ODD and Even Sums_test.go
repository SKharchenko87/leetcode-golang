package p3658

import "testing"

func Test_gcdOfOddEvenSums(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 4}, 4},
		{"Example 2", args{n: 5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcdOfOddEvenSums(tt.args.n); got != tt.want {
				t.Errorf("gcdOfOddEvenSums() = %v, want %v", got, tt.want)
			}
		})
	}
}
