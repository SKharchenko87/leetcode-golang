package p1014

import "testing"

func Test_maxScoreSightseeingPair(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{values: []int{8, 1, 5, 2, 6}}, 11},
		{"Example 2", args{values: []int{1, 2}}, 2},
		{"Example 2", args{values: []int{2, 2, 2}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScoreSightseeingPair(tt.args.values); got != tt.want {
				t.Errorf("maxScoreSightseeingPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
