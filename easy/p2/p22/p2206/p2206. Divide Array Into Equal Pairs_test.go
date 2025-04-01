package p2206

import "testing"

func Test_divideArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{nums: []int{3, 2, 3, 2, 2, 2}}, true},
		{"Example 2", args{nums: []int{1, 2, 3, 4}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divideArray(tt.args.nums); got != tt.want {
				t.Errorf("divideArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
