package p3568

import "testing"

func Test_minMoves(t *testing.T) {
	type args struct {
		classroom []string
		energy    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{classroom: []string{"S.", "XL"}, energy: 2}, 2},
		{"Example 2", args{classroom: []string{"LS", "RL"}, energy: 4}, 3},
		{"Example 3", args{classroom: []string{"L.S", "RXL"}, energy: 3}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves(tt.args.classroom, tt.args.energy); got != tt.want {
				t.Errorf("minMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
