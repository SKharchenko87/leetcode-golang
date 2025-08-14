package p1497

import "testing"

func Test_canArrange(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{arr: []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, k: 5}, true},
		{"Example 2", args{arr: []int{1, 2, 3, 4, 5, 6}, k: 7}, true},
		{"Example 3", args{arr: []int{1, 2, 3, 4, 5, 6}, k: 10}, false},
		{"Example 40", args{arr: []int{-1, 1, -2, 2, -3, 3, -4, 4}, k: 3}, true},
		{"Example 92", args{arr: []int{-4, -7, 5, 2, 9, 1, 10, 4, -8, -3}, k: 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canArrange(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("canArrange() = %v, want %v", got, tt.want)
			}
		})
	}
}
