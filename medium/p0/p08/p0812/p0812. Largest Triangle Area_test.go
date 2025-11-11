package p0812

import "testing"

func Test_largestTriangleArea(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Example 1", args{points: [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}}, 2.00000},
		{"Example 2", args{points: [][]int{{1, 0}, {0, 0}, {0, 1}}}, 0.50000},
		{"TestCase 51", args{points: [][]int{{-35, 19}, {40, 19}, {27, -20}, {35, -3}, {44, 20}, {22, -21}, {35, 33}, {-19, 42}, {11, 47}, {11, 37}}}, 1799.00000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestTriangleArea(tt.args.points); got != tt.want {
				t.Errorf("largestTriangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
