package p3021

import "testing"

func Test_flowerGame(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{n: 3, m: 2}, 3},
		{"Example 2", args{n: 1, m: 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flowerGame(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("flowerGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
