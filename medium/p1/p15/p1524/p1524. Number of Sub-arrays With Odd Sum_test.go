package p1524

import "testing"

func Test_numOfSubarrays(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{arr: []int{1, 3, 5}}, 4},
		{"Example 2", args{arr: []int{2, 4, 6}}, 0},
		{"Example 3", args{arr: []int{1, 2, 3, 4, 5, 6, 7}}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfSubarrays(tt.args.arr); got != tt.want {
				t.Errorf("numOfSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
