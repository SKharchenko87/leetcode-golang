package p2874

import "testing"

func Test_maximumTripletValue(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{nums: []int{12, 6, 1, 2, 7}}, 77},
		{"Example 2", args{nums: []int{1, 10, 3, 4, 19}}, 133},
		{"Example 3", args{nums: []int{1, 2, 3}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTripletValue(tt.args.nums); got != tt.want {
				t.Errorf("maximumTripletValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
