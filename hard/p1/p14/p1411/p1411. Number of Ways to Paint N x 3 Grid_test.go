package p1411

import "testing"

func Test_numOfWays(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 1}, 12},
		{"Example 2", args{n: 5000}, 30228214},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfWays(tt.args.n); got != tt.want {
				t.Errorf("numOfWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
