package p1

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
		{"Case 1", args{[]int{1, 2, 2, 3, 1, 4}}, 4},
		{"Case 1", args{[]int{10, 12, 11, 9, 6, 19, 11}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequencyElements(tt.args.nums); got != tt.want {
				t.Errorf("maxFrequencyElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
