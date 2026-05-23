package p1752

import "testing"

func Test_check(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{nums: []int{3, 4, 5, 1, 2}}, true},
		{"Example 2", args{nums: []int{2, 1, 3, 4}}, false},
		{"Example 3", args{nums: []int{1, 2, 3}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.nums); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func bench(b *testing.B, f func(nums []int) bool) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f([]int{38, 8, 53, 18, 57, 44, 56, 58, 57, 93, 56, 12, 89, 75, 37, 45, 57, 40, 67, 100, 6, 92, 100, 87, 47, 67, 47, 38, 16, 21, 48, 29, 21, 61, 29, 84, 9, 23, 43, 22, 84, 50, 7, 21, 70, 31, 30, 22, 2, 76, 33, 25, 11, 51, 74, 95, 94, 2, 50, 49, 55, 43, 93, 60, 44, 97, 52, 1, 63, 97, 66, 20, 84, 4, 85, 14, 85, 18, 15, 26, 51, 70, 52, 8, 95, 84, 37, 4, 41, 97, 96, 17, 20, 98, 34, 70, 93, 83, 81, 62})
	}
}

func Benchmark_check(b *testing.B) {
	bench(b, check)
}

func Benchmark_check1(b *testing.B) {
	bench(b, check1)
}

func Benchmark_check0(b *testing.B) {
	bench(b, check0)
}
