package p3718

import "testing"

func Test_missingMultiple(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{8, 2, 3, 4, 6}, 2}, 10},
		{"Example 2", args{[]int{1, 4, 7, 10, 15}, 5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingMultiple(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("missingMultiple() = %v, want %v", got, tt.want)
			}
		})
	}
}
