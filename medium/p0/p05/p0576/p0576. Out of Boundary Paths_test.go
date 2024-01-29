package p0576

import "testing"

func Test_findPaths(t *testing.T) {
	type args struct {
		m           int
		n           int
		maxMove     int
		startRow    int
		startColumn int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{2, 2, 2, 0, 0}, 6},
		{"Case 2", args{1, 3, 3, 0, 1}, 12},
		{"Case 3", args{3, 3, 2, 1, 1}, 4},
		{"Case 4", args{3, 3, 3, 1, 1}, 20},
		{"Case 5", args{3, 3, 2, 0, 1}, 5},
		{"Case 6", args{3, 3, 3, 0, 0}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPaths(tt.args.m, tt.args.n, tt.args.maxMove, tt.args.startRow, tt.args.startColumn); got != tt.want {
				t.Errorf("findPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
