package p2401

import "testing"

func Test_longestNiceSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 3, 8, 48, 10}}, 3},
		{"Example 2", args{nums: []int{3, 1, 5, 11, 13}}, 1},
		{"TestCase 34", args{nums: []int{744437702, 379056602, 145555074, 392756761, 560864007, 934981918, 113312475, 1090, 16384, 33, 217313281, 117883195, 978927664}}, 3},
		{"TestCase 61", args{nums: []int{8, 4096, 524288, 16777216, 2097152, 1024, 4194304, 32768, 1048576, 65536, 4, 536870912, 1, 134217728, 128, 256, 8388608, 2, 8192, 2048, 16384, 16, 64, 33554432, 131072, 512, 262144, 32, 67108864}}, 29},
		{"TestCase XX", args{nums: []int{8, 4096}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestNiceSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestNiceSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
