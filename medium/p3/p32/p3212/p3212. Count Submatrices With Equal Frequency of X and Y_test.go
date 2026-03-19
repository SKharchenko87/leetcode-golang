package p3212

import "testing"

func Test_numberOfSubmatrices(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{grid: [][]byte{{'X', 'Y', '.'}, {'Y', '.', '.'}}}, 3},
		{"Example 2", args{grid: [][]byte{{'X', 'X'}, {'X', 'Y'}}}, 0},
		{"Example 3", args{grid: [][]byte{{'.', '.'}, {'.', '.'}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSubmatrices(tt.args.grid); got != tt.want {
				t.Errorf("numberOfSubmatrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
