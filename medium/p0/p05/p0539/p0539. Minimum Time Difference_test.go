package p0539

import "testing"

func Test_findMinDifference(t *testing.T) {
	type args struct {
		timePoints []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{timePoints: []string{"23:59", "00:00"}}, 1},
		{"Example 2", args{timePoints: []string{"00:00", "23:59", "00:00"}}, 0},
		{"TestCase 84", args{timePoints: []string{"01:01", "02:01"}}, 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinDifference(tt.args.timePoints); got != tt.want {
				t.Errorf("findMinDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
