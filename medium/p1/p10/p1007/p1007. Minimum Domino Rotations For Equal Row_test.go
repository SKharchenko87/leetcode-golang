package p1007

import "testing"

func Test_minDominoRotations(t *testing.T) {
	type args struct {
		tops    []int
		bottoms []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{tops: []int{2, 1, 2, 4, 2, 2}, bottoms: []int{5, 2, 6, 2, 3, 2}}, 2},
		{"Example 2", args{tops: []int{3, 5, 1, 2, 3}, bottoms: []int{3, 6, 3, 3, 4}}, -1},
		{"My 1", args{tops: []int{1, 2, 1}, bottoms: []int{2, 1, 2}}, 1},
		{"My 2", args{tops: []int{1, 1, 2, 1}, bottoms: []int{2, 1, 1, 2}}, 1},
		{"My 2", args{tops: []int{2, 1, 1, 2}, bottoms: []int{1, 1, 2, 1}}, 1},
		{"My 2", args{tops: []int{1, 2, 2, 1, 2}, bottoms: []int{1, 1, 1, 2, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDominoRotations(tt.args.tops, tt.args.bottoms); got != tt.want {
				t.Errorf("minDominoRotations() = %v, want %v", got, tt.want)
			}
		})
	}
}
