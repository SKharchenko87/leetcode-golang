package p3761

import "testing"

func Test_minMirrorPairDistance(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{12, 21, 45, 33, 54}}, 1},
		{"Example 2", args{nums: []int{120, 21}}, 1},
		{"Example 3", args{nums: []int{21, 120}}, -1},
		{"TestCase 740", args{nums: []int{12, 34, 46, 21, 12}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMirrorPairDistance(tt.args.nums); got != tt.want {
				t.Errorf("minMirrorPairDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
