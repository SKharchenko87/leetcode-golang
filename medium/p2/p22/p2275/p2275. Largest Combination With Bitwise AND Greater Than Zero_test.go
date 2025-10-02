package p2275

import "testing"

func Test_largestCombination(t *testing.T) {
	type args struct {
		candidates []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{candidates: []int{16, 17, 71, 62, 12, 24, 14}}, 4},
		{"Example 2", args{candidates: []int{8, 8}}, 2},
		{"Example 3", args{candidates: []int{10, 72, 58, 33, 36, 70, 12, 88, 9, 48, 78, 97, 87, 19, 78, 9, 47, 73}}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestCombination(tt.args.candidates); got != tt.want {
				t.Errorf("largestCombination() = %v, want %v", got, tt.want)
			}
		})
	}
}
