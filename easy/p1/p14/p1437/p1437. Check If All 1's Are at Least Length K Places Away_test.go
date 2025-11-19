package p1437

import "testing"

func Test_kLengthApart(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{nums: []int{1, 0, 0, 0, 1, 0, 0, 1}, k: 2}, true},
		{"Example 2", args{nums: []int{1, 0, 0, 1, 0, 1}, k: 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kLengthApart(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("kLengthApart() = %v, want %v", got, tt.want)
			}
		})
	}
}
