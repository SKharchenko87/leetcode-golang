package p1039

import "testing"

func Test_minScoreTriangulation(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{values: []int{1, 2, 3}}, 6},
		{"Example 2", args{values: []int{3, 7, 4, 5}}, 144},
		{"Example 3", args{values: []int{1, 3, 1, 4, 1, 5}}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minScoreTriangulation(tt.args.values); got != tt.want {
				t.Errorf("minScoreTriangulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
