package p3583

import "testing"

func Test_specialTriplets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{6, 3, 6}}, 1},
		{"Example 2", args{nums: []int{0, 1, 0, 0}}, 1},
		{"Example 3", args{nums: []int{8, 4, 2, 8, 4}}, 2},
		{"TestCase 199", args{nums: []int{28, 52, 14, 28, 34, 26, 14, 52}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := specialTriplets(tt.args.nums); got != tt.want {
				t.Errorf("specialTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
