package p1792

import "testing"

func Test_maxAverageRatio(t *testing.T) {
	type args struct {
		classes       [][]int
		extraStudents int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Example 1", args{classes: [][]int{{1, 2}, {3, 5}, {2, 2}}, extraStudents: 2}, 0.78333},
		{"Example 2", args{classes: [][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}}, extraStudents: 4}, 0.53485},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAverageRatio(tt.args.classes, tt.args.extraStudents); got != tt.want {
				t.Errorf("maxAverageRatio() = %v, want %v", got, tt.want)
			}
		})
	}
}
