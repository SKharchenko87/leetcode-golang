package p3341

import "testing"

func Test_minTimeToReach(t *testing.T) {
	type args struct {
		moveTime [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{moveTime: [][]int{{0, 4}, {4, 4}}}, 6},
		{"Example 2", args{moveTime: [][]int{{0, 0, 0}, {0, 0, 0}}}, 3},
		{"Example 3", args{moveTime: [][]int{{0, 1}, {1, 2}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTimeToReach(tt.args.moveTime); got != tt.want {
				t.Errorf("minTimeToReach() = %v, want %v", got, tt.want)
			}
		})
	}
}
