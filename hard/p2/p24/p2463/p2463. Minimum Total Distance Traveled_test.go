package p2463

import "testing"

func Test_minimumTotalDistance(t *testing.T) {
	type args struct {
		robot   []int
		factory [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{robot: []int{0, 4, 6}, factory: [][]int{{2, 2}, {6, 2}}}, 4},
		{"Example 2", args{robot: []int{1, -1}, factory: [][]int{{-2, 1}, {2, 1}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotalDistance(tt.args.robot, tt.args.factory); got != tt.want {
				t.Errorf("minimumTotalDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
