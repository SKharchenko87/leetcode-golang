package p1590

import "testing"

func Test_minSubarray(t *testing.T) {
	type args struct {
		nums []int
		p    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{3, 1, 4, 2}, p: 6}, 1},
		{"Example 2", args{nums: []int{6, 3, 5, 2}, p: 9}, 2},
		{"Example 3", args{nums: []int{1, 2, 3}, p: 3}, 0},
		{"My 1", args{nums: []int{1, 2, 3}, p: 1}, 0},
		{"My 2", args{nums: []int{1, 2, 3}, p: 31}, -1},
		{"Testcase 134", args{nums: []int{8, 32, 31, 18, 34, 20, 21, 13, 1, 27, 23, 22, 11, 15, 30, 4, 2}, p: 148}, 7},
		{"Testcase 73", args{nums: []int{4, 4, 2}, p: 7}, -1},
		{"Testcase 45", args{nums: []int{4, 5, 8, 5, 4}, p: 7}, 1},
		{"My 3", args{nums: []int{4, 3, 1, 8, 6, 4}, p: 7}, 3},
		{"My 4", args{nums: []int{2, 4, 1, 1, 8, 6, 4}, p: 7}, 2},
		{"My 5", args{nums: []int{7, 7, 7, 2, 7, 7, 7, 3, 7, 7, 7}, p: 7}, 5},
		{"My 5", args{nums: []int{3, 3, 7, 7, 7, 7, 2, 7, 7, 7, 3, 7, 7, 7, 7, 8}, p: 7}, 5},
		{"My 5", args{nums: []int{3, 3, 7, 7, 7, 7, 6, 7, 7, 7, 6, 7, 7, 7, 7, 8}, p: 7}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubarray(tt.args.nums, tt.args.p); got != tt.want {
				t.Errorf("minSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_minSubarray(b *testing.B) {
	b.ReportAllocs()
	f := minSubarray
	for i := 0; i < b.N; i++ {
		f([]int{8, 32, 31, 18, 34, 20, 21, 13, 1, 27, 23, 22, 11, 15, 30, 4, 2}, 148)
	}
}

func Benchmark_minSubarray2(b *testing.B) {
	b.ReportAllocs()
	f := minSubarray2
	for i := 0; i < b.N; i++ {
		f([]int{8, 32, 31, 18, 34, 20, 21, 13, 1, 27, 23, 22, 11, 15, 30, 4, 2}, 148)
	}
}

func Benchmark_minSubarray1(b *testing.B) {
	b.ReportAllocs()
	f := minSubarray1
	for i := 0; i < b.N; i++ {
		f([]int{8, 32, 31, 18, 34, 20, 21, 13, 1, 27, 23, 22, 11, 15, 30, 4, 2}, 148)
	}
}

func Benchmark_minSubarray0(b *testing.B) {
	b.ReportAllocs()
	f := minSubarray0
	for i := 0; i < b.N; i++ {
		f([]int{8, 32, 31, 18, 34, 20, 21, 13, 1, 27, 23, 22, 11, 15, 30, 4, 2}, 148)
	}
}
