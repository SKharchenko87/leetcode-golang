package p0486

import "testing"

func Test_predictTheWinner(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{nums: []int{1, 5, 2}}, false},
		{"Example 2", args{nums: []int{1, 5, 233, 7}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := predictTheWinner(tt.args.nums); got != tt.want {
				t.Errorf("predictTheWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
