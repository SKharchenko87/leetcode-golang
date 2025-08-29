package p3000

import "testing"

func Test_areaOfMaxDiagonal(t *testing.T) {
	type args struct {
		dimensions [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{dimensions: [][]int{{9, 3}, {8, 6}}}, 48},
		{"Example 2", args{dimensions: [][]int{{3, 4}, {4, 3}}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areaOfMaxDiagonal(tt.args.dimensions); got != tt.want {
				t.Errorf("areaOfMaxDiagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}
