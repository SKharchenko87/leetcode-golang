package p2818

import "testing"

func Test_maximumScore(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{8, 3, 9, 3, 8}, k: 2}, 81},
		{"Example 2", args{nums: []int{19, 12, 14, 6, 10, 18}, k: 3}, 4788},
		{"TestCase 558", args{nums: []int{1, 7, 11, 1, 5}, k: 14}, 823751938},
		{"TestCase 560", args{nums: []int{1, 1, 2, 18, 1, 9, 3, 1}, k: 32}, 346264255},
		{"TestCase 633", args{nums: []int{3289, 2832, 14858, 22011}, k: 6}, 256720975},
		{"TestCase 654", args{nums: []int{12, 5, 1, 6, 9, 1, 17, 14}, k: 12}, 62996359},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScore(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPrimeScore1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{n: 1}, 1},
		{"Test 2", args{n: 2}, 1},
		{"Test 3", args{n: 3}, 1},
		{"Test 4", args{n: 4}, 2},
		{"Test 9", args{n: 9}, 2},
		{"Test 12", args{n: 12}, 3},
		{"Test 16", args{n: 16}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPrimeScore1(tt.args.n); got != tt.want {
				t.Errorf("getPrimeScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
