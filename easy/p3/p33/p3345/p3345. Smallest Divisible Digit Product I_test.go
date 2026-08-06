package p3345

import "testing"

func Test_smallestNumber(t *testing.T) {
	type args struct {
		n int
		t int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 10, t: 2}, 10},
		{"Example 2", args{n: 15, t: 3}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestNumber(tt.args.n, tt.args.t); got != tt.want {
				t.Errorf("smallestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
