package p3074

import "testing"

func Test_minimumBoxes(t *testing.T) {
	type args struct {
		apple    []int
		capacity []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{apple: []int{1, 3, 2}, capacity: []int{4, 3, 1, 5, 2}}, 2},
		{"Example 2", args{apple: []int{5, 5, 5}, capacity: []int{2, 4, 2, 7}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumBoxes(tt.args.apple, tt.args.capacity); got != tt.want {
				t.Errorf("minimumBoxes() = %v, want %v", got, tt.want)
			}
		})
	}
}
