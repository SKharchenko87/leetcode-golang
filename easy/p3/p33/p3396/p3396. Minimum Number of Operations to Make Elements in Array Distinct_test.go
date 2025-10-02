package p3396

import "testing"

func Test_minimumOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 2, 3, 4, 2, 3, 3, 5, 7}}, 2},
		{"Example 2", args{nums: []int{4, 5, 6, 4, 4}}, 2},
		{"Example 3", args{nums: []int{6, 7, 8, 9}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations(tt.args.nums); got != tt.want {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
