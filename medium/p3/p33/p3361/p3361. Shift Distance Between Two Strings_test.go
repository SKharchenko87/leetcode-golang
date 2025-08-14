package p3361

import "testing"

func Test_shiftDistance(t *testing.T) {
	type args struct {
		s            string
		t            string
		nextCost     []int
		previousCost []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{s: "abab", t: "baba", nextCost: []int{100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, previousCost: []int{1, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}, 2},
		{"Example 2", args{s: "leet", t: "code", nextCost: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, previousCost: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}, 31},
		{"TestCase 6", args{s: "abcdefghijklmnopqrstuvwxyz", t: "zabcdefghijklmnopqrstuvwxy", nextCost: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, previousCost: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}, 26},
		{"TestCase 27", args{s: "ccdbbcaadb", t: "aadbbdaaac", nextCost: []int{6, 6, 9, 8, 2, 4, 10, 10, 6, 4, 9, 5, 5, 5, 2, 7, 9, 7, 4, 1, 4, 10, 1, 5, 2, 4}, previousCost: []int{0, 4, 5, 6, 7, 10, 5, 5, 8, 1, 1, 10, 7, 8, 10, 8, 7, 2, 3, 3, 6, 3, 0, 6, 4, 0}}, 48},
		{"TestCase 602", args{s: "fszfsfmfs", t: "fmssfmzmf", nextCost: []int{1, 9, 10, 7, 10, 8, 1, 9, 1, 3, 3, 6, 4, 10, 10, 5, 6, 6, 8, 5, 7, 3, 10, 5, 6, 3}, previousCost: []int{7, 1, 2, 10, 2, 1, 5, 8, 6, 8, 6, 7, 4, 6, 5, 1, 1, 5, 8, 10, 1, 8, 4, 2, 3, 6}}, 386},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftDistance(tt.args.s, tt.args.t, tt.args.nextCost, tt.args.previousCost); got != tt.want {
				t.Errorf("shiftDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

//abcdefghijklmnopqrstuvwxyz
//zabcdefghijklmnopqrstuvwxy
