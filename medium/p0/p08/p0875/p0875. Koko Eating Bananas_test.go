package p0875

import "testing"

func Test_minEatingSpeed(t *testing.T) {
	type args struct {
		piles []int
		h     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{[]int{3, 6, 7, 11}, 8}, 4},
		{"Case 2", args{[]int{30, 11, 23, 4, 20}, 5}, 30},
		{"Case 3", args{[]int{30, 11, 23, 4, 20}, 6}, 23},
		{"Case 4", args{[]int{312884470}, 312884469}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEatingSpeed(tt.args.piles, tt.args.h); got != tt.want {
				t.Errorf("minEatingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
