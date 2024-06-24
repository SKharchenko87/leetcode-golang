package p0995

import "testing"

func Test_minKBitFlips(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{0, 1, 0}, k: 1}, 2},
		{"Example 2", args{nums: []int{1, 1, 0}, k: 2}, -1},
		{"Example 3", args{nums: []int{0, 0, 0, 1, 0, 1, 1, 0}, k: 3}, 3},
		{"My 1", args{nums: []int{0, 1, 0, 1, 0, 1, 0, 1}, k: 2}, 4},
		{"My 2", args{nums: []int{1, 0, 1, 0, 1, 0, 1, 0}, k: 2}, 4},
		{"My 3", args{nums: []int{0, 0, 1, 0, 0, 1, 0, 0, 1}, k: 3}, -1},
		{"My 4", args{nums: []int{0, 1, 0, 0, 0, 1, 0, 0, 1}, k: 3}, -1},
		{"My 5", args{nums: []int{1, 0, 0, 0, 0, 1, 0, 0, 1, 0}, k: 3}, 5},
		//                              0, 1, 1, 1, 1, 2  3  3  2  1
		//                                 |        |  |  |  |
		{"My 5", args{nums: []int{1, 0, 0, 0, 0, 1, 0, 0, 1}, k: 3}, -1},
		//                              0, 1, 1, 1, 1, 2, 3	 3  2
		{"My 6", args{nums: []int{1, 0, 0, 0, 1, 0, 0, 0, 1}, k: 3}, 2},
		//                              0, 1, 1, 1, 0, 1, 1	 1  0
		{"My 6", args{nums: []int{1, 0, 0, 0, 0, 0, 0, 1, 1}, k: 3}, 2},
		{"My 6", args{nums: []int{1, 0, 0, 1, 1, 1, 0, 1, 1}, k: 3}, 3},
		//                              0, 1, 1, 2, 2, 2, 1	 0  0
		{"My 7", args{nums: []int{0, 0, 1, 1, 1, 0, 1, 1, 1}, k: 3}, 3},
		{"My 7", args{nums: []int{0, 0, 1, 0, 0, 1, 1, 1, 1}, k: 3}, 2},
		{"My 7", args{nums: []int{0, 0, 0, 1, 0, 0, 0}, k: 4}, 2},
		//                              1, 1, 1, 2, 1, 1, 1
		{"My 7", args{nums: []int{0, 0, 1, 1, 0, 0}, k: 4}, 2},
		{"My 7", args{nums: []int{0, 1, 1, 1, 0}, k: 4}, 2},
		{"My 7", args{nums: []int{1, 1, 1, 1}, k: 4}, 0},
		{"My 7", args{nums: []int{0, 1, 0, 1, 0, 0, 0}, k: 4}, -1},
		//                              1, 2, 3, 4, 3, 3(-1), 3
		{"My 7", args{nums: []int{0, 1, 0, 1, 0, 1, 0}, k: 4}, 4},
		//                              1, 2, 3, 4, 3, 2, 1
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minKBitFlips(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minKBitFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
