package p3623

import "testing"

func Test_countTrapezoids(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{points: [][]int{{1, 0}, {2, 0}, {3, 0}, {2, 2}, {3, 2}}}, 3},
		{"Example 2", args{points: [][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}}}, 1},
		{"TestCase 139", args{points: [][]int{{-17, -23}, {58, 51}, {-16, -3}, {57, -23}, {-90, 51}}}, 1},
		{"TestCase 144", args{points: [][]int{{-27, 63}, {4, -71}, {79, -71}, {28, 41}, {46, 41}}}, 1},
		{"TestCase 200", args{points: [][]int{{-73, -72}, {-1, -56}, {-92, 30}, {-57, -89}, {-19, -89}, {-35, 30}, {64, -72}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTrapezoids(tt.args.points); got != tt.want {
				t.Errorf("countTrapezoids() = %v, want %v", got, tt.want)
			}
		})
	}
}
