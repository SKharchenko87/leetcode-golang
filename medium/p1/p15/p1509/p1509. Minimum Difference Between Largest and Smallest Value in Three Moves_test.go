package p1509

import "testing"

func Test_minDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{5, 3, 2, 4}}, 0},      //
		{"Example 2", args{nums: []int{1, 5, 0, 10, 14}}, 1}, // 0 1 5 10 14
		{"Example 2", args{nums: []int{0, 5, 6, 6, 11}}, 0},  // 0 1 5 10 14
		{"Example 3", args{nums: []int{3, 100, 20}}, 0},
		{"Example 2", args{nums: []int{0, 5, 6, 9, 10, 11}}, 2},
		{"Example 2", args{nums: []int{6, 6, 6, 9, 10, 11}}, 0},
		{"Example 2", args{nums: []int{6, 6, 6, 6, 10, 11}}, 0},
		{"Example 2", args{nums: []int{0, 6, 6, 6, 6, 11}}, 0},
		{"Example 2", args{nums: []int{0, 6, 6, 6, 10, 11}}, 0},
		{"Example 2", args{nums: []int{0, 5, 6, 6, 7, 11}}, 1},
		{"Example 2", args{nums: []int{0, 1, 1, 4, 5, 6, 6, 6}}, 2},
		{"Example 2", args{nums: []int{0, 1, 1, 4, 5, 10, 10, 10}}, 5},
		{"Example 2", args{nums: []int{0, 1, 1, 4, 10, 10, 10}}, 4},
		{"Example 2", args{nums: []int{0, 1, 1, 4, 4, 10, 10}}, 3}, // 0,1,1,4,4,10,10
		{"Example 2", args{nums: []int{0, 1, 7, 8, 9, 10, 10}}, 2}, //
		{"Example XXX", args{nums: []int{0, 1, 1, 4, 5, 8, 8, 10, 10, 10}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDifference(tt.args.nums); got != tt.want {
				t.Errorf("minDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
