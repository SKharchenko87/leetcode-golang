package p3394

import "testing"

func Test_checkValidCuts(t *testing.T) {
	type args struct {
		n          int
		rectangles [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{n: 5, rectangles: [][]int{{1, 0, 5, 2}, {0, 2, 2, 4}, {3, 2, 5, 3}, {0, 4, 4, 5}}}, true},
		{"Example 2", args{n: 4, rectangles: [][]int{{0, 0, 1, 1}, {2, 0, 3, 4}, {0, 2, 2, 3}, {3, 0, 4, 3}}}, true},
		{"Example 3", args{n: 4, rectangles: [][]int{{0, 2, 2, 4}, {1, 0, 3, 2}, {2, 2, 3, 4}, {3, 0, 4, 2}, {3, 2, 4, 4}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidCuts(tt.args.n, tt.args.rectangles); got != tt.want {
				t.Errorf("checkValidCuts() = %v, want %v", got, tt.want)
			}
		})
	}
}
