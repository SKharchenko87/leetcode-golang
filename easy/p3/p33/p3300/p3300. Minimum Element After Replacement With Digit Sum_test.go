package p3300

import "testing"

func Test_minElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{10, 12, 13, 14}}, 1},
		{"Example 2", args{nums: []int{1, 2, 3, 4}}, 1},
		{"Example 3", args{nums: []int{999, 19, 199}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minElement(tt.args.nums); got != tt.want {
				t.Errorf("minElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
