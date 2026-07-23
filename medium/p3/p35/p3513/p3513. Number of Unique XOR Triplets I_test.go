package p3513

import "testing"

func Test_uniqueXorTriplets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 2}}, 2},
		{"Example 2", args{nums: []int{3, 1, 2}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueXorTriplets(tt.args.nums); got != tt.want {
				t.Errorf("uniqueXorTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
