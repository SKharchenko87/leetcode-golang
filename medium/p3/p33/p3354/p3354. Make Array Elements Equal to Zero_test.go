package p3354

import "testing"

func Test_countValidSelections(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 0, 2, 0, 3}}, 2},
		{"Example 2", args{nums: []int{2, 3, 4, 0, 4, 1, 0}}, 0},
		{"TestCase 285", args{nums: []int{16, 13, 10, 0, 0, 0, 10, 6, 7, 8, 7}}, 3},
		{"TestCase 5xx", args{nums: []int{0}}, 2},
		{"TestCase 549", args{nums: []int{0, 20, 15, 0, 0, 0, 17, 0}}, 0},
		{"TestCase 579", args{nums: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countValidSelections(tt.args.nums); got != tt.want {
				t.Errorf("countValidSelections() = %v, want %v", got, tt.want)
			}
		})
	}
}
