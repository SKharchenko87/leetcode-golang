package p3625

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
		{"Example 1", args{points: [][]int{{-3, 2}, {3, 0}, {2, 3}, {3, 2}, {2, -3}}}, 2},
		{"Example 2", args{points: [][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}}}, 1},
		{"TestCase 64", args{points: [][]int{{-32, 12}, {-32, -94}, {-32, -15}, {-30, 88}}}, 0},
		{"TestCase 338", args{points: [][]int{{-100, -23}, {-3, 42}, {56, -26}, {-3, -4}, {-99, -23}, {35, -53}, {-36, 34}}}, 0},
		{"TestCase 426", args{points: [][]int{{71, -89}, {-75, -89}, {-9, 11}, {-24, -89}, {-51, -89}, {-77, -89}, {42, 11}}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTrapezoids(tt.args.points); got != tt.want {
				t.Errorf("countTrapezoids() = %v, want %v", got, tt.want)
			}
		})
	}
}
