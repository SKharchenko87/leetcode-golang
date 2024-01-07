package p1

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
		{"Case 1", args{[][]int{{9, 3}, {8, 6}}}, 48},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areaOfMaxDiagonal(tt.args.dimensions); got != tt.want {
				t.Errorf("areaOfMaxDiagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}
