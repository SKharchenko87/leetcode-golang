package p1482

import "testing"

func Test_minDays(t *testing.T) {
	type args struct {
		bloomDay []int
		m        int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{bloomDay: []int{1, 10, 3, 10, 2}, m: 3, k: 1}, 3},
		{"Example 2", args{bloomDay: []int{1, 10, 3, 10, 2}, m: 3, k: 2}, -1},
		{"Example 3", args{bloomDay: []int{7, 7, 7, 7, 12, 7, 7}, m: 2, k: 3}, 12},
		{"My 1", args{bloomDay: []int{7, 7, 7, 7, 12, 7, 7, 7}, m: 2, k: 3}, 7},
		{"Testcase 45", args{bloomDay: []int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, m: 4, k: 2}, 9},
		{"Testcase 82", args{bloomDay: []int{40, 74, 42, 94, 90, 9, 29, 45, 32, 35, 42, 71, 73, 47, 49, 83, 72, 64, 66, 100, 31, 35, 23, 24, 96, 9, 71, 37, 95, 26, 25, 54, 65, 45, 92, 88, 38, 80}, m: 2, k: 13}, 95},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDays(tt.args.bloomDay, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("minDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
