package p2

import "testing"

func Test_abs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{-1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.x); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countMatchingSubarrays(t *testing.T) {
	type args struct {
		nums    []int
		pattern []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{[]int{1, 2, 3, 4, 5, 6}, []int{1, 1}}, 4},
		{"Case 1", args{[]int{1, 4, 4, 1, 3, 5, 5, 3}, []int{1, 0, -1}}, 2},
		{"Case 1", args{[]int{872500231,190002870}, []int{-1}}, 2},
	[872500231,190002870]
	[-1]
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMatchingSubarrays(tt.args.nums, tt.args.pattern); got != tt.want {
				t.Errorf("countMatchingSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
