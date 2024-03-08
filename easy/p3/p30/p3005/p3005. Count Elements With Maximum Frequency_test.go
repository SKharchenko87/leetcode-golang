package p3005

import "testing"

func Test_maxFrequencyElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{1, 2, 2, 3, 1, 4}}, 4},
		{"Example 2", args{[]int{1, 2, 3, 4, 5}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequencyElements(tt.args.nums); got != tt.want {
				t.Errorf("maxFrequencyElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
