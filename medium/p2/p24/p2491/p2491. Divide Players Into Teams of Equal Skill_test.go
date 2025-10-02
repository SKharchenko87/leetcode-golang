package p2491

import "testing"

func Test_dividePlayers(t *testing.T) {
	type args struct {
		skill []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{skill: []int{3, 2, 5, 1, 3, 4}}, 22},
		{"Example 2", args{skill: []int{3, 4}}, 12},
		{"Example 3", args{skill: []int{1, 1, 2, 3}}, -1},
		{"Example 3", args{skill: []int{2, 2, 2, 2}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dividePlayers(tt.args.skill); got != tt.want {
				t.Errorf("dividePlayers() = %v, want %v", got, tt.want)
			}
		})
	}
}
