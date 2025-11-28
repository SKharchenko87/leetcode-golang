package p3190

import "testing"

func Test_minimumOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 2, 3, 4}}, 3},
		{"Example 2", args{nums: []int{3, 6, 9}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations(tt.args.nums); got != tt.want {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_minimumOperations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minimumOperations([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	}
}

func Benchmark_minimumOperations0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minimumOperations0([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	}
}

func Benchmark_minimumOperationsX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minimumOperationsX([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	}
}
