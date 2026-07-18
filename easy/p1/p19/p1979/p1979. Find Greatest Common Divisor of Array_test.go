package p1979

import "testing"

func Test_findGCD(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{2, 5, 6, 9, 10}}, 2},
		{"Example 2", args{nums: []int{7, 5, 6, 8, 3}}, 1},
		{"Example 3", args{nums: []int{3, 3}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findGCD(tt.args.nums); got != tt.want {
				t.Errorf("findGCD() = %v, want %v", got, tt.want)
			}
		})
	}
}
