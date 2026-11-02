package p0462

import "testing"

func Test_minMoves2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 2, 3}}, 2},
		{"Example 2", args{nums: []int{1, 10, 2, 9}}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves2(tt.args.nums); got != tt.want {
				t.Errorf("minMoves2() = %v, want %v", got, tt.want)
			}
		})
	}
}
